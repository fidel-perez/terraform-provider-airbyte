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
var _ resource.Resource = &DestinationGcsResource{}
var _ resource.ResourceWithImportState = &DestinationGcsResource{}

func NewDestinationGcsResource() resource.Resource {
	return &DestinationGcsResource{}
}

// DestinationGcsResource defines the resource implementation.
type DestinationGcsResource struct {
	client *sdk.SDK
}

// DestinationGcsResourceModel describes the resource data model.
type DestinationGcsResourceModel struct {
	Configuration   DestinationGcs `tfsdk:"configuration"`
	DefinitionID    types.String   `tfsdk:"definition_id"`
	DestinationID   types.String   `tfsdk:"destination_id"`
	DestinationType types.String   `tfsdk:"destination_type"`
	Name            types.String   `tfsdk:"name"`
	WorkspaceID     types.String   `tfsdk:"workspace_id"`
}

func (r *DestinationGcsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_gcs"
}

func (r *DestinationGcsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationGcs Resource",

		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"credential": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"hmac_key": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"credential_type": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["HMAC_KEY"]; Default: "HMAC_KEY"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"HMAC_KEY",
											),
										},
									},
									"hmac_key_access_id": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys#overview">here</a>.`,
									},
									"hmac_key_secret": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `The corresponding secret for the access ID. It is a 40-character base-64 encoded string.  Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys#secrets">here</a>.`,
									},
								},
								Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
							},
						},
						Description: `An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"format": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"avro_apache_avro": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression_codec": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"bzip2": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["bzip2"]; Default: "bzip2"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"bzip2",
															),
														},
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
											"deflate": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["Deflate"]; Default: "Deflate"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"Deflate",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Optional: true,
														MarkdownDescription: `Default: 0` + "\n" +
															`0: no compression & fastest, 9: best compression & slowest.`,
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["no compression"]; Default: "no compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"no compression",
															),
														},
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
											"snappy": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["snappy"]; Default: "snappy"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"snappy",
															),
														},
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
											"xz": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["xz"]; Default: "xz"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"xz",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Optional: true,
														MarkdownDescription: `Default: 6` + "\n" +
															`The presets 0-3 are fast presets with medium compression. The presets 4-6 are fairly slow presets with high compression. The default preset is 6. The presets 7-9 are like the preset 6 but use bigger dictionaries and have higher compressor and decompressor memory requirements. Unless the uncompressed size of the file exceeds 8 MiB, 16 MiB, or 32 MiB, it is waste of memory to use the presets 7, 8, or 9, respectively. Read more <a href="https://commons.apache.org/proper/commons-compress/apidocs/org/apache/commons/compress/compressors/xz/XZCompressorOutputStream.html#XZCompressorOutputStream-java.io.OutputStream-int-">here</a> for details.`,
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
											"zstandard": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"codec": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["zstandard"]; Default: "zstandard"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"zstandard",
															),
														},
													},
													"compression_level": schema.Int64Attribute{
														Optional: true,
														MarkdownDescription: `Default: 3` + "\n" +
															`Negative levels are 'fast' modes akin to lz4 or snappy, levels above 9 are generally for archival purposes, and levels above 18 use a lot of memory.`,
													},
													"include_checksum": schema.BoolAttribute{
														Optional: true,
														MarkdownDescription: `Default: false` + "\n" +
															`If true, include a checksum with each data block.`,
													},
												},
												Description: `The compression algorithm used to compress data. Default to no compression.`,
											},
										},
										Description: `The compression algorithm used to compress data. Default to no compression.`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"format_type": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["Avro"]; Default: "Avro"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Avro",
											),
										},
									},
								},
								Description: `Output data format. One of the following formats must be selected - <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-avro#advantages_of_avro">AVRO</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-parquet#parquet_schemas">PARQUET</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-csv#loading_csv_data_into_a_table">CSV</a> format, or <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-json#loading_json_data_into_a_new_table">JSONL</a> format.`,
							},
							"csv_comma_separated_values": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"gzip": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["GZIP"]; Default: "GZIP"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz").`,
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["No Compression"]; Default: "No Compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz").`,
											},
										},
										Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".csv.gz").`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"flattening": schema.StringAttribute{
										Optional: true,
										MarkdownDescription: `must be one of ["No flattening", "Root level flattening"]; Default: "No flattening"` + "\n" +
											`Whether the input JSON data should be normalized (flattened) in the output CSV. Please refer to docs for details.`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"No flattening",
												"Root level flattening",
											),
										},
									},
									"format_type": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["CSV"]; Default: "CSV"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"CSV",
											),
										},
									},
								},
								Description: `Output data format. One of the following formats must be selected - <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-avro#advantages_of_avro">AVRO</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-parquet#parquet_schemas">PARQUET</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-csv#loading_csv_data_into_a_table">CSV</a> format, or <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-json#loading_json_data_into_a_new_table">JSONL</a> format.`,
							},
							"json_lines_newline_delimited_json": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"compression": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"gzip": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["GZIP"]; Default: "GZIP"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"GZIP",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
											"no_compression": schema.SingleNestedAttribute{
												Optional: true,
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Optional:    true,
														Description: `must be one of ["No Compression"]; Default: "No Compression"`,
														Validators: []validator.String{
															stringvalidator.OneOf(
																"No Compression",
															),
														},
													},
												},
												Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
											},
										},
										Description: `Whether the output files should be compressed. If compression is selected, the output filename will have an extra extension (GZIP: ".jsonl.gz").`,
										Validators: []validator.Object{
											validators.ExactlyOneChild(),
										},
									},
									"format_type": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["JSONL"]; Default: "JSONL"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"JSONL",
											),
										},
									},
								},
								Description: `Output data format. One of the following formats must be selected - <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-avro#advantages_of_avro">AVRO</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-parquet#parquet_schemas">PARQUET</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-csv#loading_csv_data_into_a_table">CSV</a> format, or <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-json#loading_json_data_into_a_new_table">JSONL</a> format.`,
							},
							"parquet_columnar_storage": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"block_size_mb": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 128` + "\n" +
											`This is the size of a row group being buffered in memory. It limits the memory usage when writing. Larger values will improve the IO when reading, but consume more memory when writing. Default: 128 MB.`,
									},
									"compression_codec": schema.StringAttribute{
										Optional: true,
										MarkdownDescription: `must be one of ["UNCOMPRESSED", "SNAPPY", "GZIP", "LZO", "BROTLI", "LZ4", "ZSTD"]; Default: "UNCOMPRESSED"` + "\n" +
											`The compression algorithm used to compress data pages.`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"UNCOMPRESSED",
												"SNAPPY",
												"GZIP",
												"LZO",
												"BROTLI",
												"LZ4",
												"ZSTD",
											),
										},
									},
									"dictionary_encoding": schema.BoolAttribute{
										Optional: true,
										MarkdownDescription: `Default: true` + "\n" +
											`Default: true.`,
									},
									"dictionary_page_size_kb": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 1024` + "\n" +
											`There is one dictionary page per column per row group when dictionary encoding is used. The dictionary page size works like the page size but for dictionary. Default: 1024 KB.`,
									},
									"format_type": schema.StringAttribute{
										Optional:    true,
										Description: `must be one of ["Parquet"]; Default: "Parquet"`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"Parquet",
											),
										},
									},
									"max_padding_size_mb": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 8` + "\n" +
											`Maximum size allowed as padding to align row groups. This is also the minimum size of a row group. Default: 8 MB.`,
									},
									"page_size_kb": schema.Int64Attribute{
										Optional: true,
										MarkdownDescription: `Default: 1024` + "\n" +
											`The page size is for compression. A block is composed of pages. A page is the smallest unit that must be read fully to access a single record. If this value is too small, the compression will deteriorate. Default: 1024 KB.`,
									},
								},
								Description: `Output data format. One of the following formats must be selected - <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-avro#advantages_of_avro">AVRO</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-parquet#parquet_schemas">PARQUET</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-csv#loading_csv_data_into_a_table">CSV</a> format, or <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-json#loading_json_data_into_a_new_table">JSONL</a> format.`,
							},
						},
						Description: `Output data format. One of the following formats must be selected - <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-avro#advantages_of_avro">AVRO</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-parquet#parquet_schemas">PARQUET</a> format, <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-csv#loading_csv_data_into_a_table">CSV</a> format, or <a href="https://cloud.google.com/bigquery/docs/loading-data-cloud-storage-json#loading_json_data_into_a_new_table">JSONL</a> format.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"gcs_bucket_name": schema.StringAttribute{
						Required:    true,
						Description: `You can find the bucket name in the App Engine Admin console Application Settings page, under the label Google Cloud Storage Bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.`,
					},
					"gcs_bucket_path": schema.StringAttribute{
						Required:    true,
						Description: `GCS Bucket Path string Subdirectory under the above bucket to sync the data into.`,
					},
					"gcs_bucket_region": schema.StringAttribute{
						Optional: true,
						MarkdownDescription: `must be one of ["northamerica-northeast1", "northamerica-northeast2", "us-central1", "us-east1", "us-east4", "us-west1", "us-west2", "us-west3", "us-west4", "southamerica-east1", "southamerica-west1", "europe-central2", "europe-north1", "europe-west1", "europe-west2", "europe-west3", "europe-west4", "europe-west6", "asia-east1", "asia-east2", "asia-northeast1", "asia-northeast2", "asia-northeast3", "asia-south1", "asia-south2", "asia-southeast1", "asia-southeast2", "australia-southeast1", "australia-southeast2", "asia", "eu", "us", "asia1", "eur4", "nam4"]; Default: "us"` + "\n" +
							`Select a Region of the GCS Bucket. Read more <a href="https://cloud.google.com/storage/docs/locations">here</a>.`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"northamerica-northeast1",
								"northamerica-northeast2",
								"us-central1",
								"us-east1",
								"us-east4",
								"us-west1",
								"us-west2",
								"us-west3",
								"us-west4",
								"southamerica-east1",
								"southamerica-west1",
								"europe-central2",
								"europe-north1",
								"europe-west1",
								"europe-west2",
								"europe-west3",
								"europe-west4",
								"europe-west6",
								"asia-east1",
								"asia-east2",
								"asia-northeast1",
								"asia-northeast2",
								"asia-northeast3",
								"asia-south1",
								"asia-south2",
								"asia-southeast1",
								"asia-southeast2",
								"australia-southeast1",
								"australia-southeast2",
								"asia",
								"eu",
								"us",
								"asia1",
								"eur4",
								"nam4",
							),
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

func (r *DestinationGcsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *DestinationGcsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationGcsResourceModel
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
	res, err := r.client.Destinations.CreateDestinationGcs(ctx, request)
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

func (r *DestinationGcsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationGcsResourceModel
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
	request := operations.GetDestinationGcsRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationGcs(ctx, request)
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

func (r *DestinationGcsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationGcsResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationGcsPutRequest := data.ToUpdateSDKType()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationGcsRequest{
		DestinationGcsPutRequest: destinationGcsPutRequest,
		DestinationID:            destinationID,
	}
	res, err := r.client.Destinations.PutDestinationGcs(ctx, request)
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
	getRequest := operations.GetDestinationGcsRequest{
		DestinationID: destinationId1,
	}
	getResponse, err := r.client.Destinations.GetDestinationGcs(ctx, getRequest)
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

func (r *DestinationGcsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationGcsResourceModel
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
	request := operations.DeleteDestinationGcsRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationGcs(ctx, request)
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

func (r *DestinationGcsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
