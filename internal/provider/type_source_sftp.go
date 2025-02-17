// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSftp struct {
	Credentials *SourceSftpAuthenticationWildcard `tfsdk:"credentials"`
	FilePattern types.String                      `tfsdk:"file_pattern"`
	FileTypes   types.String                      `tfsdk:"file_types"`
	FolderPath  types.String                      `tfsdk:"folder_path"`
	Host        types.String                      `tfsdk:"host"`
	Port        types.Int64                       `tfsdk:"port"`
	User        types.String                      `tfsdk:"user"`
}
