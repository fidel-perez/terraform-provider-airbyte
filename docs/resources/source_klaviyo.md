---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_klaviyo Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceKlaviyo Resource
---

# airbyte_source_klaviyo (Resource)

SourceKlaviyo Resource

## Example Usage

```terraform
resource "airbyte_source_klaviyo" "my_source_klaviyo" {
  configuration = {
    api_key    = "...my_api_key..."
    start_date = "2017-01-25T00:00:00Z"
  }
  definition_id = "d98f81ed-eee1-4be4-a723-eeaf419bc59e"
  name          = "Joanne Murray"
  secret_id     = "...my_secret_id..."
  workspace_id  = "9e9d149f-3b04-4e32-9c64-9b6bc8e2c7d0"
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

- `api_key` (String, Sensitive) Klaviyo API Key. See our <a href="https://docs.airbyte.com/integrations/sources/klaviyo">docs</a> if you need help finding this key.

Optional:

- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. This field is optional - if not provided, all data will be replicated.


