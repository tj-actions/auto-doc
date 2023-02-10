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
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var inputsHeader = "## Inputs"
var outputsHeader = "## Outputs"
var secretsHeader = "## Secrets"
var autoDocStart = "<!-- AUTO-DOC-%s:START - Do not remove or modify this section -->"
var autoDocEnd = "<!-- AUTO-DOC-%s:END -->"
var pipeSeparator = "|"
var newLineSeparator = "\n"
var inputAutoDocStart = fmt.Sprintf(autoDocStart, "INPUT")
var inputAutoDocEnd = fmt.Sprintf(autoDocEnd, "INPUT")
var outputAutoDocStart = fmt.Sprintf(autoDocStart, "OUTPUT")
var outputAutoDocEnd = fmt.Sprintf(autoDocEnd, "OUTPUT")
var secretsAutoDocStart = fmt.Sprintf(autoDocStart, "SECRETS")
var secretsAutoDocEnd = fmt.Sprintf(autoDocEnd, "SECRETS")

var defaultInputColumns = []string{"Input", "Type", "Required", "Default", "Description"}
var defaultOutputColumns = []string{"Output", "Type", "Description"}
var defaultSecretsColumns = []string{"Secret", "Required", "Description"}

var documentFileName string
var outputFileName string
var colMaxWidth string
var colMaxWords string
var inputColumns = defaultInputColumns
var outputColumns = defaultOutputColumns
var secretsColumns = defaultSecretsColumns

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

// Secret represents the secret of reusable workflows
type Secret struct {
	Required    bool   `yaml:"required"`
	Description string `yaml:"description"`
}

// Action represents the action.yml
type Action struct {
	Inputs  map[string]Input  `yaml:"inputs,omitempty"`
	Outputs map[string]Output `yaml:"outputs,omitempty"`
}

// Reusable represents the reusable workflow yaml
type Reusable struct {
	On struct {
		WorkflowCall struct {
			Inputs  map[string]Input  `yaml:"inputs,omitempty"`
			Secrets map[string]Secret `yaml:"secrets,omitempty"`
		} `yaml:"workflow_call"`
	}
}

// Documentation is the interface for Action and Reusable
type Documentation interface {
	getData() error
	renderOutput() error
}

