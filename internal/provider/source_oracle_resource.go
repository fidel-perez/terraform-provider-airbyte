// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk"
	"context"
	"fmt"

	speakeasy_stringplanmodifier "airbyte/internal/planmodifiers/stringplanmodifier"
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/validators"
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
var _ resource.Resource = &SourceOracleResource{}
var _ resource.ResourceWithImportState = &SourceOracleResource{}

func NewSourceOracleResource() resource.Resource {
	return &SourceOracleResource{}
}

// SourceOracleResource defines the resource implementation.
type SourceOracleResource struct {
	client *sdk.SDK
}

// SourceOracleResourceModel describes the resource data model.
type SourceOracleResourceModel struct {
	Configuration SourceOracle `tfsdk:"configuration"`
	Name          types.String `tfsdk:"name"`
	SecretID      types.String `tfsdk:"secret_id"`
	SourceID      types.String `tfsdk:"source_id"`
	SourceType    types.String `tfsdk:"source_type"`
	WorkspaceID   types.String `tfsdk:"workspace_id"`
}

func (r *SourceOracleResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_oracle"
}

func (r *SourceOracleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceOracle Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"connection_data": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_oracle_connect_by_service_name": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"connection_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"service_name",
											),
										},
										Description: `must be one of ["service_name"]`,
									},
									"service_name": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Use service name`,
							},
							"source_oracle_connect_by_system_id_sid": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"connection_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"sid",
											),
										},
										Description: `must be one of ["sid"]`,
									},
									"sid": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Use SID (Oracle System Identifier)`,
							},
							"source_oracle_update_connect_by_service_name": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"connection_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"service_name",
											),
										},
										Description: `must be one of ["service_name"]`,
									},
									"service_name": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Use service name`,
							},
							"source_oracle_update_connect_by_system_id_sid": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"connection_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"sid",
											),
										},
										Description: `must be one of ["sid"]`,
									},
									"sid": schema.StringAttribute{
										Required: true,
									},
								},
								Description: `Use SID (Oracle System Identifier)`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Connect data that will be used for DB connection`,
					},
					"encryption": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"source_oracle_encryption_native_network_encryption_nne": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"encryption_algorithm": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"AES256",
												"RC4_56",
												"3DES168",
											),
										},
										MarkdownDescription: `must be one of ["AES256", "RC4_56", "3DES168"]` + "\n" +
											`This parameter defines what encryption algorithm is used.`,
									},
									"encryption_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"client_nne",
											),
										},
										Description: `must be one of ["client_nne"]`,
									},
								},
								Description: `The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.`,
							},
							"source_oracle_encryption_tls_encrypted_verify_certificate": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"encryption_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_verify_certificate",
											),
										},
										Description: `must be one of ["encrypted_verify_certificate"]`,
									},
									"ssl_certificate": schema.StringAttribute{
										Required:    true,
										Description: `Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
							},
							"source_oracle_update_encryption_native_network_encryption_nne": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"encryption_algorithm": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"AES256",
												"RC4_56",
												"3DES168",
											),
										},
										MarkdownDescription: `must be one of ["AES256", "RC4_56", "3DES168"]` + "\n" +
											`This parameter defines what encryption algorithm is used.`,
									},
									"encryption_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"client_nne",
											),
										},
										Description: `must be one of ["client_nne"]`,
									},
								},
								Description: `The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.`,
							},
							"source_oracle_update_encryption_tls_encrypted_verify_certificate": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"encryption_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"encrypted_verify_certificate",
											),
										},
										Description: `must be one of ["encrypted_verify_certificate"]`,
									},
									"ssl_certificate": schema.StringAttribute{
										Required:    true,
										Description: `Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `The encryption method with is used when communicating with the database.`,
					},
					"host": schema.StringAttribute{
						Required:    true,
						Description: `Hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"password": schema.StringAttribute{
						Optional:    true,
						Description: `The password associated with the username.`,
					},
					"port": schema.Int64Attribute{
						Required: true,
						MarkdownDescription: `Port of the database.` + "\n" +
							`Oracle Corporations recommends the following port numbers:` + "\n" +
							`1521 - Default listening port for client connections to the listener. ` + "\n" +
							`2484 - Recommended and officially registered listening port for client connections to the listener using TCP/IP with SSL`,
					},
					"schemas": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `The list of schemas to sync from. Defaults to user. Case sensitive.`,
					},
					"source_type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"oracle",
							),
						},
						Description: `must be one of ["oracle"]`,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_oracle_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_oracle_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_oracle_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_oracle_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"NO_TUNNEL",
											),
										},
										MarkdownDescription: `must be one of ["NO_TUNNEL"]` + "\n" +
											`No ssh tunnel needed to connect to database`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_oracle_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_PASSWORD_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_PASSWORD_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and password authentication`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required:    true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_oracle_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required:    true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_method": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"SSH_KEY_AUTH",
											),
										},
										MarkdownDescription: `must be one of ["SSH_KEY_AUTH"]` + "\n" +
											`Connect through a jump server tunnel host using username and ssh key`,
									},
									"tunnel_port": schema.Int64Attribute{
										Required:    true,
										Description: `Port on the proxy/jump server that accepts inbound ssh connections.`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `The username which is used to access the database.`,
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

func (r *SourceOracleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceOracleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceOracleResourceModel
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
	res, err := r.client.Sources.CreateSourceOracle(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceOracleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceOracleResourceModel
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
	request := operations.GetSourceOracleRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceOracle(ctx, request)
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

func (r *SourceOracleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceOracleResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceOraclePutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceOracleRequest{
		SourceOraclePutRequest: sourceOraclePutRequest,
		SourceID:               sourceID,
	}
	res, err := r.client.Sources.PutSourceOracle(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
	getRequest := operations.GetSourceOracleRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceOracle(ctx, getRequest)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
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
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceOracleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceOracleResourceModel
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
	request := operations.DeleteSourceOracleRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceOracle(ctx, request)
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

func (r *SourceOracleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
