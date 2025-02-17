---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_github Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGithub Resource
---

# airbyte_source_github (Resource)

SourceGithub Resource

## Example Usage

```terraform
resource "airbyte_source_github" "my_source_github" {
  configuration = {
    api_url = "https://github.company.org"
    branch  = "airbytehq/airbyte/master airbytehq/airbyte/my-branch"
    branches = [
      "...",
    ]
    credentials = {
      o_auth = {
        access_token  = "...my_access_token..."
        client_id     = "...my_client_id..."
        client_secret = "...my_client_secret..."
      }
    }
    repositories = [
      "...",
    ]
    repository        = "airbytehq/airbyte"
    requests_per_hour = 10
    start_date        = "2021-03-01T00:00:00Z"
  }
  definition_id = "e017f905-2f20-440e-8692-82dd6a12cb01"
  name          = "Bennie Stroman"
  secret_id     = "...my_secret_id..."
  workspace_id  = "aeeda058-2852-4791-bedf-cf9c9058e69d"
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

- `credentials` (Attributes) Choose how to authenticate to GitHub (see [below for nested schema](#nestedatt--configuration--credentials))
- `repositories` (List of String) List of GitHub organizations/repositories, e.g. `airbytehq/airbyte` for single repository, `airbytehq/*` for get all repositories from organization and `airbytehq/airbyte airbytehq/another-repo` for multiple repositories.

Optional:

- `api_url` (String) Default: "https://api.github.com/"
Please enter your basic URL from self-hosted GitHub instance or leave it empty to use GitHub.
- `branch` (String) (DEPRCATED) Space-delimited list of GitHub repository branches to pull commits for, e.g. `airbytehq/airbyte/master`. If no branches are specified for a repository, the default branch will be pulled.
- `branches` (List of String) List of GitHub repository branches to pull commits for, e.g. `airbytehq/airbyte/master`. If no branches are specified for a repository, the default branch will be pulled.
- `repository` (String) (DEPRCATED) Space-delimited list of GitHub organizations/repositories, e.g. `airbytehq/airbyte` for single repository, `airbytehq/*` for get all repositories from organization and `airbytehq/airbyte airbytehq/another-repo` for multiple repositories.
- `requests_per_hour` (Number) The GitHub API allows for a maximum of 5000 requests per hour (15000 for Github Enterprise). You can specify a lower value to limit your use of the API quota.
- `start_date` (String) The date from which you'd like to replicate data from GitHub in the format YYYY-MM-DDT00:00:00Z. If the date is not set, all data will be replicated.  For the streams which support this configuration, only data generated on or after the start date will be replicated. This field doesn't apply to all streams, see the <a href="https://docs.airbyte.com/integrations/sources/github">docs</a> for more info

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `o_auth` (Attributes) Choose how to authenticate to GitHub (see [below for nested schema](#nestedatt--configuration--credentials--o_auth))
- `personal_access_token` (Attributes) Choose how to authenticate to GitHub (see [below for nested schema](#nestedatt--configuration--credentials--personal_access_token))

<a id="nestedatt--configuration--credentials--o_auth"></a>
### Nested Schema for `configuration.credentials.o_auth`

Required:

- `access_token` (String, Sensitive) OAuth access token

Optional:

- `client_id` (String) OAuth Client Id
- `client_secret` (String) OAuth Client secret


<a id="nestedatt--configuration--credentials--personal_access_token"></a>
### Nested Schema for `configuration.credentials.personal_access_token`

Required:

- `personal_access_token` (String, Sensitive) Log into GitHub and then generate a <a href="https://github.com/settings/tokens">personal access token</a>. To load balance your API quota consumption across multiple API tokens, input multiple tokens separated with ","


