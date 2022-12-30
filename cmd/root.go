// Package cmd provides a cli script that parses the github action.yml file and outputs a markdown table to a specified path.
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
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var inputsHeader = "## Inputs"
var outputsHeader = "## Outputs"
var autoDocStart = "<!-- AUTO-DOC-%s:START - Do not remove or modify this section -->"
var autoDocEnd = "<!-- AUTO-DOC-%s:END -->"
var pipeSeparator = "|"
var newLineSeparator = "\n"
var inputAutoDocStart = fmt.Sprintf(autoDocStart, "INPUT")
var inputAutoDocEnd = fmt.Sprintf(autoDocEnd, "INPUT")
var outputAutoDocStart = fmt.Sprintf(autoDocStart, "OUTPUT")
var outputAutoDocEnd = fmt.Sprintf(autoDocEnd, "OUTPUT")

var defaultInputColumns = []string{"Input", "Type", "Required", "Default", "Description"}
var defaultOutputColumns = []string{"Output", "Type", "Description"}

var actionFileName string
var outputFileName string
var colMaxWidth string
var colMaxWords string
var inputColumns = defaultInputColumns
var outputColumns = defaultOutputColumns

// Input represents the input of the action.yml
type Input struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default,omitempty"`
}

// Output represents the output of the action.yml
type Output struct {
	Description string `yaml:"description"`
	Value       string `yaml:"default,omitempty"`
}

// Action represents the action.yml
type Action struct {
	Inputs  map[string]Input  `yaml:"inputs,omitempty"`
	Outputs map[string]Output `yaml:"outputs,omitempty"`
}

func (a *Action) getAction() error {
	actionYaml, err := ioutil.ReadFile(actionFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(actionYaml, &a)

	return err
}

func (a *Action) renderOutput() error {
	var err error
	maxWidth, err := strconv.Atoi(colMaxWidth)
	if err != nil {
		return err
	}

	maxWords, err := strconv.Atoi(colMaxWords)
	if err != nil {
		return err
	}

	inputTableOutput := &strings.Builder{}

	if len(a.Inputs) > 0 {
		_, err = fmt.Fprintln(inputTableOutput, inputAutoDocStart)
		if err != nil {
			return err
		}

		inputTable := tablewriter.NewWriter(inputTableOutput)
		inputTable.SetHeader(inputColumns)
		inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		inputTable.SetCenterSeparator(pipeSeparator)
		inputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(a.Inputs))
		for k := range a.Inputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		inputTable.SetColWidth(maxWidth)

		for _, key := range keys {
			var inputDefault string
			if len(a.Inputs[key].Default) > 0 {
				inputDefault = a.Inputs[key].Default
				var defaultValue string
				var parts = strings.Split(inputDefault, "\n")

				if len(parts) > 1 && inputDefault != newLineSeparator {
					for _, part := range parts {
						if part != "" {
							defaultValue += "`\"" + part + "\"`" + "<br>"
						}
					}
				} else {
					if strings.Contains(inputDefault, pipeSeparator) {
						inputDefault = strings.Replace(inputDefault, pipeSeparator, "\"\\"+pipeSeparator+"\"", -1)
					} else {
						inputDefault = fmt.Sprintf("%#v", a.Inputs[key].Default)
					}
					defaultValue = "`" + inputDefault + "`"
				}

				inputDefault = defaultValue
			}

			var row []string

			for _, col := range inputColumns {
				switch col {
				case "Input":
					row = append(row, key)
				case "Type":
					row = append(row, "string")
				case "Required":
					row = append(row, strconv.FormatBool(a.Inputs[key].Required))
				case "Default":
					row = append(row, inputDefault)
				case "Description":
					row = append(row, wordWrap(a.Inputs[key].Description, maxWords))
				default:
					return fmt.Errorf(
						"unknown input column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(defaultInputColumns, ", "),
					)
				}
			}
			inputTable.Append(row)
		}

		_, err = fmt.Fprintln(inputTableOutput)
		if err != nil {
			return err
		}

		inputTable.Render()

		_, err = fmt.Fprintln(inputTableOutput)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(inputTableOutput, inputAutoDocEnd)
		if err != nil {
			return err
		}
	}

	outputTableOutput := &strings.Builder{}

	if len(a.Outputs) > 0 {
		_, err = fmt.Fprintln(outputTableOutput, outputAutoDocStart)
		if err != nil {
			return err
		}

		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader(outputColumns)
		outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		outputTable.SetCenterSeparator(pipeSeparator)
		outputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(a.Outputs))
		for k := range a.Outputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		outputTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string

			for _, col := range outputColumns {
				switch col {
				case "Output":
					row = append(row, key)
				case "Type":
					row = append(row, "string")
				case "Description":
					row = append(row, wordWrap(a.Outputs[key].Description, maxWords))
				default:
					return fmt.Errorf(
						"unknown output column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(defaultOutputColumns, ", "),
					)
				}
			}
			outputTable.Append(row)
		}

		_, err = fmt.Fprintln(outputTableOutput)
		if err != nil {
			return err
		}
		outputTable.Render()

		_, err = fmt.Fprintln(outputTableOutput)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(outputTableOutput, outputAutoDocEnd)
		if err != nil {
			return err
		}
	}

	input, err := ioutil.ReadFile(outputFileName)

	if err != nil {
		return err
	}

	var output []byte

	hasInputsData, inputStartIndex, inputEndIndex := hasBytesInBetween(
		input,
		[]byte(inputsHeader),
		[]byte(inputAutoDocEnd),
	)

	if hasInputsData {
		inputsStr := fmt.Sprintf("%s\n\n%v", inputsHeader, inputTableOutput.String())
		output = replaceBytesInBetween(input, inputStartIndex, inputEndIndex, []byte(inputsStr))
	} else {
		inputsStr := fmt.Sprintf("%s\n\n%v", inputsHeader, inputTableOutput.String())
		output = bytes.Replace(input, []byte(inputsHeader), []byte(inputsStr), -1)
	}

	hasOutputsData, outputStartIndex, outputEndIndex := hasBytesInBetween(
		output,
		[]byte(outputsHeader),
		[]byte(outputAutoDocEnd),
	)

	if hasOutputsData {
		outputsStr := fmt.Sprintf("%s\n\n%v", outputsHeader, outputTableOutput.String())
		output = replaceBytesInBetween(output, outputStartIndex, outputEndIndex, []byte(outputsStr))
	} else {
		outputsStr := fmt.Sprintf("%s\n\n%v", outputsHeader, outputTableOutput.String())
		output = bytes.Replace(output, []byte(outputsHeader), []byte(outputsStr), -1)
	}

	if len(output) > 0 {
		if err = ioutil.WriteFile(outputFileName, output, 0666); err != nil {
			cobra.CheckErr(err)
		}
	}

	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-doc",
	Short: "Auto doc generator for your github action",
	Long:  `Auto generate documentation for your github action.`,
	RunE:  RootCmdRunE,
}

