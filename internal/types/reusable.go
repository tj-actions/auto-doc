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

	"github.com/tj-actions/auto-doc/v2/internal"
	"github.com/tj-actions/auto-doc/v2/internal/utils"
)

// ReusableInput represents the input of the reusable workflow
type ReusableInput struct {
	Description        string `yaml:"description"`
	Required           bool   `yaml:"required"`
	Default            string `yaml:"default,omitempty"`
	Type               string `yaml:"type"`
	DeprecationMessage string `yaml:"deprecationMessage,omitempty"`
}

// ReusableOutput represents the output of the reusable workflow
type ReusableOutput struct {
	Description string `yaml:"description"`
	Value       string `yaml:"value"`
}

// ReusableSecret represents the secret of reusable workflows
type ReusableSecret struct {
	Required    bool   `yaml:"required"`
	Description string `yaml:"description"`
}

// Reusable represents a reusable workflow
type Reusable struct {
	InputFileName  string
	OutputFileName string
	ColMaxWidth    string
	ColMaxWords    string
	InputColumns   []string
	OutputColumns  []string
	SecretColumns  []string
	On             struct {
		WorkflowCall struct {
			Inputs  map[string]ReusableInput  `yaml:"inputs,omitempty"`
			Secrets map[string]ReusableSecret `yaml:"secrets,omitempty"`
			Outputs map[string]ReusableOutput `yaml:"outputs,omitempty"`
		} `yaml:"workflow_call"`
	}
	InputMarkdownLinks bool
}

// GetData parses the source yaml file
func (r *Reusable) GetData() error {
	reusableYaml, err := os.ReadFile(r.InputFileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(reusableYaml, &r)

	return err
}

// WriteDocumentation write the table to the output file
func (r *Reusable) WriteDocumentation(inputTable, outputTable, secretsTable *strings.Builder) error {
	var err error
	input, err := os.ReadFile(r.OutputFileName)

	if err != nil {
		return err
	}

	var output []byte

	hasInputsData, indices := utils.HasBytesInBetween(
		input,
		[]byte(internal.InputAutoDocStart),
		[]byte(internal.InputAutoDocEnd),
	)

	output = input
	inputsStr := strings.TrimSpace(inputTable.String())

	if hasInputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(inputsStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.InputsHeader))
		output = re.ReplaceAllFunc(input, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.InputsHeader)) {
				return []byte(fmt.Sprintf("%s\n\n%v", internal.InputsHeader, inputsStr))
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
				return []byte(fmt.Sprintf("%s\n\n%v", internal.OutputsHeader, outputsStr))
			}
			return match
		})
	}

	hasSecretsData, indices := utils.HasBytesInBetween(
		output,
		[]byte(internal.SecretsAutoDocStart),
		[]byte(internal.SecretsAutoDocEnd),
	)

	secretsStr := strings.TrimSpace(secretsTable.String())

	if hasSecretsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(secretsStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.SecretsHeader))
		output = re.ReplaceAllFunc(output, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.SecretsHeader)) {
				return []byte(fmt.Sprintf("%s\n\n%v", internal.SecretsHeader, secretsStr))
			}
			return match
		})
	}

	if err = os.WriteFile(r.OutputFileName, output, 0666); err != nil {
		cobra.CheckErr(err)
	}

	return nil
}

