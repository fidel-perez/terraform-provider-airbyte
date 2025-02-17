---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sendinblue Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSendinblue Resource
---

# airbyte_source_sendinblue (Resource)

SourceSendinblue Resource

## Example Usage

```terraform
resource "airbyte_source_sendinblue" "my_source_sendinblue" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "0de87dfe-701e-4dbd-8d10-cf57eb672b8a"
  name          = "Derek Heller"
  secret_id     = "...my_secret_id..."
  workspace_id  = "3fb2a63d-a091-47a6-951f-ac3e8ec69bab"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_key` (String, Sensitive) Your API Key. See <a href="https://developers.sendinblue.com/docs/getting-started">here</a>.


