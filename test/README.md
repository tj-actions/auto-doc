# test

## Inputs
<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section --> 
|         INPUT          | REQUIRED |       DEFAULT       |          DESCRIPTION           |
|------------------------|----------|---------------------|--------------------------------|
| sha                    | true     | ${{ github.sha }}   | Specify a current commit SHA   |
|                        |          |                     | used for comparing changes     |
| base_sha               | false    |                     | Specify a base commit SHA on   |
|                        |          |                     | used for comparing changes     |
| path                   | false    |                     | Specify a relative path under  |
|                        |          |                     | $GITHUB_WORKSPACE to locate    |
|                        |          |                     | the repository                 |
| token                  | true     | ${{ github.token }} | Github token or Repo Scoped    |
|                        |          |                     | Personal Access Token          |
| separator              | true     |                     | Split character for array      |
|                        |          |                     | output                         |
| files_from_source_file | false    |                     | Source file to populate the    |
|                        |          |                     | files input                    |
| files                  | false    |                     | Check for changes using only   |
|                        |          |                     | this list of files (Defaults   |
|                        |          |                     | to the entire repo)            |
<!-- AUTO-DOC-INPUT:END -->


## Outputs
<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section --> 
|             OUTPUT             |          DESCRIPTION           | VALUE |
|--------------------------------|--------------------------------|-------|
| other_changed_files            | Return list of changed files   |       |
|                                | not listed in the files input. |       |
| other_deleted_files            | Return list of deleted files   |       |
|                                | not listed in the files input. |       |
| all_changed_and_modified_files | List of all changed files.     |       |
| modified_files                 | List of modified files.        |       |
| only_changed                   | Return true when all files     |       |
|                                | provided using the files input |       |
|                                | have changed.                  |       |
| any_deleted                    | Return true only when any      |       |
|                                | files provided using the files |       |
|                                | input have been deleted.       |       |
| added_files                    | List of added files.           |       |
| deleted_files                  | List of deleted files.         |       |
| type_changed_files             | List of files that had type    |       |
|                                | changes.                       |       |
| unmerged_files                 | List of unmerged files.        |       |
| all_modified_files             | List of all copied modified    |       |
|                                | and added files.               |       |
| any_changed                    | Return true only when any      |       |
|                                | files provided using the files |       |
|                                | input have changed.            |       |
| copied_files                   | List of copied files.          |       |
| unknown_files                  | List of unknown files.         |       |
| only_deleted                   | Return true when all files     |       |
|                                | provided using the files input |       |
|                                | have been deleted.             |       |
| renamed_files                  | List of renamed files.         |       |
<!-- AUTO-DOC-OUTPUT:END -->

