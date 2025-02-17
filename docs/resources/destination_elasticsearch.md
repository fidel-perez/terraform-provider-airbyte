---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_elasticsearch Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationElasticsearch Resource
---

# airbyte_destination_elasticsearch (Resource)

DestinationElasticsearch Resource

## Example Usage

```terraform
resource "airbyte_destination_elasticsearch" "my_destination_elasticsearch" {
  configuration = {
    authentication_method = {
      api_key_secret = {
        api_key_id     = "...my_api_key_id..."
        api_key_secret = "...my_api_key_secret..."
      }
    }
    ca_certificate = "...my_ca_certificate..."
    endpoint       = "...my_endpoint..."
    upsert         = false
  }
  definition_id = "da65ed46-5e75-48af-92ad-38ed7ed0e5e2"
  name          = "Katherine Considine"
  workspace_id  = "7d0e4e50-95ed-494b-8ecb-397d064562ef"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the destination e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `endpoint` (String) The full url of the Elasticsearch server

Optional:

- `authentication_method` (Attributes) The type of authentication to be used (see [below for nested schema](#nestedatt--configuration--authentication_method))
- `ca_certificate` (String) CA certificate
- `upsert` (Boolean) Default: true
If a primary key identifier is defined in the source, an upsert will be performed using the primary key value as the elasticsearch doc id. Does not support composite primary keys.

<a id="nestedatt--configuration--authentication_method"></a>
### Nested Schema for `configuration.authentication_method`

Optional:

- `api_key_secret` (Attributes) Use a api key and secret combination to authenticate (see [below for nested schema](#nestedatt--configuration--authentication_method--api_key_secret))
- `username_password` (Attributes) Basic auth header with a username and password (see [below for nested schema](#nestedatt--configuration--authentication_method--username_password))

<a id="nestedatt--configuration--authentication_method--api_key_secret"></a>
### Nested Schema for `configuration.authentication_method.api_key_secret`

Required:

- `api_key_id` (String) The Key ID to used when accessing an enterprise Elasticsearch instance.
- `api_key_secret` (String) The secret associated with the API Key ID.


<a id="nestedatt--configuration--authentication_method--username_password"></a>
### Nested Schema for `configuration.authentication_method.username_password`

Required:

- `password` (String, Sensitive) Basic auth password to access a secure Elasticsearch server
- `username` (String) Basic auth username to access a secure Elasticsearch server


