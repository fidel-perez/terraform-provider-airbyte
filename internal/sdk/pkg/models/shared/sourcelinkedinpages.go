// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceLinkedinPagesAuthenticationAccessTokenAuthMethod string

const (
	SourceLinkedinPagesAuthenticationAccessTokenAuthMethodAccessToken SourceLinkedinPagesAuthenticationAccessTokenAuthMethod = "access_token"
)

func (e SourceLinkedinPagesAuthenticationAccessTokenAuthMethod) ToPointer() *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod {
	return &e
}

func (e *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinPagesAuthenticationAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesAuthenticationAccessTokenAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesAuthenticationAccessToken struct {
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	AccessToken string                                                  `json:"access_token"`
	authMethod  *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod `const:"access_token" json:"auth_method,omitempty"`
}

func (s SourceLinkedinPagesAuthenticationAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesAuthenticationAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesAuthenticationAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceLinkedinPagesAuthenticationAccessToken) GetAuthMethod() *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod {
	return SourceLinkedinPagesAuthenticationAccessTokenAuthMethodAccessToken.ToPointer()
}

type SourceLinkedinPagesAuthenticationOAuth20AuthMethod string

const (
	SourceLinkedinPagesAuthenticationOAuth20AuthMethodOAuth20 SourceLinkedinPagesAuthenticationOAuth20AuthMethod = "oAuth2.0"
)

func (e SourceLinkedinPagesAuthenticationOAuth20AuthMethod) ToPointer() *SourceLinkedinPagesAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceLinkedinPagesAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinPagesAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceLinkedinPagesAuthenticationOAuth20 struct {
	authMethod *SourceLinkedinPagesAuthenticationOAuth20AuthMethod `const:"oAuth2.0" json:"auth_method,omitempty"`
	// The client ID of the LinkedIn developer application.
	ClientID string `json:"client_id"`
	// The client secret of the LinkedIn developer application.
	ClientSecret string `json:"client_secret"`
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceLinkedinPagesAuthenticationOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPagesAuthenticationOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPagesAuthenticationOAuth20) GetAuthMethod() *SourceLinkedinPagesAuthenticationOAuth20AuthMethod {
	return SourceLinkedinPagesAuthenticationOAuth20AuthMethodOAuth20.ToPointer()
}

func (o *SourceLinkedinPagesAuthenticationOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceLinkedinPagesAuthenticationOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceLinkedinPagesAuthenticationOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLinkedinPagesAuthenticationType string

const (
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20     SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_Authentication_OAuth2.0"
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_Authentication_Access token"
)

type SourceLinkedinPagesAuthentication struct {
	SourceLinkedinPagesAuthenticationOAuth20     *SourceLinkedinPagesAuthenticationOAuth20
	SourceLinkedinPagesAuthenticationAccessToken *SourceLinkedinPagesAuthenticationAccessToken

	Type SourceLinkedinPagesAuthenticationType
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesAuthenticationOAuth20(sourceLinkedinPagesAuthenticationOAuth20 SourceLinkedinPagesAuthenticationOAuth20) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesAuthenticationOAuth20: &sourceLinkedinPagesAuthenticationOAuth20,
		Type:                                     typ,
	}
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesAuthenticationAccessToken(sourceLinkedinPagesAuthenticationAccessToken SourceLinkedinPagesAuthenticationAccessToken) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesAuthenticationAccessToken: &sourceLinkedinPagesAuthenticationAccessToken,
		Type: typ,
	}
}

func (u *SourceLinkedinPagesAuthentication) UnmarshalJSON(data []byte) error {

	sourceLinkedinPagesAuthenticationAccessToken := new(SourceLinkedinPagesAuthenticationAccessToken)
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesAuthenticationAccessToken, "", true, true); err == nil {
		u.SourceLinkedinPagesAuthenticationAccessToken = sourceLinkedinPagesAuthenticationAccessToken
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken
		return nil
	}

	sourceLinkedinPagesAuthenticationOAuth20 := new(SourceLinkedinPagesAuthenticationOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceLinkedinPagesAuthenticationOAuth20, "", true, true); err == nil {
		u.SourceLinkedinPagesAuthenticationOAuth20 = sourceLinkedinPagesAuthenticationOAuth20
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinPagesAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinPagesAuthenticationOAuth20 != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesAuthenticationOAuth20, "", true)
	}

	if u.SourceLinkedinPagesAuthenticationAccessToken != nil {
		return utils.MarshalJSON(u.SourceLinkedinPagesAuthenticationAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceLinkedinPagesLinkedinPages string

const (
	SourceLinkedinPagesLinkedinPagesLinkedinPages SourceLinkedinPagesLinkedinPages = "linkedin-pages"
)

func (e SourceLinkedinPagesLinkedinPages) ToPointer() *SourceLinkedinPagesLinkedinPages {
	return &e
}

func (e *SourceLinkedinPagesLinkedinPages) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linkedin-pages":
		*e = SourceLinkedinPagesLinkedinPages(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesLinkedinPages: %v", v)
	}
}

type SourceLinkedinPages struct {
	Credentials *SourceLinkedinPagesAuthentication `json:"credentials,omitempty"`
	// Specify the Organization ID
	OrgID      string                           `json:"org_id"`
	sourceType SourceLinkedinPagesLinkedinPages `const:"linkedin-pages" json:"sourceType"`
}

func (s SourceLinkedinPages) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinPages) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinPages) GetCredentials() *SourceLinkedinPagesAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLinkedinPages) GetOrgID() string {
	if o == nil {
		return ""
	}
	return o.OrgID
}

func (o *SourceLinkedinPages) GetSourceType() SourceLinkedinPagesLinkedinPages {
	return SourceLinkedinPagesLinkedinPagesLinkedinPages
}
