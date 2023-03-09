package types

import (
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tj-actions/auto-doc/internal"
	"github.com/tj-actions/auto-doc/internal/utils"
	"gopkg.in/yaml.v3"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ReusableInput represents the input of the reusable workflow
type ReusableInput struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default,omitempty"`
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

// Reusable represents the reusable workflow yaml
type Reusable struct {
	InputFileName string
	OutputFileName string
	ColMaxWidth string
	ColMaxWords string
	InputColumns []string
	OutputColumns []string
	SecretColumns []string
	On struct {
		WorkflowCall struct {
			Inputs  map[string]ReusableInput          `yaml:"inputs,omitempty"`
			Secrets map[string]ReusableSecret         `yaml:"secrets,omitempty"`
			Outputs map[string]ReusableOutput `yaml:"outputs,omitempty"`
		} `yaml:"workflow_call"`
	}
}

func (r *Reusable) GetData() error {
	reusableYaml, err := os.ReadFile(r.InputFileName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(reusableYaml, &r)

	return err
}

func (r *Reusable) WriteDocumentation(inputTable, outputTable, secretsTable *strings.Builder) error {
	input, err := os.ReadFile(r.OutputFileName)

	if err != nil {
		return err
	}

	var output []byte

	hasInputsData, inputStartIndex, inputEndIndex := utils.HasBytesInBetween(
		input,
		[]byte(internal.InputsHeader),
		[]byte(internal.InputAutoDocEnd),
	)

	if hasInputsData {
		inputsStr := fmt.Sprintf("%s\n\n%v", internal.InputsHeader, inputTable.String())
		output = utils.ReplaceBytesInBetween(input, inputStartIndex, inputEndIndex, []byte(inputsStr))
	} else {
		inputsStr := fmt.Sprintf("%s\n\n%v", internal.InputsHeader, inputTable.String())
		output = bytes.Replace(input, []byte(internal.InputsHeader), []byte(inputsStr), -1)
	}

	hasOutputsData, outputStartIndex, outputEndIndex := utils.HasBytesInBetween(
		output,
		[]byte(internal.OutputsHeader),
		[]byte(internal.OutputAutoDocEnd),
	)

	if hasOutputsData {
		outputsStr := fmt.Sprintf("%s\n\n%v", internal.OutputsHeader, outputTable.String())
		output = utils.ReplaceBytesInBetween(output, outputStartIndex, outputEndIndex, []byte(outputsStr))
	} else {
		outputsStr := fmt.Sprintf("%s\n\n%v", internal.OutputsHeader, outputTable.String())
		output = bytes.Replace(output, []byte(internal.OutputsHeader), []byte(outputsStr), -1)
	}

	hasSecretsData, secretsStartIndex, secretsEndIndex := utils.HasBytesInBetween(
		output,
		[]byte(internal.SecretsHeader),
		[]byte(internal.SecretsAutoDocEnd),
	)

	if hasSecretsData {
		secretsStr := fmt.Sprintf("%s\n\n%v", internal.SecretsHeader, secretsTable.String())
		output = utils.ReplaceBytesInBetween(output, secretsStartIndex, secretsEndIndex, []byte(secretsStr))
	} else {
		secretsStr := fmt.Sprintf("%s\n\n%v", internal.SecretsHeader, secretsTable.String())
		output = bytes.Replace(output, []byte(internal.SecretsHeader), []byte(secretsStr), -1)
	}

	if len(output) > 0 {
		if err = os.WriteFile(r.OutputFileName, output, 0666); err != nil {
			cobra.CheckErr(err)
		}
	}

	return nil
}

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
	inputTableOutput, err := renderReusableInputTableOutput(r.On.WorkflowCall.Inputs, r.InputColumns, maxWidth, maxWords)
	if err != nil {
		return err
	}

	secretTableOutput, err := renderReusableSecretTableOutput(r.On.WorkflowCall.Secrets, r.SecretColumns, maxWidth, maxWords)
	if err != nil {
		return err
	}

	outputTableOutput, err := renderReusableOutputTableOutput(r.On.WorkflowCall.Outputs, r.OutputColumns, maxWidth, maxWords)
	if err != nil {
		return err
	}

	err = r.WriteDocumentation(inputTableOutput, outputTableOutput, secretTableOutput)
	if err != nil {
		return err
	}

	return nil
}

func renderReusableInputTableOutput(i map[string]ReusableInput, inputColumns[]string, maxWidth int, maxWords int) (*strings.Builder, error) {
	inputTableOutput := &strings.Builder{}

	if len(i) > 0 {
		_, err := fmt.Fprintln(inputTableOutput, internal.InputAutoDocStart)
		if err != nil {
			return inputTableOutput, err
		}

		inputTable := tablewriter.NewWriter(inputTableOutput)
		inputTable.SetHeader(inputColumns)
		inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		inputTable.SetCenterSeparator(internal.PipeSeparator)
		inputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(i))
		for k := range i {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		inputTable.SetColWidth(maxWidth)

		for _, key := range keys {
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
					row = append(row, utils.FormatValue(i[key].Default))
				case "Description":
					row = append(row, utils.WordWrap(i[key].Description, maxWords))
				default:
					return inputTableOutput, fmt.Errorf(
						"unknown input column: '%s'. Please specify any of the following columns: %s",
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

func renderReusableOutputTableOutput(o map[string]ReusableOutput, reusableOutputColumns[]string, maxWidth int, maxWords int) (*strings.Builder, error) {
	outputTableOutput := &strings.Builder{}

	if len(o) > 0 {
		_, err := fmt.Fprintln(outputTableOutput, internal.OutputAutoDocStart)
		if err != nil {
			return outputTableOutput, err
		}

		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader(reusableOutputColumns)
		outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		outputTable.SetCenterSeparator(internal.PipeSeparator)
		outputTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(o))
		for k := range o {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		outputTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string

			for _, col := range reusableOutputColumns {
				switch col {
				case "Output":
					row = append(row, key)
				case "Value":
					row = append(row, utils.FormatValue(o[key].Value))
				case "Description":
					row = append(row, utils.WordWrap(o[key].Description, maxWords))
				default:
					return outputTableOutput, fmt.Errorf(
						"unknown output column: '%s'. Please specify any of the following columns: %s",
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

func renderReusableSecretTableOutput(s map[string]ReusableSecret, secretColumns[]string, maxWidth int, maxWords int) (*strings.Builder, error) {
	secretTableOutput := &strings.Builder{}

	if len(s) > 0 {
		_, err := fmt.Fprintln(secretTableOutput, internal.SecretsAutoDocStart)
		if err != nil {
			return secretTableOutput, err
		}

		secretTable := tablewriter.NewWriter(secretTableOutput)
		secretTable.SetHeader(secretColumns)
		secretTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		secretTable.SetCenterSeparator(internal.PipeSeparator)
		secretTable.SetAlignment(tablewriter.ALIGN_CENTER)

		keys := make([]string, 0, len(s))
		for k := range s {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		secretTable.SetColWidth(maxWidth)
		for _, key := range keys {
			var row []string

			for _, col := range secretColumns {
				switch col {
				case "Secret":
					row = append(row, key)
				case "Required":
					row = append(row, fmt.Sprintf("%v", s[key].Required))
				case "Description":
					row = append(row, utils.WordWrap(s[key].Description, maxWords))
				default:
					return secretTableOutput, fmt.Errorf(
						"unknown secrets column: '%s'. Please specify any of the following columns: %s",
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