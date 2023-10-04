// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

type SourceDelightedDelighted string

const (
	SourceDelightedDelightedDelighted SourceDelightedDelighted = "delighted"
)

func (e SourceDelightedDelighted) ToPointer() *SourceDelightedDelighted {
	return &e
}

func (e *SourceDelightedDelighted) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delighted":
		*e = SourceDelightedDelighted(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDelightedDelighted: %v", v)
	}
}

type SourceDelighted struct {
	// A Delighted API key.
	APIKey string `json:"api_key"`
	// The date from which you'd like to replicate the data
	Since      time.Time                `json:"since"`
	sourceType SourceDelightedDelighted `const:"delighted" json:"sourceType"`
}

func (s SourceDelighted) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceDelighted) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceDelighted) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceDelighted) GetSince() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Since
}

func (o *SourceDelighted) GetSourceType() SourceDelightedDelighted {
	return SourceDelightedDelightedDelighted
}
