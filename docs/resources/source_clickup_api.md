---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_clickup_api Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceClickupAPI Resource
---

# airbyte_source_clickup_api (Resource)

SourceClickupAPI Resource

## Example Usage

```terraform
resource "airbyte_source_clickup_api" "my_source_clickupapi" {
  configuration = {
    api_token            = "...my_api_token..."
    folder_id            = "...my_folder_id..."
    include_closed_tasks = true
    list_id              = "...my_list_id..."
    space_id             = "...my_space_id..."
    team_id              = "...my_team_id..."
  }
  definition_id = "517f0e32-c2e3-402e-ade9-2b3e43098446"
  name          = "Freddie Little"
  secret_id     = "...my_secret_id..."
  workspace_id  = "e6422d15-b828-4621-a877-d2e625cdd80b"
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

- `api_token` (String, Sensitive) Every ClickUp API call required authentication. This field is your personal API token. See <a href="https://clickup.com/api/developer-portal/authentication/#personal-token">here</a>.

Optional:

- `folder_id` (String) The ID of your folder in your space. Retrieve it from the `/space/{space_id}/folder` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetFolders/">here</a>.
- `include_closed_tasks` (Boolean) Default: false
Include or exclude closed tasks. By default, they are excluded. See <a https://clickup.com/api/clickupreference/operation/GetTasks/#!in=query&path=include_closed&t=request">here</a>.
- `list_id` (String) The ID of your list in your folder. Retrieve it from the `/folder/{folder_id}/list` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetLists/">here</a>.
- `space_id` (String) The ID of your space in your workspace. Retrieve it from the `/team/{team_id}/space` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetSpaces/">here</a>.
- `team_id` (String) The ID of your team in ClickUp. Retrieve it from the `/team` of the ClickUp API. See <a href="https://clickup.com/api/clickupreference/operation/GetAuthorizedTeams/">here</a>.


