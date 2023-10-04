// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod string

const (
	SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodOauth2AccessToken SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod = "oauth2_access_token"
)

func (e SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod) ToPointer() *SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod {
	return &e
}

func (e *SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_access_token":
		*e = SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod: %v", v)
	}
}

type SourceAuth0AuthenticationMethodOAuth2AccessToken struct {
	// Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
	AccessToken string                                                               `json:"access_token"`
	authType    SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod `const:"oauth2_access_token" json:"auth_type"`
}

func (s SourceAuth0AuthenticationMethodOAuth2AccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0AuthenticationMethodOAuth2AccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0AuthenticationMethodOAuth2AccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceAuth0AuthenticationMethodOAuth2AccessToken) GetAuthType() SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethod {
	return SourceAuth0AuthenticationMethodOAuth2AccessTokenAuthenticationMethodOauth2AccessToken
}

type SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod string

const (
	SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodOauth2ConfidentialApplication SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod = "oauth2_confidential_application"
)

func (e SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod) ToPointer() *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod {
	return &e
}

func (e *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2_confidential_application":
		*e = SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod: %v", v)
	}
}

type SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication struct {
	// The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
	Audience string                                                                           `json:"audience"`
	authType SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod `const:"oauth2_confidential_application" json:"auth_type"`
	// Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientID string `json:"client_id"`
	// Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
	ClientSecret string `json:"client_secret"`
}

func (s SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) GetAudience() string {
	if o == nil {
		return ""
	}
	return o.Audience
}

func (o *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) GetAuthType() SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethod {
	return SourceAuth0AuthenticationMethodOAuth2ConfidentialApplicationAuthenticationMethodOauth2ConfidentialApplication
}

func (o *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

type SourceAuth0AuthenticationMethodType string

const (
	SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication SourceAuth0AuthenticationMethodType = "source-auth0_Authentication Method_OAuth2 Confidential Application"
	SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken             SourceAuth0AuthenticationMethodType = "source-auth0_Authentication Method_OAuth2 Access Token"
)

type SourceAuth0AuthenticationMethod struct {
	SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication *SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
	SourceAuth0AuthenticationMethodOAuth2AccessToken             *SourceAuth0AuthenticationMethodOAuth2AccessToken

	Type SourceAuth0AuthenticationMethodType
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication(sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication

	return SourceAuth0AuthenticationMethod{
		SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication: &sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication,
		Type: typ,
	}
}

func CreateSourceAuth0AuthenticationMethodSourceAuth0AuthenticationMethodOAuth2AccessToken(sourceAuth0AuthenticationMethodOAuth2AccessToken SourceAuth0AuthenticationMethodOAuth2AccessToken) SourceAuth0AuthenticationMethod {
	typ := SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken

	return SourceAuth0AuthenticationMethod{
		SourceAuth0AuthenticationMethodOAuth2AccessToken: &sourceAuth0AuthenticationMethodOAuth2AccessToken,
		Type: typ,
	}
}

func (u *SourceAuth0AuthenticationMethod) UnmarshalJSON(data []byte) error {

	sourceAuth0AuthenticationMethodOAuth2AccessToken := new(SourceAuth0AuthenticationMethodOAuth2AccessToken)
	if err := utils.UnmarshalJSON(data, &sourceAuth0AuthenticationMethodOAuth2AccessToken, "", true, true); err == nil {
		u.SourceAuth0AuthenticationMethodOAuth2AccessToken = sourceAuth0AuthenticationMethodOAuth2AccessToken
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2AccessToken
		return nil
	}

	sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication := new(SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication)
	if err := utils.UnmarshalJSON(data, &sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication, "", true, true); err == nil {
		u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication = sourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
		u.Type = SourceAuth0AuthenticationMethodTypeSourceAuth0AuthenticationMethodOAuth2ConfidentialApplication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAuth0AuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication != nil {
		return utils.MarshalJSON(u.SourceAuth0AuthenticationMethodOAuth2ConfidentialApplication, "", true)
	}

	if u.SourceAuth0AuthenticationMethodOAuth2AccessToken != nil {
		return utils.MarshalJSON(u.SourceAuth0AuthenticationMethodOAuth2AccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceAuth0Auth0 string

const (
	SourceAuth0Auth0Auth0 SourceAuth0Auth0 = "auth0"
)

func (e SourceAuth0Auth0) ToPointer() *SourceAuth0Auth0 {
	return &e
}

func (e *SourceAuth0Auth0) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "auth0":
		*e = SourceAuth0Auth0(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAuth0Auth0: %v", v)
	}
}

type SourceAuth0 struct {
	// The Authentication API is served over HTTPS. All URLs referenced in the documentation have the following base `https://YOUR_DOMAIN`
	BaseURL     string                          `json:"base_url"`
	Credentials SourceAuth0AuthenticationMethod `json:"credentials"`
	sourceType  SourceAuth0Auth0                `const:"auth0" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate *string `default:"2023-08-05T00:43:59.244Z" json:"start_date"`
}

func (s SourceAuth0) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAuth0) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAuth0) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *SourceAuth0) GetCredentials() SourceAuth0AuthenticationMethod {
	if o == nil {
		return SourceAuth0AuthenticationMethod{}
	}
	return o.Credentials
}

func (o *SourceAuth0) GetSourceType() SourceAuth0Auth0 {
	return SourceAuth0Auth0Auth0
}

func (o *SourceAuth0) GetStartDate() *string {
	if o == nil {
		return nil
	}
	return o.StartDate
}
