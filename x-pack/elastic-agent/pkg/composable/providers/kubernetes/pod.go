// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package kubernetes

import (
	"fmt"
	"sync"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/safemapstr"

	k8s "k8s.io/client-go/kubernetes"

	"github.com/elastic/beats/v7/libbeat/common/kubernetes"
	"github.com/elastic/beats/v7/libbeat/common/kubernetes/metadata"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/agent/errors"
	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/composable"
)

type pod struct {
	logger           *logp.Logger
	cleanupTimeout   time.Duration
	comm             composable.DynamicProviderComm
	scope            string
	config           *Config
	metagen          metadata.MetaGen
	namespaceWatcher kubernetes.Watcher

	// Mutex used by configuration updates not triggered by the main watcher,
	// to avoid race conditions between cross updates and deletions.
	// Other updaters must use a write lock.
	crossUpdate sync.RWMutex
}

type providerData struct {
	uid        string
	mapping    map[string]interface{}
	processors []map[string]interface{}
}

type containerInPod struct {
	id      string
	runtime string
	spec    kubernetes.Container
	status  kubernetes.PodContainerStatus
}

// podUpdaterHandlerFunc is a function that handles pod updater notifications.
type podUpdaterHandlerFunc func(interface{})

// podUpdaterStore is the interface that an object needs to implement to be
// used as a pod updater store.
type podUpdaterStore interface {
	List() []interface{}
}

// namespacePodUpdater notifies updates on pods when their namespaces are updated.
type namespacePodUpdater struct {
	handler podUpdaterHandlerFunc
	store   podUpdaterStore
	locker  sync.Locker
}

// NewPodWatcher creates a watcher that can discover and process pod objects
func NewPodWatcher(
	comm composable.DynamicProviderComm,
	cfg *Config,
	logger *logp.Logger,
	client k8s.Interface,
	scope string,
	rawConfig *common.Config) (kubernetes.Watcher, error) {
	watcher, err := kubernetes.NewWatcher(client, &kubernetes.Pod{}, kubernetes.WatchOptions{
		SyncTimeout:  cfg.SyncPeriod,
		Node:         cfg.Node,
		Namespace:    cfg.Namespace,
		HonorReSyncs: true,
	}, nil)
	if err != nil {
		return nil, errors.New(err, "couldn't create kubernetes watcher")
	}

	options := kubernetes.WatchOptions{
		SyncTimeout: cfg.SyncPeriod,
		Node:        cfg.Node,
	}
	metaConf := cfg.AddResourceMetadata
	if metaConf == nil {
		metaConf = metadata.GetDefaultResourceMetadataConfig()
	}
	nodeWatcher, err := kubernetes.NewWatcher(client, &kubernetes.Node{}, options, nil)
	if err != nil {
		logger.Errorf("couldn't create watcher for %T due to error %+v", &kubernetes.Node{}, err)
	}
	namespaceWatcher, err := kubernetes.NewWatcher(client, &kubernetes.Namespace{}, kubernetes.WatchOptions{
		SyncTimeout: cfg.SyncPeriod,
	}, nil)
	if err != nil {
		logger.Errorf("couldn't create watcher for %T due to error %+v", &kubernetes.Namespace{}, err)
	}

	metaGen := metadata.GetPodMetaGen(rawConfig, watcher, nodeWatcher, namespaceWatcher, metaConf)

	p := pod{
		logger:           logger,
		cleanupTimeout:   cfg.CleanupTimeout,
		comm:             comm,
		scope:            scope,
		config:           cfg,
		metagen:          metaGen,
		namespaceWatcher: namespaceWatcher,
	}

	watcher.AddEventHandler(&p)

	if namespaceWatcher != nil && metaConf.Namespace.Enabled() {
		updater := newNamespacePodUpdater(p.unlockedUpdate, watcher.Store(), &p.crossUpdate)
		namespaceWatcher.AddEventHandler(updater)
	}

	return watcher, nil
}

