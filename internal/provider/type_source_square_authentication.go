// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceSquareAuthentication struct {
	SourceSquareAuthenticationAPIKey                    *SourceAirtableAuthenticationPersonalAccessToken     `tfsdk:"source_square_authentication_api_key"`
	SourceSquareAuthenticationOauthAuthentication       *DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"source_square_authentication_oauth_authentication"`
	SourceSquareUpdateAuthenticationAPIKey              *SourceAirtableAuthenticationPersonalAccessToken     `tfsdk:"source_square_update_authentication_api_key"`
	SourceSquareUpdateAuthenticationOauthAuthentication *DestinationGoogleSheetsAuthenticationViaGoogleOAuth `tfsdk:"source_square_update_authentication_oauth_authentication"`
}
