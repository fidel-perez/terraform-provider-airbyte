// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SourceAmazonAdsResource{}
var _ resource.ResourceWithImportState = &SourceAmazonAdsResource{}

func NewSourceAmazonAdsResource() resource.Resource {
	return &SourceAmazonAdsResource{}
}

// SourceAmazonAdsResource defines the resource implementation.
type SourceAmazonAdsResource struct {
	client *sdk.SDK
}

// SourceAmazonAdsResourceModel describes the resource data model.
type SourceAmazonAdsResourceModel struct {
	Configuration SourceAmazonAds `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	SourceType    types.String    `tfsdk:"source_type"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

func (r *SourceAmazonAdsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_amazon_ads"
}

func (r *SourceAmazonAdsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceAmazonAds Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Required:    true,
						Description: `The client ID of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.`,
					},
					"client_secret": schema.StringAttribute{
						Required:    true,
						Description: `The client secret of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.`,
					},
					"look_back_window": schema.Int64Attribute{
						Optional: true,
						MarkdownDescription: `Default: 3` + "\n" +
							`The amount of days to go back in time to get the updated data from Amazon Ads`,
					},
					"marketplace_ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `Marketplace IDs you want to fetch data for. Note: If Profile IDs are also selected, profiles will be selected if they match the Profile ID OR the Marketplace ID.`,
					},
					"profiles": schema.ListAttribute{
						Optional:    true,
						ElementType: types.Int64Type,
						Description: `Profile IDs you want to fetch data for. See <a href="https://advertising.amazon.com/API/docs/en-us/concepts/authorization/profiles">docs</a> for more details. Note: If Marketplace IDs are also selected, profiles will be selected if they match the Profile ID OR the Marketplace ID.`,
					},
					"refresh_token": schema.StringAttribute{
						Required:    true,
						Description: `Amazon Ads refresh token. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens">docs</a> for more information on how to obtain this token.`,
					},
					"region": schema.StringAttribute{
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"NA",
								"EU",
								"FE",
							),
						},
						MarkdownDescription: `must be one of ["NA", "EU", "FE"]; Default: "NA"` + "\n" +
							`Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.`,
					},
					"report_record_types": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `Optional configuration which accepts an array of string of record types. Leave blank for default behaviour to pull all report types. Use this config option only if you want to pull specific report type(s). See <a href="https://advertising.amazon.com/API/docs/en-us/reporting/v2/report-types">docs</a> for more details`,
					},
					"start_date": schema.StringAttribute{
						Optional:    true,
						Description: `The Start date for collecting reports, should not be more than 60 days in the past. In YYYY-MM-DD format`,
					},
					"state_filter": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `Reflects the state of the Display, Product, and Brand Campaign streams as enabled, paused, or archived. If you do not populate this field, it will be ignored completely.`,
					},
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required: true,
			},
			"secret_id": schema.StringAttribute{
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

func (r *SourceAmazonAdsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceAmazonAdsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceAmazonAdsResourceModel
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
	res, err := r.client.Sources.CreateSourceAmazonAds(ctx, request)
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

func (r *SourceAmazonAdsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceAmazonAdsResourceModel
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
	request := operations.GetSourceAmazonAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceAmazonAds(ctx, request)
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

func (r *SourceAmazonAdsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceAmazonAdsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceAmazonAdsPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceAmazonAdsRequest{
		SourceAmazonAdsPutRequest: sourceAmazonAdsPutRequest,
		SourceID:                  sourceID,
	}
	res, err := r.client.Sources.PutSourceAmazonAds(ctx, request)
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
	getRequest := operations.GetSourceAmazonAdsRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceAmazonAds(ctx, getRequest)
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

func (r *SourceAmazonAdsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceAmazonAdsResourceModel
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
	request := operations.DeleteSourceAmazonAdsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceAmazonAds(ctx, request)
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

func (r *SourceAmazonAdsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
