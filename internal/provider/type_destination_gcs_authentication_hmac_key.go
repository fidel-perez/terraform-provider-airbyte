// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationGcsAuthenticationHMACKey struct {
	CredentialType  types.String `tfsdk:"credential_type"`
	HmacKeyAccessID types.String `tfsdk:"hmac_key_access_id"`
	HmacKeySecret   types.String `tfsdk:"hmac_key_secret"`
}
