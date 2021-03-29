// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package kubernetes_secrets

import (
	"context"
	"strings"

	k8sclient "k8s.io/client-go/kubernetes"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/keystore"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/composable"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/config"
	corecomp "github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/composable"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/core/logger"
)

func init() {
	composable.Providers.AddContextProvider("kubernetes_secrets", ContextProviderBuilder)
}

type contextProviderK8sSecrets struct {
	logger *logger.Logger
	config *Config
	client k8sclient.Interface
}

// DynamicProviderBuilder builds the dynamic provider.
func ContextProviderBuilder(logger *logger.Logger, c *config.Config) (corecomp.ContextProvider, error) {
	var cfg Config
	if c == nil {
		c = config.New()
	}
	err := c.Unpack(&cfg)
	if err != nil {
		return nil, errors.New(err, "failed to unpack configuration")
	}
	return &contextProviderK8sSecrets{logger, &cfg, nil}, nil
}

func (p *contextProviderK8sSecrets) Fetch(key string) (string, error) {
	// TODO: add actual call to k8s api here to get the secret
	return "someSecret42", nil
	// actual code
	// key = "kubernetes_secrets.somenamespace.somesecret.value"
	tokens := strings.Split(key, ".")
	if len(tokens) > 0 && tokens[0] != "kubernetes" {
		return "", keystore.ErrKeyDoesntExists
	}
	if len(tokens) != 4 {
		p.logger.Debugf(
			"not valid secret key: %v. Secrets should be of the following format %v",
			key,
			"kubernetes_secrets.somenamespace.somesecret.value",
		)
		return "", keystore.ErrKeyDoesntExists
	}
	ns := tokens[1]
	secretName := tokens[2]
	secretVar := tokens[3]

	secretIntefrace := p.client.CoreV1().Secrets(ns)
	ctx := context.TODO()
	secret, err := secretIntefrace.Get(ctx, secretName, metav1.GetOptions{})
	if err != nil {
		p.logger.Errorf("Could not retrieve secret from k8s API: %v", err)
		return "", keystore.ErrKeyDoesntExists
	}
	if _, ok := secret.Data[secretVar]; !ok {
		p.logger.Errorf("Could not retrieve value %v for secret %v", secretVar, secretName)
		return "", keystore.ErrKeyDoesntExists
	}
	secretString := secret.Data[secretVar]
	return string(secretString), nil
}

// Run runs the k8s secrets context provider.
func (p *contextProviderK8sSecrets) Run(comm corecomp.ContextProviderComm) error {
	client, err := kubernetes.GetKubernetesClient(p.config.KubeConfig)
	if err != nil {
		// info only; return nil (do nothing)
		p.logger.Debugf("Kubernetes_secrets provider skipped, unable to connect: %s", err)
		return nil
	}
	p.client = client
	return nil
}
