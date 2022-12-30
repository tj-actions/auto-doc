# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|     INPUT      |  TYPE  | REQUIRED |              DEFAULT               |                                       DESCRIPTION                                       |
|----------------|--------|----------|------------------------------------|-----------------------------------------------------------------------------------------|
|    base_sha    | string |  false   |                                    |               Specify a base commit SHA<br>used for comparing changes<br>               |
|     files      | string |  false   | `"README.md"`<br>`"README.md"`<br> | Check for changes using only<br>this list of files (Defaults<br>to the entire repo)<br> |
|      path      | string |  false   |                                    |     Specify a relative path under<br>$GITHUB_WORKSPACE to locate the repository<br>     |
| path_separator | string |  false   |               `"\n"`               |                                   Path separator<br>                                    |
|   separator    | string |   true   |               `"\|"`               |                          Split character for array output<br>                           |
|      sha       | string |   true   |       `"${{ github.sha }}"`        |             Specify a current commit SHA<br>used for comparing changes<br>              |
|     token      | string |   true   |      `"${{ github.token }}"`       |                Github token or Repo Scoped<br>Personal Access Token<br>                 |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |  TYPE  |
|--------------------------------|--------|
|          added_files           | string |
| all_changed_and_modified_files | string |
|       all_modified_files       | string |
|          any_changed           | string |
|          any_deleted           | string |
|          copied_files          | string |
|         deleted_files          | string |
|         modified_files         | string |
|          only_changed          | string |
|          only_deleted          | string |
|      other_deleted_files       | string |
|         renamed_files          | string |
|       type_changed_files       | string |
|         unknown_files          | string |
|         unmerged_files         | string |

<!-- AUTO-DOC-OUTPUT:END -->
