// Package internal contains all internal packages and utility functions.
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
package internal

import "fmt"

// DescriptionHeader represents the markdown header of description
const DescriptionHeader = "## Description"

// InputsHeader represents the markdown header of inputs
const InputsHeader = "## Inputs"

// OutputsHeader represents the markdown header of outputs
const OutputsHeader = "## Outputs"

// SecretsHeader represents the markdown header of secrets
const SecretsHeader = "## Secrets"

// AutoDocStart placeholder that represents the start of the input table
const AutoDocStart = "<!-- AUTO-DOC-%s:START - Do not remove or modify this section -->"

// AutoDocEnd placeholder that represents the end of the input table
const AutoDocEnd = "<!-- AUTO-DOC-%s:END -->"

// PipeSeparator represents the separator used for the distinguishing between columns
const PipeSeparator = "|"

// NewLineSeparator used for splitting lines
const NewLineSeparator = "\n"

// InputAutoDocStart is the start of the input
var DescriptionAutoDocStart = fmt.Sprintf(AutoDocStart, "DESCRIPTION")

// InputAutoDocEnd is the end of the input
var DescriptionAutoDocEnd = fmt.Sprintf(AutoDocEnd, "DESCRIPTION")

// InputAutoDocStart is the start of the input
var InputAutoDocStart = fmt.Sprintf(AutoDocStart, "INPUT")

// InputAutoDocEnd is the end of the input
var InputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "INPUT")

// NoInputsMessage is the message printed when there are no inputs
var NoInputsMessage = "No inputs."

// OutputAutoDocStart is the start of the output
var OutputAutoDocStart = fmt.Sprintf(AutoDocStart, "OUTPUT")

// OutputAutoDocEnd is the end of the output
var OutputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "OUTPUT")

// NoOutputsMessage is the message printed when there are no outputs
var NoOutputsMessage = "No outputs."

// SecretsAutoDocStart is the start of the secrets
var SecretsAutoDocStart = fmt.Sprintf(AutoDocStart, "SECRETS")

// SecretsAutoDocEnd is the end of the secrets
var SecretsAutoDocEnd = fmt.Sprintf(AutoDocEnd, "SECRETS")

// NoSecretsMessage is the message printed when there are no secrets
var NoSecretsMessage = "No secrets."

// action.yml

// DefaultActionInputColumns default values
var DefaultActionInputColumns = []string{"Input", "Type", "Required", "Default", "Description"}

// DefaultActionOutputColumns default values
var DefaultActionOutputColumns = []string{"Output", "Type", "Description"}

// Reusable workflows

// DefaultReusableSecretColumns default values
var DefaultReusableSecretColumns = []string{"Secret", "Required", "Description"}

// DefaultReusableOutputColumns default values
var DefaultReusableOutputColumns = []string{"Output", "Value", "Description"}

// DefaultReusableInputColumns default values
var DefaultReusableInputColumns = []string{"Input", "Type", "Required", "Default", "Description"}
