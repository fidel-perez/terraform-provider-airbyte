// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"airbyte/internal/sdk/pkg/utils"
)

type SourceSnapchatMarketingUpdate struct {
	// The Client ID of your Snapchat developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Snapchat developer application.
	ClientSecret string `json:"client_secret"`
	// Date in the format 2017-01-25. Any data after this date will not be replicated.
	EndDate *types.Date `json:"end_date,omitempty"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// Date in the format 2022-01-01. Any data before this date will not be replicated.
	StartDate *types.Date `default:"2022-01-01" json:"start_date"`
}

func (s SourceSnapchatMarketingUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSnapchatMarketingUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSnapchatMarketingUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSnapchatMarketingUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSnapchatMarketingUpdate) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceSnapchatMarketingUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceSnapchatMarketingUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
