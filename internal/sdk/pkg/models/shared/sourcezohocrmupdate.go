// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

// SourceZohoCrmUpdateDataCenterLocation - Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>
type SourceZohoCrmUpdateDataCenterLocation string

const (
	SourceZohoCrmUpdateDataCenterLocationUs SourceZohoCrmUpdateDataCenterLocation = "US"
	SourceZohoCrmUpdateDataCenterLocationAu SourceZohoCrmUpdateDataCenterLocation = "AU"
	SourceZohoCrmUpdateDataCenterLocationEu SourceZohoCrmUpdateDataCenterLocation = "EU"
	SourceZohoCrmUpdateDataCenterLocationIn SourceZohoCrmUpdateDataCenterLocation = "IN"
	SourceZohoCrmUpdateDataCenterLocationCn SourceZohoCrmUpdateDataCenterLocation = "CN"
	SourceZohoCrmUpdateDataCenterLocationJp SourceZohoCrmUpdateDataCenterLocation = "JP"
)

func (e SourceZohoCrmUpdateDataCenterLocation) ToPointer() *SourceZohoCrmUpdateDataCenterLocation {
	return &e
}

func (e *SourceZohoCrmUpdateDataCenterLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "AU":
		fallthrough
	case "EU":
		fallthrough
	case "IN":
		fallthrough
	case "CN":
		fallthrough
	case "JP":
		*e = SourceZohoCrmUpdateDataCenterLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCrmUpdateDataCenterLocation: %v", v)
	}
}

// SourceZohoCRMUpdateZohoCRMEdition - Choose your Edition of Zoho CRM to determine API Concurrency Limits
type SourceZohoCRMUpdateZohoCRMEdition string

const (
	SourceZohoCRMUpdateZohoCRMEditionFree         SourceZohoCRMUpdateZohoCRMEdition = "Free"
	SourceZohoCRMUpdateZohoCRMEditionStandard     SourceZohoCRMUpdateZohoCRMEdition = "Standard"
	SourceZohoCRMUpdateZohoCRMEditionProfessional SourceZohoCRMUpdateZohoCRMEdition = "Professional"
	SourceZohoCRMUpdateZohoCRMEditionEnterprise   SourceZohoCRMUpdateZohoCRMEdition = "Enterprise"
	SourceZohoCRMUpdateZohoCRMEditionUltimate     SourceZohoCRMUpdateZohoCRMEdition = "Ultimate"
)

func (e SourceZohoCRMUpdateZohoCRMEdition) ToPointer() *SourceZohoCRMUpdateZohoCRMEdition {
	return &e
}

func (e *SourceZohoCRMUpdateZohoCRMEdition) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Free":
		fallthrough
	case "Standard":
		fallthrough
	case "Professional":
		fallthrough
	case "Enterprise":
		fallthrough
	case "Ultimate":
		*e = SourceZohoCRMUpdateZohoCRMEdition(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCRMUpdateZohoCRMEdition: %v", v)
	}
}

// SourceZohoCrmUpdateEnvironment - Please choose the environment
type SourceZohoCrmUpdateEnvironment string

const (
	SourceZohoCrmUpdateEnvironmentProduction SourceZohoCrmUpdateEnvironment = "Production"
	SourceZohoCrmUpdateEnvironmentDeveloper  SourceZohoCrmUpdateEnvironment = "Developer"
	SourceZohoCrmUpdateEnvironmentSandbox    SourceZohoCrmUpdateEnvironment = "Sandbox"
)

func (e SourceZohoCrmUpdateEnvironment) ToPointer() *SourceZohoCrmUpdateEnvironment {
	return &e
}

func (e *SourceZohoCrmUpdateEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Production":
		fallthrough
	case "Developer":
		fallthrough
	case "Sandbox":
		*e = SourceZohoCrmUpdateEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZohoCrmUpdateEnvironment: %v", v)
	}
}

type SourceZohoCrmUpdate struct {
	// OAuth2.0 Client ID
	ClientID string `json:"client_id"`
	// OAuth2.0 Client Secret
	ClientSecret string `json:"client_secret"`
	// Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>
	DcRegion SourceZohoCrmUpdateDataCenterLocation `json:"dc_region"`
	// Choose your Edition of Zoho CRM to determine API Concurrency Limits
	Edition *SourceZohoCRMUpdateZohoCRMEdition `default:"Free" json:"edition"`
	// Please choose the environment
	Environment SourceZohoCrmUpdateEnvironment `json:"environment"`
	// OAuth2.0 Refresh Token
	RefreshToken string `json:"refresh_token"`
	// ISO 8601, for instance: `YYYY-MM-DD`, `YYYY-MM-DD HH:MM:SS+HH:MM`
	StartDatetime *time.Time `json:"start_datetime,omitempty"`
}

func (s SourceZohoCrmUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZohoCrmUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZohoCrmUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceZohoCrmUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceZohoCrmUpdate) GetDcRegion() SourceZohoCrmUpdateDataCenterLocation {
	if o == nil {
		return SourceZohoCrmUpdateDataCenterLocation("")
	}
	return o.DcRegion
}

func (o *SourceZohoCrmUpdate) GetEdition() *SourceZohoCRMUpdateZohoCRMEdition {
	if o == nil {
		return nil
	}
	return o.Edition
}

func (o *SourceZohoCrmUpdate) GetEnvironment() SourceZohoCrmUpdateEnvironment {
	if o == nil {
		return SourceZohoCrmUpdateEnvironment("")
	}
	return o.Environment
}

func (o *SourceZohoCrmUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceZohoCrmUpdate) GetStartDatetime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDatetime
}
