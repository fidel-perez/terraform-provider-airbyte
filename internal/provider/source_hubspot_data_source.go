// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceHubspotDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceHubspotDataSource{}

func NewSourceHubspotDataSource() datasource.DataSource {
	return &SourceHubspotDataSource{}
}

// SourceHubspotDataSource is the data source implementation.
type SourceHubspotDataSource struct {
	client *sdk.SDK
}

// SourceHubspotDataSourceModel describes the data model.
type SourceHubspotDataSourceModel struct {
	Configuration SourceHubspot `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceHubspotDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_hubspot"
}

// Schema defines the schema for the data source.
func (r *SourceHubspotDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceHubspot DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_hubspot_authentication_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_authentication_private_app": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_update_authentication_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
							"source_hubspot_update_authentication_private_app": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.`,
									},
								},
								Description: `Choose how to authenticate to HubSpot.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Choose how to authenticate to HubSpot.`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.`,
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret_id": schema.StringAttribute{
				Optional:    true,
				Description: `Optional secretID obtained through the public API OAuth redirect flow.`,
			},
			"source_id": schema.StringAttribute{
				Required: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *SourceHubspotDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SourceHubspotDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceHubspotDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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
	request := operations.GetSourceHubspotRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceHubspot(ctx, request)
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
