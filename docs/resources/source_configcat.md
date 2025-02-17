---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_configcat Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceConfigcat Resource
---

# airbyte_source_configcat (Resource)

SourceConfigcat Resource

## Example Usage

```terraform
resource "airbyte_source_configcat" "my_source_configcat" {
  configuration = {
    password = "...my_password..."
    username = "Estrella_Wilkinson70"
  }
  definition_id = "81a7466b-f78b-43b7-9ede-547fc7c1cb53"
  name          = "Ms. Luis Harris"
  secret_id     = "...my_secret_id..."
  workspace_id  = "9ddb3b3d-7401-439d-82cf-2cb416442d85"
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

- `password` (String, Sensitive) Basic auth password. See <a href="https://api.configcat.com/docs/#section/Authentication">here</a>.
- `username` (String) Basic auth user name. See <a href="https://api.configcat.com/docs/#section/Authentication">here</a>.


