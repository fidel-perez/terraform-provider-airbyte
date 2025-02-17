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
var _ resource.Resource = &DestinationPineconeResource{}
var _ resource.ResourceWithImportState = &DestinationPineconeResource{}

func NewDestinationPineconeResource() resource.Resource {
	return &DestinationPineconeResource{}
}

// DestinationPineconeResource defines the resource implementation.
type DestinationPineconeResource struct {
	client *sdk.SDK
}

// DestinationPineconeResourceModel describes the resource data model.
type DestinationPineconeResourceModel struct {
	Configuration   DestinationPinecone `tfsdk:"configuration"`
	DefinitionID    types.String        `tfsdk:"definition_id"`
	DestinationID   types.String        `tfsdk:"destination_id"`
	DestinationType types.String        `tfsdk:"destination_type"`
	Name            types.String        `tfsdk:"name"`
	WorkspaceID     types.String        `tfsdk:"workspace_id"`
}

func (r *DestinationPineconeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_pinecone"
}

func (r *DestinationPineconeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationPinecone Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"embedding": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"azure_open_ai": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_base": schema.StringAttribute{
										Required:    true,
										Description: `The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
									"deployment": schema.StringAttribute{
										Required:    true,
										Description: `The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
									"openai_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource`,
									},
								},
								Description: `Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
							"cohere": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"cohere_key": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
									},
								},
								Description: `Use the Cohere API to embed text.`,
							},
							"fake": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.`,
							},
							"open_ai": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"openai_key": schema.StringAttribute{
										Required:  true,
										Sensitive: true,
									},
								},
								Description: `Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.`,
							},
							"open_ai_compatible": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"api_key": schema.StringAttribute{
										Optional:    true,
										Sensitive:   true,
										Description: `Default: ""`,
									},
									"base_url": schema.StringAttribute{
										Required:    true,
										Description: `The base URL for your OpenAI-compatible service`,
									},
									"dimensions": schema.Int64Attribute{
										Required:    true,
										Description: `The number of dimensions the embedding model is generating`,
									},
									"model_name": schema.StringAttribute{
										Optional: true,
										MarkdownDescription: `Default: "text-embedding-ada-002"` + "\n" +
											`The name of the model to use for embedding`,
									},
								},
								Description: `Use a service that's compatible with the OpenAI API to embed text.`,
							},
						},
						Description: `Embedding configuration`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"indexing": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"index": schema.StringAttribute{
								Required:    true,
								Description: `Pinecone index in your project to load data into`,
							},
							"pinecone_environment": schema.StringAttribute{
								Required:    true,
								Description: `Pinecone Cloud environment to use`,
							},
							"pinecone_key": schema.StringAttribute{
								Required:    true,
								Sensitive:   true,
								Description: `The Pinecone API key to use matching the environment (copy from Pinecone console)`,
							},
						},
						Description: `Pinecone is a popular vector store that can be used to store and retrieve embeddings.`,
					},
					"processing": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"chunk_overlap": schema.Int64Attribute{
								Optional: true,
								MarkdownDescription: `Default: 0` + "\n" +
									`Size of overlap between chunks in tokens to store in vector store to better capture relevant context`,
							},
							"chunk_size": schema.Int64Attribute{
								Required:    true,
								Description: `Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)`,
							},
							"field_name_mappings": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from_field": schema.StringAttribute{
											Required:    true,
											Description: `The field name in the source`,
										},
										"to_field": schema.StringAttribute{
											Required:    true,
											Description: `The field name to use in the destination`,
										},
									},
								},
								Description: `List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.`,
							},
							"metadata_fields": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.`,
							},
							"text_fields": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
								Description: `List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. ` + "`" + `user.name` + "`" + ` will access the ` + "`" + `name` + "`" + ` field in the ` + "`" + `user` + "`" + ` object. It's also possible to use wildcards to access all fields in an object, e.g. ` + "`" + `users.*.name` + "`" + ` will access all ` + "`" + `names` + "`" + ` fields in all entries of the ` + "`" + `users` + "`" + ` array.`,
							},
							"text_splitter": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"by_markdown_header": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"split_level": schema.Int64Attribute{
												Optional: true,
												MarkdownDescription: `Default: 1` + "\n" +
													`Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points`,
											},
										},
										Description: `Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.`,
									},
									"by_programming_language": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"language": schema.StringAttribute{
												Required: true,
												MarkdownDescription: `must be one of ["cpp", "go", "java", "js", "php", "proto", "python", "rst", "ruby", "rust", "scala", "swift", "markdown", "latex", "html", "sol"]` + "\n" +
													`Split code in suitable places based on the programming language`,
												Validators: []validator.String{
													stringvalidator.OneOf(
														"cpp",
														"go",
														"java",
														"js",
														"php",
														"proto",
														"python",
														"rst",
														"ruby",
														"rust",
														"scala",
														"swift",
														"markdown",
														"latex",
														"html",
														"sol",
													),
												},
											},
										},
										Description: `Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.`,
									},
									"by_separator": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"keep_separator": schema.BoolAttribute{
												Optional: true,
												MarkdownDescription: `Default: false` + "\n" +
													`Whether to keep the separator in the resulting chunks`,
											},
											"separators": schema.ListAttribute{
												Optional:    true,
												ElementType: types.StringType,
												Description: `List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".`,
											},
										},
										Description: `Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.`,
									},
								},
								Description: `Split text fields into chunks based on the specified method.`,
								Validators: []validator.Object{
									validators.ExactlyOneChild(),
								},
							},
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

func (r *DestinationPineconeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationPineconeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationPineconeResourceModel
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
	res, err := r.client.Destinations.CreateDestinationPinecone(ctx, request)
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

func (r *DestinationPineconeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationPineconeResourceModel
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
	request := operations.GetDestinationPineconeRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationPinecone(ctx, request)
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

func (r *DestinationPineconeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationPineconeResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationPineconePutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationPineconeRequest{
		DestinationPineconePutRequest: destinationPineconePutRequest,
		DestinationID:                 destinationID,
	}
	res, err := r.client.Destinations.PutDestinationPinecone(ctx, request)
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
	getRequest := operations.GetDestinationPineconeRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationPinecone(ctx, getRequest)
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

func (r *DestinationPineconeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationPineconeResourceModel
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
	request := operations.DeleteDestinationPineconeRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationPinecone(ctx, request)
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

func (r *DestinationPineconeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
