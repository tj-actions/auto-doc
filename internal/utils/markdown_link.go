package utils

// format string input as markdown link
func MarkdownLink(s string) string {
	return "<a name=\"" + s + "\"></a>[" + s + "](#" + s + ")"
}
