package utils

// format string input as markdown link (named anchor)
func MarkdownLink(s string) string {
	return "<a name=\"" + s + "\"></a>[" + s + "](#" + s + ")"
}
