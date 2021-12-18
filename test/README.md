# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|   INPUT   | REQUIRED |       DEFAULT       |          DESCRIPTION           |
|-----------|----------|---------------------|--------------------------------|
| separator | true     |                     | Split character for array      |
|           |          |                     | output                         |
| files     | false    |                     | Check for changes using only   |
|           |          |                     | this list of files (Defaults   |
|           |          |                     | to the entire repo)            |
| sha       | true     | ${{ github.sha }}   | Specify a current commit SHA   |
|           |          |                     | used for comparing changes     |
| base_sha  | false    |                     | Specify a base commit SHA on   |
|           |          |                     | used for comparing changes     |
| path      | false    |                     | Specify a relative path under  |
|           |          |                     | $GITHUB_WORKSPACE to locate    |
|           |          |                     | the repository                 |
| token     | true     | ${{ github.token }} | Github token or Repo Scoped    |
|           |          |                     | Personal Access Token          |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |          DESCRIPTION           |  TYPE  |
|--------------------------------|--------------------------------|--------|
| unmerged_files                 | List of unmerged files.        | string |
| any_changed                    | Return true only when any      | string |
|                                | files provided using the files |        |
|                                | input have changed.            |        |
| other_deleted_files            | Return list of deleted files   | string |
|                                | not listed in the files input. |        |
| renamed_files                  | List of renamed files.         | string |
| all_changed_and_modified_files | List of all changed files.     | string |
| all_modified_files             | List of all copied modified    | string |
|                                | and added files.               |        |
| any_deleted                    | Return true only when any      | string |
|                                | files provided using the files |        |
|                                | input have been deleted.       |        |
| only_deleted                   | Return true when all files     | string |
|                                | provided using the files input |        |
|                                | have been deleted.             |        |
| added_files                    | List of added files.           | string |
| modified_files                 | List of modified files.        | string |
| unknown_files                  | List of unknown files.         | string |
| only_changed                   | Return true when all files     | string |
|                                | provided using the files input |        |
|                                | have changed.                  |        |
| copied_files                   | List of copied files.          | string |
| deleted_files                  | List of deleted files.         | string |
| type_changed_files             | List of files that had type    | string |
|                                | changes.                       |        |

<!-- AUTO-DOC-OUTPUT:END -->
