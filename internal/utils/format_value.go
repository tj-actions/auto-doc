package utils

import (
	"fmt"
	"github.com/tj-actions/auto-doc/internal"
	"strings"
)

func FormatValue(v string) string {
	if len(v) == 0 {
		return ""
	}

	var inputDefault = v
	var defaultValue string
	var parts = strings.Split(inputDefault, "\n")

	if len(parts) > 1 && inputDefault != internal.NewLineSeparator {
		for _, part := range parts {
			if part != "" {
				defaultValue += "`\"" + part + "\"`" + "<br>"
			}
		}
	} else {
		if strings.Contains(inputDefault, internal.PipeSeparator) {
			inputDefault = strings.Replace(inputDefault, internal.PipeSeparator, "\"\\"+internal.PipeSeparator+"\"", -1)
		} else {
			inputDefault = fmt.Sprintf("%#v", inputDefault)
		}
		defaultValue = "`" + inputDefault + "`"
	}

	return defaultValue
}