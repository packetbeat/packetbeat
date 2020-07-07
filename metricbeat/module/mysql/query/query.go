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

/*
Package status fetches MySQL server status metrics.

For more information on the query it uses, see:
http://dev.mysql.com/doc/refman/5.7/en/show-status.html
*/
package query

import (
	"context"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/helper/sql"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

func init() {
	mb.Registry.MustAddMetricSet("mysql", "query", New,
		mb.DefaultMetricSet(),
	)
}

type query struct {
	QueryNamespace string `config:"query_namespace"`
	Query          string `config:"query"`
}

// MetricSet for fetching MySQL server status.
type MetricSet struct {
	mb.BaseMetricSet
	db     *sql.DbClient
	Config struct {
		Queries        []query `config:"sql_queries" validate:"nonzero,required"`
		Namespace      string  `config:"namespace" validate:"nonzero,required"`
		ResponseFormat string  `config:"sql_response_format"`
	}
}

// New creates and returns a new MetricSet instance.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The mysql 'query' metricset is beta.")

	b := &MetricSet{BaseMetricSet: base}

	if err := base.Module().UnpackConfig(&b.Config); err != nil {
		return nil, err
	}

	return b, nil
}

// Fetch fetches status messages from a mysql host.
func (m *MetricSet) Fetch(ctx context.Context, reporter mb.ReporterV2) error {
	if m.db == nil {
		var err error
		m.db, err = sql.DB("mysql", m.HostData().URI, m.Logger())
		if err != nil {
			return errors.Wrap(err, "mysql-status fetch failed")
		}
	}

	for _, q := range m.Config.Queries {
		err := m.fetchQuery(ctx, q, reporter)
		if err != nil {
			m.Logger().Error("error doing query", err)
		}
	}

	return nil
}

func (m *MetricSet) fetchQuery(ctx context.Context, q query, reporter mb.ReporterV2) error {
	if m.Config.ResponseFormat == "table" {
		mss, err := m.db.FetchTableMode(ctx, q.Query)
		if err != nil {
			return err
		}

		for _, ms := range mss {
			event := mb.Event{ModuleFields: common.MapStr{m.Config.Namespace: common.MapStr{}}}
			if q.QueryNamespace != "" {
				event.ModuleFields[m.Config.Namespace] = common.MapStr{q.QueryNamespace: sql.ToDotKeys(ms)}
			} else {
				event.ModuleFields[m.Config.Namespace] = sql.ToDotKeys(ms)
			}
			reporter.Event(event)
		}
	} else {
		ms, err := m.db.FetchVariableMode(ctx, q.Query)
		if err != nil {
			return err
		}

		event := mb.Event{ModuleFields: common.MapStr{m.Config.Namespace: common.MapStr{}}}
		if q.QueryNamespace != "" {
			event.ModuleFields[m.Config.Namespace] = common.MapStr{q.QueryNamespace: sql.ToDotKeys(ms)}
		} else {
			event.ModuleFields[m.Config.Namespace] = sql.ToDotKeys(ms)
		}
		reporter.Event(event)
	}

	return nil
}

// Close closes the database connection and prevents future queries.
func (m *MetricSet) Close() error {
	if m.db == nil {
		return nil
	}
	return errors.Wrap(m.db.Close(), "failed to close mysql database client")
}
