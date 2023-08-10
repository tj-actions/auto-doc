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
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spf13/cobra"
)

func Test_rootCommand(t *testing.T) {
	t.Run("Missing filename flag", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		o := bytes.NewBufferString("")
		b := bytes.NewBufferString("")
		cmd.SetOut(o)
		cmd.SetErr(b)
		mdFile := filepath.Join("..", "test", "README.md")
		cmd.SetArgs([]string{"--filename", "", "--output", mdFile})
		err := cmd.Execute()

		if err == nil {
			t.Fatal("expected error got nil")
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Error: filename must be specified with --filename")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Passing positional arguments", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		o := bytes.NewBufferString("")
		b := bytes.NewBufferString("")
		cmd.SetOut(o)
		cmd.SetErr(b)
		mdFile := filepath.Join("..", "test", "README.md")
		cmd.SetArgs([]string{mdFile})
		err := cmd.Execute()

		if err == nil {
			t.Fatal("expected error got nil")
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Error: requires no positional arguments: 1 given")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Non existent action file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		o := bytes.NewBufferString("")
		b := bytes.NewBufferString("")
		cmd.SetOut(o)
		cmd.SetErr(b)
		mdFile := filepath.Join("..", "test", "README.md")
		cmd.SetArgs([]string{"--filename", "../test/invalid.yml", "--output", mdFile})
		err := cmd.Execute()

		if err == nil {
			t.Fatal("expected error got nil")
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		if runtime.GOOS == "windows" {
			exp := fmt.Sprintln("Error: open ../test/invalid.yml: The system cannot find the file specified.")
			if string(out) != exp {
				t.Fatalf(
					"expected \"%s\" got \"%s\"",
					exp,
					string(out),
				)
			}
		} else {
			exp := fmt.Sprintln("Error: open ../test/invalid.yml: no such file or directory")
			if string(out) != exp {
				t.Fatalf(
					"expected \"%s\" got \"%s\"",
					exp,
					string(out),
				)
			}
		}
	})

	t.Run("Non existent reusable workflow file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		o := bytes.NewBufferString("")
		b := bytes.NewBufferString("")
		cmd.SetOut(o)
		cmd.SetErr(b)
		mdFile := filepath.Join("..", "test", "README-reusable.md")
		cmd.SetArgs([]string{"--filename", "../test/reusable-invalid.yml", "--reusable", "--output", mdFile})
		err := cmd.Execute()

		if err == nil {
			t.Fatal("expected error got nil")
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		if runtime.GOOS == "windows" {
			exp := fmt.Sprintln("Error: open ../test/reusable-invalid.yml: The system cannot find the file specified.")
			if string(out) != exp {
				t.Fatalf(
					"expected \"%s\" got \"%s\"",
					exp,
					string(out),
				)
			}
		} else {
			exp := fmt.Sprintln("Error: open ../test/reusable-invalid.yml: no such file or directory")
			if string(out) != exp {
				t.Fatalf(
					"expected \"%s\" got \"%s\"",
					exp,
					string(out),
				)
			}
		}
	})

	t.Run("Update test/README.md using custom action file and output file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action.yml")
		mdFile := filepath.Join("..", "test", "README.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--output", mdFile})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-outputColumns.md using custom action file and output file and custom outputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action.yml")
		mdFile := filepath.Join("..", "test", "README-outputColumns.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--output", mdFile, "--outputColumns", "Output", "--outputColumns", "Type"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-inputColumns.md using custom action file and output file and custom inputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action.yml")
		mdFile := filepath.Join("..", "test", "README-inputColumns.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--output", mdFile, "--inputColumns", "Input", "--inputColumns", "Type", "--inputColumns", "Description"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-reusable.md using custom action file and output file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "reusable-action.yml")
		mdFile := filepath.Join("..", "test", "README-reusable.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-outputColumns.md using custom action file and output file and custom reusableOutputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "reusable-action.yml")
		mdFile := filepath.Join("..", "test", "README-reusable-outputColumns.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile, "--reusableOutputColumns", "Output", "--reusableOutputColumns", "Value"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-inputColumns.md using custom action file and output file and custom reusableInputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "reusable-action.yml")
		mdFile := filepath.Join("..", "test", "README-reusable-inputColumns.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile, "--reusableInputColumns", "Input", "--reusableInputColumns", "Type", "--reusableInputColumns", "Description"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-secretColumns.md using custom action file and output file and custom reusableSecretColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "reusable-action.yml")
		mdFile := filepath.Join("..", "test", "README-reusable-secretColumns.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile, "--reusableSecretColumns", "Secret", "--reusableSecretColumns", "Description"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})

	t.Run("Update test/README-markdownLinks.md using custom action file and output file and markdownLinks flag", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action.yml")
		mdFile := filepath.Join("..", "test", "README-markdownLinks.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--output", mdFile, "--markdownLinks"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})
	t.Run("Update test/README-reusable.md using custom action file and output file and markdownLinks flag", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "reusable-action.yml")
		mdFile := filepath.Join("..", "test", "README-reusable-markdownLinks.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile, "-m"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})
	t.Run("Update test/README-markdownLinks.md using custom action file and output file and markdownLinks flag", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action.yml")
		mdFile := filepath.Join("..", "test", "README-markdownLinks.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--output", mdFile, "--markdownLinks"})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})
	t.Run("Update test/README-action-empty-markers.md with action without inputs and outputs", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		inputFile := filepath.Join("..", "test", "action-no-inputs-no-outputs.yml")
		mdFile := filepath.Join("..", "test", "README-action-empty-markers.md")
		cmd.SetArgs([]string{"--filename", inputFile, "--reusable", "--output", mdFile})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})
	t.Run("Update test/README-workflow-empty-markers.md with workflow without inputs and outputs", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		workflowFile := filepath.Join("..", "test", "reusable-workflow-no-inputs-no-outputs.yml")
		mdFile := filepath.Join("..", "test", "README-workflow-empty-markers.md")
		cmd.SetArgs([]string{"--filename", workflowFile, "--reusable", "--output", mdFile})
		err := cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		exp := fmt.Sprintln("Successfully generated documentation")

		if string(out) != exp {
			t.Fatalf(
				"expected \"%s\" got \"%s\"",
				exp,
				string(out),
			)
		}
	})
}
