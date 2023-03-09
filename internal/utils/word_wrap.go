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