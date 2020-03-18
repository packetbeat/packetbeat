package ccr

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/metricbeat/module/elasticsearch"

	mbtest "github.com/elastic/beats/v7/metricbeat/mb/testing"
)

func startESServer(esVersion, license string, ccrAvailable bool) *httptest.Server {

	nodesLocalHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"nodes": { "foobar": {}}}`))
	}
	clusterStateMasterHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"master_node": "foobar"}`))
	}
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		}
		w.Write([]byte(`{"version": { "number": "` + esVersion + `" } }`))
	}
	licenseHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "license": { "type": "` + license + `" } }`))
	}
	xpackHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{ "features": { "ccr": { "available": ` + strconv.FormatBool(ccrAvailable) + `}}}`))
	}
	ccrStatsHandler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "this should never have been called", 418)
	}

	mux := http.NewServeMux()
	mux.Handle("/_nodes/_local/nodes", http.HandlerFunc(nodesLocalHandler))
	mux.Handle("/_cluster/state/master_node", http.HandlerFunc(clusterStateMasterHandler))
	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/_license", http.HandlerFunc(licenseHandler))       // for 7.0 and above
	mux.Handle("/_xpack/license", http.HandlerFunc(licenseHandler)) // for before 7.0
	mux.Handle("/_xpack", http.HandlerFunc(xpackHandler))
	mux.Handle("/_ccr/stats", http.HandlerFunc(ccrStatsHandler))

	return httptest.NewServer(mux)
}

func TestCCRNotAvailable(t *testing.T) {
	tests := map[string]struct {
		esVersion    string
		license      string
		ccrAvailable bool
	}{
		"old_version": {
			"6.4.0",
			"platinum",
			true,
		},
		"low_license": {
			"7.6.0",
			"basic",
			true,
		},
		"feature_unavailable": {
			"7.6.0",
			"platinum",
			false,
		},
	}

	// Disable license caching for these tests
	elasticsearch.LicenseEnableCaching = false
	defer func() { elasticsearch.LicenseEnableCaching = true }()

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			server := startESServer(test.esVersion, test.license, test.ccrAvailable)
			defer server.Close()

			ms := mbtest.NewReportingMetricSetV2Error(t, getConfig(server.URL))
			events, errs := mbtest.ReportingFetchV2Error(ms)

			require.Empty(t, errs)
			require.Empty(t, events)
		})
	}
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     elasticsearch.ModuleName,
		"metricsets": []string{"ccr"},
		"hosts":      []string{host},
	}
}
