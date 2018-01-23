package kubernetes

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"

	"github.com/elastic/beats/libbeat/logp"
)

const defaultNode = "localhost"

// GetKubernetesClient returns a kubernetes client. If inCluster is true, it returns an
// in cluster configuration based on the secrets mounted in the Pod. If kubeConfig is passed,
// it parses the config file to get the config required to build a client.
func GetKubernetesClient(inCluster bool, kubeConfig string) (client *k8s.Client, err error) {
	if inCluster == true {
		client, err = k8s.NewInClusterClient()
		if err != nil {
			return nil, fmt.Errorf("Unable to get in cluster configuration: %v", err)
		}
	} else {
		data, err := ioutil.ReadFile(kubeConfig)
		if err != nil {
			return nil, fmt.Errorf("read kubeconfig: %v", err)
		}

		// Unmarshal YAML into a Kubernetes config object.
		var config k8s.Config
		if err = yaml.Unmarshal(data, &config); err != nil {
			return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
		}
		client, err = k8s.NewClient(&config)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// DiscoverKubernetesNode figures out the Kubernetes node to use.
// If host is provided in the config use it directly.
// If beat is deployed in k8s cluster, use hostname of pod which is pod name to query pod meta for node name.
// If beat is deployed outside k8s cluster, use machine-id to match against k8s nodes for node name.
func DiscoverKubernetesNode(host string, inCluster bool, client *k8s.Client) (node string) {
	if host != "" {
		logp.Info("kubernetes: Using node %s provided in the config", host)
		return host
	}

	if inCluster {
		ns, err := inClusterNamespace()
		if err != nil {
			logp.Err("kubernetes: Couldn't get namespace when beat is in cluster with error: ", err.Error())
			return defaultNode
		}
		podName, err := os.Hostname()
		if err != nil {
			logp.Err("kubernetes: Couldn't get hostname as beat pod name in cluster with error: ", err.Error())
			return defaultNode
		}
		logp.Info("kubernetes: Using pod name %s and namespace %s to discover kubernetes node", podName, ns)
		pod, err := client.CoreV1().GetPod(context.TODO(), podName, ns)
		if err != nil {
			logp.Err("kubernetes: Querying for pod failed with error: ", err.Error())
			return defaultNode
		}
		logp.Info("kubernetes: Using node %s discovered by in cluster pod node query", pod.Spec.GetNodeName())
		return pod.Spec.GetNodeName()
	}

	mid := machineID()
	if mid == "" {
		logp.Err("kubernetes: Couldn't collect info from any of the files in /etc/machine-id /var/lib/dbus/machine-id")
		return defaultNode
	}

	nodes, err := client.CoreV1().ListNodes(context.TODO())
	if err != nil {
		logp.Err("kubernetes: Querying for nodes failed with error: ", err.Error())
		return defaultNode
	}
	for _, n := range nodes.Items {
		if n.GetStatus().GetNodeInfo().GetMachineID() == mid {
			logp.Info("kubernetes: Using node %s discovered by machine-id matching", n.GetMetadata().GetName())
			return n.GetMetadata().GetName()
		}
	}

	logp.Warn("kubernetes: Couldn't discover node, using localhost as default")
	return defaultNode
}

// machineID borrowed from cadvisor.
func machineID() string {
	for _, file := range []string{
		"/etc/machine-id",
		"/var/lib/dbus/machine-id",
	} {
		id, err := ioutil.ReadFile(file)
		if err == nil {
			return strings.TrimSpace(string(id))
		}
	}
	return ""
}

// inClusterNamespace gets namespace from serviceaccount when beat is in cluster.
// code borrowed from client-go with some changes.
func inClusterNamespace() (string, error) {
	// get namespace associated with the service account token, if available
	data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
