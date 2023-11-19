## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->
```yaml
- uses: tj-actions/changed-files@v40.1.1
  id: changed-files
  with:
    # Specify a different base commit SHA used for comparing 
    # changes 
    # Type: string
    base_sha: ''

    # Exclude changes outside the current directory and show path 
    # names relative to it. **NOTE:** This requires you to 
    # specify the top level directory via the `path` input. 
    # Type: string
    diff_relative: ''

    # Output unique changed directories instead of filenames. **NOTE:** This 
    # returns `.` for changed files located in the root 
    # of the project. 
    # Type: boolean
    # Default: "false"
    dir_names: ''

    # Maximum depth of directories to output. e.g `test/test1/test2` with 
    # max depth of `2` returns `test/test1`. 
    # Type: string
    dir_names_max_depth: ''

    # Depth of additional branch history fetched. **NOTE**: This can 
    # be adjusted to resolve errors with insufficient history. 
    # Type: string
    # Default: "50"
    fetch_depth: ''

    # File and directory patterns to detect changes using only 
    # these list of file(s) (Defaults to the entire repo) **NOTE:** Multiline file/directory patterns 
    # should not include quotes. 
    # Type: string
    # Default: a.txt
    #          b.txt
    #          test.txt
    #          
    files: ''

    # Source file(s) used to populate the `files` input.
    # Type: string
    files_from_source_file: ''

    # Ignore changes to these file(s) **NOTE:** Multiline file/directory patterns 
    # should not include quotes. 
    # Type: string
    files_ignore: ''

    # Source file(s) used to populate the `files_ignore` input
    # Type: string
    files_ignore_from_source_file: ''

    # Separator used to split the `files_ignore` input
    # Type: string
    # Default: "\n"
    files_ignore_separator: ''

    # Separator used to split the `files` input
    # Type: string
    # Default: "\n"
    files_separator: ''

    # Include `all_old_new_renamed_files` output. Note this can generate a large 
    # output See: [#501](https://github.com/tj-actions/changed-files/issues/501). 
    # Type: boolean
    # Default: "false"
    include_all_old_new_renamed_files: ''

    # Output list of changed files in a JSON formatted 
    # string which can be used for matrix jobs. 
    # Type: boolean
    # Default: "false"
    json: ''

    # Output list of changed files in a raw format 
    # which means that the output will not be surrounded 
    # by quotes and special characters will not be escaped. 
    # Type: boolean
    # Default: "false"
    # Deprecated: Use `json_unescaped` instead.
    json_raw_format: ''

    # Output list of changed files in a JSON formatted 
    # string without escaping special characters. 
    # Type: boolean
    # Default: "false"
    json_unescaped: ''

    # Boolean indicating whether to output input, output and secret 
    # names as markdown links 
    # Type: boolean
    # Default: "false"
    markdown_links: ''

    # Indicates whether to include match directories
    # Type: boolean
    # Default: "true"
    match_directories: ''

    # Split character for old and new renamed filename pairs.
    # Type: string
    # Default: " "
    old_new_files_separator: ''

    # Split character for old and new filename pairs.
    # Type: string
    # Default: ","
    old_new_separator: ''

    # Directory to store output files.
    # Type: string
    # Default: ".github/outputs"
    output_dir: ''

    # Specify a relative path under `$GITHUB_WORKSPACE` to locate the 
    # repository. 
    # Type: string
    # Default: "."
    path: ''

    # Use non ascii characters to match files and output 
    # the filenames completely verbatim by setting this to `false` 
    # Type: boolean
    # Default: "true"
    quotepath: ''

    # Split character for output strings.
    # Type: string
    # Default: "|"
    separator: ''

    # Specify a different commit SHA used for comparing changes
    # Type: string
    sha: ''

    # Get changed files for commits whose timestamp is older 
    # than the given time. 
    # Type: string
    since: ''

    # Use the last commit on the remote branch as 
    # the `base_sha`. Defaults to the last non merge commit 
    # on the target branch for pull request events and 
    # the previous remote commit of the current branch for 
    # push events. 
    # Type: boolean
    # Default: "false"
    since_last_remote_commit: ''

    # The GitHub token to use for authentication.
    # Type: string
    # Default: "${{ github.token }}"
    token: ''

    # Get changed files for commits whose timestamp is earlier 
    # than the given time. 
    # Type: string
    until: ''

    # Write outputs to files in the `.github/outputs` folder by 
    # default. 
    # Type: boolean
    # Default: "false"
    write_output_files: ''

```
<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |  TYPE  |                                                                                                                                   DESCRIPTION                                                                                                                                    |
|--------------------------------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|          added_files           | string |                                                                                                                   Returns only files that are Added <br>(A).                                                                                                                     |
| all_changed_and_modified_files | string |                                                                                                  Returns all changed and modified files <br>i.e. *a combination of (ACMRDTUX)*                                                                                                   |
|       all_changed_files        | string |                                                                                     Returns all changed files i.e. *a combination of all added, copied, modified and renamed files (ACMR)*                                                                                       |
|       all_modified_files       | string |                                                                                Returns all changed files i.e. *a combination of all added, copied, modified, renamed and deleted files (ACMRD)*.                                                                                 |
|   all_old_new_renamed_files    | string |                                                       Returns only files that are Renamed <br>and list their old and new <br>names. **NOTE:** This requires setting `include_all_old_new_renamed_files` <br>to `true` (R)                                                        |
|          any_changed           | string |         Returns `true` when any of the <br>filenames provided using the `files` input <br>has changed. If no `files` have <br>been specified,an empty string `''` is <br>returned. i.e. *using a combination of all added, copied, modified and renamed files (ACMR)*.           |
|          any_deleted           | string |                                               Returns `true` when any of the <br>filenames provided using the `files` input <br>has been deleted. If no `files` <br>have been specified,an empty string `''` <br>is returned. (D)                                                |
|          any_modified          | string | Returns `true` when any of the <br>filenames provided using the `files` input <br>has been modified. If no `files` <br>have been specified,an empty string `''` <br>is returned. i.e. *using a combination of all added, copied, modified, renamed, and deleted files (ACMRD)*.  |
|          copied_files          | string |                                                                                                                   Returns only files that are Copied <br>(C).                                                                                                                    |
|         deleted_files          | string |                                                                                                                  Returns only files that are Deleted <br>(D).                                                                                                                    |
|         modified_files         | string |                                                                                                                  Returns only files that are Modified <br>(M).                                                                                                                   |
|          only_changed          | string |              Returns `true` when only files provided <br>using the `files` input has changed. <br>If no `files` have been specified,an <br>empty string `''` is returned. i.e. <br>*using a combination of all added, copied, modified and renamed files (ACMR)*.                |
|          only_deleted          | string |                                                    Returns `true` when only files provided <br>using the `files` input has been <br>deleted. If no `files` have been <br>specified,an empty string `''` is returned. <br>(D)                                                     |
|         only_modified          | string |                                                   Returns `true` when only files provided <br>using the `files` input has been <br>modified. If no `files` have been <br>specified,an empty string `''` is returned.(ACMRD).                                                     |
|      other_changed_files       | string |                                                            Returns all other changed files not <br>listed in the files input i.e. <br>*using a combination of all added, copied, modified and renamed files (ACMR)*.                                                             |
|      other_deleted_files       | string |                                                                              Returns all other deleted files not <br>listed in the files input i.e. <br>*a  combination of all deleted files (D)*                                                                                |
|      other_modified_files      | string |                                                             Returns all other modified files not <br>listed in the files input i.e. <br>*a  combination of all added, copied, modified, and deleted files (ACMRD)*                                                               |
|         renamed_files          | string |                                                                                                                  Returns only files that are Renamed <br>(R).                                                                                                                    |
|       type_changed_files       | string |                                                                                                          Returns only files that have their <br>file type changed (T).                                                                                                           |
|         unknown_files          | string |                                                                                                                  Returns only files that are Unknown <br>(X).                                                                                                                    |
|         unmerged_files         | string |                                                                                                                  Returns only files that are Unmerged <br>(U).                                                                                                                   |

<!-- AUTO-DOC-OUTPUT:END -->
