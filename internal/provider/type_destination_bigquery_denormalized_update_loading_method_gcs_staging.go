// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging struct {
	Credential           DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential `tfsdk:"credential"`
	FileBufferCount      types.Int64                                                            `tfsdk:"file_buffer_count"`
	GcsBucketName        types.String                                                           `tfsdk:"gcs_bucket_name"`
	GcsBucketPath        types.String                                                           `tfsdk:"gcs_bucket_path"`
	KeepFilesInGcsBucket types.String                                                           `tfsdk:"keep_files_in_gcs_bucket"`
}
