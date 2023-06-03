[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge\&logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Mac OS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge\&logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge\&logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Public workflows that use this action.](https://img.shields.io/endpoint?style=for-the-badge\&url=https%3A%2F%2Fused-by.vercel.app%2Fapi%2Fgithub-actions%2Fused-by%3Faction%3Dtj-actions%2Fauto-doc%26badge%3Dtrue%26package_id%3DUGFja2FnZS0yOTU3MDU0NDI1)](https://github.com/search?o=desc\&q=tj-actions+auto-doc+language%3AYAML\&s=\&type=Code)

![Coverage](https://img.shields.io/badge/Coverage-80.6%25-brightgreen)
[![codecov](https://codecov.io/github/tj-actions/auto-doc/branch/main/graph/badge.svg?token=TNXW4QRRJD)](https://codecov.io/github/tj-actions/auto-doc)
[![Go Reference](https://pkg.go.dev/badge/github.com/tj-actions/auto-doc/v2.svg)](https://pkg.go.dev/github.com/tj-actions/auto-doc/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/tj-actions/auto-doc)](https://goreportcard.com/report/github.com/tj-actions/auto-doc)
[![CI](https://github.com/tj-actions/auto-doc/workflows/CI/badge.svg)](https://github.com/tj-actions/auto-doc/actions?query=workflow%3ACI)
[![Update release version.](https://github.com/tj-actions/auto-doc/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/auto-doc/actions?query=workflow%3A%22Update+release+version.%22)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bb5dc10c1d2645c0894fa6774300639b)](https://app.codacy.com/gh/tj-actions/auto-doc?utm_source=github.com\&utm_medium=referral\&utm_content=tj-actions/auto-doc\&utm_campaign=Badge_Grade_Settings)

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

## auto-doc

GitHub Action that generates beautiful, easy-to-read markdown tables with just a few lines of code. Say goodbye to manual table creation and hello to streamlined documentation that is always up-to-date.

## Features

*   Document your custom [Github action](https://docs.github.com/en/actions/creating-actions) using the `action.yml` file.
*   Document [reusable workflows](https://docs.github.com/en/actions/using-workflows/reusing-workflows) by specifying the filename.
*   Easy to understand markdown table of all inputs, outputs, secrets.
*   Fast and always up-to-date documentation.

## Table of Contents

*   [Usage](#usage)
*   [Inputs](#inputs)
*   [Examples](#examples)
*   [CLI](#cli)
    *   [Installation](#installation)
        *   [Install using Go:](#install-using-go)
        *   [Install using Homebrew:](#install-using-homebrew)
        *   [Install using Chocolatey (Windows):](#install-using-chocolatey-windows)
    *   [Synopsis](#synopsis)
    *   [Flags](#flags)
*   [Credits](#credits)
*   [Report Bugs](#report-bugs)
*   [Contributors ✨](#contributors-)

## Usage

Add the `Inputs` and/or `Outputs` and/or `Secrets` (only supported by reusable workflows) [`H2` header](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#headers) to any markdown file.

```yaml
...
    steps:
      - uses: actions/checkout@v2
      - name: Run auto-doc
        uses: tj-actions/auto-doc@v2
```

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|          INPUT          |  TYPE  | REQUIRED |    DEFAULT     |                                           DESCRIPTION                                            |
|-------------------------|--------|----------|----------------|--------------------------------------------------------------------------------------------------|
|        bin\_path         | string |  false   |                |                                   Path to the auto-doc binary                                    |
|      col\_max\_width      | string |  false   |    `"1000"`    |                                      Max width of a column                                       |
|      col\_max\_words      | string |  false   |     `"5"`      |                          Max number of words per line <br>in a column                            |
|        filename         | string |  false   | `"action.yml"` |                                      Path to the yaml file                                       |
|      input\_columns      | string |  false   |                |    List of action.yml **input** columns names <br>to display, default (display all columns)      |
|     markdown\_links      | string |  false   |   `"false"`    |  Boolean indicating whether to output input, <br>output and secret names as markdown <br>links   |
|         output          | string |  false   | `"README.md"`  |                                     Path to the output file                                      |
|     output\_columns      | string |  false   |                |    List of action.yml **output** column names <br>to display, default (display all columns)      |
|        reusable         | string |  false   |                |                 Boolean Indicating whether the file is <br>a reusable workflow                   |
| reusable\_input\_columns  | string |  false   |                | List of reusable workflow **input** column <br>names to display, default (display all columns)   |
| reusable\_output\_columns | string |  false   |                | List of reusable workflow **output** column <br>names to display, default (display all columns)  |
| reusable\_secret\_columns | string |  false   |                | List of reusable workflow **secret** column <br>names to display, default (display all columns)  |
|         version         | string |  false   |                |                                    The version number to run                                     |

<!-- AUTO-DOC-INPUT:END -->

**👆 This is generated 👆 using :point\_right: [action.yml](./action.yml)**

## Examples

Create a pull request each time the action.yml inputs/outputs change

```yaml
name: Update README.md with the latest actions.yml

on:
  push:
    branches:
      - main

jobs:
  update-doc:
     runs-on: ubuntu-latest
     steps:
       - name: Checkout
         uses: actions/checkout@v2.4.0
         with:
           fetch-depth: 0  # otherwise, you will failed to push refs to dest repo

       - name: Run auto-doc
         uses: tj-actions/auto-doc@v2

       - name: Verify Changed files
         uses: tj-actions/verify-changed-files@v8.6
         id: verify-changed-files
         with:
           files: |
             README.md

       - name: Create Pull Request
         if: steps.verify-changed-files.outputs.files_changed == 'true'
         uses: peter-evans/create-pull-request@v3
         with:
           base: "main"
           title: "auto-doc: Updated README.md"
           branch: "chore/auto-doc-update-readme"
           commit-message: "auto-doc: Updated README.md"
           body: "auto-doc: Updated README.md"
```

## CLI

### Installation

#### Install using Go:

```shell script
go get -u github.com/tj-actions/auto-doc/v2
```

#### Install using Homebrew:

```shell script
brew install tj-actions/tap/auto-doc
```

#### Install using Chocolatey (Windows):

```shell script
choco install auto-doc
```

### Synopsis

Automatically generate documentation for your custom github action or reusable workflow.

    auto-doc [flags]

### Flags

    --colMaxWidth string                  Max width of a column (default "1000")
    --colMaxWords string                  Max number of words per line in a column (default "6")
    -f, --filename string                 config file
    -h, --help                            help for auto-doc
    --inputColumns stringArray            list of input column names (default [Input,Type,Required,Default,Description])
    -o, --output string                   Output file (default "README.md")
    --outputColumns stringArray           list of output column names (default [Output,Type,Description])
    -r, --reusable                        A reusable workflow
    --reusableInputColumns stringArray    list of reusable input column names (default [Input,Type,Required,Default,Description])
    --reusableOutputColumns stringArray   list of reusable output column names (default [Output,Value,Description])
    --reusableSecretColumns stringArray   list of reusable secrets column names (default [Secret,Required,Description])

*   Free software: [Apache License 2.0](LICENSE)

If you feel generous and want to show some extra appreciation:

[![Buy me a coffee][buymeacoffee-shield]][buymeacoffee]

[buymeacoffee]: https://www.buymeacoffee.com/jackton1

[buymeacoffee-shield]: https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png

## Credits

This package was created with [Cookiecutter](https://github.com/cookiecutter/cookiecutter) using [cookiecutter-action](https://github.com/tj-actions/cookiecutter-action)

*   [cobra](https://github.com/spf13/cobra)
*   [goreleaser](https://github.com/goreleaser/goreleaser/)

## Report Bugs

Report bugs at https://github.com/tj-actions/auto-doc/issues.

If you are reporting a bug, please include:

*   Your operating system name and version.
*   Any details about your workflow that might be helpful in troubleshooting.
*   Detailed steps to reproduce the bug.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->

<!-- prettier-ignore-start -->

<!-- markdownlint-disable -->

<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/andreas-aman"><img src="https://avatars.githubusercontent.com/u/118175695?v=4?s=100" width="100px;" alt="Andreas Åman"/><br /><sub><b>Andreas Åman</b></sub></a><br /><a href="https://github.com/tj-actions/auto-doc/commits?author=andreas-aman" title="Code">💻</a> <a href="https://github.com/tj-actions/auto-doc/commits?author=andreas-aman" title="Documentation">📖</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/ViacheslavKudinov"><img src="https://avatars.githubusercontent.com/u/56436734?v=4?s=100" width="100px;" alt="Viacheslav Kudinov"/><br /><sub><b>Viacheslav Kudinov</b></sub></a><br /><a href="https://github.com/tj-actions/auto-doc/commits?author=ViacheslavKudinov" title="Code">💻</a> <a href="https://github.com/tj-actions/auto-doc/commits?author=ViacheslavKudinov" title="Tests">⚠️</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->

<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
