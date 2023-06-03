// Package cmd provides a cli script that parses the GitHub action.yml and reusable workflow files and outputs a Markdown table to a specified path.
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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tj-actions/auto-doc/v2/internal"
	"github.com/tj-actions/auto-doc/v2/internal/types"
)

var fileName string
var outputFileName string
var colMaxWidth string
var colMaxWords string

// action.yml
var inputColumns []string
var outputColumns []string

// Reusable workflows
var reusableInputColumns []string
var reusableOutputColumns []string
var reusableSecretColumns []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-doc",
	Short: "Auto doc generator for your custom github action or reusable workflow",
	Long:  "Automatically generate documentation for your custom github action or reusable workflow",
	RunE:  RootCmdRunE,
}

// RootCmdRunE runs the root commands RunE function	and handles invalid options and prints the help message
// if the command is called with no arguments.
func RootCmdRunE(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("requires no positional arguments: %d given", len(args))
	}

	if fileName == "" {
		return fmt.Errorf("filename must be specified with --filename")
	}

	reusable, err := cmd.Flags().GetBool("reusable")
	if err != nil {
		return err
	}

	markdownLinks, err := cmd.Flags().GetBool("markdownLinks")
	if err != nil {
		return err
	}

	var documentation types.Documentation

	if reusable {
		documentation = &types.Reusable{
			InputFileName:      fileName,
			OutputFileName:     outputFileName,
			ColMaxWidth:        colMaxWidth,
			ColMaxWords:        colMaxWords,
			InputColumns:       reusableInputColumns,
			OutputColumns:      reusableOutputColumns,
			SecretColumns:      reusableSecretColumns,
			InputMarkdownLinks: markdownLinks,
		}
	} else {
		documentation = &types.Action{
			InputFileName:      fileName,
			OutputFileName:     outputFileName,
			ColMaxWidth:        colMaxWidth,
			ColMaxWords:        colMaxWords,
			InputColumns:       inputColumns,
			OutputColumns:      outputColumns,
			InputMarkdownLinks: markdownLinks,
		}
	}

	err = documentation.GetData()
	if err != nil {
		return err
	}

	err = documentation.RenderOutput()
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
	cmd.Flags().StringVarP(
		&fileName,
		"filename",
		"f",
		"",
		"config file",
	)
	cmd.Flags().BoolP(
		"reusable",
		"r",
		false,
		"A reusable workflow",
	)
	cmd.Flags().StringVarP(
		&outputFileName,
		"output",
		"o",
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
		internal.DefaultActionInputColumns,
		"list of input column names",
	)
	cmd.Flags().StringArrayVar(
		&outputColumns,
		"outputColumns",
		internal.DefaultActionOutputColumns,
		"list of output column names",
	)
	cmd.Flags().StringArrayVar(
		&reusableInputColumns,
		"reusableInputColumns",
		internal.DefaultReusableInputColumns,
		"list of reusable input column names",
	)
	cmd.Flags().StringArrayVar(
		&reusableOutputColumns,
		"reusableOutputColumns",
		internal.DefaultReusableOutputColumns,
		"list of reusable output column names",
	)
	cmd.Flags().StringArrayVar(
		&reusableSecretColumns,
		"reusableSecretColumns",
		internal.DefaultReusableSecretColumns,
		"list of reusable secrets column names",
	)
	cmd.Flags().BoolP(
		"markdownLinks",
		"m",
		false,
		"Names of inputs, outputs and secrets as markdown links",
	)
}

func init() {
	RootCmdFlags(rootCmd)
}