func (p *pod) emitRunning(pod *kubernetes.Pod) {

	namespaceAnnotations := podNamespaceAnnotations(pod, p.namespaceWatcher)

	data := generatePodData(pod, p.config, p.metagen, namespaceAnnotations)
	data.mapping["scope"] = p.scope
	// Emit the pod
	// We emit Pod + containers to ensure that configs matching Pod only
	// get Pod metadata (not specific to any container)
	p.comm.AddOrUpdate(data.uid, PodPriority, data.mapping, data.processors)

	// Emit all containers in the pod
	p.emitContainers(pod, pod.Spec.Containers, pod.Status.ContainerStatuses)

	// TODO: deal with init containers stopping after initialization
	p.emitContainers(pod, pod.Spec.InitContainers, pod.Status.InitContainerStatuses)

	// Get ephemeral containers and their status
	ephContainers, ephContainersStatuses := getEphemeralContainers(pod)
	p.emitContainers(pod, ephContainers, ephContainersStatuses)

}

func (p *pod) emitContainers(
	pod *kubernetes.Pod,
	containers []kubernetes.Container,
	containerstatuses []kubernetes.PodContainerStatus) {
	generateContainerData(p.comm, pod, containers, containerstatuses, p.config, p.metagen)
}

func (p *pod) emitStopped(pod *kubernetes.Pod) {
	p.comm.Remove(string(pod.GetUID()))

	for _, c := range pod.Spec.Containers {
		// ID is the combination of pod UID + container name
		eventID := fmt.Sprintf("%s.%s", pod.GetObjectMeta().GetUID(), c.Name)
		p.comm.Remove(eventID)
	}

	for _, c := range pod.Spec.InitContainers {
		// ID is the combination of pod UID + container name
		eventID := fmt.Sprintf("%s.%s", pod.GetObjectMeta().GetUID(), c.Name)
		p.comm.Remove(eventID)
	}
}

// OnAdd ensures processing of pod objects that are newly added
func (p *pod) OnAdd(obj interface{}) {
	p.crossUpdate.RLock()
	defer p.crossUpdate.RUnlock()

	p.logger.Debugf("pod add: %+v", obj)
	p.emitRunning(obj.(*kubernetes.Pod))
}

// OnUpdate emits events for a given pod depending on the state of the pod,
// if it is terminating, a stop event is scheduled, if not, a stop and a start
// events are sent sequentially to recreate the resources assotiated to the pod.
func (p *pod) OnUpdate(obj interface{}) {
	p.crossUpdate.RLock()
	defer p.crossUpdate.RUnlock()

	p.unlockedUpdate(obj)
}

func (p *pod) unlockedUpdate(obj interface{}) {
	p.logger.Debugf("Watcher Pod update: %+v", obj)
	pod := obj.(*kubernetes.Pod)

	p.logger.Debugf("pod update for pod: %+v, status: %+v", pod.Name, pod.Status.Phase)

	if podTerminated(pod, getContainersInPod(pod)) {
		time.AfterFunc(p.cleanupTimeout, func() { p.emitStopped(pod) })
		return
	}

	p.logger.Debugf("pod update: %+v", obj)
	p.emitRunning(pod)
}

// OnDelete stops pod objects that are deleted
func (p *pod) OnDelete(obj interface{}) {
	p.crossUpdate.RLock()
	defer p.crossUpdate.RUnlock()

	p.logger.Debugf("pod delete: %+v", obj)
	pod := obj.(*kubernetes.Pod)
	time.AfterFunc(p.cleanupTimeout, func() { p.emitStopped(pod) })
}

func generatePodData(
	pod *kubernetes.Pod,
	cfg *Config,
	kubeMetaGen metadata.MetaGen,
	namespaceAnnotations common.MapStr) providerData {

	meta := kubeMetaGen.Generate(pod)
	kubemetaMap, err := meta.GetValue("kubernetes")
	if err != nil {
		return providerData{}
	}

	// k8sMapping includes only the metadata that fall under kubernetes.*
	// and these are available as dynamic vars through the provider
	k8sMapping := map[string]interface{}(kubemetaMap.(common.MapStr).Clone())

	if len(namespaceAnnotations) != 0 {
		k8sMapping["namespace_annotations"] = namespaceAnnotations
	}

	// Pass annotations to all events so that it can be used in templating and by annotation builders.
	annotations := common.MapStr{}
	for k, v := range pod.GetObjectMeta().GetAnnotations() {
		safemapstr.Put(annotations, k, v)
	}
	k8sMapping["annotations"] = annotations

	processors := []map[string]interface{}{}
	// meta map includes metadata that go under kubernetes.*
	// but also other ECS fields like orchestrator.*
	for field, metaMap := range meta {
		processor := map[string]interface{}{
			"add_fields": map[string]interface{}{
				"fields": metaMap,
				"target": field,
			},
		}
		processors = append(processors, processor)
	}

	return providerData{
		uid:        string(pod.GetUID()),
		mapping:    k8sMapping,
		processors: processors,
	}
}

