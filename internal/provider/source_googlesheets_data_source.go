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
var _ datasource.DataSource = &SourceGoogleSheetsDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceGoogleSheetsDataSource{}

func NewSourceGoogleSheetsDataSource() datasource.DataSource {
	return &SourceGoogleSheetsDataSource{}
}

// SourceGoogleSheetsDataSource is the data source implementation.
type SourceGoogleSheetsDataSource struct {
	client *sdk.SDK
}

// SourceGoogleSheetsDataSourceModel describes the data model.
type SourceGoogleSheetsDataSourceModel struct {
	Configuration SourceGoogleSheets `tfsdk:"configuration"`
	Name          types.String       `tfsdk:"name"`
	SecretID      types.String       `tfsdk:"secret_id"`
	SourceID      types.String       `tfsdk:"source_id"`
	WorkspaceID   types.String       `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceGoogleSheetsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_google_sheets"
}

// Schema defines the schema for the data source.
func (r *SourceGoogleSheetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceGoogleSheets DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"source_google_sheets_authentication_authenticate_via_google_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of ["Client"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's Client Secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's refresh token`,
									},
								},
								Description: `Credentials for connecting to the Google Sheets API`,
							},
							"source_google_sheets_authentication_service_account_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
										Description: `must be one of ["Service"]`,
									},
									"service_account_info": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format`,
									},
								},
								Description: `Credentials for connecting to the Google Sheets API`,
							},
							"source_google_sheets_update_authentication_authenticate_via_google_o_auth": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Client",
											),
										},
										Description: `must be one of ["Client"]`,
									},
									"client_id": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's Client ID`,
									},
									"client_secret": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's Client Secret`,
									},
									"refresh_token": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google application's refresh token`,
									},
								},
								Description: `Credentials for connecting to the Google Sheets API`,
							},
							"source_google_sheets_update_authentication_service_account_key_authentication": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"auth_type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Service",
											),
										},
										Description: `must be one of ["Service"]`,
									},
									"service_account_info": schema.StringAttribute{
										Computed:    true,
										Description: `Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format`,
									},
								},
								Description: `Credentials for connecting to the Google Sheets API`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Credentials for connecting to the Google Sheets API`,
					},
					"row_batch_size": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of rows fetched when making a Google Sheet API call. Defaults to 200.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"google-sheets",
							),
						},
						Description: `must be one of ["google-sheets"]`,
					},
					"spreadsheet_id": schema.StringAttribute{
						Computed:    true,
						Description: `Enter the link to the Google spreadsheet you want to sync`,
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

func (r *SourceGoogleSheetsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceGoogleSheetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceGoogleSheetsDataSourceModel
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
	request := operations.GetSourceGoogleSheetsRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceGoogleSheets(ctx, request)
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
