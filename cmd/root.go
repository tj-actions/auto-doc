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
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var actionFileName string
var outputFileName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auto-doc",
	Short: "Auto doc generator for your github action",
	Long: `Auto generate documentation for your github action.`,
	Run: func(cmd *cobra.Command, args []string) {
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

	    fmt.Printf("Working on %s \n", outputFileName)
		fmt.Printf("Action file: %s \n", actionFileName)
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
