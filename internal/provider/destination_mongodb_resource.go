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
var _ resource.Resource = &DestinationMongodbResource{}
var _ resource.ResourceWithImportState = &DestinationMongodbResource{}

func NewDestinationMongodbResource() resource.Resource {
	return &DestinationMongodbResource{}
}

// DestinationMongodbResource defines the resource implementation.
type DestinationMongodbResource struct {
	client *sdk.SDK
}

// DestinationMongodbResourceModel describes the resource data model.
type DestinationMongodbResourceModel struct {
	Configuration   DestinationMongodb `tfsdk:"configuration"`
	DefinitionID    types.String       `tfsdk:"definition_id"`
	DestinationID   types.String       `tfsdk:"destination_id"`
	DestinationType types.String       `tfsdk:"destination_type"`
	Name            types.String       `tfsdk:"name"`
	WorkspaceID     types.String       `tfsdk:"workspace_id"`
}

func (r *DestinationMongodbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_mongodb"
}

func (r *DestinationMongodbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationMongodb Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"auth_type": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"login_password": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `Password associated with the username.`,
									},
									"username": schema.StringAttribute{
										Required:    true,
										Description: `Username to use to access the database.`,
									},
								},
								Description: `Login/Password.`,
							},
							"none": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `None.`,
							},
						},
						Description: `Authorization type.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"database": schema.StringAttribute{
						Required:    true,
						Description: `Name of the database.`,
					},
					"instance_type": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mongo_db_atlas": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"cluster_url": schema.StringAttribute{
										Required:    true,
										Description: `URL of a cluster to connect to.`,
									},
									"instance": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["atlas"]; Default: "atlas"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"atlas",
											),
										},
									},
								},
								Description: `MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.`,
							},
							"replica_set": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"instance": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["replica"]; Default: "replica"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"replica",
											),
										},
									},
									"replica_set": schema.StringAttribute{
										Optional:    true,
										Description: `A replica set name.`,
									},
									"server_addresses": schema.StringAttribute{
										Required:    true,
										Description: `The members of a replica set. Please specify ` + "`" + `host` + "`" + `:` + "`" + `port` + "`" + ` of each member seperated by comma.`,
									},
								},
								Description: `MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.`,
							},
							"standalone_mongo_db_instance": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Required:    true,
										Description: `The Host of a Mongo database to be replicated.`,
									},
									"instance": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["standalone"]; Default: "standalone"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"standalone",
											),
										},
									},
									"port": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 27017` + "\n" +
											`The Port of a Mongo database to be replicated.`,
									},
								},
								Description: `MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.`,
							},
						},
						Description: `MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"no_tunnel": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 22` + "\n" +
											`Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 22` + "\n" +
											`Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.`,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(),
				},
				Required:    true,
				Description: `Name of the destination e.g. dev-mysql-instance.`,
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

func (r *DestinationMongodbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationMongodbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationMongodbResourceModel
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
	res, err := r.client.Destinations.CreateDestinationMongodb(ctx, request)
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
	data.RefreshFromCreateResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationMongodbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationMongodbResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationMongodbRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationMongodb(ctx, request)
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

func (r *DestinationMongodbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationMongodbResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationMongodbPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationMongodbRequest{
		DestinationMongodbPutRequest: destinationMongodbPutRequest,
		DestinationID:                destinationID,
	}
	res, err := r.client.Destinations.PutDestinationMongodb(ctx, request)
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
	destinationId1 := data.DestinationID.ValueString()
	getRequest := operations.GetDestinationMongodbRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationMongodb(ctx, getRequest)
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
	if getResponse.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationMongodbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationMongodbResourceModel
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

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationMongodbRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationMongodb(ctx, request)
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

func (r *DestinationMongodbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
