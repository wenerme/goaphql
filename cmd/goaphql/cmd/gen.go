// Copyright © 2018 wener <wenermail@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// goaphql generate graphql-java -d ./gen -D SchemaName=Test -D JavaPackage=me.wener.test test.graphqls

var generateConf = &_generateConf{}

type _generateConf struct {
	Templates   []string
	Destination string
	Schemas     []string
	Includes    []string

	OptionMap map[string]string
	Options   []string
}

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g", "gen"},
	Short:   "Code generate tools",
	Long:    `Generate code for other implementations`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("gen called")
	//},
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.PersistentFlags().StringVarP(&generateConf.Destination, "dest", "d", "out", "Output dir")
	genCmd.PersistentFlags().StringSliceVarP(&generateConf.Templates, "template", "t", nil, "Custom templates")
	genCmd.PersistentFlags().StringSliceVarP(&generateConf.Schemas, "schema", "s", nil, "GraphQL schemas")
	genCmd.PersistentFlags().StringSliceVarP(&generateConf.Includes, "include", "i", nil, "Include extra schemas but not generate")
	genCmd.PersistentFlags().StringSliceVar(&generateConf.Options, "D", nil, "Custom option")

	genCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		generateConf.OptionMap = sliceToMap(generateConf.Options)
		return nil
	}
}
