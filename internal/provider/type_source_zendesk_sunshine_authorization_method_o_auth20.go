// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskSunshineAuthorizationMethodOAuth20 struct {
	AccessToken  types.String `tfsdk:"access_token"`
	AuthMethod   types.String `tfsdk:"auth_method"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
}
