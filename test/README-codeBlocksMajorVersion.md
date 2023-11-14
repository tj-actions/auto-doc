## Inputs

<!-- AUTO-DOC:START - Do not remove or modify this section -->
```yaml
- uses: tj-actions/changed-files@v40
  with:
    # Specify a different base commit SHA used for comparing 
    # changes 
    base_sha: ''

    # Exclude changes outside the current directory and show path 
    # names relative to it. **NOTE:** This requires you to 
    # specify the top level directory via the `path` input. 
    diff_relative: ''

    # Output unique changed directories instead of filenames. **NOTE:** This 
    # returns `.` for changed files located in the root 
    # of the project. 
    # Default: false
    dir_names: ''

    # Maximum depth of directories to output. e.g `test/test1/test2` with 
    # max depth of `2` returns `test/test1`. 
    dir_names_max_depth: ''

    # Depth of additional branch history fetched. **NOTE**: This can 
    # be adjusted to resolve errors with insufficient history. 
    # Default: 50
    fetch_depth: ''

    # File and directory patterns to detect changes using only 
    # these list of file(s) (Defaults to the entire repo) **NOTE:** Multiline file/directory patterns 
    # should not include quotes. 
    files: ''

    # Source file(s) used to populate the `files` input.
    files_from_source_file: ''

    # Ignore changes to these file(s) **NOTE:** Multiline file/directory patterns 
    # should not include quotes. 
    files_ignore: ''

    # Source file(s) used to populate the `files_ignore` input
    files_ignore_from_source_file: ''

    # Separator used to split the `files_ignore` input
    # Default: 

    files_ignore_separator: ''

    # Separator used to split the `files` input
    # Default: 

    files_separator: ''

    # Include `all_old_new_renamed_files` output. Note this can generate a large 
    # output See: [#501](https://github.com/tj-actions/changed-files/issues/501). 
    # Default: false
    include_all_old_new_renamed_files: ''

    # Output list of changed files in a JSON formatted 
    # string which can be used for matrix jobs. 
    # Default: false
    json: ''

    # Output list of changed files in a raw format 
    # which means that the output will not be surrounded 
    # by quotes and special characters will not be escaped. 
    # Default: false
    # Deprecated: Use `json_unescaped` instead.
    json_raw_format: ''

    # Output list of changed files in a JSON formatted 
    # string without escaping special characters. 
    # Default: false
    json_unescaped: ''

    # Boolean indicating whether to output input, output and secret 
    # names as markdown links 
    # Default: false
    markdown_links: ''

    # Indicates whether to include match directories
    # Default: true
    match_directories: ''

    # Split character for old and new renamed filename pairs.
    # Default:  
    old_new_files_separator: ''

    # Split character for old and new filename pairs.
    # Default: ,
    old_new_separator: ''

    # Directory to store output files.
    # Default: .github/outputs
    output_dir: ''

    # Specify a relative path under `$GITHUB_WORKSPACE` to locate the 
    # repository. 
    # Default: .
    path: ''

    # Use non ascii characters to match files and output 
    # the filenames completely verbatim by setting this to `false` 
    # Default: true
    quotepath: ''

    # Split character for output strings.
    # Default:  
    separator: ''

    # Specify a different commit SHA used for comparing changes
    sha: ''

    # Get changed files for commits whose timestamp is older 
    # than the given time. 
    since: ''

    # Use the last commit on the remote branch as 
    # the `base_sha`. Defaults to the last non merge commit 
    # on the target branch for pull request events and 
    # the previous remote commit of the current branch for 
    # push events. 
    # Default: false
    since_last_remote_commit: ''

    # The GitHub token to use for authentication.
    # Default: ${{ github.token }}
    token: ''

    # Get changed files for commits whose timestamp is earlier 
    # than the given time. 
    until: ''

    # Write outputs to files in the `.github/outputs` folder by 
    # default. 
    # Default: false
    write_output_files: ''

```
<!-- AUTO-DOC:END -->