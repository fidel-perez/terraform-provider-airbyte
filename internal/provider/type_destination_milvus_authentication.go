// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type DestinationMilvusAuthentication struct {
	APIToken         *DestinationMilvusAPIToken                                `tfsdk:"api_token"`
	NoAuth           *DestinationAzureBlobStorageJSONLinesNewlineDelimitedJSON `tfsdk:"no_auth"`
	UsernamePassword *UsernamePassword                                         `tfsdk:"username_password"`
}
