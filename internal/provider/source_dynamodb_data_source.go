// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceDynamodbDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceDynamodbDataSource{}

func NewSourceDynamodbDataSource() datasource.DataSource {
	return &SourceDynamodbDataSource{}
}

// SourceDynamodbDataSource is the data source implementation.
type SourceDynamodbDataSource struct {
	client *sdk.SDK
}

// SourceDynamodbDataSourceModel describes the data model.
type SourceDynamodbDataSourceModel struct {
	Configuration SourceDynamodb1 `tfsdk:"configuration"`
	Name          types.String    `tfsdk:"name"`
	SecretID      types.String    `tfsdk:"secret_id"`
	SourceID      types.String    `tfsdk:"source_id"`
	WorkspaceID   types.String    `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceDynamodbDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_dynamodb"
}

// Schema defines the schema for the data source.
func (r *SourceDynamodbDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceDynamodb DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						Computed:    true,
						Description: `The access key id to access Dynamodb. Airbyte requires read permissions to the database`,
					},
					"endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `the URL of the Dynamodb database`,
					},
					"region": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"",
								"us-east-1",
								"us-east-2",
								"us-west-1",
								"us-west-2",
								"af-south-1",
								"ap-east-1",
								"ap-south-1",
								"ap-northeast-1",
								"ap-northeast-2",
								"ap-northeast-3",
								"ap-southeast-1",
								"ap-southeast-2",
								"ca-central-1",
								"cn-north-1",
								"cn-northwest-1",
								"eu-central-1",
								"eu-north-1",
								"eu-south-1",
								"eu-west-1",
								"eu-west-2",
								"eu-west-3",
								"sa-east-1",
								"me-south-1",
								"us-gov-east-1",
								"us-gov-west-1",
							),
						},
						MarkdownDescription: `must be one of ["", "us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]` + "\n" +
							`The region of the Dynamodb database`,
					},
					"reserved_attribute_names": schema.StringAttribute{
						Computed:    true,
						Description: `Comma separated reserved attribute names present in your tables`,
					},
					"secret_access_key": schema.StringAttribute{
						Computed:    true,
						Description: `The corresponding secret to the access key id.`,
					},
					"source_type": schema.StringAttribute{
						Computed: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"dynamodb",
							),
						},
						Description: `must be one of ["dynamodb"]`,
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

func (r *SourceDynamodbDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceDynamodbDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceDynamodbDataSourceModel
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
	request := operations.GetSourceDynamodbRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceDynamodb(ctx, request)
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
