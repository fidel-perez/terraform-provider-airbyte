// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Smaily string

const (
	SmailySmaily Smaily = "smaily"
)

func (e Smaily) ToPointer() *Smaily {
	return &e
}

func (e *Smaily) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smaily":
		*e = Smaily(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Smaily: %v", v)
	}
}

type SourceSmaily struct {
	// API user password. See https://smaily.com/help/api/general/create-api-user/
	APIPassword string `json:"api_password"`
	// API Subdomain. See https://smaily.com/help/api/general/create-api-user/
	APISubdomain string `json:"api_subdomain"`
	// API user username. See https://smaily.com/help/api/general/create-api-user/
	APIUsername string `json:"api_username"`
	sourceType  Smaily `const:"smaily" json:"sourceType"`
}

func (s SourceSmaily) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSmaily) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSmaily) GetAPIPassword() string {
	if o == nil {
		return ""
	}
	return o.APIPassword
}

func (o *SourceSmaily) GetAPISubdomain() string {
	if o == nil {
		return ""
	}
	return o.APISubdomain
}

func (o *SourceSmaily) GetAPIUsername() string {
	if o == nil {
		return ""
	}
	return o.APIUsername
}

func (o *SourceSmaily) GetSourceType() Smaily {
	return SmailySmaily
}
