// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceLeverHiringSchemasAuthType string

const (
	SourceLeverHiringSchemasAuthTypeAPIKey SourceLeverHiringSchemasAuthType = "Api Key"
)

func (e SourceLeverHiringSchemasAuthType) ToPointer() *SourceLeverHiringSchemasAuthType {
	return &e
}

func (e *SourceLeverHiringSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Api Key":
		*e = SourceLeverHiringSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringSchemasAuthType: %v", v)
	}
}

// SourceLeverHiringAuthenticateViaLeverAPIKey - Choose how to authenticate to Lever Hiring.
type SourceLeverHiringAuthenticateViaLeverAPIKey struct {
	// The Api Key of your Lever Hiring account.
	APIKey   string                            `json:"api_key"`
	authType *SourceLeverHiringSchemasAuthType `const:"Api Key" json:"auth_type,omitempty"`
}

func (s SourceLeverHiringAuthenticateViaLeverAPIKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLeverHiringAuthenticateViaLeverAPIKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLeverHiringAuthenticateViaLeverAPIKey) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceLeverHiringAuthenticateViaLeverAPIKey) GetAuthType() *SourceLeverHiringSchemasAuthType {
	return SourceLeverHiringSchemasAuthTypeAPIKey.ToPointer()
}

type SourceLeverHiringAuthType string

const (
	SourceLeverHiringAuthTypeClient SourceLeverHiringAuthType = "Client"
)

func (e SourceLeverHiringAuthType) ToPointer() *SourceLeverHiringAuthType {
	return &e
}

func (e *SourceLeverHiringAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceLeverHiringAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringAuthType: %v", v)
	}
}

// SourceLeverHiringAuthenticateViaLeverOAuth - Choose how to authenticate to Lever Hiring.
type SourceLeverHiringAuthenticateViaLeverOAuth struct {
	authType *SourceLeverHiringAuthType `const:"Client" json:"auth_type,omitempty"`
	// The Client ID of your Lever Hiring developer application.
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your Lever Hiring developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
	// The token for obtaining new access token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceLeverHiringAuthenticateViaLeverOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLeverHiringAuthenticateViaLeverOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLeverHiringAuthenticateViaLeverOAuth) GetAuthType() *SourceLeverHiringAuthType {
	return SourceLeverHiringAuthTypeClient.ToPointer()
}

func (o *SourceLeverHiringAuthenticateViaLeverOAuth) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceLeverHiringAuthenticateViaLeverOAuth) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceLeverHiringAuthenticateViaLeverOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLeverHiringAuthenticationMechanismType string

const (
	SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverOAuth  SourceLeverHiringAuthenticationMechanismType = "source-lever-hiring_Authenticate via Lever (OAuth)"
	SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverAPIKey SourceLeverHiringAuthenticationMechanismType = "source-lever-hiring_Authenticate via Lever (Api Key)"
)

type SourceLeverHiringAuthenticationMechanism struct {
	SourceLeverHiringAuthenticateViaLeverOAuth  *SourceLeverHiringAuthenticateViaLeverOAuth
	SourceLeverHiringAuthenticateViaLeverAPIKey *SourceLeverHiringAuthenticateViaLeverAPIKey

	Type SourceLeverHiringAuthenticationMechanismType
}

func CreateSourceLeverHiringAuthenticationMechanismSourceLeverHiringAuthenticateViaLeverOAuth(sourceLeverHiringAuthenticateViaLeverOAuth SourceLeverHiringAuthenticateViaLeverOAuth) SourceLeverHiringAuthenticationMechanism {
	typ := SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverOAuth

	return SourceLeverHiringAuthenticationMechanism{
		SourceLeverHiringAuthenticateViaLeverOAuth: &sourceLeverHiringAuthenticateViaLeverOAuth,
		Type: typ,
	}
}

