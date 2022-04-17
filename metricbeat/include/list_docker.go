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

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

//go:build linux || darwin || windows
// +build linux darwin windows

package include

import (
	// Import packages that need to register themselves.
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/container"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/cpu"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/diskio"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/event"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/healthcheck"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/image"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/info"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/memory"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/network"
	_ "github.com/menderesk/beats/v7/metricbeat/module/docker/network_summary"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/apiserver"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/container"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/controllermanager"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/event"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/node"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/pod"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/proxy"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/scheduler"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_container"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_cronjob"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_daemonset"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_deployment"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_job"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_node"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_persistentvolume"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_persistentvolumeclaim"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_pod"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_replicaset"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_resourcequota"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_service"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_statefulset"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/state_storageclass"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/system"
	_ "github.com/menderesk/beats/v7/metricbeat/module/kubernetes/volume"
)
