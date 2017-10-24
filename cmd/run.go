// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"github.com/gaocegege/tosknight/crawler"
	"github.com/gaocegege/tosknight/source"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SOURCEFILE is the name of source file config.
const SOURCEFILE = "sourceFile"

var sourceFile string

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run the spider logic",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&sourceFile, SOURCEFILE, "s", "", "Source directory to read from")
}

func run() error {
	// TODO
	if sourceFile == "" {
		log.Fatalln("There is no source file given")
	}
	sourceManager := source.NewManager()
	sourceManager.AddSource(source.Source{
		URL: "http://localhost:9080/text.html",
	})
	contentCrawler := crawler.New(sourceManager, "/home/ist/go/src/github.com/gaocegege/tosknight/storage")
	contentCrawler.Run()
	log.Println("Run called.")
	return nil
}
