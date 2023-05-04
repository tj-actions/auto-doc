package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Maximum depth of directories to output.<br>e.g `test/test1/test2` with max depth of `2` returns `test/test1`."

	codeBlockRegex := regexp.MustCompile("`[^`]*`")
	codeBlockIndices := codeBlockRegex.FindAllStringIndex(text, -1)

	fmt.Println("Code block indices:", codeBlockIndices)
}
