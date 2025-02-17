---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_iterable Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceIterable Resource
---

# airbyte_source_iterable (Resource)

SourceIterable Resource

## Example Usage

```terraform
resource "airbyte_source_iterable" "my_source_iterable" {
  configuration = {
    api_key    = "...my_api_key..."
    start_date = "2021-04-01T00:00:00Z"
  }
  definition_id = "00977793-827c-406d-986b-4fbde6ae5395"
  name          = "Katherine Bashirian"
  secret_id     = "...my_secret_id..."
  workspace_id  = "d8df8fdd-acae-4826-9af8-b9bb4850d654"
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

- `api_key` (String, Sensitive) Iterable API Key. See the <a href="https://docs.airbyte.com/integrations/sources/iterable">docs</a> for more information on how to obtain this key.
- `start_date` (String) The date from which you'd like to replicate data for Iterable, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.