func CreateSourceLeverHiringAuthenticationMechanismSourceLeverHiringAuthenticateViaLeverAPIKey(sourceLeverHiringAuthenticateViaLeverAPIKey SourceLeverHiringAuthenticateViaLeverAPIKey) SourceLeverHiringAuthenticationMechanism {
	typ := SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverAPIKey

	return SourceLeverHiringAuthenticationMechanism{
		SourceLeverHiringAuthenticateViaLeverAPIKey: &sourceLeverHiringAuthenticateViaLeverAPIKey,
		Type: typ,
	}
}

func (u *SourceLeverHiringAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	sourceLeverHiringAuthenticateViaLeverAPIKey := new(SourceLeverHiringAuthenticateViaLeverAPIKey)
	if err := utils.UnmarshalJSON(data, &sourceLeverHiringAuthenticateViaLeverAPIKey, "", true, true); err == nil {
		u.SourceLeverHiringAuthenticateViaLeverAPIKey = sourceLeverHiringAuthenticateViaLeverAPIKey
		u.Type = SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverAPIKey
		return nil
	}

	sourceLeverHiringAuthenticateViaLeverOAuth := new(SourceLeverHiringAuthenticateViaLeverOAuth)
	if err := utils.UnmarshalJSON(data, &sourceLeverHiringAuthenticateViaLeverOAuth, "", true, true); err == nil {
		u.SourceLeverHiringAuthenticateViaLeverOAuth = sourceLeverHiringAuthenticateViaLeverOAuth
		u.Type = SourceLeverHiringAuthenticationMechanismTypeSourceLeverHiringAuthenticateViaLeverOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLeverHiringAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceLeverHiringAuthenticateViaLeverOAuth != nil {
		return utils.MarshalJSON(u.SourceLeverHiringAuthenticateViaLeverOAuth, "", true)
	}

	if u.SourceLeverHiringAuthenticateViaLeverAPIKey != nil {
		return utils.MarshalJSON(u.SourceLeverHiringAuthenticateViaLeverAPIKey, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceLeverHiringEnvironment - The environment in which you'd like to replicate data for Lever. This is used to determine which Lever API endpoint to use.
type SourceLeverHiringEnvironment string

const (
	SourceLeverHiringEnvironmentProduction SourceLeverHiringEnvironment = "Production"
	SourceLeverHiringEnvironmentSandbox    SourceLeverHiringEnvironment = "Sandbox"
)

func (e SourceLeverHiringEnvironment) ToPointer() *SourceLeverHiringEnvironment {
	return &e
}

func (e *SourceLeverHiringEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Production":
		fallthrough
	case "Sandbox":
		*e = SourceLeverHiringEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLeverHiringEnvironment: %v", v)
	}
}

type LeverHiring string

const (
	LeverHiringLeverHiring LeverHiring = "lever-hiring"
)

func (e LeverHiring) ToPointer() *LeverHiring {
	return &e
}

func (e *LeverHiring) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "lever-hiring":
		*e = LeverHiring(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LeverHiring: %v", v)
	}
}

type SourceLeverHiring struct {
	// Choose how to authenticate to Lever Hiring.
	Credentials *SourceLeverHiringAuthenticationMechanism `json:"credentials,omitempty"`
	// The environment in which you'd like to replicate data for Lever. This is used to determine which Lever API endpoint to use.
	Environment *SourceLeverHiringEnvironment `default:"Sandbox" json:"environment"`
	sourceType  LeverHiring                   `const:"lever-hiring" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. Note that it will be used only in the following incremental streams: comments, commits, and issues.
	StartDate string `json:"start_date"`
}

func (s SourceLeverHiring) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLeverHiring) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLeverHiring) GetCredentials() *SourceLeverHiringAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLeverHiring) GetEnvironment() *SourceLeverHiringEnvironment {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *SourceLeverHiring) GetSourceType() LeverHiring {
	return LeverHiringLeverHiring
}

func (o *SourceLeverHiring) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
