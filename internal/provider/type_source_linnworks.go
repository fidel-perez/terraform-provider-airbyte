// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLinnworks struct {
	ApplicationID     types.String `tfsdk:"application_id"`
	ApplicationSecret types.String `tfsdk:"application_secret"`
	StartDate         types.String `tfsdk:"start_date"`
	Token             types.String `tfsdk:"token"`
}
