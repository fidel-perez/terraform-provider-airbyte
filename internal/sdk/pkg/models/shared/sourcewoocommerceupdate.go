// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"airbyte/internal/sdk/pkg/utils"
)

type SourceWoocommerceUpdate struct {
	// Customer Key for API in WooCommerce shop
	APIKey string `json:"api_key"`
	// Customer Secret for API in WooCommerce shop
	APISecret string `json:"api_secret"`
	// The name of the store. For https://EXAMPLE.com, the shop name is 'EXAMPLE.com'.
	Shop string `json:"shop"`
	// The date you would like to replicate data from. Format: YYYY-MM-DD
	StartDate types.Date `json:"start_date"`
}

func (s SourceWoocommerceUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceWoocommerceUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceWoocommerceUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceWoocommerceUpdate) GetAPISecret() string {
	if o == nil {
		return ""
	}
	return o.APISecret
}

func (o *SourceWoocommerceUpdate) GetShop() string {
	if o == nil {
		return ""
	}
	return o.Shop
}

func (o *SourceWoocommerceUpdate) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}
