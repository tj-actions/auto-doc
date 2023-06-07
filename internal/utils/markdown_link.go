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

// MarkdownLink formats input, output and secret name to Markdown link (named anchor) and return the link
func MarkdownLink(s string, t string) string {
	var link string
	var typeName = t + "_" + s

	link = "<a name=\"" + typeName + "\"></a>[" + s + "](#" + typeName + ")"

	return link
}
