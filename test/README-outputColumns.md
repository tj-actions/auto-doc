# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|               INPUT               |  TYPE  | REQUIRED |       DEFAULT       |                                                                                                              DESCRIPTION                                                                                                              |
|-----------------------------------|--------|----------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|             base_sha              | string |  false   |                     |                                                                                 Specify a different base commit<br>SHA used for comparing changes<br>                                                                                 |
|           diff_relative           | string |  false   |                     |                         Exclude changes outside the current<br>directory and show path names<br>relative to it. **NOTE:** This<br>requires you to specify the<br>top level directory via the<br>`path` input.                         |
|             dir_names             | string |  false   |      `"false"`      |                                           Output unique changed directories instead<br>of filenames. **NOTE:** This returns<br>`.` for changed files locatedin the root of the<br>project.                                            |
|        dir_names_max_depth        | string |  false   |                     |                                                             Maximum depth of directories to<br>output. e.g `test/test1/test2` with maxdepth of `2` returns `test/test1`.                                                              |
|            fetch_depth            | string |  false   |       `"50"`        |                                                   Depth of additional branch history<br>fetched. **NOTE**: This can be<br>adjusted to resolve errors with<br>insufficient history.                                                    |
|               files               | string |  false   |                     |                  File and directory patterns to<br>detect changes using only these<br>list of file(s) (Defaults to<br>the entire repo) **NOTE:** Multiline<br>file/directory patterns should not include<br>quotes.                   |
|      files_from_source_file       | string |  false   |                     |                                                                                         Source file(s) used to populate<br>the `files` input.                                                                                         |
|           files_ignore            | string |  false   |                     |                                                             Ignore changes to these file(s)<br>**NOTE:** Multiline file/directory patterns should<br>not include quotes.                                                              |
|   files_ignore_from_source_file   | string |  false   |                     |                                                                                      Source file(s) used to populate<br>the `files_ignore` input                                                                                      |
|      files_ignore_separator       | string |  false   |       `"\n"`        |                                                                                          Separator used to split the<br>`files_ignore` input                                                                                          |
|          files_separator          | string |  false   |       `"\n"`        |                                                                                             Separator used to split the<br>`files` input                                                                                              |
| include_all_old_new_renamed_files | string |  false   |      `"false"`      |                                       Include `all_old_new_renamed_files` output. Note thiscan generate a large output<br>See: [#501](https://github.com/tj-actions/changed-files/issues/501).                                        |
|               json                | string |  false   |      `"false"`      |                                                                  Output list of changed files<br>in a JSON formatted string<br>which can be used for<br>matrix jobs.                                                                  |
|          json_raw_format          | string |  false   |      `"false"`      |                                 Output list of changed files<br>in a raw format which<br>means that the output will<br>not be surrounded by quotes<br>and special characters will not<br>be escaped.                                  |
|         match_directories         | string |  false   |      `"true"`       |                                                                                           Indicates whether to include match<br>directories                                                                                           |
|      old_new_files_separator      | string |  false   |        `" "`        |                                                                                      Split character for old and<br>new renamed filename pairs.                                                                                       |
|         old_new_separator         | string |  false   |        `","`        |                                                                                          Split character for old and<br>new filename pairs.                                                                                           |
|            output_dir             | string |  false   | `".github/outputs"` |                                                                                                 Directory to store output files.<br>                                                                                                  |
|               path                | string |  false   |        `"."`        |                                                                            Specify a relative path under<br>`$GITHUB_WORKSPACE` to locate the repository.                                                                             |
|             quotepath             | string |  false   |      `"true"`       |                                                       Use non ascii characters to<br>match files and output the<br>filenames completely verbatim by setting<br>this to `false`                                                        |
|             separator             | string |  false   |        `" "`        |                                                                                                Split character for output strings<br>                                                                                                 |
|                sha                | string |  false   |                     |                                                                                     Specify a different commit SHA<br>used for comparing changes                                                                                      |
|               since               | string |  false   |                     |                                                                           Get changed files for commits<br>whose timestamp is older than<br>the given time.                                                                           |
|     since_last_remote_commit      | string |   true   |      `"false"`      | Use the last commit on<br>the remote branch as the<br>`base_sha`. Defaults to the lastnon merge commit on the<br>target branch for pull request<br>events and the previous remote<br>commit of the current branch<br>for push events. |
|               until               | string |  false   |                     |                                                                          Get changed files for commits<br>whose timestamp is earlier than<br>the given time.                                                                          |
|        write_output_files         | string |  false   |      `"false"`      |                                                                                 Write outputs to files in<br>the `.github/outputs` folder by default.                                                                                 |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |  TYPE  |
|--------------------------------|--------|
|          added_files           | string |
| all_changed_and_modified_files | string |
|       all_changed_files        | string |
|       all_modified_files       | string |
|   all_old_new_renamed_files    | string |
|          any_changed           | string |
|          any_deleted           | string |
|          any_modified          | string |
|          copied_files          | string |
|         deleted_files          | string |
|         modified_files         | string |
|          only_changed          | string |
|          only_deleted          | string |
|         only_modified          | string |
|      other_changed_files       | string |
|      other_deleted_files       | string |
|      other_modified_files      | string |
|         renamed_files          | string |
|       type_changed_files       | string |
|         unknown_files          | string |
|         unmerged_files         | string |

<!-- AUTO-DOC-OUTPUT:END -->
