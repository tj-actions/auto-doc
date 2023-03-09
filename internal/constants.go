package internal

import "fmt"

const InputsHeader = "## Inputs"
const OutputsHeader = "## Outputs"
const SecretsHeader = "## Secrets"
const AutoDocStart = "<!-- AUTO-DOC-%s:START - Do not remove or modify this section -->"
const AutoDocEnd = "<!-- AUTO-DOC-%s:END -->"
const PipeSeparator = "|"
const NewLineSeparator = "\n"
var InputAutoDocStart = fmt.Sprintf(AutoDocStart, "INPUT")
var InputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "INPUT")
var OutputAutoDocStart = fmt.Sprintf(AutoDocStart, "OUTPUT")
var OutputAutoDocEnd = fmt.Sprintf(AutoDocEnd, "OUTPUT")
var SecretsAutoDocStart = fmt.Sprintf(AutoDocStart, "SECRETS")
var SecretsAutoDocEnd = fmt.Sprintf(AutoDocEnd, "SECRETS")

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