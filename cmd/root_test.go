package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func Test_rootCommand(t *testing.T) {
	cmd := rootCmd
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
}
