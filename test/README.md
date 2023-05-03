# test

Test text ## Inputs

## Inputs

## Inputs

## Inputs

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|             OUTPUT             |  TYPE  |                                                                                                                                      DESCRIPTION                                                                                                                                       |
|--------------------------------|--------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
|          added_files           | string |                                                                                                                       Returns only files that are Added<br>(A).                                                                                                                        |
| all_changed_and_modified_files | string |                                                                                                      Returns all changed and modified files<br>i.e. *a combination of (ACMRDTUX)*                                                                                                      |
|       all_changed_files        | string |                                                                                     Returns all changed files i.e. *a<br> combination of all added, copied, modified<br>and renamed files (ACMR)*                                                                                      |
|       all_modified_files       | string |                                                                                Returns all changed files i.e. *a<br> combination of all added, copied, modified,<br>renamed and deleted files (ACMRD)*.                                                                                |
|   all_old_new_renamed_files    | string |                                                           Returns only files that are Renamed<br> and list their old and new<br> names. **NOTE:** This requires setting `include_all_old_new_renamed_files`<br>to `true` (R)                                                           |
|          any_changed           | string |         Returns `true` when any of the<br> filenames provided using the `files` input<br> has changed. If no `files` have<br> been specified,an empty string `''` is<br> returned. i.e. *using a combination of<br> all added, copied, modified and renamed<br>files (ACMR)*.          |
|          any_deleted           | string |                                                   Returns `true` when any of the<br> filenames provided using the `files` input<br> has been deleted. If no `files`<br> have been specified,an empty string `''`<br>is returned. (D)                                                   |
|          any_modified          | string | Returns `true` when any of the<br> filenames provided using the `files` input<br> has been modified. If no `files`<br> have been specified,an empty string `''`<br> is returned. i.e. *using a combination<br> of all added, copied, modified, renamed,<br>and deleted files (ACMRD)*. |
|          copied_files          | string |                                                                                                                       Returns only files that are Copied<br>(C).                                                                                                                       |
|         deleted_files          | string |                                                                                                                      Returns only files that are Deleted<br>(D).                                                                                                                       |
|         modified_files         | string |                                                                                                                      Returns only files that are Modified<br>(M).                                                                                                                      |
|          only_changed          | string |                Returns `true` when only files provided<br> using the `files` input has changed.<br> If no `files` have been specified,an<br> empty string `''` is returned. i.e.<br> *using a combination of all added,<br>copied, modified and renamed files (ACMR)*.                 |
|          only_deleted          | string |                                                        Returns `true` when only files provided<br> using the `files` input has been<br> deleted. If no `files` have been<br> specified,an empty string `''` is returned.<br>(D)                                                        |
|         only_modified          | string |                                                       Returns `true` when only files provided<br> using the `files` input has been<br> modified. If no `files` have been<br>specified,an empty string `''` is returned.(ACMRD).                                                        |
|      other_changed_files       | string |                                                              Returns all other changed files not<br> listed in the files input i.e.<br> *using a combination of all added,<br>copied, modified and renamed files (ACMR)*.                                                              |
|      other_deleted_files       | string |                                                                                 Returns all other deleted files not<br> listed in the files input i.e.<br> *a combination of all deleted files<br>(D)*                                                                                 |
|      other_modified_files      | string |                                                                Returns all other modified files not<br> listed in the files input i.e.<br> *a combination of all added, copied,<br>modified, and deleted files (ACMRD)*                                                                |
|         renamed_files          | string |                                                                                                                      Returns only files that are Renamed<br>(R).                                                                                                                       |
|       type_changed_files       | string |                                                                                                              Returns only files that have their<br>file type changed (T).                                                                                                              |
|         unknown_files          | string |                                                                                                                      Returns only files that are Unknown<br>(X).                                                                                                                       |
|         unmerged_files         | string |                                                                                                                      Returns only files that are Unmerged<br>(U).                                                                                                                      |

<!-- AUTO-DOC-OUTPUT:END -->
