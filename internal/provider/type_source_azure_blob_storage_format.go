// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceAzureBlobStorageFormat struct {
	AvroFormat                         *AvroFormat                                               `tfsdk:"avro_format"`
	CSVFormat                          *CSVFormat                                                `tfsdk:"csv_format"`
	DocumentFileTypeFormatExperimental *DocumentFileTypeFormatExperimental                       `tfsdk:"document_file_type_format_experimental"`
	JsonlFormat                        *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON `tfsdk:"jsonl_format"`
	ParquetFormat                      *ParquetFormat                                            `tfsdk:"parquet_format"`
}
