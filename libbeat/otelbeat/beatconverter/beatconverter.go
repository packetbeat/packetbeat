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

package beatconverter

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"go.opentelemetry.io/collector/confmap"

	"github.com/elastic/beats/v7/libbeat/cloudid"
	"github.com/elastic/beats/v7/libbeat/outputs/elasticsearch"
	"github.com/elastic/elastic-agent-libs/config"
)

// list of supported beatreceivers
var supportedReceivers = []string{"filebeatreceiver"} // Add more beat receivers to this list when we add support

type converter struct{}

// NewFactory returns a factory for a  confmap.Converter,
func NewFactory() confmap.ConverterFactory {
	return confmap.NewConverterFactory(newConverter)
}

func newConverter(set confmap.ConverterSettings) confmap.Converter {
	return converter{}
}

// Convert converts [beatreceiver].output to OTel config here
func (c converter) Convert(_ context.Context, conf *confmap.Conf) error {

	for _, beatreceiver := range supportedReceivers {
		var out map[string]any
		var accessString = "receivers::" + beatreceiver

		// check if supported beat receiver is configured. Skip translation logic if not
		if v := conf.Get(accessString); v == nil {
			continue
		}

		// handle cloud id if set
		if conf.IsSet("receivers::" + beatreceiver + "::cloud") {
			if err := handleCloudId(accessString, conf); err != nil {
				return fmt.Errorf("error handling cloud id %w", err)
			}
		}

		// handle http pprof
		if err := httpPPROFEndpoint(accessString, conf); err != nil {
			return fmt.Errorf("error handling http pprof %w", err)
		}

		receiverCfg, _ := conf.Sub(accessString)

		output, _ := receiverCfg.Sub("output")
		if len(output.ToStringMap()) > 1 {
			return fmt.Errorf("multiple outputs are not supported")
		}

		for key, output := range output.ToStringMap() {
			switch key {
			case "elasticsearch":
				esConfig := config.MustNewConfigFrom(output)
				esOTelConfig, err := elasticsearch.ToOTelConfig(esConfig)
				if err != nil {
					return fmt.Errorf("cannot convert elasticsearch config: %w", err)
				}
				out = map[string]any{
					"service::pipelines::logs::exporters": []string{"elasticsearch"},
					"exporters": map[string]any{
						"elasticsearch": esOTelConfig,
					},
				}
				err = conf.Merge(confmap.NewFromStringMap(out))
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("output type %q is unsupported in OTel mode", key)
			}
		}

		// Replace output.[configured-output] with output.otelconsumer
		out = map[string]any{
			accessString + "::output": nil,
		}
		err := conf.Merge(confmap.NewFromStringMap(out))
		if err != nil {
			return err
		}
		out = map[string]any{
			accessString + "::output::otelconsumer": nil,
		}

		err = conf.Merge(confmap.NewFromStringMap(out))
		s, _ := json.MarshalIndent(conf.ToStringMap(), "", " ")
		fmt.Println(string(s), "from outside handling cloud id")
		if err != nil {
			return err
		}
	}

	return nil
}

func handleCloudId(accessString string, conf *confmap.Conf) error {

	receiverCfg, _ := conf.Sub(accessString)
	beatCfg := config.MustNewConfigFrom(receiverCfg.ToStringMap())

	// Handle cloud.id the same way Beats does, this will also handle
	// extracting the Kibana URL
	if err := cloudid.OverwriteSettings(beatCfg); err != nil {
		return fmt.Errorf("cannot read cloudid: %w", err)
	}

	var beatOutput map[string]any
	err := beatCfg.Unpack(&beatOutput)
	if err != nil {
		return err
	}

	out := map[string]any{
		accessString: beatOutput,
	}
	err = conf.Merge(confmap.NewFromStringMap(out))
	if err != nil {
		return err
	}

	// we set this to nil to ensure cloudid check does not throw error when output is next set to otelconsumer
	out = map[string]any{
		accessString + "::cloud": nil,
	}
	err = conf.Merge(confmap.NewFromStringMap(out))
	if err != nil {
		return err
	}

	return nil
}

func httpPPROFEndpoint(accessString string, conf *confmap.Conf) error {
	if v := conf.Get(accessString + "::http::pprof::enabled"); v != nil {
		if v.(bool) {
			var httpHost string
			var httpPort int
			if v := conf.Get(accessString + "::http::host"); v != nil {
				httpHost = v.(string)
			} else {
				httpHost = "localhost"
			}

			if v := conf.Get(accessString + "::http::port"); v != nil {
				httpPort = v.(int)
			} else {
				httpPort = 5066
			}

			out := map[string]any{
				"extensions": map[string]any{
					"pprof": map[string]any{
						"endpoint": httpHost + ":" + strconv.Itoa(httpPort),
					},
				},
				"service": map[string]any{
					"extensions": []string{"pprof"}},
			}

			err := conf.Merge(confmap.NewFromStringMap(out))

			s, _ := json.MarshalIndent(conf.ToStringMap(), "", " ")
			fmt.Println(string(s), "from outside handling cloud id")

			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil

}
