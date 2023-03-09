# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|    INPUT    |  TYPE  |      DESCRIPTION       |
|-------------|--------|------------------------|
| config-path | string | The configuration path |
|  username   | string |        Username        |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|   OUTPUT   |                    VALUE                    |       DESCRIPTION        |
|------------|---------------------------------------------|--------------------------|
| firstword  | `"${{ jobs.example_job.outputs.output1 }}"` | The first output string  |
| secondword | `"${{ jobs.example_job.outputs.output2 }}"` | The second output string |

<!-- AUTO-DOC-OUTPUT:END -->

## Secrets

<!-- AUTO-DOC-SECRETS:START - Do not remove or modify this section -->

| SECRET | REQUIRED |      DESCRIPTION      |
|--------|----------|-----------------------|
| token  |   true   | Repo scoped PAT token |

<!-- AUTO-DOC-SECRETS:END -->
