// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleDrive struct {
	Credentials SourceGoogleDriveAuthentication          `tfsdk:"credentials"`
	FolderURL   types.String                             `tfsdk:"folder_url"`
	StartDate   types.String                             `tfsdk:"start_date"`
	Streams     []SourceGoogleDriveFileBasedStreamConfig `tfsdk:"streams"`
}
