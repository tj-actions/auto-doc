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
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/tj-actions/auto-doc/v3/internal"
	"github.com/tj-actions/auto-doc/v3/internal/utils"
)

// CodeBlockInput represents the input of the action.yml
type CodeBlockInput struct {
	Description        string `yaml:"description"`
	Required           bool   `yaml:"required"`
	Default            string `yaml:"default,omitempty"`
	DeprecationMessage string `yaml:"deprecationMessage,omitempty"`
}

// CodeBlock represents the action.yml outputted as a code block
type CodeBlock struct {
	Repository      string
	Token           string
	UseMajorVersion bool
	InputFileName   string
	OutputFileName  string
	InputColumns    []string
	OutputColumns   []string
	Inputs          map[string]CodeBlockInput `yaml:"inputs,omitempty"`
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

// WriteDocumentation write the table to the output file
func (c *CodeBlock) WriteDocumentation(codeBlock *strings.Builder) error {
	var err error
	input, err := os.ReadFile(c.OutputFileName)
	// coverage:ignore
	if err != nil {
		return err
	}

	var output []byte

	hasInputsData, indices := utils.HasBytesInBetween(
		input,
		[]byte(internal.AutoDocCodeBlockStart),
		[]byte(internal.AutoDocCodeBlockEnd),
	)

	output = input
	codeBlockStr := strings.TrimSpace(codeBlock.String())

	if hasInputsData {
		output = utils.ReplaceBytesInBetween(output, indices, []byte(codeBlockStr))
	} else {
		re := regexp.MustCompile(fmt.Sprintf("(?m)^%s", internal.InputsHeader))
		output = re.ReplaceAllFunc(input, func(match []byte) []byte {
			if bytes.HasPrefix(match, []byte(internal.InputsHeader)) {
				if codeBlockStr != "" {
					return []byte(fmt.Sprintf("%s\n\n%v", internal.InputsHeader, codeBlockStr))
				} else {
					return []byte(internal.InputsHeader)
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

func (c *CodeBlock) GetLatestTagForRepository() (string, error) {
	fmt.Println("Downloading the latest release")

	tag, err := utils.GetLatestRepositoryTag(c.Repository, c.Token, c.UseMajorVersion)
	// coverage:ignore
	if err != nil {
		return "", err
	}

	return tag, nil
}

// RenderOutput renders the output and writes it to the given output
func (c *CodeBlock) RenderOutput() error {
	// Output this as a code block
	codeBlock := &strings.Builder{}

	_, err := fmt.Fprintln(codeBlock, internal.AutoDocCodeBlockStart)
	// coverage:ignore
	if err != nil {
		return err
	}

	keys := make([]string, 0, len(c.Inputs))
	for k := range c.Inputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	if len(keys) > 0 {
		tag, err := c.GetLatestTagForRepository()
		// coverage:ignore
		if err != nil {
			return err
		}

		codeBlock.WriteString("```yaml\n")
		codeBlock.WriteString(fmt.Sprintf("- uses: %s@%s\n", c.Repository, tag))
		codeBlock.WriteString("  with:\n")

		for _, key := range keys {
			inputKey := key
			codeBlock.WriteString(fmt.Sprintf("    # %s\n", utils.WordWrap(c.Inputs[key].Description, 9, "\n    # ")))
			if c.Inputs[key].Default != "" {
				codeBlock.WriteString(fmt.Sprintf("    # Default: %s\n", c.Inputs[key].Default))
			}
			if c.Inputs[key].DeprecationMessage != "" {
				codeBlock.WriteString(fmt.Sprintf("    # Deprecated: %s\n", c.Inputs[key].DeprecationMessage))
			}
			codeBlock.WriteString(fmt.Sprintf("    %s: ''\n", inputKey))
			codeBlock.WriteString("\n")
		}
		codeBlock.WriteString("```\n")
	}

	_, err = fmt.Fprintln(codeBlock, internal.AutoDocCodeBlockEnd)
	// coverage:ignore
	if err != nil {
		return err
	}

	err = c.WriteDocumentation(codeBlock)
	// coverage:ignore
	if err != nil {
		return err
	}

	return nil
}
