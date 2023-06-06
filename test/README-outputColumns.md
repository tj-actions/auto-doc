# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|                         INPUT                          |  TYPE  | REQUIRED |       DEFAULT       |                                                                                                                  DESCRIPTION                                                                                                                  |
|--------------------------------------------------------|--------|----------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|                        base_sha                        | string |  false   |                     |                                                                                      Specify a different base commit SHA <br>used for comparing changes                                                                                       |
|                     diff_relative                      | string |  false   |                     |                            Exclude changes outside the current directory <br>and show path names relative to <br>it. **NOTE:** This requires you to <br>specify the top level directory via <br>the `path` input.                             |
|                       dir_names                        | string |  false   |      `"false"`      |                                             Output unique changed directories instead of <br>filenames. **NOTE:** This returns `.` for <br>changed files located in the root <br>of the project.                                              |
|                  dir_names_max_depth                   | string |  false   |                     |                                                              Maximum depth of directories to output. <br>e.g `test/test1/test2` with max depth of <br>`2` returns `test/test1`.                                                               |
|                      fetch_depth                       | string |  false   |       `"50"`        |                                                       Depth of additional branch history fetched. <br>**NOTE**: This can be adjusted to <br>resolve errors with insufficient history.                                                         |
|                         files                          | string |  false   |                     |                       File and directory patterns to detect <br>changes using only these list of <br>file(s) (Defaults to the entire repo) **NOTE:** Multiline file/directory patterns <br>should not include quotes.                         |
|                 files_from_source_file                 | string |  false   |                     |                                                                                            Source file(s) used to populate the <br>`files` input.                                                                                             |
|                      files_ignore                      | string |  false   |                     |                                                                Ignore changes to these file(s) **NOTE:** <br>Multiline file/directory patterns should not include <br>quotes.                                                                 |
|             files_ignore_from_source_file              | string |  false   |                     |                                                                                         Source file(s) used to populate the <br>`files_ignore` input                                                                                          |
|                 files_ignore_separator                 | string |  false   |       `"\n"`        |                                                                                             Separator used to split the `files_ignore` <br>input                                                                                              |
|                    files_separator                     | string |  false   |       `"\n"`        |                                                                                                Separator used to split the `files` <br>input                                                                                                  |
|           include_all_old_new_renamed_files            | string |  false   |      `"false"`      |                                          Include `all_old_new_renamed_files` output. Note this can <br>generate a large output See: [#501](https://github.com/tj-actions/changed-files/issues/501).                                           |
|                          json                          | string |  false   |      `"false"`      |                                                                      Output list of changed files in <br>a JSON formatted string which can <br>be used for matrix jobs.                                                                       |
| ~~json_raw_format~~ <br> Use `json_unescaped` instead. | string |  false   |      `"false"`      |                            **Deprecated:** Output list of changed files <br>in a raw format which means <br>that the output will not be <br>surrounded by quotes and special characters <br>will not be escaped.                              |
|                     json_unescaped                     | string |  false   |      `"false"`      |                                                                     Output list of changed files in <br>a JSON formatted string without escaping <br>special characters.                                                                      |
|                     markdown_links                     | string |  false   |      `"false"`      |                                                                        Boolean indicating whether to output input, <br>output and secret names as markdown <br>links                                                                          |
|                   match_directories                    | string |  false   |      `"true"`       |                                                                                                Indicates whether to include match directories                                                                                                 |
|                old_new_files_separator                 | string |  false   |        `" "`        |                                                                                         Split character for old and new <br>renamed filename pairs.                                                                                           |
|                   old_new_separator                    | string |  false   |        `","`        |                                                                                             Split character for old and new <br>filename pairs.                                                                                               |
|                       output_dir                       | string |  false   | `".github/outputs"` |                                                                                                       Directory to store output files.                                                                                                        |
|                          path                          | string |  false   |        `"."`        |                                                                               Specify a relative path under `$GITHUB_WORKSPACE` <br>to locate the repository.                                                                                 |
|                       quotepath                        | string |  false   |      `"true"`       |                                                           Use non ascii characters to match <br>files and output the filenames completely <br>verbatim by setting this to `false`                                                             |
|                       separator                        | string |  false   |        `" "`        |                                                                                                      Split character for output strings.                                                                                                      |
|                          sha                           | string |  false   |                     |                                                                                        Specify a different commit SHA used <br>for comparing changes                                                                                          |
|                         since                          | string |  false   |                     |                                                                             Get changed files for commits whose <br>timestamp is older than the given <br>time.                                                                               |
|                since_last_remote_commit                | string |   true   |      `"false"`      | Use the last commit on the <br>remote branch as the `base_sha`. Defaults <br>to the last non merge commit <br>on the target branch for pull <br>request events and the previous remote <br>commit of the current branch for <br>push events.  |
|                         until                          | string |  false   |                     |                                                                            Get changed files for commits whose <br>timestamp is earlier than the given <br>time.                                                                              |
|                   write_output_files                   | string |  false   |      `"false"`      |                                                                                    Write outputs to files in the <br>`.github/outputs` folder by default.                                                                                     |

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
