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
var _ datasource.DataSource = &DestinationElasticsearchDataSource{}
var _ datasource.DataSourceWithConfigure = &DestinationElasticsearchDataSource{}

func NewDestinationElasticsearchDataSource() datasource.DataSource {
	return &DestinationElasticsearchDataSource{}
}

// DestinationElasticsearchDataSource is the data source implementation.
type DestinationElasticsearchDataSource struct {
	client *sdk.SDK
}

// DestinationElasticsearchDataSourceModel describes the data model.
type DestinationElasticsearchDataSourceModel struct {
	Configuration DestinationElasticsearch `tfsdk:"configuration"`
	DestinationID types.String             `tfsdk:"destination_id"`
	Name          types.String             `tfsdk:"name"`
	WorkspaceID   types.String             `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *DestinationElasticsearchDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_elasticsearch"
}

// Schema defines the schema for the data source.
func (r *DestinationElasticsearchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationElasticsearch DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"authentication_method": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"destination_elasticsearch_authentication_method_api_key_secret": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Key ID to used when accessing an enterprise Elasticsearch instance.`,
									},
									"api_key_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The secret associated with the API Key ID.`,
									},
								},
								Description: `Use a api key and secret combination to authenticate`,
							},
							"destination_elasticsearch_authentication_method_username_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Basic auth password to access a secure Elasticsearch server`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Description: `Basic auth username to access a secure Elasticsearch server`,
									},
								},
								Description: `Basic auth header with a username and password`,
							},
							"destination_elasticsearch_update_authentication_method_api_key_secret": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"api_key_id": schema.StringAttribute{
										Computed:    true,
										Description: `The Key ID to used when accessing an enterprise Elasticsearch instance.`,
									},
									"api_key_secret": schema.StringAttribute{
										Computed:    true,
										Description: `The secret associated with the API Key ID.`,
									},
								},
								Description: `Use a api key and secret combination to authenticate`,
							},
							"destination_elasticsearch_update_authentication_method_username_password": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Computed:    true,
										Description: `Basic auth password to access a secure Elasticsearch server`,
									},
									"username": schema.StringAttribute{
										Computed:    true,
										Description: `Basic auth username to access a secure Elasticsearch server`,
									},
								},
								Description: `Basic auth header with a username and password`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The type of authentication to be used`,
					},
					"ca_certificate": schema.StringAttribute{
						Computed:    true,
						Description: `CA certificate`,
					},
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `The full url of the Elasticsearch server`,
					},
					"upsert": schema.BoolAttribute{
						Computed: true,
						MarkdownDescription: `Default: true` + "\n" +
							`If a primary key identifier is defined in the source, an upsert will be performed using the primary key value as the elasticsearch doc id. Does not support composite primary keys.`,
					},
				},
			},
			"destination_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *DestinationElasticsearchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *DestinationElasticsearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DestinationElasticsearchDataSourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationElasticsearchRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationElasticsearch(ctx, request)
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
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
