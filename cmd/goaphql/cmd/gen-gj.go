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
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wenerme/goaphql/cmd/goaphql/cmd/goaphqltmpl"
	"github.com/wenerme/goaphql/gqlg"
	"github.com/wenerme/goaphql/gqll"
	"github.com/wenerme/goaphql/gqlp"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var generateGJConf = &_generateGJConf{}

type _generateGJConf struct {
	JavaPackage    string
	ImportPackages []string
	JavaTypeMap    map[string]string
	JavaTypeMaps   []string
	SchemaName     string
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

		if len(generateConf.Schemas) > 0 {
			bytes, err := ioutil.ReadFile(generateConf.Schemas[0])
			if err != nil {
				logrus.WithError(err).Fatal("failed to open schema")
			}
			if doc, err := gqlp.ParseContent(string(bytes)); err != nil {
				logrus.WithError(err).Fatal("failed parse schema")
			} else {
				config.TypeSystem, err = gqll.NewTypeSystem(doc)
				if err != nil {
					logrus.WithError(err).Fatal("failed to build type system")
				}
			}
		} else {
			logrus.Fatal("no schema")
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
		if err = g.ScanTemplate(); err != nil {
			logrus.WithError(err).Fatal("failed to scan template file")
		}

		if err := g.GenerateGraphQLJava(); err != nil {
			logrus.WithError(err).Fatal("failed generate code")
		}

		for _, v := range g.Files {
			fn := filepath.Join(generateConf.Destination, strings.Replace(g.Config.JavaPackage, ".", "/", -1), v.Name)
			dir := filepath.Dir(fn)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				logrus.WithError(err).WithField("Dir", dir).Fatal("failed to make dir")
			}
			if err := ioutil.WriteFile(fn, v.Buf.Bytes(), os.ModePerm); err != nil {
				logrus.WithError(err).WithField("File", fn).Fatal("failed to write file")
			}
			logrus.WithField("File", fn).Info("write file")
		}

	},
}

func init() {
	genCmd.AddCommand(genGJCmd)

	genGJCmd.PersistentFlags().StringVarP(&generateGJConf.JavaPackage, "java-package", "p", "", "Target java package")
	genGJCmd.PersistentFlags().StringSliceVarP(&generateGJConf.ImportPackages, "java-import", "i", nil, "Extra imports")
	genGJCmd.PersistentFlags().StringSliceVarP(&generateGJConf.JavaTypeMaps, "java-type", "m", nil, "Type mapping")
	genGJCmd.PersistentFlags().StringVarP(&generateGJConf.SchemaName, "schema-name", "n", "", "Schema name")

}
