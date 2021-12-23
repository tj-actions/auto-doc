# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|   INPUT   |  TYPE  | REQUIRED |       DEFAULT       |                                  DESCRIPTION                                  |
|-----------|--------|----------|---------------------|-------------------------------------------------------------------------------|
| base_sha  | string | false    |                     | Specify a base commit SHA on used for comparing changes                       |
| files     | string | false    |                     | Check for changes using only this list of files (Defaults to the entire repo) |
| path      | string | false    |                     | Specify a relative path under $GITHUB_WORKSPACE to locate the repository      |
| separator | string | true     |                     | Split character for array output                                              |
| sha       | string | true     | ${{ github.sha }}   | Specify a current commit SHA used for comparing changes                       |
| token     | string | true     | ${{ github.token }} | Github token or Repo Scoped Personal Access Token                             |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |                                    DESCRIPTION                                    |  TYPE  |
|--------------------------------|-----------------------------------------------------------------------------------|--------|
| added_files                    | List of added files.                                                              | string |
| all_changed_and_modified_files | List of all changed files.                                                        | string |
| all_modified_files             | List of all copied modified and added files.                                      | string |
| any_changed                    | Return true only when any files provided using the files input have changed.      | string |
| any_deleted                    | Return true only when any files provided using the files input have been deleted. | string |
| copied_files                   | List of copied files.                                                             | string |
| deleted_files                  | List of deleted files.                                                            | string |
| modified_files                 | List of modified files.                                                           | string |
| only_changed                   | Return true when all files provided using the files input have changed.           | string |
| only_deleted                   | Return true when all files provided using the files input have been deleted.      | string |
| other_deleted_files            | Return list of deleted files not listed in the files input.                       | string |
| renamed_files                  | List of renamed files.                                                            | string |
| type_changed_files             | List of files that had type changes.                                              | string |
| unknown_files                  | List of unknown files.                                                            | string |
| unmerged_files                 | List of unmerged files.                                                           | string |

<!-- AUTO-DOC-OUTPUT:END -->
