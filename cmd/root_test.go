package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"testing"
)

func Test_rootCommand(t *testing.T) {
	t.Run("Update test/README.md using custom action file and output file", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--action", "../test/action.yml", "--output", "../test/README.md"})
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

	t.Run("Update test/README.md using custom action file and output file and custom outputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--action", "../test/action.yml", "--output", "../test/README-outputColumns.md", "--outputColumns", "Output", "--outputColumns", "Type"})
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

	t.Run("Update test/README.md using custom action file and output file and custom inputColumns", func(t *testing.T) {
		cmd := &cobra.Command{Use: "auto-doc", RunE: RootCmdRunE}
		RootCmdFlags(cmd)
		b := bytes.NewBufferString("")
		cmd.SetOut(b)
		cmd.SetArgs([]string{"--action", "../test/action.yml", "--output", "../test/README-inputColumns.md", "--inputColumns", "Input", "--inputColumns", "Type", "--inputColumns", "Description"})
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
