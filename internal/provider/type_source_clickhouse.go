// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceClickhouse struct {
	Database     types.String                          `tfsdk:"database"`
	Host         types.String                          `tfsdk:"host"`
	Password     types.String                          `tfsdk:"password"`
	Port         types.Int64                           `tfsdk:"port"`
	TunnelMethod *DestinationClickhouseSSHTunnelMethod `tfsdk:"tunnel_method"`
	Username     types.String                          `tfsdk:"username"`
}
