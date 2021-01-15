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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/wenerme/goaphql/cmd/goaphql/cmd/goaphqltmpl"
)

// goaphql g dump -t tmpl/graphql-java -d ./tmpl
var dumpConf = &_dumpConf{}

type _dumpConf struct {
	ListOnly    bool
	Prefix      string
	Matches     []string
	Destination string
}

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump binded assets",
	Long:  `Dump assets to local for easy to modify and test`,
	Run: func(cmd *cobra.Command, args []string) {
		if dumpConf.ListOnly {
			if len(dumpConf.Matches) > 0 {
				match := dumpConf.Matches[0]
				for _, n := range goaphqltmpl.AssetNames() {
					if matchAsset(match, n) {
						fmt.Println(n)
					}
				}
			} else {
				for _, n := range goaphqltmpl.AssetNames() {
					fmt.Println(n)
				}
			}

			return
		}

		files := make(map[string][]byte)

		if len(dumpConf.Matches) > 0 {
			match := dumpConf.Matches[0]
			for _, n := range goaphqltmpl.AssetNames() {
				if matchAsset(match, n) {
					files[n] = goaphqltmpl.MustAsset(n)
				}
			}
		}

		for k, v := range files {
			fn := filepath.Join(dumpConf.Destination, strings.TrimLeft(k, dumpConf.Prefix))
			dir := filepath.Dir(fn)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				logrus.WithError(err).WithField("Dir", dir).Fatal("failed to make dir")
			}
			if err := ioutil.WriteFile(fn, v, os.ModePerm); err != nil {
				logrus.WithError(err).WithField("File", k).Fatal("failed to write file")
			}
			logrus.WithField("File", fn).Info("dump file")
		}
	},
}

func matchAsset(p string, n string) bool {
	if strings.HasPrefix(n, p) {
		return true
	} else if m, _ := filepath.Match(p, n); m {
		return true
	}
	return false
}

func init() {
	RootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().BoolVarP(&dumpConf.ListOnly, "list", "l", false, "List all assets and exit")
	dumpCmd.Flags().StringVarP(&dumpConf.Prefix, "prefix", "p", "", "Prefix to trim")
	dumpCmd.Flags().StringVarP(&dumpConf.Destination, "dest", "d", "", "Output dir")
	dumpCmd.Flags().StringSliceVarP(&dumpConf.Matches, "match", "m", nil, "Matches")
}
