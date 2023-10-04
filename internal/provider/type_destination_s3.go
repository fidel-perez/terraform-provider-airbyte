// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationS3 struct {
	AccessKeyID     types.String              `tfsdk:"access_key_id"`
	FileNamePattern types.String              `tfsdk:"file_name_pattern"`
	Format          DestinationS3OutputFormat `tfsdk:"format"`
	S3BucketName    types.String              `tfsdk:"s3_bucket_name"`
	S3BucketPath    types.String              `tfsdk:"s3_bucket_path"`
	S3BucketRegion  types.String              `tfsdk:"s3_bucket_region"`
	S3Endpoint      types.String              `tfsdk:"s3_endpoint"`
	S3PathFormat    types.String              `tfsdk:"s3_path_format"`
	SecretAccessKey types.String              `tfsdk:"secret_access_key"`
}
