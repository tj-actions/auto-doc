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

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/tj-actions/auto-doc/v3/internal"
	"github.com/tj-actions/auto-doc/v3/internal/utils"
)

// CodeBlock represents the action.yml outputted as a code block
type CodeBlock struct {
	Repository         string
	Token              string
	UseMajorVersion    bool
	OutputColumns      []string
	InputMarkdownLinks bool
	ColMaxWidth        string
	ColMaxWords        string
	InputFileName      string
	OutputFileName     string
	Inputs             map[string]ActionInput  `yaml:"inputs,omitempty"`
	Outputs            map[string]ActionOutput `yaml:"outputs,omitempty"`
}

// GetData parses the source yaml file
func (c *CodeBlock) GetData() error {
	actionYaml, err := os.ReadFile(c.InputFileName)
	// coverage:ignore
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(actionYaml, &c)
	return err
}

// writeDocumentation write the table to the output file
func (c *CodeBlock) writeDocumentation(inputCodeBlock, outputCodeBlock *strings.Builder) error {
	var err error
	input, err := os.ReadFile(c.OutputFileName)
	// coverage:ignore
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
	inputCodeBlockStr := strings.TrimSpace(inputCodeBlock.String())

	if hasInputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(inputCodeBlockStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.InputsHeader))
		output = re.ReplaceAllFunc(input, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.InputsHeader)) {
				if inputCodeBlockStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.InputsHeader, inputCodeBlockStr))
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

	outputCodeBlockStr := strings.TrimSpace(outputCodeBlock.String())

	if hasOutputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(outputCodeBlockStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.OutputsHeader))
		output = re.ReplaceAllFunc(output, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.OutputsHeader)) {
				if outputCodeBlockStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.OutputsHeader, outputCodeBlockStr))
				} else {
					return []byte(internal.OutputsHeader)
				}
			}
			return match
		})
	}

	if err = os.WriteFile(c.OutputFileName, output, 0o666); err != nil {
		cobra.CheckErr(err)
	}

	return nil
}

// getLatestTagForRepository returns the latest tag of a repository
func (c *CodeBlock) getLatestTagForRepository() (string, error) {
	fmt.Println("Downloading the latest release")

	tag, err := utils.GetLatestRepositoryTag(c.Repository, c.Token, c.UseMajorVersion)
	// coverage:ignore
	if err != nil {
		return "", err
	}

	return tag, nil
}

// renderCodeBlockActionInputs renders the inputs as a code block
func renderCodeBlockActionInputs(inputs map[string]ActionInput, repository, tag string) (*strings.Builder, error) {
	// Output this as a code block
	codeBlock := &strings.Builder{}

	_, err := fmt.Fprintln(codeBlock, internal.InputAutoDocStart)
	// coverage:ignore
	if err != nil {
		return codeBlock, err
	}

	keys := make([]string, 0, len(inputs))
	for k := range inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if len(keys) > 0 {
		codeBlock.WriteString("```yaml\n")
		codeBlock.WriteString(fmt.Sprintf("- uses: %s@%s\n", repository, tag))
		codeBlock.WriteString(fmt.Sprintf("  id: %s\n", strings.Split(repository, "/")[1]))
		codeBlock.WriteString("  with:\n")

		for _, key := range keys {
			codeBlock.WriteString(fmt.Sprintf("    # %s\n", utils.WordWrap(inputs[key].Description, 9, "\n    # ")))
			if inputs[key].Default == "false" || inputs[key].Default == "true" {
				codeBlock.WriteString("    # Type: boolean\n")
			} else {
				codeBlock.WriteString("    # Type: string\n")
			}
			if inputs[key].Default != "" {
				codeBlock.WriteString(fmt.Sprintf("    # Default: %s\n", utils.FormatValue(inputs[key].Default, false, "\n    #          ")))
			}
			if inputs[key].DeprecationMessage != "" {
				codeBlock.WriteString(fmt.Sprintf("    # Deprecated: %s\n", inputs[key].DeprecationMessage))
			}
			codeBlock.WriteString(fmt.Sprintf("    %s: ''\n", key))
			codeBlock.WriteString("\n")
		}
		codeBlock.WriteString("```\n")
	} else {
		_, err := fmt.Fprintln(codeBlock, internal.NoInputsMessage)
		if err != nil {
			return codeBlock, err
		}
	}

	_, err = fmt.Fprintln(codeBlock, internal.InputAutoDocEnd)
	// coverage:ignore
	if err != nil {
		return codeBlock, err
	}

	return codeBlock, nil
}

// RenderOutput renders the output and writes it to the given output
func (c *CodeBlock) RenderOutput() error {
	maxWidth, err := strconv.Atoi(c.ColMaxWidth)
	if err != nil {
		return err
	}

	maxWords, err := strconv.Atoi(c.ColMaxWords)
	if err != nil {
		return err
	}

	tag, err := c.getLatestTagForRepository()
	// coverage:ignore
	if err != nil {
		return err
	}

	inputCodeBlockOutput, err := renderCodeBlockActionInputs(c.Inputs, c.Repository, tag)

	// coverage:ignore
	if err != nil {
		return err
	}

	outputTableOutput, err := renderActionOutputTableOutput(c.Outputs, c.OutputColumns, c.InputMarkdownLinks, maxWidth, maxWords)
	if err != nil {
		return err
	}

	err = c.writeDocumentation(inputCodeBlockOutput, outputTableOutput)
	// coverage:ignore
	if err != nil {
		return err
	}

	return nil
}
