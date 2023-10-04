// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SourceOmnisendOmnisend string

const (
	SourceOmnisendOmnisendOmnisend SourceOmnisendOmnisend = "omnisend"
)

func (e SourceOmnisendOmnisend) ToPointer() *SourceOmnisendOmnisend {
	return &e
}

func (e *SourceOmnisendOmnisend) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "omnisend":
		*e = SourceOmnisendOmnisend(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOmnisendOmnisend: %v", v)
	}
}

type SourceOmnisend struct {
	// API Key
	APIKey     string                 `json:"api_key"`
	sourceType SourceOmnisendOmnisend `const:"omnisend" json:"sourceType"`
}

func (s SourceOmnisend) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOmnisend) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOmnisend) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceOmnisend) GetSourceType() SourceOmnisendOmnisend {
	return SourceOmnisendOmnisendOmnisend
}
