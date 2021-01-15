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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wenerme/goaphql/cmd/goaphql/cmd/goaphqltmpl"
	"github.com/wenerme/goaphql/pkg/gqlg"
	"github.com/wenerme/goaphql/pkg/gqll"
	"github.com/wenerme/goaphql/pkg/gqlp"
)

var generateGJConf = &_generateGJConf{}

type _generateGJConf struct {
	JavaPackage    string
	ImportPackages []string
	JavaTypeMap    map[string]string
	JavaTypeMaps   []string
	SchemaName     string
	Tags           []string
}

// genGjCmd represents the genGj command
var genGJCmd = &cobra.Command{
	Use:     "graphql-java",
	Aliases: []string{"gj", "jraphql"},
	Short:   "Generate code for graphql-java",
	Long:    `Generate code for graphql-java`,
	Run: func(cmd *cobra.Command, args []string) {

		config := gqlg.JavaGeneratorConfig{
			JavaPackage:    generateGJConf.JavaPackage,
			ImportPackages: generateGJConf.ImportPackages,
			JavaTypeMap:    generateGJConf.JavaTypeMap,
			SchemaName:     generateGJConf.SchemaName,
		}
		var definitions []gqll.Definition
		var includes []gqll.Definition
		var err error
		for _, schema := range generateConf.Schemas {
			logrus := logrus.WithField("Schema", schema)

			bytes, err := ioutil.ReadFile(schema)
			if err != nil {
				logrus.WithError(err).Fatal("open schema")
			}
			if doc, err := gqlp.ParseContent(string(bytes)); err != nil {
				logrus.WithError(err).Fatal("parse schema")
			} else {
				logrus.WithField("Definitions", len(doc.Definitions)).Debug("parse schema")
				definitions = append(definitions, doc.Definitions...)
			}
		}

		for _, schema := range generateConf.Includes {
			logrus := logrus.WithField("Include", schema)

			bytes, err := ioutil.ReadFile(schema)
			if err != nil {
				logrus.WithError(err).Fatal("open include schema")
			}
			if doc, err := gqlp.ParseContent(string(bytes)); err != nil {
				logrus.WithError(err).Fatal("parse include schema")
			} else {
				logrus.WithField("Definitions", len(doc.Definitions)).Debug("parse schema")
				includes = append(includes, doc.Definitions...)
			}
		}

		for _, v := range includes {
			if gqll.TypeOf(v).IsTypeDefinition() {
				config.Ignores = append(config.Ignores, gqll.NameOf(v))
			}
		}

		// build type system
		definitions = append(definitions, includes...)
		if config.TypeSystem, err = gqll.NewTypeSystem(definitions); err != nil {
			logrus.WithError(err).Fatal("failed to build type system")
		}

		g, err := gqlg.NewGraphQLJavaGenerator(config)
		if err != nil {
			logrus.WithError(err).Fatal("failed to create generator")
		}
		if len(generateConf.Templates) != 0 {
			tpl := generateConf.Templates[0]
			if stat, err := os.Stat(tpl); err != nil {
				if _, err = g.Template.ParseGlob(tpl); err != nil {
					logrus.WithError(err).WithField("Template", tpl).Fatal("failed parse template glob")
				}
			} else if stat.IsDir() {
				if _, err = g.Template.ParseGlob(tpl + "/**.tmpl"); err != nil {
					logrus.WithError(err).WithField("Template", tpl).Fatal("failed parse template dir")
				}
			} else {
				bytes, err := ioutil.ReadFile(tpl)
				if err != nil {
					logrus.Fatal("failed to open template file")
				}
				if _, err = g.Template.New(filepath.Base(tpl)).Parse(string(bytes)); err != nil {
					logrus.WithError(err).WithField("Template", tpl).Fatal("failed parse template file")
				}
			}
		} else {
			// Use binddata
			match := "tmpl/graphql-java/"
			for _, v := range goaphqltmpl.AssetNames() {
				if matchAsset(match, v) {
					if _, err = g.Template.New(filepath.Base(v)).Parse(string(goaphqltmpl.MustAsset(v))); err != nil {
						logrus.WithError(err).WithField("Template", v).Fatal("failed parse asset template")
					}
				}
			}
		}

		if err := g.GenerateGraphQLJava(); err != nil {
			logrus.WithError(err).Fatal("failed generate code")
		}

		for _, v := range g.Files {
			var pkg string
			if v, ok := v.(*gqlg.GenerateJavaFile); ok {
				pkg = v.JavaPackage
			}
			fn := filepath.Join(generateConf.Destination, strings.Replace(pkg, ".", "/", -1), v.Name())
			dir := filepath.Dir(fn)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				logrus.WithError(err).WithField("Dir", dir).Fatal("failed to make dir")
			}
			if err := ioutil.WriteFile(fn, v.Buf().Bytes(), os.ModePerm); err != nil {
				logrus.WithError(err).WithField("File", fn).Fatal("failed to write file")
			}
			logrus.WithField("File", fn).Info("write file")
		}
	},
}

func init() {
	genCmd.AddCommand(genGJCmd)

	genGJCmd.PersistentFlags().StringVarP(&generateGJConf.JavaPackage, "package", "P", "", "Target java package")
	genGJCmd.PersistentFlags().StringSliceVarP(&generateGJConf.ImportPackages, "import", "I", nil, "Extra imports")
	genGJCmd.PersistentFlags().StringSliceVarP(&generateGJConf.JavaTypeMaps, "type-map", "M", nil, "Type mapping")
	genGJCmd.PersistentFlags().StringVarP(&generateGJConf.SchemaName, "schema-name", "N", "", "Schema name")
	genGJCmd.PersistentFlags().StringSliceVarP(&generateGJConf.Tags, "tag", "T", nil, "Select tag [WIP]")

	genGJCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		generateGJConf.JavaTypeMap = sliceToMap(generateGJConf.JavaTypeMaps)
		return nil
	}
}
