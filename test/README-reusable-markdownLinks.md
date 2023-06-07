# test

## Inputs

<!-- AUTO-DOC-INPUT:START - Do not remove or modify this section -->

|                                              INPUT                                              |  TYPE   | REQUIRED |   DEFAULT   |         DESCRIPTION          |
|-------------------------------------------------------------------------------------------------|---------|----------|-------------|------------------------------|
| ~~<a name="input_bool_tested"></a>[bool_tested](#input_bool_tested)~~ <br> Use `tested` instead | boolean |  false   |   `true`    | **Deprecated:** Test of bool |
|                <a name="input_config-path"></a>[config-path](#input_config-path)                | string  |   true   |             |    The configuration path    |
|                       <a name="input_tested"></a>[tested](#input_tested)                        | boolean |  false   |   `false`   |         Test of bool         |
|                    <a name="input_username"></a>[username](#input_username)                     | string  |  false   | `"example"` |           Username           |

<!-- AUTO-DOC-INPUT:END -->

## Outputs

<!-- AUTO-DOC-OUTPUT:START - Do not remove or modify this section -->

|                              OUTPUT                              |                    VALUE                    |       DESCRIPTION        |
|------------------------------------------------------------------|---------------------------------------------|--------------------------|
|  <a name="output_firstword"></a>[firstword](#output_firstword)   | `"${{ jobs.example_job.outputs.output1 }}"` | The first output string  |
| <a name="output_secondword"></a>[secondword](#output_secondword) | `"${{ jobs.example_job.outputs.output2 }}"` | The second output string |

<!-- AUTO-DOC-OUTPUT:END -->

## Secrets

<!-- AUTO-DOC-SECRETS:START - Do not remove or modify this section -->

|                      SECRET                       | REQUIRED |      DESCRIPTION      |
|---------------------------------------------------|----------|-----------------------|
| <a name="secret_token"></a>[token](#secret_token) |   true   | Repo scoped PAT token |

<!-- AUTO-DOC-SECRETS:END -->
