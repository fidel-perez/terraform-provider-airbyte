---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_convex Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceConvex Resource
---

# airbyte_source_convex (Resource)

SourceConvex Resource

## Example Usage

```terraform
resource "airbyte_source_convex" "my_source_convex" {
  configuration = {
    access_key     = "...my_access_key..."
    deployment_url = "https://murky-swan-635.convex.cloud"
  }
  definition_id = "581ee677-0fa8-4ec1-ba80-4bd6457a40e8"
  name          = "Corey Braun"
  secret_id     = "...my_secret_id..."
  workspace_id  = "541ba6f5-d90d-45a8-a349-e2072bdff381"
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

- `access_key` (String, Sensitive) API access key used to retrieve data from Convex.
- `deployment_url` (String)


