//Package utils is a package that contains all the utility functions
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
package utils

import (
	"fmt"
	"strings"

	"github.com/tj-actions/auto-doc/v2/internal"
)

// FormatValue formats a string that would be outputed as markdown
func FormatValue(v string) string {
	if len(v) == 0 {
		return ""
	}

	var inputDefault = v
	var defaultValue string
	var parts = strings.Split(inputDefault, internal.NewLineSeparator)

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
