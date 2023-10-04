// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod string

const (
	SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethodAPIToken SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod = "api_token"
)

func (e SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod) ToPointer() *SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod {
	return &e
}

func (e *SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod: %v", v)
	}
}

type SourceZendeskSunshineUpdateAuthorizationMethodAPIToken struct {
	// API Token. See the <a href="https://docs.airbyte.com/integrations/sources/zendesk_sunshine">docs</a> for information on how to generate this key.
	APIToken   string                                                            `json:"api_token"`
	authMethod *SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod `const:"api_token" json:"auth_method"`
	// The user email for your Zendesk account
	Email string `json:"email"`
}

func (s SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) GetAuthMethod() *SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethod {
	return SourceZendeskSunshineUpdateAuthorizationMethodAPITokenAuthMethodAPIToken.ToPointer()
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

type SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod string

const (
	SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethodOauth20 SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod = "oauth2.0"
)

func (e SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod) ToPointer() *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod {
	return &e
}

func (e *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod: %v", v)
	}
}

type SourceZendeskSunshineUpdateAuthorizationMethodOAuth20 struct {
	// Long-term access Token for making authenticated requests.
	AccessToken string                                                           `json:"access_token"`
	authMethod  *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod `const:"oauth2.0" json:"auth_method"`
	// The Client ID of your OAuth application.
	ClientID string `json:"client_id"`
	// The Client Secret of your OAuth application.
	ClientSecret string `json:"client_secret"`
}

func (s SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) GetAuthMethod() *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethod {
	return SourceZendeskSunshineUpdateAuthorizationMethodOAuth20AuthMethodOauth20.ToPointer()
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

type SourceZendeskSunshineUpdateAuthorizationMethodType string

const (
	SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodOAuth20  SourceZendeskSunshineUpdateAuthorizationMethodType = "source-zendesk-sunshine-update_Authorization Method_OAuth2.0"
	SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodAPIToken SourceZendeskSunshineUpdateAuthorizationMethodType = "source-zendesk-sunshine-update_Authorization Method_API Token"
)

type SourceZendeskSunshineUpdateAuthorizationMethod struct {
	SourceZendeskSunshineUpdateAuthorizationMethodOAuth20  *SourceZendeskSunshineUpdateAuthorizationMethodOAuth20
	SourceZendeskSunshineUpdateAuthorizationMethodAPIToken *SourceZendeskSunshineUpdateAuthorizationMethodAPIToken

	Type SourceZendeskSunshineUpdateAuthorizationMethodType
}

func CreateSourceZendeskSunshineUpdateAuthorizationMethodSourceZendeskSunshineUpdateAuthorizationMethodOAuth20(sourceZendeskSunshineUpdateAuthorizationMethodOAuth20 SourceZendeskSunshineUpdateAuthorizationMethodOAuth20) SourceZendeskSunshineUpdateAuthorizationMethod {
	typ := SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodOAuth20

	return SourceZendeskSunshineUpdateAuthorizationMethod{
		SourceZendeskSunshineUpdateAuthorizationMethodOAuth20: &sourceZendeskSunshineUpdateAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceZendeskSunshineUpdateAuthorizationMethodSourceZendeskSunshineUpdateAuthorizationMethodAPIToken(sourceZendeskSunshineUpdateAuthorizationMethodAPIToken SourceZendeskSunshineUpdateAuthorizationMethodAPIToken) SourceZendeskSunshineUpdateAuthorizationMethod {
	typ := SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodAPIToken

	return SourceZendeskSunshineUpdateAuthorizationMethod{
		SourceZendeskSunshineUpdateAuthorizationMethodAPIToken: &sourceZendeskSunshineUpdateAuthorizationMethodAPIToken,
		Type: typ,
	}
}

func (u *SourceZendeskSunshineUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {

	sourceZendeskSunshineUpdateAuthorizationMethodAPIToken := new(SourceZendeskSunshineUpdateAuthorizationMethodAPIToken)
	if err := utils.UnmarshalJSON(data, &sourceZendeskSunshineUpdateAuthorizationMethodAPIToken, "", true, true); err == nil {
		u.SourceZendeskSunshineUpdateAuthorizationMethodAPIToken = sourceZendeskSunshineUpdateAuthorizationMethodAPIToken
		u.Type = SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodAPIToken
		return nil
	}

	sourceZendeskSunshineUpdateAuthorizationMethodOAuth20 := new(SourceZendeskSunshineUpdateAuthorizationMethodOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceZendeskSunshineUpdateAuthorizationMethodOAuth20, "", true, true); err == nil {
		u.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20 = sourceZendeskSunshineUpdateAuthorizationMethodOAuth20
		u.Type = SourceZendeskSunshineUpdateAuthorizationMethodTypeSourceZendeskSunshineUpdateAuthorizationMethodOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskSunshineUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20 != nil {
		return utils.MarshalJSON(u.SourceZendeskSunshineUpdateAuthorizationMethodOAuth20, "", true)
	}

	if u.SourceZendeskSunshineUpdateAuthorizationMethodAPIToken != nil {
		return utils.MarshalJSON(u.SourceZendeskSunshineUpdateAuthorizationMethodAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceZendeskSunshineUpdate struct {
	Credentials *SourceZendeskSunshineUpdateAuthorizationMethod `json:"credentials,omitempty"`
	// The date from which you'd like to replicate data for Zendesk Sunshine API, in the format YYYY-MM-DDT00:00:00Z.
	StartDate time.Time `json:"start_date"`
	// The subdomain for your Zendesk Account.
	Subdomain string `json:"subdomain"`
}

func (s SourceZendeskSunshineUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSunshineUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSunshineUpdate) GetCredentials() *SourceZendeskSunshineUpdateAuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskSunshineUpdate) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}

func (o *SourceZendeskSunshineUpdate) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