func generateContainerData(
	comm composable.DynamicProviderComm,
	pod *kubernetes.Pod,
	containers []kubernetes.Container,
	containerstatuses []kubernetes.PodContainerStatus,
	cfg *Config,
	kubeMetaGen metadata.MetaGen) {

	containerIDs := map[string]string{}
	runtimes := map[string]string{}
	for _, c := range containerstatuses {
		cid, runtime := kubernetes.ContainerIDWithRuntime(c)
		containerIDs[c.Name] = cid
		runtimes[c.Name] = runtime
	}

	// Pass annotations to all events so that it can be used in templating and by annotation builders.
	annotations := common.MapStr{}
	for k, v := range pod.GetObjectMeta().GetAnnotations() {
		safemapstr.Put(annotations, k, v)
	}

	for _, c := range containers {
		// If it doesn't have an ID, container doesn't exist in
		// the runtime, emit only an event if we are stopping, so
		// we are sure of cleaning up configurations.
		cid := containerIDs[c.Name]
		if cid == "" {
			continue
		}

		// ID is the combination of pod UID + container name
		eventID := fmt.Sprintf("%s.%s", pod.GetObjectMeta().GetUID(), c.Name)

		meta := kubeMetaGen.Generate(pod, metadata.WithFields("container.name", c.Name))
		kubemetaMap, err := meta.GetValue("kubernetes")
		if err != nil {
			continue
		}

		// k8sMapping includes only the metadata that fall under kubernetes.*
		// and these are available as dynamic vars through the provider
		k8sMapping := map[string]interface{}(kubemetaMap.(common.MapStr).Clone())

		// add annotations to be discoverable by templates
		k8sMapping["annotations"] = annotations

		//container ECS fields
		cmeta := common.MapStr{
			"id":      cid,
			"runtime": runtimes[c.Name],
			"image": common.MapStr{
				"name": c.Image,
			},
		}

		processors := []map[string]interface{}{
			{
				"add_fields": map[string]interface{}{
					"fields": cmeta,
					"target": "container",
				},
			},
		}
		// meta map includes metadata that go under kubernetes.*
		// but also other ECS fields like orchestrator.*
		for field, metaMap := range meta {
			processor := map[string]interface{}{
				"add_fields": map[string]interface{}{
					"fields": metaMap,
					"target": field,
				},
			}
			processors = append(processors, processor)
		}

		// add container metadata under kubernetes.container.* to
		// make them available to dynamic var resolution
		containerMeta := common.MapStr{
			"id":      cid,
			"name":    c.Name,
			"image":   c.Image,
			"runtime": runtimes[c.Name],
		}
		if len(c.Ports) > 0 {
			for _, port := range c.Ports {
				containerMeta.Put("port", fmt.Sprintf("%v", port.ContainerPort))
				containerMeta.Put("port_name", port.Name)
				k8sMapping["container"] = containerMeta
				comm.AddOrUpdate(eventID, ContainerPriority, k8sMapping, processors)
			}
		} else {
			k8sMapping["container"] = containerMeta
			comm.AddOrUpdate(eventID, ContainerPriority, k8sMapping, processors)
		}
	}
}

func getEphemeralContainers(pod *kubernetes.Pod) ([]kubernetes.Container, []kubernetes.PodContainerStatus) {
	var ephContainers []kubernetes.Container
	var ephContainersStatuses []kubernetes.PodContainerStatus
	for _, c := range pod.Spec.EphemeralContainers {
		c := kubernetes.Container(c.EphemeralContainerCommon)
		ephContainers = append(ephContainers, c)
	}
	for _, s := range pod.Status.EphemeralContainerStatuses {
		ephContainersStatuses = append(ephContainersStatuses, s)
	}
	return ephContainers, ephContainersStatuses
}

