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
)

// HasBytesInBetween checks if a byte array has a start and end byte array and returns true if and all occurrences of start and end
func HasBytesInBetween(value, start, end []byte) (found bool, startIndexes []int, endIndexes []int) {
	startRegexp := regexp.MustCompile("(?m)^" + string(start))
	endRegexp := regexp.MustCompile("(?m)^" + string(end))

	// Find all start and end indexes
	for i := 0; i < len(value); i++ {
		startLoc := startRegexp.FindIndex(value[i:])
		endLoc := endRegexp.FindIndex(value[i:])
		if len(startLoc) > 0 && len(endLoc) > 0 {
			startIndex := startLoc[0] + i
			endIndex := endLoc[1] + i

			if startIndex < endIndex {
				startIndexes = append(startIndexes, startIndex)
				endIndexes = append(endIndexes, endIndex)
			}
			i += endIndex // skip the content between end and next start
		}
	}

	if len(startIndexes) == 0 || len(endIndexes) == 0 {
		return false, nil, nil
	}

	return true, startIndexes, endIndexes
}

// ReplaceBytesInBetween replaces a byte array between a start and end index with a new byte array
func ReplaceBytesInBetween(value []byte, startIndex int, endIndex int, new []byte) []byte {
	t := make([]byte, len(value)+len(new))
	w := 0

	w += copy(t[:startIndex], value[:startIndex])
	w += copy(t[w:w+len(new)], new)
	w += copy(t[w:], value[endIndex:])
	return t[0:w]
}
