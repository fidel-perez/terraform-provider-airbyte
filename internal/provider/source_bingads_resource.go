// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"

	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceBingAdsResource{}
var _ resource.ResourceWithImportState = &SourceBingAdsResource{}

func NewSourceBingAdsResource() resource.Resource {
	return &SourceBingAdsResource{}
}

// SourceBingAdsResource defines the resource implementation.
type SourceBingAdsResource struct {
	client *sdk.SDK
}

// SourceBingAdsResourceModel describes the resource data model.
type SourceBingAdsResourceModel struct {
	Configuration SourceBingAds `tfsdk:"configuration"`
	DefinitionID  types.String  `tfsdk:"definition_id"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	SourceType    types.String  `tfsdk:"source_type"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

func (r *SourceBingAdsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_bing_ads"
}

func (r *SourceBingAdsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceBingAds Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: `The Client ID of your Microsoft Advertising developer application.`,
					},
					"client_secret": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: `Default: ""` + "\n" +
							`The Client Secret of your Microsoft Advertising developer application.`,
					},
					"custom_reports": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Required:    true,
									Description: `The name of the custom report, this name would be used as stream name`,
								},
								"report_aggregation": schema.StringAttribute{
									Optional: true,
									MarkdownDescription: `Default: "[Hourly]"` + "\n" +
										`A list of available aggregations.`,
								},
								"report_columns": schema.ListAttribute{
									Required:    true,
									ElementType: types.StringType,
									Description: `A list of available report object columns. You can find it in description of reporting object that you want to add to custom report.`,
								},
								"reporting_object": schema.StringAttribute{
									Required: true,
									MarkdownDescription: `must be one of ["AccountPerformanceReportRequest", "AdDynamicTextPerformanceReportRequest", "AdExtensionByAdReportRequest", "AdExtensionByKeywordReportRequest", "AdExtensionDetailReportRequest", "AdGroupPerformanceReportRequest", "AdPerformanceReportRequest", "AgeGenderAudienceReportRequest", "AudiencePerformanceReportRequest", "CallDetailReportRequest", "CampaignPerformanceReportRequest", "ConversionPerformanceReportRequest", "DestinationUrlPerformanceReportRequest", "DSAAutoTargetPerformanceReportRequest", "DSACategoryPerformanceReportRequest", "DSASearchQueryPerformanceReportRequest", "GeographicPerformanceReportRequest", "GoalsAndFunnelsReportRequest", "HotelDimensionPerformanceReportRequest", "HotelGroupPerformanceReportRequest", "KeywordPerformanceReportRequest", "NegativeKeywordConflictReportRequest", "ProductDimensionPerformanceReportRequest", "ProductMatchCountReportRequest", "ProductNegativeKeywordConflictReportRequest", "ProductPartitionPerformanceReportRequest", "ProductPartitionUnitPerformanceReportRequest", "ProductSearchQueryPerformanceReportRequest", "ProfessionalDemographicsAudienceReportRequest", "PublisherUsagePerformanceReportRequest", "SearchCampaignChangeHistoryReportRequest", "SearchQueryPerformanceReportRequest", "ShareOfVoiceReportRequest", "UserLocationPerformanceReportRequest"]` + "\n" +
										`The name of the the object derives from the ReportRequest object. You can find it in Bing Ads Api docs - Reporting API - Reporting Data Objects.`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"AccountPerformanceReportRequest",
											"AdDynamicTextPerformanceReportRequest",
											"AdExtensionByAdReportRequest",
											"AdExtensionByKeywordReportRequest",
											"AdExtensionDetailReportRequest",
											"AdGroupPerformanceReportRequest",
											"AdPerformanceReportRequest",
											"AgeGenderAudienceReportRequest",
											"AudiencePerformanceReportRequest",
											"CallDetailReportRequest",
											"CampaignPerformanceReportRequest",
											"ConversionPerformanceReportRequest",
											"DestinationUrlPerformanceReportRequest",
											"DSAAutoTargetPerformanceReportRequest",
											"DSACategoryPerformanceReportRequest",
											"DSASearchQueryPerformanceReportRequest",
											"GeographicPerformanceReportRequest",
											"GoalsAndFunnelsReportRequest",
											"HotelDimensionPerformanceReportRequest",
											"HotelGroupPerformanceReportRequest",
											"KeywordPerformanceReportRequest",
											"NegativeKeywordConflictReportRequest",
											"ProductDimensionPerformanceReportRequest",
											"ProductMatchCountReportRequest",
											"ProductNegativeKeywordConflictReportRequest",
											"ProductPartitionPerformanceReportRequest",
											"ProductPartitionUnitPerformanceReportRequest",
											"ProductSearchQueryPerformanceReportRequest",
											"ProfessionalDemographicsAudienceReportRequest",
											"PublisherUsagePerformanceReportRequest",
											"SearchCampaignChangeHistoryReportRequest",
											"SearchQueryPerformanceReportRequest",
											"ShareOfVoiceReportRequest",
											"UserLocationPerformanceReportRequest",
										),
									},
								},
							},
						},
						Description: `You can add your Custom Bing Ads report by creating one.`,
					},
					"developer_token": schema.StringAttribute{
						Required:    true,
						Sensitive:   true,
						Description: `Developer token associated with user. See more info <a href="https://docs.microsoft.com/en-us/advertising/guides/get-started?view=bingads-13#get-developer-token"> in the docs</a>.`,
					},
					"lookback_window": schema.Int64Attribute{
						Optional: true,
						MarkdownDescription: `Default: 0` + "\n" +
							`Also known as attribution or conversion window. How far into the past to look for records (in days). If your conversion window has an hours/minutes granularity, round it up to the number of days exceeding. Used only for performance report streams in incremental mode without specified Reports Start Date.`,
					},
					"refresh_token": schema.StringAttribute{
						Required:    true,
						Sensitive:   true,
						Description: `Refresh Token to renew the expired Access Token.`,
					},
					"reports_start_date": schema.StringAttribute{
						Optional:    true,
						Description: `The start date from which to begin replicating report data. Any data generated before this date will not be replicated in reports. This is a UTC date in YYYY-MM-DD format. If not set, data from previous and current calendar year will be replicated.`,
						Validators: []validator.String{
							validators.IsValidDate(),
						},
					},
					"tenant_id": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: `Default: "common"` + "\n" +
							`The Tenant ID of your Microsoft Advertising developer application. Set this to "common" unless you know you need a different value.`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required:    true,
				Description: `Name of the source e.g. dev-mysql-instance.`,
			},
			"secret_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"source_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
		},
	}
}

func (r *SourceBingAdsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceBingAdsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceBingAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToCreateSDKType()
	res, err := r.client.Sources.CreateSourceBingAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceBingAdsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceBingAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.GetSourceBingAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceBingAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceBingAdsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceBingAdsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceBingAdsPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceBingAdsRequest{
		SourceBingAdsPutRequest: sourceBingAdsPutRequest,
		SourceID:                sourceID,
	}
	res, err := r.client.Sources.PutSourceBingAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceBingAdsRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceBingAds(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if getResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", getResponse))
		return
	}
	if getResponse.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", getResponse.StatusCode), debugResponse(getResponse.RawResponse))
		return
	}
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceBingAdsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceBingAdsResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	sourceID := data.SourceID.ValueString()
	request := operations.DeleteSourceBingAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceBingAds(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *SourceBingAdsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("source_id"), req.ID)...)
}
