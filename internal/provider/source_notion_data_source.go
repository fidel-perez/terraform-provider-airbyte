// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceNotionDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceNotionDataSource{}

func NewSourceNotionDataSource() datasource.DataSource {
	return &SourceNotionDataSource{}
}

// SourceNotionDataSource is the data source implementation.
type SourceNotionDataSource struct {
	client *sdk.SDK
}

// SourceNotionDataSourceModel describes the data model.
type SourceNotionDataSourceModel struct {
	Configuration SourceNotion `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceNotionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_notion"
}

// Schema defines the schema for the data source.
func (r *SourceNotionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceNotion DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_notion_authenticate_using_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"token",
											),
										},
										Description: `must be one of ["token"]`,
									},
									"token": schema.StringAttribute{
										Computed:    true,
										Description: `Notion API access token, see the <a href="https://developers.notion.com/docs/authorization">docs</a> for more information on how to obtain this token.`,
									},
								},
								Description: `Pick an authentication method.`,
							},
							"source_notion_authenticate_using_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token is a token you received by complete the OauthWebFlow of Notion.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of ["OAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The ClientID of your Notion integration.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The ClientSecret of your Notion integration.`,
									},
								},
								Description: `Pick an authentication method.`,
							},
							"source_notion_update_authenticate_using_access_token": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"token",
											),
										},
										Description: `must be one of ["token"]`,
									},
									"token": schema.StringAttribute{
										Computed:    true,
										Description: `Notion API access token, see the <a href="https://developers.notion.com/docs/authorization">docs</a> for more information on how to obtain this token.`,
									},
								},
								Description: `Pick an authentication method.`,
							},
							"source_notion_update_authenticate_using_o_auth2_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"access_token": schema.StringAttribute{
										Computed:    true,
										Description: `Access Token is a token you received by complete the OauthWebFlow of Notion.`,
									},
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"OAuth2.0",
											),
										},
										Description: `must be one of ["OAuth2.0"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `The ClientID of your Notion integration.`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The ClientSecret of your Notion integration.`,
									},
								},
								Description: `Pick an authentication method.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Pick an authentication method.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"notion",
							),
						},
						Description: `must be one of ["notion"]`,
					},
					"start_date": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							validators.IsRFC3339(),
						},
						Description: `UTC date and time in the format 2017-01-25T00:00:00.000Z. Any data before this date will not be replicated.`,
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

func (r *SourceNotionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceNotionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceNotionDataSourceModel
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
	request := operations.GetSourceNotionRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceNotion(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
