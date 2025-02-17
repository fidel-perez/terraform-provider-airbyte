---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_launchdarkly Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceLaunchdarkly Resource
---

# airbyte_source_launchdarkly (Resource)

SourceLaunchdarkly Resource

## Example Usage

```terraform
resource "airbyte_source_launchdarkly" "my_source_launchdarkly" {
  configuration = {
    access_token = "...my_access_token..."
  }
  definition_id = "422849b5-8575-49fd-b9d7-4aa20ea69f1b"
  name          = "Jodi Marquardt"
  secret_id     = "...my_secret_id..."
  workspace_id  = "dd1b5a02-95b1-497b-bb02-27d625c3155f"
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

- `access_token` (String, Sensitive) Your Access token. See <a href="https://apidocs.launchdarkly.com/#section/Overview/Authentication">here</a>.


