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
var _ resource.Resource = &SourceAlloydbResource{}
var _ resource.ResourceWithImportState = &SourceAlloydbResource{}

func NewSourceAlloydbResource() resource.Resource {
	return &SourceAlloydbResource{}
}

// SourceAlloydbResource defines the resource implementation.
type SourceAlloydbResource struct {
	client *sdk.SDK
}

// SourceAlloydbResourceModel describes the resource data model.
type SourceAlloydbResourceModel struct {
	Configuration SourceAlloydb `tfsdk:"configuration"`
	Name          types.String  `tfsdk:"name"`
	SecretID      types.String  `tfsdk:"secret_id"`
	SourceID      types.String  `tfsdk:"source_id"`
	SourceType    types.String  `tfsdk:"source_type"`
	WorkspaceID   types.String  `tfsdk:"workspace_id"`
}

func (r *SourceAlloydbResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_alloydb"
}

func (r *SourceAlloydbResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "SourceAlloydb Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required:    true,
						Description: `Name of the database.`,
					},
					"host": schema.StringAttribute{
						Required:    true,
						Description: `Hostname of the database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (Eg. key1=value1&key2=value2&key3=value3). For more information read about <a href="https://jdbc.postgresql.org/documentation/head/connect.html">JDBC URL parameters</a>.`,
					},
					"password": schema.StringAttribute{
						Optional:    true,
						Description: `Password associated with the username.`,
					},
					"port": schema.Int64Attribute{
						Optional: true,
						MarkdownDescription: `Default: 5432` + "\n" +
							`Port of the database.`,
					},
					"replication_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"initial_waiting_seconds": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 300` + "\n" +
											`The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-5-optional-set-up-initial-waiting-time">initial waiting time</a>.`,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										MarkdownDescription: `must be one of ["While reading Data", "After loading Data in the destination"]; Default: "After loading Data in the destination"` + "\n" +
											`Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"plugin": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										MarkdownDescription: `must be one of ["pgoutput"]; Default: "pgoutput"` + "\n" +
											`A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Required:    true,
										Description: `A Postgres publication used for consuming changes. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-4-create-publications-and-replication-identities-for-tables">publications and replication identities</a>.`,
									},
									"queue_size": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 10000` + "\n" +
											`The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.`,
									},
									"replication_slot": schema.StringAttribute{
										Required:    true,
										Description: `A plugin logical replication slot. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-3-create-replication-slot">replication slots</a>.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_alloydb_replication_method_standard": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_alloydb_replication_method_standard_xmin": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Xmin replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_alloydb_update_replication_method_logical_replication_cdc": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"initial_waiting_seconds": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 300` + "\n" +
											`The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-5-optional-set-up-initial-waiting-time">initial waiting time</a>.`,
									},
									"lsn_commit_behaviour": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"While reading Data",
												"After loading Data in the destination",
											),
										},
										MarkdownDescription: `must be one of ["While reading Data", "After loading Data in the destination"]; Default: "After loading Data in the destination"` + "\n" +
											`Determines when Airbtye should flush the LSN of processed WAL logs in the source database. ` + "`" + `After loading Data in the destination` + "`" + ` is default. If ` + "`" + `While reading Data` + "`" + ` is selected, in case of a downstream failure (while loading data into the destination), next sync would result in a full sync.`,
									},
									"plugin": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"pgoutput",
											),
										},
										MarkdownDescription: `must be one of ["pgoutput"]; Default: "pgoutput"` + "\n" +
											`A logical decoding plugin installed on the PostgreSQL server.`,
									},
									"publication": schema.StringAttribute{
										Required:    true,
										Description: `A Postgres publication used for consuming changes. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-4-create-publications-and-replication-identities-for-tables">publications and replication identities</a>.`,
									},
									"queue_size": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 10000` + "\n" +
											`The size of the internal queue. This may interfere with memory consumption and efficiency of the connector, please be careful.`,
									},
									"replication_slot": schema.StringAttribute{
										Required:    true,
										Description: `A plugin logical replication slot. Read about <a href="https://docs.airbyte.com/integrations/sources/postgres#step-3-create-replication-slot">replication slots</a>.`,
									},
								},
								Description: `Logical replication uses the Postgres write-ahead log (WAL) to detect inserts, updates, and deletes. This needs to be configured on the source database itself. Only available on Postgres 10 and above. Read the <a href="https://docs.airbyte.com/integrations/sources/postgres">docs</a>.`,
							},
							"source_alloydb_update_replication_method_standard": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Standard replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
							"source_alloydb_update_replication_method_standard_xmin": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Xmin replication requires no setup on the DB side but will not be able to represent deletions incrementally.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Replication method for extracting data from the database.`,
					},
					"schemas": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
						Description: `The list of schemas (case sensitive) to sync from. Defaults to public.`,
					},
					"ssl_mode": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_ssl_modes_allow": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_alloydb_ssl_modes_disable": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_alloydb_ssl_modes_prefer": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_alloydb_ssl_modes_require": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_alloydb_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"ca_certificate": schema.StringAttribute{
										Required:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Optional:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Optional:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_alloydb_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"ca_certificate": schema.StringAttribute{
										Required:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Optional:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Optional:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
							"source_alloydb_update_ssl_modes_allow": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Enables encryption only when required by the source database.`,
							},
							"source_alloydb_update_ssl_modes_disable": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Disables encryption of communication between Airbyte and source database.`,
							},
							"source_alloydb_update_ssl_modes_prefer": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Allows unencrypted connection only if the source database does not support encryption.`,
							},
							"source_alloydb_update_ssl_modes_require": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
								},
								Description: `Always require encryption. If the source database server does not support encryption, connection will fail.`,
							},
							"source_alloydb_update_ssl_modes_verify_ca": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"ca_certificate": schema.StringAttribute{
										Required:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Optional:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Optional:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
								},
								Description: `Always require encryption and verifies that the source database server has a valid SSL certificate.`,
							},
							"source_alloydb_update_ssl_modes_verify_full": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											validators.IsValidJSON(),
										},
										Description: `Parsed as JSON.`,
									},
									"ca_certificate": schema.StringAttribute{
										Required:    true,
										Description: `CA certificate`,
									},
									"client_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Client certificate`,
									},
									"client_key": schema.StringAttribute{
										Optional:    true,
										Description: `Client key`,
									},
									"client_key_password": schema.StringAttribute{
										Optional:    true,
										Description: `Password for keystorage. If you do not add it - the password will be generated automatically.`,
									},
								},
								Description: `This is the most secure mode. Always require encryption and verifies the identity of the source database server.`,
							},
						},
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						MarkdownDescription: `SSL connection modes. ` + "\n" +
							`  Read more <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.`,
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"source_alloydb_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_alloydb_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
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
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_alloydb_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
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
							"source_alloydb_update_ssh_tunnel_method_no_tunnel": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_alloydb_update_ssh_tunnel_method_password_authentication": schema.SingleNestedAttribute{
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
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
							},
							"source_alloydb_update_ssh_tunnel_method_ssh_key_authentication": schema.SingleNestedAttribute{
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
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `Username to access the database.`,
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

func (r *SourceAlloydbResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SourceAlloydbResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SourceAlloydbResourceModel
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
	res, err := r.client.Sources.CreateSourceAlloydb(ctx, request)
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
	data.RefreshFromCreateResponse(res.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAlloydbResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SourceAlloydbResourceModel
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
	request := operations.GetSourceAlloydbRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.GetSourceAlloydb(ctx, request)
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

func (r *SourceAlloydbResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SourceAlloydbResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	sourceAlloydbPutRequest := data.ToUpdateSDKType()
	sourceID := data.SourceID.ValueString()
	request := operations.PutSourceAlloydbRequest{
		SourceAlloydbPutRequest: sourceAlloydbPutRequest,
		SourceID:                sourceID,
	}
	res, err := r.client.Sources.PutSourceAlloydb(ctx, request)
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
	sourceId1 := data.SourceID.ValueString()
	getRequest := operations.GetSourceAlloydbRequest{
		SourceID: sourceId1,
	}
	getResponse, err := r.client.Sources.GetSourceAlloydb(ctx, getRequest)
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
	if getResponse.SourceResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(getResponse.RawResponse))
		return
	}
	data.RefreshFromGetResponse(getResponse.SourceResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SourceAlloydbResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SourceAlloydbResourceModel
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
	request := operations.DeleteSourceAlloydbRequest{
		SourceID: sourceID,
	}
	res, err := r.client.Sources.DeleteSourceAlloydb(ctx, request)
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

func (r *SourceAlloydbResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("source_id"), req, resp)
}
