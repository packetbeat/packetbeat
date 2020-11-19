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

// +build !integration

package index

import (
	"github.com/elastic/beats/v7/metricbeat/helper"
	"github.com/elastic/beats/v7/metricbeat/mb"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
	"github.com/elastic/beats/v7/metricbeat/module/elasticsearch"
)

var info = elasticsearch.Info{
	ClusterID:   "1234",
	ClusterName: "helloworld",
}

func TestMapper(t *testing.T) {
	t.Skip("Skipping until fixing test with stats.623.json, stats.700-alpha1 and stats.800.bench.json")

	mux := createEsMuxer("7.6.0", "platinum", false)

	server := httptest.NewServer(mux)
	defer server.Close()

	httpClient, err := helper.NewHTTPFromConfig(helper.Config{
		ConnectTimeout: 30 * time.Second,
		Timeout:        30 * time.Second,
	}, mb.HostData{
		URI:          "http://localhost:9200",
		SanitizedURI: "http://localhost:9200",
		Host:         "http://localhost:9200",
	})
	if err != nil {
		t.Fatal(err)
	}

	elasticsearch.TestMapperWithHttpHelper(t, "../index/_meta/test/stats.*.json", httpClient, eventsMapping)
}

func TestEmpty(t *testing.T) {
	httpClient, err := helper.NewHTTPFromConfig(helper.Config{}, mb.HostData{})
	if err != nil {
		t.Fatal(err)
	}

	input, err := ioutil.ReadFile("./_meta/test/empty.512.json")
	require.NoError(t, err)

	reporter := &mbtest.CapturingReporterV2{}
	eventsMapping(reporter, httpClient, info, input)
	require.Equal(t, 0, len(reporter.GetEvents()))
}

func createEsMuxer(esVersion, license string, ccrEnabled bool) *http.ServeMux {
	nodesLocalHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"nodes": { "foobar": {}}}`))
	}
	clusterStateMasterHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"master_node": "foobar"}`))
	}
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "_stats") {
			input, _ := ioutil.ReadFile("./_meta/test/stats.800.snapshot.20201118.json")
			w.Write(input)
			return
		} else if r.URL.Path != "/" {
			input, _ := ioutil.ReadFile("./_meta/test/settings.json")
			w.Write(input)
			return
		}

		w.Write([]byte(`{"name":"a14cf47ef7f2","cluster_name":"docker-cluster","cluster_uuid":"8l_zoGznQRmtoX9iSC-goA","version":{"number":"8.0.0-SNAPSHOT","build_flavor":"default","build_type":"docker","build_hash":"43884496262f71aa3f33b34ac2f2271959dbf12a","build_date":"2020-10-28T09:54:14.068503Z","build_snapshot":true,"lucene_version":"8.7.0","minimum_wire_compatibility_version":"7.11.0","minimum_index_compatibility_version":"7.0.0"},"tagline":"You Know, for Search"}`))
	}
	licenseHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "license": { "type": "` + license + `" } }`))
	}

	mux := http.NewServeMux()
	mux.Handle("/_nodes/_local/nodes", http.HandlerFunc(nodesLocalHandler))
	mux.Handle("/_cluster/state/master_node", http.HandlerFunc(clusterStateMasterHandler))
	mux.Handle("/_license", http.HandlerFunc(licenseHandler))       // for 7.0 and above
	mux.Handle("/_xpack/license", http.HandlerFunc(licenseHandler)) // for before 7.0

	mux.Handle("/_xpack/usage", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			input, _ := ioutil.ReadFile("./_meta/test/xpack-usage.710.json")
			w.Write(input)
		}))

	mux.Handle("/_cluster/state/metadata,routing_table", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			input, _ := ioutil.ReadFile("./_meta/test/cluster_state.710.json")
			w.Write(input)
		}))

	mux.Handle("/", http.HandlerFunc(rootHandler))

	return mux
}

func TestData(t *testing.T) {
	mux := createEsMuxer("7.6.0", "platinum", false)

	server := httptest.NewServer(mux)
	defer server.Close()

	ms := mbtest.NewReportingMetricSetV2Error(t, getConfig(server.URL))
	if err := mbtest.WriteEventsReporterV2Error(ms, t, ""); err != nil {
		t.Fatal("errors writing events to data.json file", err)
	}
}
func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     elasticsearch.ModuleName,
		"metricsets": []string{"index"},
		"hosts":      []string{host},
	}
}
