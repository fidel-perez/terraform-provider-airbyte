---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_convex Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationConvex DataSource
---

# airbyte_destination_convex (Data Source)

DestinationConvex DataSource

## Example Usage

```terraform
data "airbyte_destination_convex" "my_destination_convex" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (String) Parsed as JSON.
The values required to configure the destination.
- `destination_type` (String)
- `name` (String)
- `workspace_id` (String)


