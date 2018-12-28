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

package virtualmachine

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/common/cfgwarn"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"

	"github.com/pkg/errors"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

var logger = logp.NewLogger("vsphere")

func init() {
	mb.Registry.MustAddMetricSet("vsphere", "virtualmachine", New,
		mb.DefaultMetricSet(),
	)
}

type MetricSet struct {
	mb.BaseMetricSet
	HostURL         *url.URL
	Insecure        bool
	GetCustomFields bool
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The vsphere virtualmachine metricset is beta")

	config := struct {
		Username        string `config:"username"`
		Password        string `config:"password"`
		Insecure        bool   `config:"insecure"`
		GetCustomFields bool   `config:"get_custom_fields"`
	}{
		GetCustomFields: false,
	}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	u, err := url.Parse(base.HostData().URI)
	if err != nil {
		return nil, err
	}

	u.User = url.UserPassword(config.Username, config.Password)

	return &MetricSet{
		BaseMetricSet:   base,
		HostURL:         u,
		Insecure:        config.Insecure,
		GetCustomFields: config.GetCustomFields,
	}, nil
}

func (m *MetricSet) Fetch() ([]common.MapStr, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var events []common.MapStr

	client, err := govmomi.NewClient(ctx, m.HostURL, m.Insecure)
	if err != nil {
		return nil, err
	}

	defer client.Logout(ctx)

	c := client.Client

	// Get custom fields (attributes) names if get_custom_fields is true.
	customFieldsMap := make(map[int32]string)
	if m.GetCustomFields {
		var err error
		customFieldsMap, err = setCustomFieldsMap(ctx, c)
		if err != nil {
			return nil, err
		}
	}

	// Create view of VirtualMachine objects
	mgr := view.NewManager(c)

	v, err := mgr.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"VirtualMachine"}, true)
	if err != nil {
		return nil, err
	}

	defer v.Destroy(ctx)

	// Retrieve summary property for all machines
	var vmt []mo.VirtualMachine
	err = v.Retrieve(ctx, []string{"VirtualMachine"}, []string{"summary"}, &vmt)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, vm := range vmt {

		wg.Add(1)

		go func(vm mo.VirtualMachine, c *vim25.Client) {

			defer wg.Done()

			freeMemory := (int64(vm.Summary.Config.MemorySizeMB) * 1024 * 1024) - (int64(vm.Summary.QuickStats.GuestMemoryUsage) * 1024 * 1024)

			event := common.MapStr{}

			event["name"] = vm.Summary.Config.Name
			event.Put("cpu.used.mhz", vm.Summary.QuickStats.OverallCpuUsage)
			event.Put("memory.used.guest.bytes", int64(vm.Summary.QuickStats.GuestMemoryUsage)*1024*1024)
			event.Put("memory.used.host.bytes", int64(vm.Summary.QuickStats.HostMemoryUsage)*1024*1024)
			event.Put("memory.total.guest.bytes", int64(vm.Summary.Config.MemorySizeMB)*1024*1024)
			event.Put("memory.free.guest.bytes", freeMemory)

			if vm.Summary.Runtime.Host != nil {
				event["host"] = vm.Summary.Runtime.Host.Value
			} else {
				logger.Debug("'Host', 'Runtime' or 'Summary' data not found. This is either a parsing error " +
					"from vsphere library, an error trying to reach host/guest or incomplete information returned " +
					"from host/guest")
			}

			// Get custom fields (attributes) values if get_custom_fields is true.
			if m.GetCustomFields && vm.Summary.CustomValue != nil {
				customFields := getCustomFields(vm.Summary.CustomValue, customFieldsMap)

				if len(customFields) > 0 {
					event["custom_fields"] = customFields
				}
			} else {
				logger.Debug("custom fields not activated or custom values not found/parse in Summary data. This " +
					"is either a parsing error from vsphere library, an error trying to reach host/guest or incomplete " +
					"information returned from host/guest")
			}

			if vm.Summary.Vm != nil {
				networkNames, err := getNetworkNames(c, vm.Summary.Vm.Reference())
				if err != nil {
					logger.Debug(err.Error())
				} else {
					if len(networkNames) > 0 {
						event["network_names"] = networkNames
					}
				}
			}

			mutex.Lock()
			events = append(events, event)
			mutex.Unlock()
		}(vm, c)
	}

	wg.Wait()

	return events, nil
}

func getCustomFields(customFields []types.BaseCustomFieldValue, customFieldsMap map[int32]string) common.MapStr {
	outputFields := common.MapStr{}
	for _, v := range customFields {
		customFieldString := v.(*types.CustomFieldStringValue)
		key, ok := customFieldsMap[v.GetCustomFieldValue().Key]
		if ok {
			// If key has '.', is replaced with '_' to be compatible with ES2.x.
			fmtKey := strings.Replace(key, ".", "_", -1)
			outputFields.Put(fmtKey, customFieldString.Value)
		}
	}

	return outputFields
}

func getNetworkNames(c *vim25.Client, ref types.ManagedObjectReference) ([]string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var outputNetworkNames []string

	pc := property.DefaultCollector(c)

	var vm mo.VirtualMachine
	err := pc.RetrieveOne(ctx, ref, []string{"network"}, &vm)
	if err != nil {
		return nil, fmt.Errorf("error retrieving virtual machine information: %v", err)
	}

	if len(vm.Network) == 0 {
		return nil, errors.New("no networks found")
	}

	var networkRefs []types.ManagedObjectReference
	for _, obj := range vm.Network {
		if obj.Type == "Network" {
			networkRefs = append(networkRefs, obj)
		}
	}

	// If only "Distributed port group" was found, for example.
	if len(networkRefs) == 0 {
		return nil, errors.New("no networks found")
	}

	var nets []mo.Network
	err = pc.Retrieve(ctx, networkRefs, []string{"name"}, &nets)
	if err != nil {
		return nil, fmt.Errorf("error retrieving network from virtual machine: %v", err)
	}

	for _, net := range nets {
		name := strings.Replace(net.Name, ".", "_", -1)
		outputNetworkNames = append(outputNetworkNames, name)
	}

	return outputNetworkNames, nil
}

func setCustomFieldsMap(ctx context.Context, client *vim25.Client) (map[int32]string, error) {
	customFieldsMap := make(map[int32]string)

	customFieldsManager, err := object.GetCustomFieldsManager(client)

	if err != nil {
		return nil, errors.Wrap(err, "failed to get custom fields manager")
	} else {
		field, err := customFieldsManager.Field(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get custom fields")
		}

		for _, def := range field {
			customFieldsMap[def.Key] = def.Name
		}
	}

	return customFieldsMap, nil
}
