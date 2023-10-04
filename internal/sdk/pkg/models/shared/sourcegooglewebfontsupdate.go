// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceGoogleWebfontsUpdate struct {
	// Optional, Available params- json, media, proto
	Alt *string `json:"alt,omitempty"`
	// API key is required to access google apis, For getting your's goto google console and generate api key for Webfonts
	APIKey string `json:"api_key"`
	// Optional, boolean type
	PrettyPrint *string `json:"prettyPrint,omitempty"`
	// Optional, to find how to sort
	Sort *string `json:"sort,omitempty"`
}

func (o *SourceGoogleWebfontsUpdate) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *SourceGoogleWebfontsUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceGoogleWebfontsUpdate) GetPrettyPrint() *string {
	if o == nil {
		return nil
	}
	return o.PrettyPrint
}

func (o *SourceGoogleWebfontsUpdate) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}
