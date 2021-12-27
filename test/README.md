# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|   INPUT   |  TYPE  | REQUIRED |        DEFAULT        |                                     DESCRIPTION                                     |
|-----------|--------|----------|-----------------------|-------------------------------------------------------------------------------------|
| base_sha  | string | false    |                       | Specify a base commit SHA used<br>for comparing changes                             |
| files     | string | false    |                       | Check for changes using only this<br>list of files (Defaults to the<br>entire repo) |
| path      | string | false    |                       | Specify a relative path under $GITHUB_WORKSPACE<br>to locate the repository         |
| separator | string | true     | ` `                   | Split character for array output                                                    |
| sha       | string | true     | `${{ github.sha }}`   | Specify a current commit SHA used<br>for comparing changes                          |
| token     | string | true     | `${{ github.token }}` | Github token or Repo Scoped Personal<br>Access Token                                |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |  TYPE  |                                       DESCRIPTION                                       |
|--------------------------------|--------|-----------------------------------------------------------------------------------------|
| added_files                    | string | List of added files.                                                                    |
| all_changed_and_modified_files | string | List of all changed files.                                                              |
| all_modified_files             | string | List of all copied modified and<br>added files.                                         |
| any_changed                    | string | Return true only when any files<br>provided using the files input have<br>changed.      |
| any_deleted                    | string | Return true only when any files<br>provided using the files input have<br>been deleted. |
| copied_files                   | string | List of copied files.                                                                   |
| deleted_files                  | string | List of deleted files.                                                                  |
| modified_files                 | string | List of modified files.                                                                 |
| only_changed                   | string | Return true when all files provided<br>using the files input have changed.<br>          |
| only_deleted                   | string | Return true when all files provided<br>using the files input have been<br>deleted.      |
| other_deleted_files            | string | Return list of deleted files not<br>listed in the files input.                          |
| renamed_files                  | string | List of renamed files.                                                                  |
| type_changed_files             | string | List of files that had type<br>changes.                                                 |
| unknown_files                  | string | List of unknown files.                                                                  |
| unmerged_files                 | string | List of unmerged files.                                                                 |

<!-- AUTO-DOC-OUTPUT:END -->
