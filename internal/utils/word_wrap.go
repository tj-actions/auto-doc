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
	"unicode"
)

// WordWrap wraps text at the specified number of columns
func WordWrap(s string, limit int) string {
	if strings.TrimSpace(s) == "" {
		return s
	}

	var (
		linkIndices      = getAllLinksIndex(s)
		codeBlockIndices = getAllCodeBlocksIndex(s)
		codeIndices      = getAllCodeIndex(s)
		parenthesisIndices         = getAllParenthesisIndex(s)
		italicIndices              = getAllItalicIndex(s)
		start                      = 0
	)

	// split the string into words that aren't between any of the links, code blocks, code and parenthesis
	strSlice := strings.FieldsFunc(s, func(r rune) bool {
		shouldExclude := isWithin(start, linkIndices) || isWithin(start, codeBlockIndices) || isWithin(start, codeIndices) || isWithin(start, parenthesisIndices) || isWithin(start, italicIndices)
		start++
		return !shouldExclude && unicode.IsSpace(r)
	})

	var result = ""

	for len(strSlice) >= 1 {
		// convert slice/array back to string
		// but insert \r\n at specified limit
		if len(strSlice) < limit {
			limit = len(strSlice)
		}

		result = result + strings.Join(strSlice[:limit], " ") + "<br>"

		// discard the elements that were copied over to result
		strSlice = strSlice[limit:]
	}

	// Trim the last <br> tag
	result = strings.TrimSuffix(result, "<br>")

	return result
}

func isWithin(index int, ranges [][]int) bool {
	for _, r := range ranges {
		if index >= r[0] && index < r[1] {
			return true
		}
	}

	return false
}

func getAllLinksIndex(s string) [][]int {
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	return linkRegex.FindAllStringIndex(s, -1)
}

func getAllCodeBlocksIndex(s string) [][]int {
	codeBlockRegex := regexp.MustCompile("```[^`]*```")

	return codeBlockRegex.FindAllStringIndex(s, -1)
}

func getAllCodeIndex(s string) [][]int {
	codeRegex := regexp.MustCompile("`[^`]*`")

	return codeRegex.FindAllStringIndex(s, -1)
}

func getAllParenthesisIndex(s string) [][]int {
	parenthesisRegex := regexp.MustCompile(`\((.*?)\)`)

	return parenthesisRegex.FindAllStringIndex(s, -1)
}

func getAllItalicIndex(s string) [][]int {
	italicRegex := regexp.MustCompile(`\*(.*?)\*`)

	return italicRegex.FindAllStringIndex(s, -1)
}