// RootCmdRunE runs the root commands RunE function	and handles invalid options and prints the help message
// if the command is called with no arguments.
func RootCmdRunE(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("requires no positional arguments: %d given", len(args))
	}

	var action Action

	err := action.getAction()
	if err != nil {
		return err
	}

	err = action.renderOutput()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(
		cmd.OutOrStdout(),
		"Successfully generated documentation",
	)

	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// RootCmdFlags adds the flags to the root command.
func RootCmdFlags(cmd *cobra.Command) {
	// Custom flags
	cmd.Flags().StringVar(
		&actionFileName,
		"action",
		"action.yml",
		"action config file",
	)
	cmd.Flags().StringVar(
		&outputFileName,
		"output",
		"README.md",
		"Output file",
	)
	cmd.Flags().StringVar(
		&colMaxWidth,
		"colMaxWidth",
		"1000",
		"Max width of a column",
	)
	cmd.Flags().StringVar(
		&colMaxWords,
		"colMaxWords",
		"6",
		"Max number of words per line in a column",
	)
	cmd.Flags().StringArrayVar(
		&inputColumns,
		"inputColumns",
		defaultInputColumns,
		"list of input column names",
	)
	cmd.Flags().StringArrayVar(
		&outputColumns,
		"outputColumns",
		defaultOutputColumns,
		"list of output column names",
	)
}

func init() {
	RootCmdFlags(rootCmd)
}

func hasBytesInBetween(value, start, end []byte) (found bool, startIndex int, endIndex int) {
	s := bytes.Index(value, start)

	if s == -1 {
		return false, -1, -1
	}

	e := bytes.Index(value, end)

	if e == -1 {
		return false, -1, -1
	}

	return true, s, e + len(end)
}

func replaceBytesInBetween(value []byte, startIndex int, endIndex int, new []byte) []byte {
	t := make([]byte, len(value)+len(new))
	w := 0

	w += copy(t[:startIndex], value[:startIndex])
	w += copy(t[w:w+len(new)], new)
	w += copy(t[w:], value[endIndex:])
	return t[0:w]
}

func wordWrap(s string, limit int) string {
	if strings.TrimSpace(s) == "" {
		return s
	}
	// compile regular expressions for Markdown links and code blocks and code
	linkRegex := regexp.MustCompile(`\[.*]\(.*\)`)
	codeBlockRegex := regexp.MustCompile(`\` + "```" + `.*` + "```" + `\s*`)

	// convert string to slice
	strSlice := strings.Fields(s)
	currentLimit := limit

	var result string

	for len(strSlice) >= 1 {
		// convert slice/array back to string
		// but insert <br> at specified limit
		// unless the current slice contains a Markdown link or code block or code
		hasMore := len(strSlice) > currentLimit

		if hasMore && len(result) > 0 {
			result += " "
		}

		if len(strSlice) < currentLimit {
			currentLimit = len(strSlice)
			result = result + strings.Join(strSlice[:currentLimit], " ")
		} else if currentLimit == limit && !linkRegex.MatchString(strings.Join(strSlice[:currentLimit], " ")) && !codeBlockRegex.MatchString(strings.Join(strSlice[:currentLimit], " ")) {
			result = result + strings.Join(strSlice[:currentLimit], " ") + "<br>"
		} else {
			result = result + strings.Join(strSlice[:currentLimit], " ")
		}

		// discard the elements that were copied over to result
		strSlice = strSlice[currentLimit:]

		// change the limit
		// to cater for the last few words in
		//
		if len(strSlice) < currentLimit {
			currentLimit = len(strSlice)
		}
	}

	return strings.TrimSpace(result)
}
