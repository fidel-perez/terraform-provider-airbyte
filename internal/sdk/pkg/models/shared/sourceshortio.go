// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SourceShortioShortio string

const (
	SourceShortioShortioShortio SourceShortioShortio = "shortio"
)

func (e SourceShortioShortio) ToPointer() *SourceShortioShortio {
	return &e
}

func (e *SourceShortioShortio) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "shortio":
		*e = SourceShortioShortio(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceShortioShortio: %v", v)
	}
}

type SourceShortio struct {
	DomainID string `json:"domain_id"`
	// Short.io Secret Key
	SecretKey  string               `json:"secret_key"`
	sourceType SourceShortioShortio `const:"shortio" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}

func (s SourceShortio) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceShortio) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceShortio) GetDomainID() string {
	if o == nil {
		return ""
	}
	return o.DomainID
}

func (o *SourceShortio) GetSecretKey() string {
	if o == nil {
		return ""
	}
	return o.SecretKey
}

func (o *SourceShortio) GetSourceType() SourceShortioShortio {
	return SourceShortioShortioShortio
}

func (o *SourceShortio) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
