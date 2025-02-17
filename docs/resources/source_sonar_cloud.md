---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sonar_cloud Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSonarCloud Resource
---

# airbyte_source_sonar_cloud (Resource)

SourceSonarCloud Resource

## Example Usage

```terraform
resource "airbyte_source_sonar_cloud" "my_source_sonarcloud" {
  configuration = {
    component_keys = [
      "{ \"see\": \"documentation\" }",
    ]
    end_date     = "YYYY-MM-DD"
    organization = "airbyte"
    start_date   = "YYYY-MM-DD"
    user_token   = "...my_user_token..."
  }
  definition_id = "d259943d-fa52-4a9e-875a-bffba2c1e7b6"
  name          = "Jose Lindgren"
  secret_id     = "...my_secret_id..."
  workspace_id  = "d761f19b-60aa-4080-8c97-1e60235dc09f"
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

- `component_keys` (List of String, Sensitive) Comma-separated list of component keys.
- `organization` (String) Organization key. See <a href="https://docs.sonarcloud.io/appendices/project-information/#project-and-organization-keys">here</a>.
- `user_token` (String, Sensitive) Your User Token. See <a href="https://docs.sonarcloud.io/advanced-setup/user-accounts/">here</a>. The token is case sensitive.

Optional:

- `end_date` (String) To retrieve issues created before the given date (inclusive).
- `start_date` (String) To retrieve issues created after the given date (inclusive).


