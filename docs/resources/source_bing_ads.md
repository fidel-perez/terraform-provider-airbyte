---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_bing_ads Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBingAds Resource
---

# airbyte_source_bing_ads (Resource)

SourceBingAds Resource

## Example Usage

```terraform
resource "airbyte_source_bing_ads" "my_source_bingads" {
  configuration = {
    client_id     = "...my_client_id..."
    client_secret = "...my_client_secret..."
    custom_reports = [
      {
        name               = "AdDynamicTextPerformanceReport"
        report_aggregation = "...my_report_aggregation..."
        report_columns = [
          "...",
        ]
        reporting_object = "ShareOfVoiceReportRequest"
      },
    ]
    developer_token    = "...my_developer_token..."
    lookback_window    = 3
    refresh_token      = "...my_refresh_token..."
    reports_start_date = "2022-08-17"
    tenant_id          = "...my_tenant_id..."
  }
  definition_id = "f49be625-99f1-47b5-861c-8d2f7dd6ee9c"
  name          = "Delia Kub Sr."
  secret_id     = "...my_secret_id..."
  workspace_id  = "90282195-430f-4896-8a32-1f431fb3aad0"
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

- `client_id` (String) The Client ID of your Microsoft Advertising developer application.
- `developer_token` (String, Sensitive) Developer token associated with user. See more info <a href="https://docs.microsoft.com/en-us/advertising/guides/get-started?view=bingads-13#get-developer-token"> in the docs</a>.
- `refresh_token` (String, Sensitive) Refresh Token to renew the expired Access Token.

Optional:

- `client_secret` (String) Default: ""
The Client Secret of your Microsoft Advertising developer application.
- `custom_reports` (Attributes List) You can add your Custom Bing Ads report by creating one. (see [below for nested schema](#nestedatt--configuration--custom_reports))
- `lookback_window` (Number) Default: 0
Also known as attribution or conversion window. How far into the past to look for records (in days). If your conversion window has an hours/minutes granularity, round it up to the number of days exceeding. Used only for performance report streams in incremental mode without specified Reports Start Date.
- `reports_start_date` (String) The start date from which to begin replicating report data. Any data generated before this date will not be replicated in reports. This is a UTC date in YYYY-MM-DD format. If not set, data from previous and current calendar year will be replicated.
- `tenant_id` (String) Default: "common"
The Tenant ID of your Microsoft Advertising developer application. Set this to "common" unless you know you need a different value.

<a id="nestedatt--configuration--custom_reports"></a>
### Nested Schema for `configuration.custom_reports`

Required:

- `name` (String) The name of the custom report, this name would be used as stream name
- `report_columns` (List of String) A list of available report object columns. You can find it in description of reporting object that you want to add to custom report.
- `reporting_object` (String) must be one of ["AccountPerformanceReportRequest", "AdDynamicTextPerformanceReportRequest", "AdExtensionByAdReportRequest", "AdExtensionByKeywordReportRequest", "AdExtensionDetailReportRequest", "AdGroupPerformanceReportRequest", "AdPerformanceReportRequest", "AgeGenderAudienceReportRequest", "AudiencePerformanceReportRequest", "CallDetailReportRequest", "CampaignPerformanceReportRequest", "ConversionPerformanceReportRequest", "DestinationUrlPerformanceReportRequest", "DSAAutoTargetPerformanceReportRequest", "DSACategoryPerformanceReportRequest", "DSASearchQueryPerformanceReportRequest", "GeographicPerformanceReportRequest", "GoalsAndFunnelsReportRequest", "HotelDimensionPerformanceReportRequest", "HotelGroupPerformanceReportRequest", "KeywordPerformanceReportRequest", "NegativeKeywordConflictReportRequest", "ProductDimensionPerformanceReportRequest", "ProductMatchCountReportRequest", "ProductNegativeKeywordConflictReportRequest", "ProductPartitionPerformanceReportRequest", "ProductPartitionUnitPerformanceReportRequest", "ProductSearchQueryPerformanceReportRequest", "ProfessionalDemographicsAudienceReportRequest", "PublisherUsagePerformanceReportRequest", "SearchCampaignChangeHistoryReportRequest", "SearchQueryPerformanceReportRequest", "ShareOfVoiceReportRequest", "UserLocationPerformanceReportRequest"]
The name of the the object derives from the ReportRequest object. You can find it in Bing Ads Api docs - Reporting API - Reporting Data Objects.

Optional:

- `report_aggregation` (String) Default: "[Hourly]"
A list of available aggregations.


