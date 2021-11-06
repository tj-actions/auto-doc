// +build !windows

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
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func (a *Action) renderToOutput() {
	var err error
	inputTableOutput := &strings.Builder{}

	if len(a.Inputs) > 0 {
		_, err = fmt.Fprintln(inputTableOutput, inputAutoDocStart)
		if err != nil {
			cobra.CheckErr(err)
		}

		inputTable := tablewriter.NewWriter(inputTableOutput)
		inputTable.SetHeader([]string{"Input", "Required", "Default", "Description"})
		inputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		inputTable.SetCenterSeparator("|")

		keys := make([]string, 0, len(a.Inputs))
		for k := range a.Inputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			row := []string{key, strconv.FormatBool(a.Inputs[key].Required), a.Inputs[key].Default, a.Inputs[key].Description}
			inputTable.Append(row)
		}

		_, err = fmt.Fprintln(inputTableOutput)
		if err != nil {
			cobra.CheckErr(err)
		}

		inputTable.Render()

		_, err = fmt.Fprintln(inputTableOutput)
		if err != nil {
			cobra.CheckErr(err)
		}

		_, err = fmt.Fprint(inputTableOutput, inputAutoDocEnd)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	outputTableOutput := &strings.Builder{}

	if len(a.Outputs) > 0 {
		_, err = fmt.Fprintln(outputTableOutput, outputAutoDocStart)
		if err != nil {
			cobra.CheckErr(err)
		}

		outputTable := tablewriter.NewWriter(outputTableOutput)
		outputTable.SetHeader([]string{"Output", "Description", "Value"})
		outputTable.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		outputTable.SetCenterSeparator("|")

		keys := make([]string, 0, len(a.Outputs))
		for k := range a.Outputs {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, key := range keys {
			row := []string{key, a.Outputs[key].Description, a.Outputs[key].Value}
			outputTable.Append(row)
		}

		_, err = fmt.Fprintln(outputTableOutput)
		if err != nil {
			cobra.CheckErr(err)
		}

		outputTable.Render()

		_, err = fmt.Fprintln(outputTableOutput)
		if err != nil {
			cobra.CheckErr(err)
		}

		_, err = fmt.Fprint(outputTableOutput, outputAutoDocEnd)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	input, err := ioutil.ReadFile(outputFileName)

	if err != nil {
		cobra.CheckErr(err)
	}

	var output = []byte("")

	hasInputsData, inputStartIndex, inputEndIndex := HasBytesInBetween(
		input,
		[]byte(InputsHeader),
		[]byte(inputAutoDocEnd),
	)

	if hasInputsData {
		inputsStr := fmt.Sprintf("%s\n\n%v", InputsHeader, inputTableOutput.String())
		output = ReplaceBytesInBetween(input, inputStartIndex, inputEndIndex, []byte(inputsStr))
	} else {
		inputsStr := fmt.Sprintf("%s\n\n%v", InputsHeader, inputTableOutput.String())
		output = bytes.Replace(input, []byte(InputsHeader), []byte(inputsStr), -1)
	}

	hasOutputsData, outputStartIndex, outputEndIndex := HasBytesInBetween(
		output,
		[]byte(OutputsHeader),
		[]byte(outputAutoDocEnd),
	)

	if hasOutputsData {
		outputsStr := fmt.Sprintf("%s\n\n%v", OutputsHeader, outputTableOutput.String())
		output = ReplaceBytesInBetween(output, outputStartIndex, outputEndIndex, []byte(outputsStr))
	} else {
		outputsStr := fmt.Sprintf("%s\n\n%v", OutputsHeader, outputTableOutput.String())
		output = bytes.Replace(output, []byte(OutputsHeader), []byte(outputsStr), -1)
	}

	if len(output) > 0 {
		if err = ioutil.WriteFile(outputFileName, output, 0666); err != nil {
			cobra.CheckErr(err)
		}
	}
}
