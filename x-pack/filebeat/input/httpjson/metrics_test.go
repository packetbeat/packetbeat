package httpjson

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/elastic/beats/v7/libbeat/monitoring/inputmon"
	beattest "github.com/elastic/beats/v7/libbeat/publisher/testing"
	conf "github.com/elastic/elastic-agent-libs/config"
	"github.com/elastic/elastic-agent-libs/monitoring"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestMetrics(t *testing.T) {
	testCases := []struct {
		name           string
		setupServer    func(*testing.T, http.HandlerFunc, map[string]interface{})
		baseConfig     map[string]interface{}
		handler        http.HandlerFunc
		expectedEvents []string
		assertMetrics  func(reg *monitoring.Registry) error
	}{
		{
			name: "Test pagination metrics",
			setupServer: func(t *testing.T, h http.HandlerFunc, config map[string]interface{}) {
				registerPaginationTransforms()
				registerResponseTransforms()
				t.Cleanup(func() { registeredTransforms = newRegistry() })
				server := httptest.NewServer(h)
				config["request.url"] = server.URL
				t.Cleanup(server.Close)
			},
			baseConfig: map[string]interface{}{
				"interval":       time.Millisecond,
				"request.method": http.MethodGet,
				"response.split": map[string]interface{}{
					"target": "body.items",
					"transforms": []interface{}{
						map[string]interface{}{
							"set": map[string]interface{}{
								"target": "body.page",
								"value":  "[[.last_response.page]]",
							},
						},
					},
				},
				"response.pagination": []interface{}{
					map[string]interface{}{
						"set": map[string]interface{}{
							"target":                 "url.params.page",
							"value":                  "[[.last_response.body.nextPageToken]]",
							"fail_on_template_error": true,
						},
					},
				},
			},
			handler: paginationHandler(),
			expectedEvents: []string{
				`{"foo":"a","page":"0"}`, `{"foo":"b","page":"1"}`, `{"foo":"c","page":"0"}`, `{"foo":"d","page":"0"}`,
				`{"foo":"a","page":"0"}`, `{"foo":"b","page":"1"}`, `{"foo":"c","page":"0"}`, `{"foo":"d","page":"0"}`,
			},
			assertMetrics: func(reg *monitoring.Registry) error {
				checkHistogramCount := func(c int64) func(v interface{}) bool {
					return func(v interface{}) bool {
						m := v.(map[string]interface{})
						h := m["histogram"].(map[string]interface{})
						return h["count"].(int64) == c
					}
				}
				checkValue := func(c int64) func(v interface{}) bool {
					return func(v interface{}) bool {
						return v.(int64) == c
					}
				}

				snapshot := monitoring.CollectStructSnapshot(reg, monitoring.Full, true)

				conds := map[string]func(interface{}) bool{
					"http_request_body_size":                 checkHistogramCount(8),
					"http_request_get":                       checkValue(8),
					"http_request_total":                     checkValue(8),
					"http_response_2xx":                      checkValue(8),
					"http_response_body_size":                checkHistogramCount(8),
					"http_response_total":                    checkValue(8),
					"http_round_trip_time":                   checkHistogramCount(8),
					"httpjson_interval_execution_time":       checkHistogramCount(6),
					"httpjson_interval_pages_execution_time": checkHistogramCount(8),
					"httpjson_interval_total":                checkValue(6),
				}

				for m, f := range conds {
					if !f(snapshot[m]) {
						return fmt.Errorf("unexpected metric value %v for metric %s", snapshot[m], m)
					}
				}

				return nil
			},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			tc.setupServer(t, tc.handler, tc.baseConfig)

			cfg := conf.MustNewConfigFrom(tc.baseConfig)

			conf := defaultConfig()
			assert.NoError(t, cfg.Unpack(&conf))

			chanClient := beattest.NewChanClient(len(tc.expectedEvents))
			t.Cleanup(func() { _ = chanClient.Close() })

			ctx, cancel := newV2Context()
			t.Cleanup(cancel)

			reg, unreg := inputmon.NewInputRegistry("httpjson-test", ctx.ID, nil)
			t.Cleanup(unreg)

			var g errgroup.Group
			g.Go(func() error {
				pub := statelessPublisher{wrapped: chanClient}
				return run(ctx, conf, pub, nil, reg)
			})

			timeout := time.NewTimer(5 * time.Second)
			t.Cleanup(func() { _ = timeout.Stop() })

			if len(tc.expectedEvents) == 0 {
				select {
				case <-timeout.C:
				case got := <-chanClient.Channel:
					t.Errorf("unexpected event: %v", got)
				}
				cancel()
				assert.NoError(t, g.Wait())
				return
			}

			var receivedCount int
		wait:
			for {
				select {
				case <-timeout.C:
					t.Errorf("timed out waiting for %d events", len(tc.expectedEvents))
					cancel()
					return
				case got := <-chanClient.Channel:
					val, err := got.Fields.GetValue("message")
					assert.NoError(t, err)
					assert.JSONEq(t, tc.expectedEvents[receivedCount], val.(string))
					receivedCount += 1
					if receivedCount == len(tc.expectedEvents) {
						cancel()
						break wait
					}
				}
			}

			assert.NoError(t, g.Wait())
			assert.NoError(t, tc.assertMetrics(reg))
		})
	}
}
