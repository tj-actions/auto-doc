# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|                         INPUT                         |  TYPE  | REQUIRED |   DEFAULT   |      DESCRIPTION       |
|-------------------------------------------------------|--------|----------|-------------|------------------------|
| <a name="bool_tested"></a>[bool_tested](#bool_tested) |  bool  |  false   |   `true`    |      Test of bool      |
| <a name="config-path"></a>[config-path](#config-path) | string |   true   |             | The configuration path |
|     <a name="username"></a>[username](#username)      | string |  false   | `"example"` |        Username        |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|                       OUTPUT                       |                    VALUE                    |       DESCRIPTION        |
|----------------------------------------------------|---------------------------------------------|--------------------------|
|  <a name="firstword"></a>[firstword](#firstword)   | `"${{ jobs.example_job.outputs.output1 }}"` | The first output string  |
| <a name="secondword"></a>[secondword](#secondword) | `"${{ jobs.example_job.outputs.output2 }}"` | The second output string |

<!-- AUTO-DOC-OUTPUT:END -->

## Secrets

<!-- AUTO-DOC-SECRETS:START - Do not remove or modify this section -->

|               SECRET                | REQUIRED |      DESCRIPTION      |
|-------------------------------------|----------|-----------------------|
| <a name="token"></a>[token](#token) |   true   | Repo scoped PAT token |

<!-- AUTO-DOC-SECRETS:END -->
