package utils

// format string input as markdown link (named anchor)
func MarkdownLink(s string, t string) string {
	var link string
	var typeName string = t + "_" + s

	link = "<a name=\"" + typeName + "\"></a>[" + s + "](#" + typeName + ")"

	return link
}
