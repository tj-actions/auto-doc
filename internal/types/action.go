//Package types contains all defined types
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
package types

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/tj-actions/auto-doc/v3/internal"
	"github.com/tj-actions/auto-doc/v3/internal/utils"
)

// ActionInput represents the input of the action.yml
type ActionInput struct {
	Description        string `yaml:"description"`
	Required           bool   `yaml:"required"`
	Default            string `yaml:"default,omitempty"`
	DeprecationMessage string `yaml:"deprecationMessage,omitempty"`
}

// ActionOutput represents the output of the action.yml
type ActionOutput struct {
	Description string `yaml:"description"`
	Value       string `yaml:"default,omitempty"`
}

// Action represents the action.yml
type Action struct {
	InputFileName      string
	OutputFileName     string
	ColMaxWidth        string
	ColMaxWords        string
	InputColumns       []string
	OutputColumns      []string
	Inputs             map[string]ActionInput  `yaml:"inputs,omitempty"`
	Outputs            map[string]ActionOutput `yaml:"outputs,omitempty"`
	Description        string                  `yaml:"description,omitempty"`
	InputMarkdownLinks bool
}

// GetData parses the source yaml file
func (a *Action) GetData() error {
	actionYaml, err := os.ReadFile(a.InputFileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(actionYaml, &a)
	return err
}

// WriteDocumentation write the table to the output file
func (a *Action) WriteDocumentation(inputTable, outputTable, description *strings.Builder) error {
	var err error
	input, err := os.ReadFile(a.OutputFileName)
	if err != nil {
		return err
	}

	var output []byte

	hasDescriptionData, indices := utils.HasBytesInBetween(
		input,
		[]byte(internal.DescriptionAutoDocStart),
		[]byte(internal.DescriptionAutoDocEnd),
	)
	output = input

	descriptionStr := strings.TrimSpace(description.String())

	if hasDescriptionData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(descriptionStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.DescriptionHeader))
		output = re.ReplaceAllFunc(input, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.DescriptionHeader)) {
				if descriptionStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.DescriptionHeader, descriptionStr))
				}
				return []byte(internal.DescriptionHeader)
			}
			return match
		})
	}

	hasInputsData, indices := utils.HasBytesInBetween(
		output,
		[]byte(internal.InputAutoDocStart),
		[]byte(internal.InputAutoDocEnd),
	)

	inputsStr := strings.TrimSpace(inputTable.String())

	if hasInputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(inputsStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.InputsHeader))
		output = re.ReplaceAllFunc(input, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.InputsHeader)) {
				if inputsStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.InputsHeader, inputsStr))
				} else {
					return []byte(internal.InputsHeader)
				}
			}
			return match
		})
	}
	hasOutputsData, indices := utils.HasBytesInBetween(
		output,
		[]byte(internal.OutputAutoDocStart),
		[]byte(internal.OutputAutoDocEnd),
	)

	outputsStr := strings.TrimSpace(outputTable.String())

	if hasOutputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(outputsStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.OutputsHeader))
		output = re.ReplaceAllFunc(output, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.OutputsHeader)) {
				if outputsStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.OutputsHeader, outputsStr))
				} else {
					return []byte(internal.OutputsHeader)
				}
			}
			return match
		})
	}

	if err = os.WriteFile(a.OutputFileName, output, 0o666); err != nil {
		cobra.CheckErr(err)
	}

	return nil
}

