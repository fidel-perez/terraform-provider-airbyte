// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	"airbyte/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectionResource{}
var _ resource.ResourceWithImportState = &ConnectionResource{}

func NewConnectionResource() resource.Resource {
	return &ConnectionResource{}
}

// ConnectionResource defines the resource implementation.
type ConnectionResource struct {
	client *sdk.SDK
}

// ConnectionResourceModel describes the resource data model.
type ConnectionResourceModel struct {
	Configurations                   *StreamConfigurations `tfsdk:"configurations"`
	ConnectionID                     types.String          `tfsdk:"connection_id"`
	DataResidency                    types.String          `tfsdk:"data_residency"`
	DestinationID                    types.String          `tfsdk:"destination_id"`
	Name                             types.String          `tfsdk:"name"`
	NamespaceDefinition              types.String          `tfsdk:"namespace_definition"`
	NamespaceFormat                  types.String          `tfsdk:"namespace_format"`
	NonBreakingSchemaUpdatesBehavior types.String          `tfsdk:"non_breaking_schema_updates_behavior"`
	Prefix                           types.String          `tfsdk:"prefix"`
	Schedule                         *ConnectionSchedule   `tfsdk:"schedule"`
	SourceID                         types.String          `tfsdk:"source_id"`
	Status                           types.String          `tfsdk:"status"`
	WorkspaceID                      types.String          `tfsdk:"workspace_id"`
}

func (r *ConnectionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection"
}

func (r *ConnectionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Connection Resource",

		Attributes: map[string]schema.Attribute{
			"configurations": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"streams": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cursor_field": schema.ListAttribute{
									Computed:    true,
									Optional:    true,
									ElementType: types.StringType,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
								"primary_key": schema.ListAttribute{
									Computed: true,
									Optional: true,
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"sync_mode": schema.StringAttribute{
									Computed: true,
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"full_refresh_overwrite",
											"full_refresh_append",
											"incremental_append",
											"incremental_deduped_history",
										),
									},
								},
							},
						},
					},
				},
				Description: `A list of configured stream options for a connection.`,
			},
			"connection_id": schema.StringAttribute{
				Computed: true,
			},
			"data_residency": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"auto",
						"us",
						"eu",
					),
				},
			},
			"destination_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"namespace_definition": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"source",
						"destination",
						"custom_format",
					),
				},
				Description: `Define the location where the data will be stored in the destination`,
			},
			"namespace_format": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"non_breaking_schema_updates_behavior": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ignore",
						"disable_connection",
					),
				},
				Description: `Set how Airbyte handles syncs when it detects a non-breaking schema change in the source`,
			},
			"prefix": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"schedule": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"basic_timing": schema.StringAttribute{
						Computed: true,
					},
					"schedule_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"manual",
								"cron",
							),
						},
					},
				},
				Description: `schedule for when the the connection should run, per the schedule type`,
			},
			"source_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
			},
			"status": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"active",
						"inactive",
						"deprecated",
					),
				},
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectionResourceModel
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

	request := *data.ToCreateSDKType()
	res, err := r.client.Connections.CreateConnection(ctx, request)
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
	if res.ConnectionResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.ConnectionResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectionResourceModel
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

	connectionID := data.ConnectionID.ValueString()
	request := operations.GetConnectionRequest{
		ConnectionID: connectionID,
	}
	res, err := r.client.Connections.GetConnection(ctx, request)
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
	if res.ConnectionResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(res.ConnectionResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectionResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	connectionPatchRequest := *data.ToUpdateSDKType()
	connectionID := data.ConnectionID.ValueString()
	request := operations.PatchConnectionRequest{
		ConnectionPatchRequest: connectionPatchRequest,
		ConnectionID:           connectionID,
	}
	res, err := r.client.Connections.PatchConnection(ctx, request)
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
	if res.ConnectionResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromUpdateResponse(res.ConnectionResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectionResourceModel
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

	connectionID := data.ConnectionID.ValueString()
	request := operations.DeleteConnectionRequest{
		ConnectionID: connectionID,
	}
	res, err := r.client.Connections.DeleteConnection(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *ConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
