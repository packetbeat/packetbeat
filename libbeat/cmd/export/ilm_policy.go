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

	"github.com/elastic/beats/libbeat/index"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/elastic/beats/libbeat/cmd/instance"

	"github.com/spf13/cobra"
)

// GenGetILMPolicyCmd is the command used to export the ilm policy.
func GenGetILMPolicyCmd(settings instance.Settings, name, idxPrefix, beatVersion string) *cobra.Command {
	genILMConfigCmd := &cobra.Command{
		Use:   "ilm-policy",
		Short: "Export ILM policies to stdout",
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

			var indicesCfg index.Configs
			if err := b.Config.Indices.Unpack(&indicesCfg); err != nil {
				fmt.Fprintf(os.Stderr, "unpacking indices config fails: %v", err)
				os.Exit(1)
			}

			if _, err = indicesCfg.PrintILMPolicies(b.Info); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}

			logp.Info("Loaded ILM policies.")
		},
	}

	genILMConfigCmd.Flags().String("version", beatVersion, "Beat version")
	genILMConfigCmd.Flags().String("index", idxPrefix, "Base index name")

	return genILMConfigCmd
}
