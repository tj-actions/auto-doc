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
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

func Test_rootCommand(t *testing.T) {
	t.Run("Update test/README.md using custom action file and output file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--filename", "../test/action.yml", "--output", "../test/README.md"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/action.yml", "--output", "../test/README-outputColumns.md", "--outputColumns", "Output", "--outputColumns", "Type"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/action.yml", "--output", "../test/README-inputColumns.md", "--inputColumns", "Input", "--inputColumns", "Type", "--inputColumns", "Description"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/reusable-action.yml", "--reusable", "--output", "../test/README-reusable.md"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/reusable-action.yml", "--reusable", "--output", "../test/README-reusable-outputColumns.md", "--reusableOutputColumns", "Output", "--reusableOutputColumns", "Value"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/reusable-action.yml", "--reusable", "--output", "../test/README-reusable-inputColumns.md", "--reusableInputColumns", "Input", "--reusableInputColumns", "Type", "--reusableInputColumns", "Description"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/reusable-action.yml", "--reusable", "--output", "../test/README-reusable-secretColumns.md", "--reusableSecretColumns", "Secret", "--reusableSecretColumns", "Description"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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

	t.Run("Update test/README-MarkdownLinks.md using custom action file and output file and markdownLinks flag", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--filename", "../test/action.yml", "--output", "../test/README-MarkdownLinks.md", "--markdownLinks"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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
		cmd.SetArgs([]string{"--filename", "../test/reusable-action.yml", "--reusable", "--output", "../test/README-reusable-MarkdownLinks.md", "-m"})
		err := cmd.Execute()

		if err != nil {
			t.Fatal(err)
		}

		out, err := ioutil.ReadAll(b)
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

func TestGetMarkdownLinksFlag(t *testing.T) {
	cmd := &cobra.Command{}
	cmd.Flags().Bool("markdownLinks", true, "Names of inputs, outputs and secrets as markdown links")
	expected := true

	actual, err := cmd.Flags().GetBool("markdownLinks")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}