// RenderOutput renders the output and writes it to the given output
func (a *Action) RenderOutput() error {
	var err error
	maxWidth, err := strconv.Atoi(a.ColMaxWidth)
	if err != nil {
		return err
	}

	maxWords, err := strconv.Atoi(a.ColMaxWords)
	if err != nil {
		return err
	}

	descriptionOutput, err := renderDescription(a.Description)
	if err != nil {
		return err
	}

	inputTableOutput, err := renderActionInputTableOutput(a.Inputs, a.InputColumns, a.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	outputTableOutput, err := renderActionOutputTableOutput(a.Outputs, a.OutputColumns, a.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	err = a.WriteDocumentation(inputTableOutput, outputTableOutput, descriptionOutput)
	if err != nil {
		return err
	}

	return nil
}

// renderDescription renders the description
func renderDescription(description string) (*strings.Builder, error) {
	descriptionOutput := &strings.Builder{}
	_, err := fmt.Fprintln(descriptionOutput, internal.DescriptionAutoDocStart)
	if err != nil {
		return descriptionOutput, err
	}

	_, err = fmt.Fprintln(descriptionOutput)
	if err != nil {
		return descriptionOutput, err
	}
	descriptionOutput.WriteString(description)
	descriptionOutput.WriteString("\n\n")

	_, err = fmt.Fprint(descriptionOutput, internal.DescriptionAutoDocEnd)
	if err != nil {
		return descriptionOutput, err
	}
	return descriptionOutput, nil
}

// renderActionOutputTableOutput renders the action input table

func renderActionInputTableOutput(inputs map[string]ActionInput, inputColumns []string, markdownLinks bool, maxWidth int, maxWords int) (*strings.Builder, error) {
	inputTableOutput := &strings.Builder{}

	_, err := fmt.Fprintln(inputTableOutput, internal.InputAutoDocStart)
	if err != nil {
		return inputTableOutput, err
	}

	if len(inputs) > 0 {
		inputTable := tablewriter.NewWriter(inputTableOutput)
		inputTable.SetHeader(inputColumns)
		inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		inputTable.SetCenterSeparator(internal.PipeSeparator)
		inputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(inputs))
		for k := range inputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		inputTable.SetColWidth(maxWidth)

		for _, key := range keys {
			var row []string
			inputKey := key

			if markdownLinks {
				inputKey = utils.MarkdownLink(inputKey, "input")
			}

			for _, col := range inputColumns {
				switch col {
				case "Input":
					if inputs[key].DeprecationMessage != "" {
						row = append(row, fmt.Sprintf("~~%s~~ <br> %s", inputKey, inputs[key].DeprecationMessage))
					} else {
						row = append(row, inputKey)
					}
				case "Type":
					row = append(row, "string")
				case "Required":
					row = append(row, strconv.FormatBool(inputs[key].Required))
				case "Default":
					row = append(row, utils.FormatValue(inputs[key].Default, true, "<br>"))
				case "Description":
					if inputs[key].DeprecationMessage != "" {
						row = append(row, utils.WordWrap(fmt.Sprintf("**Deprecated:** %s", inputs[key].Description), maxWords, "<br>"))
					} else {
						row = append(row, utils.WordWrap(inputs[key].Description, maxWords, "<br>"))
					}
				default:
					return inputTableOutput, fmt.Errorf(
						"unknown inputs column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(internal.DefaultActionInputColumns, ", "),
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

	} else {
		_, err := fmt.Fprintln(inputTableOutput, internal.NoInputsMessage)
		if err != nil {
			return inputTableOutput, err
		}
	}
	_, err = fmt.Fprint(inputTableOutput, internal.InputAutoDocEnd)
	if err != nil {
		return inputTableOutput, err
	}
	return inputTableOutput, nil
}

// renderActionOutputTableOutput renders the action output table
func renderActionOutputTableOutput(outputs map[string]ActionOutput, outputColumns []string, markdownLinks bool, maxWidth int, maxWords int) (*strings.Builder, error) {
	outputTableOutput := &strings.Builder{}

	_, err := fmt.Fprintln(outputTableOutput, internal.OutputAutoDocStart)
	if err != nil {
		return outputTableOutput, err
	}

	if len(outputs) > 0 {
		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader(outputColumns)
		outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		outputTable.SetCenterSeparator(internal.PipeSeparator)
		outputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(outputs))
		for k := range outputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		outputTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string
			outputKey := key

			if markdownLinks {
				outputKey = utils.MarkdownLink(outputKey, "output")
			}

			for _, col := range outputColumns {
				switch col {
				case "Output":
					row = append(row, outputKey)
				case "Type":
					row = append(row, "string")
				case "Description":
					row = append(row, utils.WordWrap(outputs[key].Description, maxWords, "<br>"))
				default:
					return outputTableOutput, fmt.Errorf(
						"unknown outputs column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(internal.DefaultActionOutputColumns, ", "),
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
	} else {
		_, err = fmt.Fprintln(outputTableOutput, internal.NoOutputsMessage)
		if err != nil {
			return outputTableOutput, err
		}
	}

	_, err = fmt.Fprint(outputTableOutput, internal.OutputAutoDocEnd)
	if err != nil {
		return outputTableOutput, err
	}

	return outputTableOutput, nil
}
