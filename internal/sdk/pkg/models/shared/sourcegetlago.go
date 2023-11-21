// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Getlago string

const (
	GetlagoGetlago Getlago = "getlago"
)

func (e Getlago) ToPointer() *Getlago {
	return &e
}

func (e *Getlago) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "getlago":
		*e = Getlago(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Getlago: %v", v)
	}
}

type SourceGetlago struct {
	// Your API Key. See <a href="https://doc.getlago.com/docs/api/intro">here</a>.
	APIKey string `json:"api_key"`
	// Your Lago API URL
	APIURL     *string `default:"https://api.getlago.com/api/v1" json:"api_url"`
	sourceType Getlago `const:"getlago" json:"sourceType"`
}

func (s SourceGetlago) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGetlago) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGetlago) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceGetlago) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *SourceGetlago) GetSourceType() Getlago {
	return GetlagoGetlago
}
