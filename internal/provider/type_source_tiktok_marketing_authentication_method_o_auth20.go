// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTiktokMarketingAuthenticationMethodOAuth20 struct {
	AccessToken  types.String `tfsdk:"access_token"`
	AdvertiserID types.String `tfsdk:"advertiser_id"`
	AppID        types.String `tfsdk:"app_id"`
	Secret       types.String `tfsdk:"secret"`
}
