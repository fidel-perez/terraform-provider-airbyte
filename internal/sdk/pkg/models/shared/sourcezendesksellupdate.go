// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceZendeskSellUpdate struct {
	// The API token for authenticating to Zendesk Sell
	APIToken string `json:"api_token"`
}

func (o *SourceZendeskSellUpdate) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}
