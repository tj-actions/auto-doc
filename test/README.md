# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|         INPUT          | REQUIRED |       DEFAULT       |          DESCRIPTION           |
|------------------------|----------|---------------------|--------------------------------|
| base_sha               | false    |                     | Specify a base commit SHA on   |
|                        |          |                     | used for comparing changes     |
| files                  | false    |                     | Check for changes using only   |
|                        |          |                     | this list of files (Defaults   |
|                        |          |                     | to the entire repo)            |
| files_from_source_file | false    |                     | Source file to populate the    |
|                        |          |                     | files input                    |
| path                   | false    |                     | Specify a relative path under  |
|                        |          |                     | $GITHUB_WORKSPACE to locate    |
|                        |          |                     | the repository                 |
| separator              | true     |                     | Split character for array      |
|                        |          |                     | output                         |
| sha                    | true     | ${{ github.sha }}   | Specify a current commit SHA   |
|                        |          |                     | used for comparing changes     |
| token                  | true     | ${{ github.token }} | Github token or Repo Scoped    |
|                        |          |                     | Personal Access Token          |

<!-- AUTO-DOC-INPUT:END -->

## Outputs
