// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSalesforce struct {
	AuthType        types.String                      `tfsdk:"auth_type"`
	ClientID        types.String                      `tfsdk:"client_id"`
	ClientSecret    types.String                      `tfsdk:"client_secret"`
	IsSandbox       types.Bool                        `tfsdk:"is_sandbox"`
	RefreshToken    types.String                      `tfsdk:"refresh_token"`
	SourceType      types.String                      `tfsdk:"source_type"`
	StartDate       types.String                      `tfsdk:"start_date"`
	StreamsCriteria []SourceSalesforceStreamsCriteria `tfsdk:"streams_criteria"`
}