func (r *Reusable) getData() error {
	reusableYaml, err := ioutil.ReadFile(documentFileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(reusableYaml, &r)

	return err
}

func (a *Action) getData() error {
	actionYaml, err := ioutil.ReadFile(documentFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(actionYaml, &a)
	return err
}

func renderInputOutput(i map[string]Input, maxWidth int, maxWords int) (*strings.Builder, error) {
	inputTableOutput := &strings.Builder{}

	if len(i) > 0 {
		_, err := fmt.Fprintln(inputTableOutput, inputAutoDocStart)
		if err != nil {
			return inputTableOutput, err
		}

		inputTable := tablewriter.NewWriter(inputTableOutput)
		inputTable.SetHeader(inputColumns)
		inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		inputTable.SetCenterSeparator(pipeSeparator)
		inputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(i))
		for k := range i {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		inputTable.SetColWidth(maxWidth)

		for _, key := range keys {
			var inputDefault string
			if len(i[key].Default) > 0 {
				inputDefault = i[key].Default
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
						inputDefault = fmt.Sprintf("%#v", i[key].Default)
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
					row = append(row, strconv.FormatBool(i[key].Required))
				case "Default":
					row = append(row, inputDefault)
				case "Description":
					row = append(row, wordWrap(i[key].Description, maxWords))
				default:
					return inputTableOutput, fmt.Errorf(
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
			return inputTableOutput, err
		}

		inputTable.Render()

		_, err = fmt.Fprintln(inputTableOutput)
		if err != nil {
			return inputTableOutput, err
		}

		_, err = fmt.Fprint(inputTableOutput, inputAutoDocEnd)
		if err != nil {
			return inputTableOutput, err
		}
	}
	return inputTableOutput, nil
}

func renderOutputOutput(o map[string]Output, maxWidth int, maxWords int) (*strings.Builder, error) {
	outputTableOutput := &strings.Builder{}

	if len(o) > 0 {
		_, err := fmt.Fprintln(outputTableOutput, outputAutoDocStart)
		if err != nil {
			return outputTableOutput, err
		}

		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader(outputColumns)
		outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		outputTable.SetCenterSeparator(pipeSeparator)
		outputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(o))
		for k := range o {
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
					row = append(row, wordWrap(o[key].Description, maxWords))
				default:
					return outputTableOutput, fmt.Errorf(
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
			return outputTableOutput, err
		}
		outputTable.Render()

		_, err = fmt.Fprintln(outputTableOutput)
		if err != nil {
			return outputTableOutput, err
		}

		_, err = fmt.Fprint(outputTableOutput, outputAutoDocEnd)
		if err != nil {
			return outputTableOutput, err
		}
	}
	return outputTableOutput, nil
}

func renderSecretOutput(s map[string]Secret, maxWidth int, maxWords int) (*strings.Builder, error) {
	secretTableOutput := &strings.Builder{}

	if len(s) > 0 {
		_, err := fmt.Fprintln(secretTableOutput, secretsAutoDocStart)
		if err != nil {
			return secretTableOutput, err
		}

		secretTable := tablewriter.NewWriter(secretTableOutput)
		secretTable.SetHeader(secretsColumns)
		secretTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		secretTable.SetCenterSeparator(pipeSeparator)
		secretTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(s))
		for k := range s {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		secretTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string

			for _, col := range secretsColumns {
				switch col {
				case "Secret":
					row = append(row, key)
				case "Required":
					row = append(row, fmt.Sprintf("%v", s[key].Required))
				case "Description":
					row = append(row, s[key].Description)
				default:
					return secretTableOutput, fmt.Errorf(
						"unknown secrets column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(defaultSecretsColumns, ", "),
					)
				}
			}
			secretTable.Append(row)
		}
		_, err = fmt.Fprintln(secretTableOutput)
		if err != nil {
			return secretTableOutput, err
		}
		secretTable.Render()

		_, err = fmt.Fprintln(secretTableOutput)
		if err != nil {
			return secretTableOutput, err
		}

		_, err = fmt.Fprint(secretTableOutput, secretsAutoDocEnd)
		if err != nil {
			return secretTableOutput, err
		}
	}
	return secretTableOutput, nil
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

	inputTableOutput, err := renderInputOutput(a.Inputs, maxWidth, maxWords)
	outputTableOutput, err := renderOutputOutput(a.Outputs, maxWidth, maxWords)
	err = writeDocumentation(inputTableOutput, outputTableOutput)
	if err != nil {
		return err
	}
	return nil
}

func (r *Reusable) renderOutput() error {
	var err error
	maxWidth, err := strconv.Atoi(colMaxWidth)
	if err != nil {
		return err
	}

	maxWords, err := strconv.Atoi(colMaxWords)
	if err != nil {
		return err
	}
	inputTableOutput, err := renderInputOutput(r.On.WorkflowCall.Inputs, maxWidth, maxWords)
	secretTableOutput, err := renderSecretOutput(r.On.WorkflowCall.Secrets, maxWidth, maxWords)
	err = writeDocumentation(inputTableOutput, secretTableOutput)
	if err != nil {
		return err
	}
	return nil
}

func writeDocumentation(in1, in2 *strings.Builder) error {
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
		inputsStr := fmt.Sprintf("%s\n\n%v", inputsHeader, in1.String())
		output = replaceBytesInBetween(input, inputStartIndex, inputEndIndex, []byte(inputsStr))
	} else {
		inputsStr := fmt.Sprintf("%s\n\n%v", inputsHeader, in1.String())
		output = bytes.Replace(input, []byte(inputsHeader), []byte(inputsStr), -1)
	}

	hasOutputsData, outputStartIndex, outputEndIndex := hasBytesInBetween(
		output,
		[]byte(outputsHeader),
		[]byte(outputAutoDocEnd),
	)

	if hasOutputsData {
		outputsStr := fmt.Sprintf("%s\n\n%v", outputsHeader, in2.String())
		output = replaceBytesInBetween(output, outputStartIndex, outputEndIndex, []byte(outputsStr))
	} else {
		outputsStr := fmt.Sprintf("%s\n\n%v", outputsHeader, in2.String())
		output = bytes.Replace(output, []byte(outputsHeader), []byte(outputsStr), -1)
	}

	hasSecretsData, secretsStartIndex, secretsEndIndex := hasBytesInBetween(
		output,
		[]byte(secretsHeader),
		[]byte(secretsAutoDocEnd),
	)

	if hasSecretsData {
		secretsStr := fmt.Sprintf("%s\n\n%v", secretsHeader, in2.String())
		output = replaceBytesInBetween(output, secretsStartIndex, secretsEndIndex, []byte(secretsStr))
	} else {
		secretsStr := fmt.Sprintf("%s\n\n%v", secretsHeader, in2.String())
		output = bytes.Replace(output, []byte(secretsHeader), []byte(secretsStr), -1)
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

	var documentation Documentation

	if cmd.Flags().Changed("reusable") {
		documentation = &Reusable{}
	} else {
		documentation = &Action{}
	}

	err := documentation.getData()
	if err != nil {
		return err
	}

	err = documentation.renderOutput()
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
		&documentFileName,
		"filename",
		"action.yml",
		"config file",
	)
	cmd.Flags().Bool(
		"reusable",
		false,
		"A reusable workflow",
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
	cmd.Flags().StringArrayVar(
		&secretsColumns,
		"secretsColumns",
		defaultSecretsColumns,
		"list of secrets column names",
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
		// to cater for the last few words in the line
		if len(strSlice) < currentLimit {
			currentLimit = len(strSlice)
		}
	}

	// Remove trailing <br> if any
	result = strings.TrimSuffix(result, "<br>")

	return strings.TrimSpace(result)
}