// RenderOutput renders the output and writes it to the given output
func (r *Reusable) RenderOutput() error {
	var err error
	maxWidth, err := strconv.Atoi(r.ColMaxWidth)
	if err != nil {
		return err
	}

	maxWords, err := strconv.Atoi(r.ColMaxWords)
	if err != nil {
		return err
	}
	inputTableOutput, err := renderReusableInputTableOutput(r.On.WorkflowCall.Inputs, r.InputColumns, r.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	secretTableOutput, err := renderReusableSecretTableOutput(r.On.WorkflowCall.Secrets, r.SecretColumns, r.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	outputTableOutput, err := renderReusableOutputTableOutput(r.On.WorkflowCall.Outputs, r.OutputColumns, r.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	err = r.WriteDocumentation(inputTableOutput, outputTableOutput, secretTableOutput)
	if err != nil {
		return err
	}

	return nil
}

// renderReusableInputTableOutput renders the reusable workflow input table
func renderReusableInputTableOutput(inputs map[string]ReusableInput, inputColumns []string, markdownLinks bool, maxWidth int, maxWords int) (*strings.Builder, error) {
	inputTableOutput := &strings.Builder{}

	if len(inputs) > 0 {
		_, err := fmt.Fprintln(inputTableOutput, internal.InputAutoDocStart)
		if err != nil {
			return inputTableOutput, err
		}

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

			for _, col := range inputColumns {
				switch col {
				case "Input":
					var modifiedKey = key
					if markdownLinks {
						modifiedKey = utils.MarkdownLink(key, "input")
					}
					if inputs[key].DeprecationMessage != "" {
						row = append(row, fmt.Sprintf("~~%s~~ <br> %s", modifiedKey, inputs[key].DeprecationMessage))
					} else {
						row = append(row, modifiedKey)
					}
				case "Type":
					row = append(row, inputs[key].Type)
				case "Required":
					row = append(row, strconv.FormatBool(inputs[key].Required))
				case "Default":
					switch inputs[key].Type {
					case "string":
						row = append(row, utils.FormatValue(inputs[key].Default))
					default:
						row = append(row, "`"+inputs[key].Default+"`")
					}
				case "Description":
					if inputs[key].DeprecationMessage != "" {
						row = append(row, utils.WordWrap(fmt.Sprintf("**Deprecated:** %s", inputs[key].Description), maxWords))
					} else {
						row = append(row, utils.WordWrap(inputs[key].Description, maxWords))
					}
				default:
					return inputTableOutput, fmt.Errorf(
						"unknown inputs column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(internal.DefaultReusableInputColumns, ", "),
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

		_, err = fmt.Fprint(inputTableOutput, internal.InputAutoDocEnd)
		if err != nil {
			return inputTableOutput, err
		}
	}
	return inputTableOutput, nil
}

// renderReusableOutputTableOutput renders the reusable workflow output table
func renderReusableOutputTableOutput(outputs map[string]ReusableOutput, reusableOutputColumns []string, markdownLinks bool, maxWidth int, maxWords int) (*strings.Builder, error) {
	outputTableOutput := &strings.Builder{}

	if len(outputs) > 0 {
		_, err := fmt.Fprintln(outputTableOutput, internal.OutputAutoDocStart)
		if err != nil {
			return outputTableOutput, err
		}

		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader(reusableOutputColumns)
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

			for _, col := range reusableOutputColumns {
				switch col {
				case "Output":
					if markdownLinks {
						row = append(row, utils.MarkdownLink(key, "output"))
					} else {
						row = append(row, key)
					}
				case "Value":
					row = append(row, utils.FormatValue(outputs[key].Value))
				case "Description":
					row = append(row, utils.WordWrap(outputs[key].Description, maxWords))
				default:
					return outputTableOutput, fmt.Errorf(
						"unknown outputs column: '%s'. Please specify any of the following columns: %s",
						col,
						strings.Join(internal.DefaultReusableOutputColumns, ", "),
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

		_, err = fmt.Fprint(outputTableOutput, internal.OutputAutoDocEnd)
		if err != nil {
			return outputTableOutput, err
		}
	}
	return outputTableOutput, nil
}

// renderReusableSecretTableOutput renders the reusable workflow secret table
func renderReusableSecretTableOutput(secrets map[string]ReusableSecret, secretColumns []string, markdownLinks bool, maxWidth int, maxWords int) (*strings.Builder, error) {
	secretTableOutput := &strings.Builder{}

	if len(secrets) > 0 {
		_, err := fmt.Fprintln(secretTableOutput, internal.SecretsAutoDocStart)
		if err != nil {
			return secretTableOutput, err
		}

		secretTable := tablewriter.NewWriter(secretTableOutput)
		secretTable.SetHeader(secretColumns)
		secretTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		secretTable.SetCenterSeparator(internal.PipeSeparator)
		secretTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(secrets))
		for k := range secrets {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		secretTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string

			for _, col := range secretColumns {
				switch col {
				case "Secret":
					if markdownLinks {
						row = append(row, utils.MarkdownLink(key, "secret"))
					} else {
						row = append(row, key)
					}
				case "Required":
					row = append(row, fmt.Sprintf("%v", secrets[key].Required))
				case "Description":
					row = append(row, utils.WordWrap(secrets[key].Description, maxWords))
				default:
					return secretTableOutput, fmt.Errorf(
						"unknown secrets column: '%secrets'. Please specify any of the following columns: %secrets",
						col,
						strings.Join(internal.DefaultReusableSecretColumns, ", "),
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

		_, err = fmt.Fprint(secretTableOutput, internal.SecretsAutoDocEnd)
		if err != nil {
			return secretTableOutput, err
		}
	}
	return secretTableOutput, nil
}
