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
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var InputsHeader = "## Inputs"
var OutputsHeader = "## Outputs"
var AutoDocStart = "<!-- AUTO-DOC-%s:START - Do not remove or modify this section --> \n"
var AutoDocEnd = "<!-- AUTO-DOC-%s:END -->"
var inputAutoDocStart = fmt.Sprintf(AutoDocStart, "INPUT")
var inputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "INPUT")
var outputAutoDocStart = fmt.Sprintf(AutoDocStart, "OUTPUT")
var outputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "OUTPUT")

var actionFileName string
var outputFileName string

type Input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default,omitempty"`
}

type Output struct {
	Description string `yaml:"description"`
	Value       string `yaml:"default,omitempty"`
}

type Action struct {
	Inputs  map[string]Input  `yaml:"inputs,omitempty"`
	Outputs map[string]Output `yaml:"outputs,omitempty"`
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

		inputTableOutput := &strings.Builder{}
		if len(action.Inputs) > 0 {
			_, err = fmt.Fprintln(inputTableOutput, inputAutoDocStart)

			if err != nil {
				cobra.CheckErr(err)
			}

			inputTable := tablewriter.NewWriter(inputTableOutput)
			inputTable.SetHeader([]string{"Input", "Required", "Default", "Description"})
			inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			inputTable.SetCenterSeparator("|")

			for key, input := range action.Inputs {
				row := []string{key, strconv.FormatBool(input.Required), input.Default, input.Description}
				inputTable.Append(row)
			}

			inputTable.Render()

			_, err = fmt.Fprintln(inputTableOutput, inputAutoDocEnd)
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		outputTableOutput := &strings.Builder{}

		if len(action.Outputs) > 0 {
			_, err = fmt.Fprintln(outputTableOutput, outputAutoDocStart)

			if err != nil {
				cobra.CheckErr(err)
			}

			outputTable := tablewriter.NewWriter(outputTableOutput)
			outputTable.SetHeader([]string{"Output", "Description", "Value"})
			outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			outputTable.SetCenterSeparator("|")

			for key, output := range action.Outputs {
				row := []string{key, output.Description, output.Value}
				outputTable.Append(row)
			}

			outputTable.Render()

			_, err = fmt.Fprintln(outputTableOutput, outputAutoDocEnd)
			if err != nil {
				cobra.CheckErr(err)
			}
		}

		input, err := ioutil.ReadFile(outputFileName)

		if err != nil {
			cobra.CheckErr(err)
		}

		var output = []byte("")

		hasInputsData, _, _ := HasBytesInBetween(
			input,
			[]byte(InputsHeader),
			[]byte(inputAutoDocEnd),
		)

		if hasInputsData {
			inputsStr := fmt.Sprintf("%s\n%v\n", InputsHeader, inputTableOutput.String())
			output = ReplaceBytesInBetween(input, inputStartIndex, inputEndIndex, []byte(inputsStr))
			fmt.Println(output)
		} else {
			inputsStr := fmt.Sprintf("%s\n%v\n", InputsHeader, inputTableOutput.String())
			output = bytes.Replace(input, []byte(InputsHeader), []byte(inputsStr), -1)
		}

		hasOutputsData, _, _ := HasBytesInBetween(
			output,
			[]byte(OutputsHeader),
			[]byte(outputAutoDocEnd),
		)

		if hasOutputsData {
			outputsStr := fmt.Sprintf("%s\n%v\n", OutputsHeader, outputTableOutput.String())
			fmt.Println(outputsStr)
			//output = ReplaceBytesInBetween(output, outputStartIndex, outputEndIndex, []byte(outputsStr))
		} else {
			outputsStr := fmt.Sprintf("%s\n%v\n", OutputsHeader, outputTableOutput.String())
			output = bytes.Replace(output, []byte(OutputsHeader), []byte(outputsStr), -1)
		}

		if len(output) > 0 {
			if err = ioutil.WriteFile(outputFileName, output, 0666); err != nil {
				cobra.CheckErr(err)
			}
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

func HasBytesInBetween(value, start, end []byte) (found bool, startIndex int, endIndex int) {
	s := bytes.Index(value, start)

	if s == -1 {
		return false, -1, -1
	}

	s += len(start)
	e := bytes.Index(value[s:], end)

	if e == -1 {
		return false, -1, -1
	}

	e += s + e - 1
	return true, s, e
}

func ReplaceBytesInBetween(value []byte, startIndex int, endIndex int, new []byte) []byte {
	t := make([]byte, len(value)+len(new))

	copy(t[:startIndex-1], value[:startIndex-1])
	copy(t[startIndex:endIndex], new)
	copy(t[endIndex+1:], value[endIndex+1:])
	return t
}
