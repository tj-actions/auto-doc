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
	"regexp"
	"strings"
)

func WordWrap(s string, limit int) string {
	if strings.TrimSpace(s) == "" {
		return s
	}
	// compile regular expressions for Markdown links and code blocks and code
	linkRegex := regexp.MustCompile(`\[.*]\(.*\)`)
	codeBlockRegex := regexp.MustCompile(`\` + "```" + `.*` + "```" + `\s*`)

	// convert string to slice
	strSlice := strings.Fields(s)
	currentLimit := limit

	var result string

	for len(strSlice) >= 1 {
		// convert slice/array back to string
		// but insert <br> at specified limit
		// unless the current slice contains a Markdown link or code block or code
		hasMore := len(strSlice) > currentLimit

		if hasMore && len(result) > 0 {
			result += " "
		}

		if len(strSlice) < currentLimit {
			currentLimit = len(strSlice)
			result = result + strings.Join(strSlice[:currentLimit], " ")
		} else if currentLimit == limit && !linkRegex.MatchString(strings.Join(strSlice[:currentLimit], " ")) && !codeBlockRegex.MatchString(strings.Join(strSlice[:currentLimit], " ")) {
			result = result + strings.Join(strSlice[:currentLimit], " ") + "<br>"
		} else {
			result = result + strings.Join(strSlice[:currentLimit], " ")
		}

		// discard the elements that were copied over to result
		strSlice = strSlice[currentLimit:]

		// change the limit
		// to cater for the last few words in the line
		if len(strSlice) < currentLimit {
			currentLimit = len(strSlice)
		}
	}

	// Remove trailing <br> if any
	result = strings.TrimSuffix(result, "<br>")

	return strings.TrimSpace(result)
}