// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceMicrosoftTeamsUpdateSchemasAuthType string

const (
	SourceMicrosoftTeamsUpdateSchemasAuthTypeToken SourceMicrosoftTeamsUpdateSchemasAuthType = "Token"
)

func (e SourceMicrosoftTeamsUpdateSchemasAuthType) ToPointer() *SourceMicrosoftTeamsUpdateSchemasAuthType {
	return &e
}

func (e *SourceMicrosoftTeamsUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceMicrosoftTeamsUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsUpdateSchemasAuthType: %v", v)
	}
}

// AuthenticateViaMicrosoft - Choose how to authenticate to Microsoft
type AuthenticateViaMicrosoft struct {
	authType *SourceMicrosoftTeamsUpdateSchemasAuthType `const:"Token" json:"auth_type"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

func (a AuthenticateViaMicrosoft) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaMicrosoft) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaMicrosoft) GetAuthType() *SourceMicrosoftTeamsUpdateSchemasAuthType {
	return SourceMicrosoftTeamsUpdateSchemasAuthTypeToken.ToPointer()
}

func (o *AuthenticateViaMicrosoft) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *AuthenticateViaMicrosoft) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *AuthenticateViaMicrosoft) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

type SourceMicrosoftTeamsUpdateAuthType string

const (
	SourceMicrosoftTeamsUpdateAuthTypeClient SourceMicrosoftTeamsUpdateAuthType = "Client"
)

func (e SourceMicrosoftTeamsUpdateAuthType) ToPointer() *SourceMicrosoftTeamsUpdateAuthType {
	return &e
}

func (e *SourceMicrosoftTeamsUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceMicrosoftTeamsUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMicrosoftTeamsUpdateAuthType: %v", v)
	}
}

// AuthenticateViaMicrosoftOAuth20 - Choose how to authenticate to Microsoft
type AuthenticateViaMicrosoftOAuth20 struct {
	authType *SourceMicrosoftTeamsUpdateAuthType `const:"Client" json:"auth_type"`
	// The Client ID of your Microsoft Teams developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Teams developer application.
	ClientSecret string `json:"client_secret"`
	// A Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// A globally unique identifier (GUID) that is different than your organization name or domain. Follow these steps to obtain: open one of the Teams where you belong inside the Teams Application -> Click on the … next to the Team title -> Click on Get link to team -> Copy the link to the team and grab the tenant ID form the URL
	TenantID string `json:"tenant_id"`
}

func (a AuthenticateViaMicrosoftOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthenticateViaMicrosoftOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AuthenticateViaMicrosoftOAuth20) GetAuthType() *SourceMicrosoftTeamsUpdateAuthType {
	return SourceMicrosoftTeamsUpdateAuthTypeClient.ToPointer()
}

func (o *AuthenticateViaMicrosoftOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *AuthenticateViaMicrosoftOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *AuthenticateViaMicrosoftOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *AuthenticateViaMicrosoftOAuth20) GetTenantID() string {
	if o == nil {
		return ""
	}
	return o.TenantID
}

type SourceMicrosoftTeamsUpdateAuthenticationMechanismType string

const (
	SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoftOAuth20 SourceMicrosoftTeamsUpdateAuthenticationMechanismType = "Authenticate via Microsoft (OAuth 2.0)"
	SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoft        SourceMicrosoftTeamsUpdateAuthenticationMechanismType = "Authenticate via Microsoft"
)

type SourceMicrosoftTeamsUpdateAuthenticationMechanism struct {
	AuthenticateViaMicrosoftOAuth20 *AuthenticateViaMicrosoftOAuth20
	AuthenticateViaMicrosoft        *AuthenticateViaMicrosoft

	Type SourceMicrosoftTeamsUpdateAuthenticationMechanismType
}

func CreateSourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoftOAuth20(authenticateViaMicrosoftOAuth20 AuthenticateViaMicrosoftOAuth20) SourceMicrosoftTeamsUpdateAuthenticationMechanism {
	typ := SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoftOAuth20

	return SourceMicrosoftTeamsUpdateAuthenticationMechanism{
		AuthenticateViaMicrosoftOAuth20: &authenticateViaMicrosoftOAuth20,
		Type:                            typ,
	}
}

func CreateSourceMicrosoftTeamsUpdateAuthenticationMechanismAuthenticateViaMicrosoft(authenticateViaMicrosoft AuthenticateViaMicrosoft) SourceMicrosoftTeamsUpdateAuthenticationMechanism {
	typ := SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoft

	return SourceMicrosoftTeamsUpdateAuthenticationMechanism{
		AuthenticateViaMicrosoft: &authenticateViaMicrosoft,
		Type:                     typ,
	}
}

func (u *SourceMicrosoftTeamsUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {

	authenticateViaMicrosoft := new(AuthenticateViaMicrosoft)
	if err := utils.UnmarshalJSON(data, &authenticateViaMicrosoft, "", true, true); err == nil {
		u.AuthenticateViaMicrosoft = authenticateViaMicrosoft
		u.Type = SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoft
		return nil
	}

	authenticateViaMicrosoftOAuth20 := new(AuthenticateViaMicrosoftOAuth20)
	if err := utils.UnmarshalJSON(data, &authenticateViaMicrosoftOAuth20, "", true, true); err == nil {
		u.AuthenticateViaMicrosoftOAuth20 = authenticateViaMicrosoftOAuth20
		u.Type = SourceMicrosoftTeamsUpdateAuthenticationMechanismTypeAuthenticateViaMicrosoftOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMicrosoftTeamsUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.AuthenticateViaMicrosoftOAuth20 != nil {
		return utils.MarshalJSON(u.AuthenticateViaMicrosoftOAuth20, "", true)
	}

	if u.AuthenticateViaMicrosoft != nil {
		return utils.MarshalJSON(u.AuthenticateViaMicrosoft, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMicrosoftTeamsUpdate struct {
	// Choose how to authenticate to Microsoft
	Credentials *SourceMicrosoftTeamsUpdateAuthenticationMechanism `json:"credentials,omitempty"`
	// Specifies the length of time over which the Team Device Report stream is aggregated. The supported values are: D7, D30, D90, and D180.
	Period string `json:"period"`
}

func (o *SourceMicrosoftTeamsUpdate) GetCredentials() *SourceMicrosoftTeamsUpdateAuthenticationMechanism {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceMicrosoftTeamsUpdate) GetPeriod() string {
	if o == nil {
		return ""
	}
	return o.Period
}
