// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	newrelic_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/newrelic"

	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/spf13/cobra"
)

func newCmdNewRelicImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "newrelic",
		Short: "Import current state to Terraform configuration from New Relic",
		Long:  "Import current state to Terraform configuration from New Relic",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newNewRelicProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newNewRelicProvider()))
	cmd.PersistentFlags().BoolVarP(&options.Connect, "connect", "c", false, "")
	cmd.PersistentFlags().StringSliceVarP(&options.Resources, "resources", "r", []string{}, "alert")
	cmd.PersistentFlags().StringVarP(&options.PathPattern, "path-pattern", "p", DefaultPathPattern, "{output}/{provider}/custom/{service}/")
	cmd.PersistentFlags().StringVarP(&options.PathOutput, "path-output", "o", DefaultPathOutput, "")
	cmd.PersistentFlags().StringVarP(&options.State, "state", "s", DefaultState, "local or bucket")
	cmd.PersistentFlags().StringVarP(&options.Bucket, "bucket", "b", "", "gs://terraform-state")
	return cmd
}

func newNewRelicProvider() terraform_utils.ProviderGenerator {
	return &newrelic_terraforming.NewRelicProvider{}
}
