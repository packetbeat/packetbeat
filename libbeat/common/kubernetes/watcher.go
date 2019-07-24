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

package kubernetes

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/elastic/beats/libbeat/logp"
)

const (
	add    = "add"
	update = "update"
	delete = "delete"
)

// Watcher watches Kubernetes resources events
type Watcher interface {
	// Start watching Kubernetes API for new events after resources were listed
	Start() error

	// Stop watching Kubernetes API for new events
	Stop()

	// AddEventHandler add event handlers for corresponding event type watched
	AddEventHandler(ResourceEventHandler)
}

// WatchOptions controls watch behaviors
type WatchOptions struct {
	// SyncTimeout is a timeout for listing historical resources
	SyncTimeout time.Duration
	// Node is used for filtering watched resource to given node, use "" for all nodes
	Node string
	// Namespace is used for filtering watched resource to given namespace, use "" for all namespaces
	Namespace string
}

type item struct {
	object interface{}
	state  string
}

type watcher struct {
	client   kubernetes.Interface
	informer cache.SharedInformer
	store    cache.Store
	queue    workqueue.RateLimitingInterface
	ctx      context.Context
	stop     context.CancelFunc
	handler  ResourceEventHandler
	logger   *logp.Logger
}

func tweakOptions(options *metav1.ListOptions, opt WatchOptions) {
	if opt.Node != "" {
		options.FieldSelector = "spec.nodeName=" + opt.Node
	}
}

// NewWatcher initializes the watcher client to provide a events handler for
// resource from the cluster (filtered to the given node)
func NewWatcher(client kubernetes.Interface, resource Resource, opts WatchOptions) (Watcher, error) {
	ctx, cancel := context.WithCancel(context.Background())

	var informer cache.SharedInformer
	var store cache.Store
	var queue workqueue.RateLimitingInterface
	var objType string

	var listwatch *cache.ListWatch
	switch resource.(type) {
	case *Pod:
		p := client.CoreV1().Pods(opts.Namespace)
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				tweakOptions(&options, opts)
				return p.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				tweakOptions(&options, opts)
				return p.Watch(options)
			},
		}

		objType = "pod"
	case *Event:
		e := client.CoreV1().Events(opts.Namespace)
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return e.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return e.Watch(options)
			},
		}

		objType = "event"
	case *Node:
		n := client.CoreV1().Nodes()
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return n.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return n.Watch(options)
			},
		}

		objType = "node"
	case *Deployment:
		d := client.AppsV1().Deployments(opts.Namespace)
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return d.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return d.Watch(options)
			},
		}

		objType = "deployment"
	case *ReplicaSet:
		rs := client.AppsV1().ReplicaSets(opts.Namespace)
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return rs.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return rs.Watch(options)
			},
		}

		objType = "replicaset"
	case *StatefulSet:
		ss := client.AppsV1().ReplicaSets(opts.Namespace)
		listwatch = &cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return ss.List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return ss.Watch(options)
			},
		}

		objType = "statefulset"
	default:
		return nil, fmt.Errorf("unsupported resource type for watching %T", resource)
	}

	informer = cache.NewSharedInformer(listwatch, resource, opts.SyncTimeout)
	store = informer.GetStore()
	queue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), objType)

	w := &watcher{
		client:   client,
		informer: informer,
		store:    store,
		queue:    queue,
		ctx:      ctx,
		stop:     cancel,
		logger:   logp.NewLogger("kubernetes"),
	}

	w.informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(o interface{}) {
			w.enqueue(o, add)
		},
		DeleteFunc: func(o interface{}) {
			w.enqueue(o, delete)
		},
		// TODO: When we move to one watcher per object model, we should check resource verion of
		// TODO: the old vs new and avoid reconciliation during relisting.
		UpdateFunc: func(_, o interface{}) {
			w.enqueue(o, update)
		},
	})

	return w, nil
}

func (w *watcher) enqueue(obj interface{}, state string) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		return
	}

	w.queue.Add(&item{key, state})
}

func (w *watcher) AddEventHandler(h ResourceEventHandler) {
	w.handler = h
}

// Start watching pods
func (w *watcher) Start() error {
	defer utilruntime.HandleCrash()
	defer w.queue.ShutDown()

	go w.informer.Run(w.ctx.Done())

	if !cache.WaitForCacheSync(w.ctx.Done(), w.informer.HasSynced) {
		return fmt.Errorf("kubernetes informer unable to sync cache")
	}

	// TODO: Do we run parallel workers for this? It is useful when we run metricbeat as one instance per cluster
	go func() {
		for w.process(w.ctx) {
		}
	}()

	return nil
}

func (w *watcher) process(ctx context.Context) bool {
	keyObj, quit := w.queue.Get()
	if quit {
		return false
	}

	err := func(obj interface{}) error {
		defer w.queue.Done(obj)

		var entry *item
		var ok bool
		if entry, ok = obj.(*item); !ok {
			w.queue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected *item in workqueue but got %#v", obj))
			return nil
		}

		key := entry.object.(string)

		o, exists, err := w.store.GetByKey(key)
		if err != nil {
			return nil
		}
		if !exists {
			return nil
		}

		switch entry.state {
		case add:
			w.handler.OnAdd(o)
		case update:
			w.handler.OnUpdate(o)
		case delete:
			w.handler.OnDelete(o)
		}

		return nil
	}(keyObj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (w *watcher) Stop() {
	w.stop()
}
