// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SourceK6CloudK6Cloud string

const (
	SourceK6CloudK6CloudK6Cloud SourceK6CloudK6Cloud = "k6-cloud"
)

func (e SourceK6CloudK6Cloud) ToPointer() *SourceK6CloudK6Cloud {
	return &e
}

func (e *SourceK6CloudK6Cloud) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "k6-cloud":
		*e = SourceK6CloudK6Cloud(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceK6CloudK6Cloud: %v", v)
	}
}

type SourceK6Cloud struct {
	// Your API Token. See <a href="https://k6.io/docs/cloud/integrations/token/">here</a>. The key is case sensitive.
	APIToken   string               `json:"api_token"`
	sourceType SourceK6CloudK6Cloud `const:"k6-cloud" json:"sourceType"`
}

func (s SourceK6Cloud) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceK6Cloud) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceK6Cloud) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceK6Cloud) GetSourceType() SourceK6CloudK6Cloud {
	return SourceK6CloudK6CloudK6Cloud
}
