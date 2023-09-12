// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceIntercom struct {
	AccessToken  types.String `tfsdk:"access_token"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	SourceType   types.String `tfsdk:"source_type"`
	StartDate    types.String `tfsdk:"start_date"`
}
