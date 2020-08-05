// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build linux darwin windows

package kubernetes

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/autodiscover"
	"github.com/elastic/beats/v7/libbeat/autodiscover/template"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/bus"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes/k8skeystore"
	"github.com/elastic/beats/v7/libbeat/keystore"
	"github.com/elastic/beats/v7/libbeat/logp"
)

func init() {
	autodiscover.Registry.AddProvider("kubernetes", AutodiscoverBuilder)
}

// Eventer allows defining ways in which kubernetes resource events are observed and processed
type Eventer interface {
	kubernetes.ResourceEventHandler
	GenerateHints(event bus.Event) bus.Event
	Start() error
	Stop()
}

// Provider implements autodiscover provider for docker containers
type Provider struct {
	config               *Config
	bus                  bus.Bus
	templates            template.Mapper
	builders             autodiscover.Builders
	appenders            autodiscover.Appenders
	logger               *logp.Logger
	eventer              Eventer
	leaderElection       leaderelection.LeaderElectionConfig
	cancelLeaderElection context.CancelFunc
}

// AutodiscoverBuilder builds and returns an autodiscover provider
func AutodiscoverBuilder(bus bus.Bus, uuid uuid.UUID, c *common.Config, keystore keystore.Keystore) (autodiscover.Provider, error) {
	logger := logp.NewLogger("autodiscover")

	errWrap := func(err error) error {
		return errors.Wrap(err, "error setting up kubernetes autodiscover provider")
	}

	config := defaultConfig()
	err := c.Unpack(&config)
	if err != nil {
		return nil, errWrap(err)
	}

	client, err := kubernetes.GetKubernetesClient(config.KubeConfig)
	if err != nil {
		return nil, errWrap(err)
	}

	k8sKeystoreProvider := k8skeystore.NewKubernetesKeystoresRegistry(logger, client)

	mapper, err := template.NewConfigMapper(config.Templates, keystore, k8sKeystoreProvider)
	if err != nil {
		return nil, errWrap(err)
	}

	builders, err := autodiscover.NewBuilders(config.Builders, config.Hints, k8sKeystoreProvider)
	if err != nil {
		return nil, errWrap(err)
	}

	appenders, err := autodiscover.NewAppenders(config.Appenders)
	if err != nil {
		return nil, errWrap(err)
	}

	p := &Provider{
		config:    config,
		bus:       bus,
		templates: mapper,
		builders:  builders,
		appenders: appenders,
		logger:    logger,
	}

	if p.config.Unique {
		p.initLeaderElectionConfig(client, uuid.String())
	} else {
		switch config.Resource {
		case "pod":
			p.eventer, err = NewPodEventer(uuid, c, client, p.publish)
		case "node":
			p.eventer, err = NewNodeEventer(uuid, c, client, p.publish)
		case "service":
			p.eventer, err = NewServiceEventer(uuid, c, client, p.publish)
		default:
			return nil, fmt.Errorf("unsupported autodiscover resource %s", config.Resource)
		}

		if err != nil {
			return nil, errWrap(err)
		}
	}

	return p, nil
}

// Start for Runner interface.
func (p *Provider) Start() {
	if p.config.Unique {
		ctx, cancel := context.WithCancel(context.Background())
		p.cancelLeaderElection = cancel
		p.StartLeaderElector(ctx, p.leaderElection)
	} else {
		if err := p.eventer.Start(); err != nil {
			p.logger.Errorf("Error starting kubernetes autodiscover provider: %s", err)
		}
	}
}

// StartLeaderElector starts a Leader Elector in the background with the provided config
func (p *Provider) StartLeaderElector(ctx context.Context, lec leaderelection.LeaderElectionConfig) {
	le, err := leaderelection.NewLeaderElector(lec)
	if err != nil {
		p.logger.Errorf("leader election lock GAINED, id %v", err)
	}
	if lec.WatchDog != nil {
		lec.WatchDog.SetLeaderElection(le)
	}
	p.logger.Debugf("Starting Leader Elector")
	go le.Run(ctx)
}

// Stop signals the stop channel to force the watch loop routine to stop.
func (p *Provider) Stop() {
	if p.eventer != nil {
		p.eventer.Stop()
	}
	if p.cancelLeaderElection != nil {
		p.cancelLeaderElection()
	}
}

// String returns a description of kubernetes autodiscover provider.
func (p *Provider) String() string {
	return "kubernetes"
}

func (p *Provider) publish(event bus.Event) {
	// Try to match a config
	if config := p.templates.GetConfig(event); config != nil {
		event["config"] = config
	} else {
		// If there isn't a default template then attempt to use builders
		e := p.eventer.GenerateHints(event)
		if config := p.builders.GetConfig(e); config != nil {
			event["config"] = config
		}
	}

	// Call all appenders to append any extra configuration
	p.appenders.Append(event)
	p.bus.Publish(event)
}

func (p *Provider) startLeading(uuid string, eventID string) {
	event := bus.Event{
		"start":    true,
		"provider": uuid,
		"id":       eventID,
		"unique":   "true",
	}
	if config := p.templates.GetConfig(event); config != nil {
		event["config"] = config
	}
	p.bus.Publish(event)
}

func (p *Provider) stopLeading(uuid string, eventID string) {
	event := bus.Event{
		"stop":     true,
		"provider": uuid,
		"id":       eventID,
		"unique":   "true",
	}
	if config := p.templates.GetConfig(event); config != nil {
		event["config"] = config
	}
	p.bus.Publish(event)
}

func (p *Provider) initLeaderElectionConfig(client k8s.Interface, uuid string) {
	var id string
	if p.config.Node != "" {
		id = "beats-leader-" + p.config.Node
	} else {
		id = "beats-leader-" + uuid
	}
	lease := metav1.ObjectMeta{
		Name:      p.config.LeaderLease,
		Namespace: "default",
	}
	metaUID := lease.GetObjectMeta().GetUID()
	p.leaderElection = leaderelection.LeaderElectionConfig{
		Lock: &resourcelock.LeaseLock{
			LeaseMeta: lease,
			Client:    client.CoordinationV1(),
			LockConfig: resourcelock.ResourceLockConfig{
				Identity: id,
			},
		},
		ReleaseOnCancel: true,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				p.logger.Debugf("leader election lock GAINED, id %v", id)
				eventID := fmt.Sprintf("%v-%v", metaUID, time.Now().UnixNano())
				p.startLeading(uuid, eventID)
			},
			OnStoppedLeading: func() {
				p.logger.Debugf("leader election lock LOST, id %v", id)
				eventID := fmt.Sprintf("%v-%v", metaUID, time.Now().UnixNano())
				p.stopLeading(uuid, eventID)
			},
		},
	}
}
