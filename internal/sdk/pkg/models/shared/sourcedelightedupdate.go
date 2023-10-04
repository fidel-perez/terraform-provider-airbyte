// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceDelightedUpdate struct {
	// A Delighted API key.
	APIKey string `json:"api_key"`
	// The date from which you'd like to replicate the data
	Since time.Time `json:"since"`
}

func (s SourceDelightedUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDelightedUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDelightedUpdate) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceDelightedUpdate) GetSince() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Since
}
