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

package export

import (
	"fmt"
	"os"

	"github.com/elastic/beats/libbeat/cmd/instance"
	"github.com/elastic/beats/libbeat/index"
	"github.com/spf13/cobra"
)

func GenTemplateConfigCmd(settings instance.Settings, name, idxPrefix, beatVersion string) *cobra.Command {
	genTemplateConfigCmd := &cobra.Command{
		Use:   "template",
		Short: "Export index templates to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			version, _ := cmd.Flags().GetString("es.version")
			idx, _ := cmd.Flags().GetString("index")

			b, err := instance.NewBeat(name, idx, version)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error initializing beat: %s\n", err)
				os.Exit(1)
			}
			err = b.InitWithSettings(settings)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error initializing beat: %s\n", err)
				os.Exit(1)
			}

			if err != nil {
				fmt.Fprintf(os.Stderr, "error fetching version: %s\n", err)
				os.Exit(1)
			}

			var indexCfgs index.IndexConfigs
			if b.Config.Indices != nil {
				if err := b.Config.Indices.Unpack(&indexCfgs); err != nil {
					fmt.Fprintf(os.Stderr, "unpacking indices config fails: %v", err)
					os.Exit(1)
				}
			}
			if len(indexCfgs) == 0 {
				indexCfgs, err = index.DeprecatedConfigs(b.Config.Template)
				if err != nil {
					fmt.Fprintf(os.Stderr, "unpacking template config fails: %v", err)
					os.Exit(1)
				}
			}
			if err = indexCfgs.LoadTemplates(nil, b.Info); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
		},
	}

	genTemplateConfigCmd.Flags().String("es.version", beatVersion, "Elasticsearch version")
	genTemplateConfigCmd.Flags().String("index", idxPrefix, "Base index name")

	return genTemplateConfigCmd
}