// podNamespaceAnnotations returns the annotations of the namespace of the pod
func podNamespaceAnnotations(pod *kubernetes.Pod, watcher kubernetes.Watcher) common.MapStr {
	if watcher == nil {
		return nil
	}

	rawNs, ok, err := watcher.Store().GetByKey(pod.Namespace)
	if !ok || err != nil {
		return nil
	}

	namespace, ok := rawNs.(*kubernetes.Namespace)
	if !ok {
		return nil
	}

	annotations := common.MapStr{}
	for k, v := range namespace.GetAnnotations() {
		safemapstr.Put(annotations, k, v)
	}
	return annotations
}

// newNamespacePodUpdater creates a namespacePodUpdater
func newNamespacePodUpdater(handler podUpdaterHandlerFunc, store podUpdaterStore, locker sync.Locker) *namespacePodUpdater {
	return &namespacePodUpdater{
		handler: handler,
		store:   store,
		locker:  locker,
	}
}

// OnUpdate handles update events on namespaces.
func (n *namespacePodUpdater) OnUpdate(obj interface{}) {
	ns, ok := obj.(*kubernetes.Namespace)
	if !ok {
		return
	}

	// n.store.List() returns a snapshot at this point. If a delete is received
	// from the main watcher, this loop may generate an update event after the
	// delete is processed, leaving configurations that would never be deleted.
	// Also this loop can miss updates, what could leave outdated configurations.
	// Avoid these issues by locking the processing of events from the main watcher.
	if n.locker != nil {
		n.locker.Lock()
		defer n.locker.Unlock()
	}
	for _, pod := range n.store.List() {
		pod, ok := pod.(*kubernetes.Pod)
		if ok && pod.Namespace == ns.Name {
			n.handler(pod)
		}
	}
}

// OnAdd handles add events on namespaces. Nothing to do, if pods are added to this
// namespace they will generate their own add events.
func (*namespacePodUpdater) OnAdd(interface{}) {}

// OnDelete handles delete events on namespaces. Nothing to do, if pods are deleted from this
// namespace they will generate their own delete events.
func (*namespacePodUpdater) OnDelete(interface{}) {}

// podTerminating returns true if a pod is marked for deletion or is in a phase beyond running.
func podTerminating(pod *kubernetes.Pod) bool {
	if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
		return true
	}

	switch pod.Status.Phase {
	case kubernetes.PodRunning, kubernetes.PodPending:
	default:
		return true
	}

	return false
}

// podTerminated returns true if a pod is terminated, this method considers a
// pod as terminated if none of its containers are running (or going to be running).
func podTerminated(pod *kubernetes.Pod, containers []*containerInPod) bool {
	// Pod is not marked for termination, so it is not terminated.
	if !podTerminating(pod) {
		return false
	}

	// If any container is running, the pod is not terminated yet.
	for _, container := range containers {
		if container.status.State.Running != nil {
			return false
		}
	}

	return true
}

// getContainersInPod returns all the containers defined in a pod and their statuses.
// It includes init and ephemeral containers.
func getContainersInPod(pod *kubernetes.Pod) []*containerInPod {
	var containers []*containerInPod
	for _, c := range pod.Spec.Containers {
		containers = append(containers, &containerInPod{spec: c})
	}
	for _, c := range pod.Spec.InitContainers {
		containers = append(containers, &containerInPod{spec: c})
	}
	for _, c := range pod.Spec.EphemeralContainers {
		c := kubernetes.Container(c.EphemeralContainerCommon)
		containers = append(containers, &containerInPod{spec: c})
	}

	statuses := make(map[string]*kubernetes.PodContainerStatus)
	mapStatuses := func(s []kubernetes.PodContainerStatus) {
		for i := range s {
			statuses[s[i].Name] = &s[i]
		}
	}
	mapStatuses(pod.Status.ContainerStatuses)
	mapStatuses(pod.Status.InitContainerStatuses)
	mapStatuses(pod.Status.EphemeralContainerStatuses)
	for _, c := range containers {
		if s, ok := statuses[c.spec.Name]; ok {
			c.id, c.runtime = kubernetes.ContainerIDWithRuntime(*s)
			c.status = *s
		}
	}

	return containers
}
