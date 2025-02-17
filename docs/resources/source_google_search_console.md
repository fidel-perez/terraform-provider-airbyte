---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_search_console Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGoogleSearchConsole Resource
---

# airbyte_source_google_search_console (Resource)

SourceGoogleSearchConsole Resource

## Example Usage

```terraform
resource "airbyte_source_google_search_console" "my_source_googlesearchconsole" {
  configuration = {
    authorization = {
      source_google_search_console_o_auth = {
        access_token  = "...my_access_token..."
        client_id     = "...my_client_id..."
        client_secret = "...my_client_secret..."
        refresh_token = "...my_refresh_token..."
      }
    }
    custom_reports = "...my_custom_reports..."
    custom_reports_array = [
      {
        dimensions = [
          "device",
        ]
        name = "Ms. Randy Gorczany V"
      },
    ]
    data_state = "all"
    end_date   = "2021-12-12"
    site_urls = [
      "...",
    ]
    start_date = "2020-03-18"
  }
  definition_id = "165bc484-0e7f-4b5d-b254-77f370b0ec7c"
  name          = "Wendell Rempel"
  secret_id     = "...my_secret_id..."
  workspace_id  = "0cb9d8df-c27a-48c7-ac3e-b5dc55714db0"
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

- `authorization` (Attributes) (see [below for nested schema](#nestedatt--configuration--authorization))
- `site_urls` (List of String) The URLs of the website property attached to your GSC account. Learn more about properties <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.

Optional:

- `custom_reports` (String) (DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our <a href='https://docs.airbyte.com/integrations/sources/google-search-console'>documentation</a> for more information on formulating custom reports.
- `custom_reports_array` (Attributes List) You can add your Custom Analytics report by creating one. (see [below for nested schema](#nestedatt--configuration--custom_reports_array))
- `data_state` (String) must be one of ["final", "all"]; Default: "final"
If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
- `end_date` (String) UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.
- `start_date` (String) Default: "2021-01-01"
UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.

<a id="nestedatt--configuration--authorization"></a>
### Nested Schema for `configuration.authorization`

Optional:

- `o_auth` (Attributes) (see [below for nested schema](#nestedatt--configuration--authorization--o_auth))
- `service_account_key_authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--authorization--service_account_key_authentication))

<a id="nestedatt--configuration--authorization--o_auth"></a>
### Nested Schema for `configuration.authorization.o_auth`

Required:

- `client_id` (String) The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
- `client_secret` (String) The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
- `refresh_token` (String, Sensitive) The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.

Optional:

- `access_token` (String, Sensitive) Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.


<a id="nestedatt--configuration--authorization--service_account_key_authentication"></a>
### Nested Schema for `configuration.authorization.service_account_key_authentication`

Required:

- `email` (String) The email of the user which has permissions to access the Google Workspace Admin APIs.
- `service_account_info` (String) The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.



<a id="nestedatt--configuration--custom_reports_array"></a>
### Nested Schema for `configuration.custom_reports_array`

Required:

- `dimensions` (List of String) A list of available dimensions. Please note, that for technical reasons `date` is the default dimension which will be included in your query whether you specify it or not. Primary key will consist of your custom dimensions and the default dimension along with `site_url` and `search_type`.
- `name` (String) The name of the custom report, this name would be used as stream name


