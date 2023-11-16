[![Ubuntu](https://img.shields.io/badge/Ubuntu-E95420?style=for-the-badge\&logo=ubuntu\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Mac OS](https://img.shields.io/badge/mac%20os-000000?style=for-the-badge\&logo=macos\&logoColor=F0F0F0)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge\&logo=windows\&logoColor=white)](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions#jobsjob_idruns-on)
[![Public workflows that use this action.](https://img.shields.io/endpoint?style=for-the-badge\&url=https%3A%2F%2Fused-by.vercel.app%2Fapi%2Fgithub-actions%2Fused-by%3Faction%3Dtj-actions%2Fauto-doc%26badge%3Dtrue%26package_id%3DUGFja2FnZS0yOTU3MDU0NDI1)](https://github.com/search?o=desc\&q=tj-actions+auto-doc+language%3AYAML\&s=\&type=Code)

[![CI](https://github.com/tj-actions/auto-doc/workflows/CI/badge.svg)](https://github.com/tj-actions/auto-doc/actions?query=workflow%3ACI)
![Coverage](https://img.shields.io/badge/Coverage-85.7%25-brightgreen)
[![Update release version.](https://github.com/tj-actions/auto-doc/workflows/Update%20release%20version./badge.svg)](https://github.com/tj-actions/auto-doc/actions?query=workflow%3A%22Update+release+version.%22)
[![codecov](https://codecov.io/github/tj-actions/auto-doc/branch/main/graph/badge.svg?token=TNXW4QRRJD)](https://codecov.io/github/tj-actions/auto-doc)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/bb5dc10c1d2645c0894fa6774300639b)](https://app.codacy.com/gh/tj-actions/auto-doc?utm_source=github.com\&utm_medium=referral\&utm_content=tj-actions/auto-doc\&utm_campaign=Badge_Grade_Settings)
[![Go Report Card](https://goreportcard.com/badge/github.com/tj-actions/auto-doc)](https://goreportcard.com/report/github.com/tj-actions/auto-doc)
[![Go Reference](https://pkg.go.dev/badge/github.com/tj-actions/auto-doc/v3.svg)](https://pkg.go.dev/github.com/tj-actions/auto-doc/v3)

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

## auto-doc

GitHub Action that generates beautiful, easy-to-read markdown tables with just a few lines of code. Say goodbye to manual table creation and hello to streamlined documentation that is always up-to-date.

## Features

*   Document your custom [Github action](https://docs.github.com/en/actions/creating-actions) using the `action.yml` file.
*   Document [reusable workflows](https://docs.github.com/en/actions/using-workflows/reusing-workflows) by specifying the filename.
*   Easy to understand markdown table of all inputs, outputs, secrets.
*   Show deprecated inputs.
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

Add the `Inputs` and/or `Outputs`
and/or `Secrets` (only supported by reusable workflows)
and/or `Description` (only supported by actions)
[`H2` header](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet#headers) to any markdown file.

```yaml
...
    steps:
      - uses: actions/checkout@v4
      - name: Run auto-doc
        uses: tj-actions/auto-doc@v3
```

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

```yaml
- uses: tj-actions/auto-doc@v3
  id: auto-doc
  with:
    # Optionally pass a path to the auto-doc binary
    bin_path: ''

    # Max width of a column
    # Default: "1000"
    col_max_width: ''

    # Max number of words per line in a column
    # Default: "5"
    col_max_words: ''

    # Path to the yaml file
    # Default: "action.yml"
    filename: ''

    # List of action.yml **input** columns names to display, default 
    # (display all columns) 
    input_columns: ''

    # Boolean indicating whether to output input, output and secret 
    # names as markdown links 
    # Default: "true"
    markdown_links: ''

    # Path to the output file
    # Default: "README.md"
    output: ''

    # List of action.yml **output** column names to display, default 
    # (display all columns) 
    output_columns: ''

    # Repository name with owner. For example, tj-actions/auto-doc
    # Default: "${{ github.repository }}"
    repository: ''

    # Boolean Indicating whether the file is a reusable workflow
    reusable: ''

    # List of reusable workflow **input** column names to display, 
    # default (display all columns) 
    reusable_input_columns: ''

    # List of reusable workflow **output** column names to display, 
    # default (display all columns) 
    reusable_output_columns: ''

    # List of reusable workflow **secret** column names to display, 
    # default (display all columns) 
    reusable_secret_columns: ''

    # GitHub token or Personal Access Token used to fetch 
    # the repository latest tag. 
    # Default: "${{ github.token }}"
    token: ''

    # Enable code block documentation
    # Default: "false"
    use_code_blocks: ''

    # Use the major version of the repository tag e.g 
    # v1.0.0 -> v1 
    # Default: "false"
    use_major_version: ''

    # The version number to run
    version: ''

```

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
         uses: actions/checkout@v4
         with:
           fetch-depth: 0  # otherwise, you will failed to push refs to dest repo

       - name: Run auto-doc
         uses: tj-actions/auto-doc@v3

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
go get -u github.com/tj-actions/auto-doc/v3
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

        --colMaxWidth string                  Max width of a column. (default "1000")
        --colMaxWords string                  Max number of words per line in a column. (default "6")
    -f, --filename string                 config file.
    -h, --help                            help for auto-doc
        --inputColumns stringArray            list of input column names. (default [Input,Type,Required,Default,Description])
    -m, --markdownLinks                   Names of inputs, outputs and secrets as markdown links.
    -o, --output string                   Output file. (default "README.md")
        --outputColumns stringArray           list of output column names. (default [Output,Type,Description])
        --repository string                   Repository name with owner. Example: tj-actions/auto-doc
    -r, --reusable                        A reusable workflow.
        --reusableInputColumns stringArray    list of reusable input column names. (default [Input,Type,Required,Default,Description])
        --reusableOutputColumns stringArray   list of reusable output column names. (default [Output,Value,Description])
        --reusableSecretColumns stringArray   list of reusable secrets column names. (default [Secret,Required,Description])
        --token string                        GitHub token or Personal Access Token used to fetch the repository latest tag.
        --useCodeBlocks                       Enable code block documentation.
        --useMajorVersion                     Use the major version of the repository tag. Example: v1.0.0 -> v                 Use the major version of the repository tag. e.g v1.0.0 -> v1

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
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/looztra"><img src="https://avatars.githubusercontent.com/u/246526?v=4?s=100" width="100px;" alt="Christophe Furmaniak"/><br /><sub><b>Christophe Furmaniak</b></sub></a><br /><a href="https://github.com/tj-actions/auto-doc/commits?author=looztra" title="Code">💻</a> <a href="https://github.com/tj-actions/auto-doc/commits?author=looztra" title="Tests">⚠️</a></td>
      <td align="center" valign="top" width="14.28%"><a href="https://github.com/mw-root"><img src="https://avatars.githubusercontent.com/u/12434761?v=4?s=100" width="100px;" alt="Mike W"/><br /><sub><b>Mike W</b></sub></a><br /><a href="https://github.com/tj-actions/auto-doc/commits?author=mw-root" title="Code">💻</a> <a href="https://github.com/tj-actions/auto-doc/commits?author=mw-root" title="Documentation">📖</a></td>
    </tr>
  </tbody>
</table>

<!-- markdownlint-restore -->

<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
