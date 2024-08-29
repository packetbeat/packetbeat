// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/module_include_list/module_include_list.go - DO NOT EDIT.

package include

import (
	// Import packages that perform 'func init()'.
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/activemq"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/airflow"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/aws"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/aws/awshealth"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/aws/billing"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/aws/cloudwatch"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/awsfargate"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/awsfargate/task_stats"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/app_insights"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/billing"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/monitor"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/azure/storage"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/cloudfoundry"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/cloudfoundry/container"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/cloudfoundry/counter"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/cloudfoundry/value"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/cockroachdb"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/containerd"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/containerd/blkio"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/containerd/cpu"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/containerd/memory"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/coredns"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/coredns/stats"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/enterprisesearch"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/enterprisesearch/health"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/enterprisesearch/stats"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/gcp"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/gcp/billing"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/gcp/carbon"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/gcp/metrics"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/ibmmq"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/iis"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/iis/application_pool"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio/citadel"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio/galley"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio/mesh"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio/mixer"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/istio/pilot"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/mssql"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/mssql/performance"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/mssql/transaction_log"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/oracle"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/oracle/performance"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/oracle/sysmetric"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/oracle/tablespace"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/panos"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/panos/licenses"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/prometheus"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/prometheus/collector"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/prometheus/remote_write"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/redisenterprise"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/sql"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/sql/query"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/stan"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/stan/channels"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/stan/stats"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/stan/subscriptions"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/statsd"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/statsd/server"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/syncgateway"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/syncgateway/db"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/syncgateway/memory"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/syncgateway/replication"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/syncgateway/resources"
	_ "github.com/elastic/beats/v7/x-pack/metricbeat/module/tomcat"
)
