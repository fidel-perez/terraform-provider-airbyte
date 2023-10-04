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
var _ datasource.DataSource = &SourceTheGuardianAPIDataSource{}
var _ datasource.DataSourceWithConfigure = &SourceTheGuardianAPIDataSource{}

func NewSourceTheGuardianAPIDataSource() datasource.DataSource {
	return &SourceTheGuardianAPIDataSource{}
}

// SourceTheGuardianAPIDataSource is the data source implementation.
type SourceTheGuardianAPIDataSource struct {
	client *sdk.SDK
}

// SourceTheGuardianAPIDataSourceModel describes the data model.
type SourceTheGuardianAPIDataSourceModel struct {
	Configuration SourceTheGuardianAPI `tfsdk:"configuration"`
	Name          types.String         `tfsdk:"name"`
	SecretID      types.String         `tfsdk:"secret_id"`
	SourceID      types.String         `tfsdk:"source_id"`
	WorkspaceID   types.String         `tfsdk:"workspace_id"`
}

// Metadata returns the data source type name.
func (r *SourceTheGuardianAPIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_the_guardian_api"
}

// Schema defines the schema for the data source.
func (r *SourceTheGuardianAPIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceTheGuardianAPI DataSource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"api_key": schema.StringAttribute{
						Computed:    true,
						Description: `Your API Key. See <a href="https://open-platform.theguardian.com/access/">here</a>. The key is case sensitive.`,
					},
					"end_date": schema.StringAttribute{
						Computed:    true,
						Description: `(Optional) Use this to set the maximum date (YYYY-MM-DD) of the results. Results newer than the end_date will not be shown. Default is set to the current date (today) for incremental syncs.`,
					},
					"query": schema.StringAttribute{
						Computed:    true,
						Description: `(Optional) The query (q) parameter filters the results to only those that include that search term. The q parameter supports AND, OR and NOT operators.`,
					},
					"section": schema.StringAttribute{
						Computed:    true,
						Description: `(Optional) Use this to filter the results by a particular section. See <a href="https://content.guardianapis.com/sections?api-key=test">here</a> for a list of all sections, and <a href="https://open-platform.theguardian.com/documentation/section">here</a> for the sections endpoint documentation.`,
					},
					"start_date": schema.StringAttribute{
						Computed:    true,
						Description: `Use this to set the minimum date (YYYY-MM-DD) of the results. Results older than the start_date will not be shown.`,
					},
					"tag": schema.StringAttribute{
						Computed:    true,
						Description: `(Optional) A tag is a piece of data that is used by The Guardian to categorise content. Use this parameter to filter results by showing only the ones matching the entered tag. See <a href="https://content.guardianapis.com/tags?api-key=test">here</a> for a list of all tags, and <a href="https://open-platform.theguardian.com/documentation/tag">here</a> for the tags endpoint documentation.`,
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

func (r *SourceTheGuardianAPIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *SourceTheGuardianAPIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *SourceTheGuardianAPIDataSourceModel
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
	request := operations.GetSourceTheGuardianAPIRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceTheGuardianAPI(ctx, request)
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
