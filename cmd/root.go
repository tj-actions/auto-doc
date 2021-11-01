// Package cmd
/*
Copyright © 2021 Tonye Jack <jtonye@ymail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var actionFileName string
var outputFileName string

type Input struct {
	Description string    `yaml:"description"`
	Required    bool 	  `yaml:"required"`
	Default     string    `yaml:"default,omitempty"`
}

type Output struct {
	Description string    `yaml:"description"`
	Value       string    `yaml:"default,omitempty"`
}

type Action struct {
	Inputs      map[string]Input `yaml:"inputs,omitempty"`
	Outputs     map[string]Output `yaml:"outputs,omitempty"`
}

func (a *Action) getAction() *Action {
	actionYaml, err := ioutil.ReadFile(actionFileName)
	if err != nil {
		cobra.CheckErr(err)
	}

	err = yaml.Unmarshal(actionYaml, a)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return a
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-doc",
	Short: "Auto doc generator for your github action",
	Long:  `Auto generate documentation for your github action.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if len(args) > 0 {
			_, err := fmt.Fprintf(
				os.Stderr,
				"'%d' invalid arguments passed.\n",
				len(args),
			)
			if err != nil {
				cobra.CheckErr(err)
			}
			return
		}

		var action Action
		action.getAction()

		if len(action.Inputs) > 0 {
			inputTable := tablewriter.NewWriter(os.Stdout)
			inputTable.SetHeader([]string{"Input", "Required", "Default", "Description"})

			for key, input := range action.Inputs {
				row := []string{key, strconv.FormatBool(input.Required), input.Default, input.Description}
				inputTable.Append(row)
			}

			inputTable.Render()
		}

		if len(action.Outputs) > 0 {
			outputTable := tablewriter.NewWriter(os.Stdout)
			outputTable.SetHeader([]string{"Output", "Description", "Value"})

			for key, output := range action.Outputs {
				row := []string{key, output.Description, output.Value}
				outputTable.Append(row)
			}

			outputTable.Render()
		}

		outputFile, err := os.Open(outputFileName)

		if err != nil {
			cobra.CheckErr(err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				cobra.CheckErr(err)
			}
		}(outputFile)

		scanner := bufio.NewScanner(outputFile)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			cobra.CheckErr(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Custom flags
	rootCmd.PersistentFlags().StringVar(
		&actionFileName,
		"action",
		"action.yml",
		"action config file",
	)
	rootCmd.PersistentFlags().StringVar(
		&outputFileName,
		"output",
		"README.md",
		"Output file",
	)
}
