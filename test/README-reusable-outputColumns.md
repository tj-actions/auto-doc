# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|                     INPUT                     |  TYPE   | REQUIRED |   DEFAULT   |         DESCRIPTION          |
|-----------------------------------------------|---------|----------|-------------|------------------------------|
| ~~bool_tested~~ <br> **Use `tested` instead** | boolean |  false   |   `true`    | **Deprecated:** Test of bool |
|                  config-path                  | string  |   true   |             |    The configuration path    |
|                    tested                     | boolean |  false   |   `false`   |         Test of bool         |
|                   username                    | string  |  false   | `"example"` |           Username           |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|   OUTPUT   |                    VALUE                    |
|------------|---------------------------------------------|
| firstword  | `"${{ jobs.example_job.outputs.output1 }}"` |
| secondword | `"${{ jobs.example_job.outputs.output2 }}"` |

<!-- AUTO-DOC-OUTPUT:END -->

## Secrets

<!-- AUTO-DOC-SECRETS:START - Do not remove or modify this section -->

| SECRET | REQUIRED |      DESCRIPTION      |
|--------|----------|-----------------------|
| token  |   true   | Repo scoped PAT token |

<!-- AUTO-DOC-SECRETS:END -->
