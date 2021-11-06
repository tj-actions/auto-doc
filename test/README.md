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

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section --> 

|             OUTPUT             |          DESCRIPTION           | VALUE |
|--------------------------------|--------------------------------|-------|
| added_files                    | List of added files.           |       |
| all_changed_and_modified_files | List of all changed files.     |       |
| all_modified_files             | List of all copied modified    |       |
|                                | and added files.               |       |
| any_changed                    | Return true only when any      |       |
|                                | files provided using the files |       |
|                                | input have changed.            |       |
| any_deleted                    | Return true only when any      |       |
|                                | files provided using the files |       |
|                                | input have been deleted.       |       |
| copied_files                   | List of copied files.          |       |
| deleted_files                  | List of deleted files.         |       |
| modified_files                 | List of modified files.        |       |
| only_changed                   | Return true when all files     |       |
|                                | provided using the files input |       |
|                                | have changed.                  |       |
| only_deleted                   | Return true when all files     |       |
|                                | provided using the files input |       |
|                                | have been deleted.             |       |
| other_changed_files            | Return list of changed files   |       |
|                                | not listed in the files input. |       |
| other_deleted_files            | Return list of deleted files   |       |
|                                | not listed in the files input. |       |
| renamed_files                  | List of renamed files.         |       |
| type_changed_files             | List of files that had type    |       |
|                                | changes.                       |       |
| unknown_files                  | List of unknown files.         |       |
| unmerged_files                 | List of unmerged files.        |       |

<!-- AUTO-DOC-OUTPUT:END -->















