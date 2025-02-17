// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

// SourceHubspotSchemasAuthType - Name of the credentials set
type SourceHubspotSchemasAuthType string

const (
	SourceHubspotSchemasAuthTypePrivateAppCredentials SourceHubspotSchemasAuthType = "Private App Credentials"
)

func (e SourceHubspotSchemasAuthType) ToPointer() *SourceHubspotSchemasAuthType {
	return &e
}

func (e *SourceHubspotSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Private App Credentials":
		*e = SourceHubspotSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotSchemasAuthType: %v", v)
	}
}

// SourceHubspotPrivateApp - Choose how to authenticate to HubSpot.
type SourceHubspotPrivateApp struct {
	// HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.
	AccessToken string `json:"access_token"`
	// Name of the credentials set
	credentialsTitle SourceHubspotSchemasAuthType `const:"Private App Credentials" json:"credentials_title"`
}

func (s SourceHubspotPrivateApp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHubspotPrivateApp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceHubspotPrivateApp) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceHubspotPrivateApp) GetCredentialsTitle() SourceHubspotSchemasAuthType {
	return SourceHubspotSchemasAuthTypePrivateAppCredentials
}

// SourceHubspotAuthType - Name of the credentials
type SourceHubspotAuthType string

const (
	SourceHubspotAuthTypeOAuthCredentials SourceHubspotAuthType = "OAuth Credentials"
)

func (e SourceHubspotAuthType) ToPointer() *SourceHubspotAuthType {
	return &e
}

func (e *SourceHubspotAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceHubspotAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotAuthType: %v", v)
	}
}

// SourceHubspotOAuth - Choose how to authenticate to HubSpot.
type SourceHubspotOAuth struct {
	// The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.
	ClientID string `json:"client_id"`
	// The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.
	ClientSecret string `json:"client_secret"`
	// Name of the credentials
	credentialsTitle SourceHubspotAuthType `const:"OAuth Credentials" json:"credentials_title"`
	// Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceHubspotOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHubspotOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceHubspotOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceHubspotOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceHubspotOAuth) GetCredentialsTitle() SourceHubspotAuthType {
	return SourceHubspotAuthTypeOAuthCredentials
}

func (o *SourceHubspotOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceHubspotAuthenticationType string

const (
	SourceHubspotAuthenticationTypeSourceHubspotOAuth      SourceHubspotAuthenticationType = "source-hubspot_OAuth"
	SourceHubspotAuthenticationTypeSourceHubspotPrivateApp SourceHubspotAuthenticationType = "source-hubspot_Private App"
)

type SourceHubspotAuthentication struct {
	SourceHubspotOAuth      *SourceHubspotOAuth
	SourceHubspotPrivateApp *SourceHubspotPrivateApp

	Type SourceHubspotAuthenticationType
}

func CreateSourceHubspotAuthenticationSourceHubspotOAuth(sourceHubspotOAuth SourceHubspotOAuth) SourceHubspotAuthentication {
	typ := SourceHubspotAuthenticationTypeSourceHubspotOAuth

	return SourceHubspotAuthentication{
		SourceHubspotOAuth: &sourceHubspotOAuth,
		Type:               typ,
	}
}

func CreateSourceHubspotAuthenticationSourceHubspotPrivateApp(sourceHubspotPrivateApp SourceHubspotPrivateApp) SourceHubspotAuthentication {
	typ := SourceHubspotAuthenticationTypeSourceHubspotPrivateApp

	return SourceHubspotAuthentication{
		SourceHubspotPrivateApp: &sourceHubspotPrivateApp,
		Type:                    typ,
	}
}

func (u *SourceHubspotAuthentication) UnmarshalJSON(data []byte) error {

	sourceHubspotPrivateApp := new(SourceHubspotPrivateApp)
	if err := utils.UnmarshalJSON(data, &sourceHubspotPrivateApp, "", true, true); err == nil {
		u.SourceHubspotPrivateApp = sourceHubspotPrivateApp
		u.Type = SourceHubspotAuthenticationTypeSourceHubspotPrivateApp
		return nil
	}

	sourceHubspotOAuth := new(SourceHubspotOAuth)
	if err := utils.UnmarshalJSON(data, &sourceHubspotOAuth, "", true, true); err == nil {
		u.SourceHubspotOAuth = sourceHubspotOAuth
		u.Type = SourceHubspotAuthenticationTypeSourceHubspotOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHubspotAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceHubspotOAuth != nil {
		return utils.MarshalJSON(u.SourceHubspotOAuth, "", true)
	}

	if u.SourceHubspotPrivateApp != nil {
		return utils.MarshalJSON(u.SourceHubspotPrivateApp, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Hubspot string

const (
	HubspotHubspot Hubspot = "hubspot"
)

func (e Hubspot) ToPointer() *Hubspot {
	return &e
}

func (e *Hubspot) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hubspot":
		*e = Hubspot(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hubspot: %v", v)
	}
}

type SourceHubspot struct {
	// Choose how to authenticate to HubSpot.
	Credentials SourceHubspotAuthentication `json:"credentials"`
	sourceType  Hubspot                     `const:"hubspot" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceHubspot) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceHubspot) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceHubspot) GetCredentials() SourceHubspotAuthentication {
	if o == nil {
		return SourceHubspotAuthentication{}
	}
	return o.Credentials
}

func (o *SourceHubspot) GetSourceType() Hubspot {
	return HubspotHubspot
}

func (o *SourceHubspot) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
