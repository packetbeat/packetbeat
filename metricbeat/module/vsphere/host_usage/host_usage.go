package host_usage

import (
	"context"
	"net/url"

	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/metricbeat/mb"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

func init() {
	if err := mb.Registry.AddMetricSet("vsphere", "host_usage", New); err != nil {
		panic(err)
	}
}

type MetricSet struct {
	mb.BaseMetricSet
	hostUrl  *url.URL
	insecure bool
}

func New(base mb.BaseMetricSet) (mb.MetricSet, error) {

	config := struct {
		Username string `config:"username"`
		Password string `config:"password"`
		Insecure bool   `config:"insecure"`
	}{}

	if err := base.Module().UnpackConfig(&config); err != nil {
		return nil, err
	}

	u, err := url.Parse(base.HostData().URI)
	if err != nil {
		return nil, err
	}

	u.User = url.UserPassword(config.Username, config.Password)

	return &MetricSet{
		BaseMetricSet: base,
		hostUrl:       u,
		insecure:      config.Insecure,
	}, nil
}

func (m *MetricSet) Fetch() ([]common.MapStr, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := govmomi.NewClient(ctx, m.hostUrl, m.insecure)
	if err != nil {
		return nil, err
	}

	f := find.NewFinder(c.Client, true)
	if f == nil {
		return nil, err
	}

	// Get all datacenters
	dcs, err := f.DatacenterList(ctx, "*")
	if err != nil {
		return nil, err
	}

	events := []common.MapStr{}

	for _, dc := range dcs {

		f.SetDatacenter(dc)

		// Get all hosts
		hss, err := f.HostSystemList(ctx, "*")
		if err != nil {
			return nil, err
		}

		pc := property.DefaultCollector(c.Client)

		// Convert hosts into list of references
		var refs []types.ManagedObjectReference
		for _, hs := range hss {
			refs = append(refs, hs.Reference())
		}

		// Get summary property (HostListSummary)
		var hst []mo.HostSystem
		err = pc.Retrieve(ctx, refs, []string{"summary"}, &hst)
		if err != nil {
			return nil, err
		}

		for _, hs := range hst {
			totalCpu := int64(hs.Summary.Hardware.CpuMhz) * int64(hs.Summary.Hardware.NumCpuCores)
			freeCpu := int64(totalCpu) - int64(hs.Summary.QuickStats.OverallCpuUsage)
			freeMemory := int64(hs.Summary.Hardware.MemorySize) - (int64(hs.Summary.QuickStats.OverallMemoryUsage) * 1024)

			event := common.MapStr{
				"datacenter": dc.Name(),
				"name":       hs.Summary.Config.Name,
				"cpu": common.MapStr{
					"used": common.MapStr{
						"mhz": hs.Summary.QuickStats.OverallCpuUsage,
					},
					"total": common.MapStr{
						"mhz": totalCpu,
					},
					"free": common.MapStr{
						"mhz": freeCpu,
					},
				},
				"memory": common.MapStr{
					"used": common.MapStr{
						"bytes": hs.Summary.QuickStats.OverallMemoryUsage * 1024,
					},
					"total": common.MapStr{
						"bytes": hs.Summary.Hardware.MemorySize,
					},
					"free": common.MapStr{
						"bytes": freeMemory,
					},
				},
			}

			events = append(events, event)
		}
	}

	return events, nil
}
