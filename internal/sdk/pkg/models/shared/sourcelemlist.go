// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SourceLemlistLemlist string

const (
	SourceLemlistLemlistLemlist SourceLemlistLemlist = "lemlist"
)

func (e SourceLemlistLemlist) ToPointer() *SourceLemlistLemlist {
	return &e
}

func (e *SourceLemlistLemlist) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lemlist":
		*e = SourceLemlistLemlist(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLemlistLemlist: %v", v)
	}
}

type SourceLemlist struct {
	// Lemlist API key,
	APIKey     string               `json:"api_key"`
	sourceType SourceLemlistLemlist `const:"lemlist" json:"sourceType"`
}

func (s SourceLemlist) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLemlist) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLemlist) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceLemlist) GetSourceType() SourceLemlistLemlist {
	return SourceLemlistLemlistLemlist
}
