// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"airbyte/internal/sdk/pkg/models/operations"
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SourceQualarooDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceQualarooDataSource{}

func NewSourceQualarooDataSource() datasource.DataSource {
	return &SourceQualarooDataSource{}
}

// SourceQualarooDataSource is the data source implementation.
type SourceQualarooDataSource struct {
	client *sdk.SDK
}

// SourceQualarooDataSourceModel describes the data model.
type SourceQualarooDataSourceModel struct {
	Configuration SourceQualaroo `tfsdk:"configuration"`
	Name          types.String   `tfsdk:"name"`
	SecretID      types.String   `tfsdk:"secret_id"`
	SourceID      types.String   `tfsdk:"source_id"`
	WorkspaceID   types.String   `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceQualarooDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_qualaroo"
}

// Schema defines the schema for the data source.
func (r *SourceQualarooDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceQualaroo DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Computed:    true,
						Description: `A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.`,
					},
					"start_date": schema.StringAttribute{
						Computed:    true,
						Description: `UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.`,
					},
					"survey_ids": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `IDs of the surveys from which you'd like to replicate data. If left empty, data from all surveys to which you have access will be replicated.`,
					},
					"token": schema.StringAttribute{
						Computed:    true,
						Description: `A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.`,
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

func (r *SourceQualarooDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceQualarooDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceQualarooDataSourceModel
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
	request := operations.GetSourceQualarooRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceQualaroo(ctx, request)
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
