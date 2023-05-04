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
	"regexp"
)

// HasBytesInBetween checks if a byte array has a start and end byte array and returns true if and all occurrences of start and end
func HasBytesInBetween(value, start, end []byte) (found bool, Indices [][]int) {
	// Multiline regex
	findRegex := regexp.MustCompile(fmt.Sprintf(`(?s)%s(.*?)%s`, regexp.QuoteMeta(string(start)), regexp.QuoteMeta(string(end))))
	Indices = findRegex.FindAllIndex(value, -1)
	
	if len(Indices) == 0 {
		return false, Indices
	}

	return true, Indices
}

// ReplaceBytesInBetween replaces a byte array between an array of start and end Indices with a new byte array
func ReplaceBytesInBetween(value []byte, indices [][]int, new []byte) []byte {
	t := make([]byte, 0, len(value)+len(new)*len(indices))
	prevIndex := 0

	for _, v := range indices {
		t = append(t, value[prevIndex:v[0]]...)
		t = append(t, new...)
		prevIndex = v[1]
	}

	t = append(t, value[prevIndex:]...)
	return t
}
