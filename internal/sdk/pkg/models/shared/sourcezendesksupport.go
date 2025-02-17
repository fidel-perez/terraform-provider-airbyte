// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceZendeskSupportSchemasCredentials string

const (
	SourceZendeskSupportSchemasCredentialsAPIToken SourceZendeskSupportSchemasCredentials = "api_token"
)

func (e SourceZendeskSupportSchemasCredentials) ToPointer() *SourceZendeskSupportSchemasCredentials {
	return &e
}

func (e *SourceZendeskSupportSchemasCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskSupportSchemasCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportSchemasCredentials: %v", v)
	}
}

// SourceZendeskSupportAPIToken - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportAPIToken struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The value of the API token generated. See our <a href="https://docs.airbyte.com/integrations/sources/zendesk-support#setup-guide">full documentation</a> for more information on generating this token.
	APIToken    string                                  `json:"api_token"`
	credentials *SourceZendeskSupportSchemasCredentials `const:"api_token" json:"credentials,omitempty"`
	// The user email for your Zendesk account.
	Email string `json:"email"`
}

func (s SourceZendeskSupportAPIToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupportAPIToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupportAPIToken) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskSupportAPIToken) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceZendeskSupportAPIToken) GetCredentials() *SourceZendeskSupportSchemasCredentials {
	return SourceZendeskSupportSchemasCredentialsAPIToken.ToPointer()
}

func (o *SourceZendeskSupportAPIToken) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

type SourceZendeskSupportCredentials string

const (
	SourceZendeskSupportCredentialsOauth20 SourceZendeskSupportCredentials = "oauth2.0"
)

func (e SourceZendeskSupportCredentials) ToPointer() *SourceZendeskSupportCredentials {
	return &e
}

func (e *SourceZendeskSupportCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskSupportCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportCredentials: %v", v)
	}
}

// SourceZendeskSupportOAuth20 - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportOAuth20 struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The OAuth access token. See the <a href="https://developer.zendesk.com/documentation/ticketing/working-with-oauth/creating-and-using-oauth-tokens-with-the-api/">Zendesk docs</a> for more information on generating this token.
	AccessToken string `json:"access_token"`
	// The OAuth client's ID. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientID *string `json:"client_id,omitempty"`
	// The OAuth client secret. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientSecret *string                          `json:"client_secret,omitempty"`
	credentials  *SourceZendeskSupportCredentials `const:"oauth2.0" json:"credentials,omitempty"`
}

func (s SourceZendeskSupportOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupportOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupportOAuth20) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceZendeskSupportOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceZendeskSupportOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *SourceZendeskSupportOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *SourceZendeskSupportOAuth20) GetCredentials() *SourceZendeskSupportCredentials {
	return SourceZendeskSupportCredentialsOauth20.ToPointer()
}

type SourceZendeskSupportAuthenticationType string

const (
	SourceZendeskSupportAuthenticationTypeSourceZendeskSupportOAuth20  SourceZendeskSupportAuthenticationType = "source-zendesk-support_OAuth2.0"
	SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAPIToken SourceZendeskSupportAuthenticationType = "source-zendesk-support_API Token"
)

type SourceZendeskSupportAuthentication struct {
	SourceZendeskSupportOAuth20  *SourceZendeskSupportOAuth20
	SourceZendeskSupportAPIToken *SourceZendeskSupportAPIToken

	Type SourceZendeskSupportAuthenticationType
}

func CreateSourceZendeskSupportAuthenticationSourceZendeskSupportOAuth20(sourceZendeskSupportOAuth20 SourceZendeskSupportOAuth20) SourceZendeskSupportAuthentication {
	typ := SourceZendeskSupportAuthenticationTypeSourceZendeskSupportOAuth20

	return SourceZendeskSupportAuthentication{
		SourceZendeskSupportOAuth20: &sourceZendeskSupportOAuth20,
		Type:                        typ,
	}
}

func CreateSourceZendeskSupportAuthenticationSourceZendeskSupportAPIToken(sourceZendeskSupportAPIToken SourceZendeskSupportAPIToken) SourceZendeskSupportAuthentication {
	typ := SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAPIToken

	return SourceZendeskSupportAuthentication{
		SourceZendeskSupportAPIToken: &sourceZendeskSupportAPIToken,
		Type:                         typ,
	}
}

func (u *SourceZendeskSupportAuthentication) UnmarshalJSON(data []byte) error {

	sourceZendeskSupportAPIToken := new(SourceZendeskSupportAPIToken)
	if err := utils.UnmarshalJSON(data, &sourceZendeskSupportAPIToken, "", true, true); err == nil {
		u.SourceZendeskSupportAPIToken = sourceZendeskSupportAPIToken
		u.Type = SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAPIToken
		return nil
	}

	sourceZendeskSupportOAuth20 := new(SourceZendeskSupportOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceZendeskSupportOAuth20, "", true, true); err == nil {
		u.SourceZendeskSupportOAuth20 = sourceZendeskSupportOAuth20
		u.Type = SourceZendeskSupportAuthenticationTypeSourceZendeskSupportOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskSupportAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskSupportOAuth20 != nil {
		return utils.MarshalJSON(u.SourceZendeskSupportOAuth20, "", true)
	}

	if u.SourceZendeskSupportAPIToken != nil {
		return utils.MarshalJSON(u.SourceZendeskSupportAPIToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type ZendeskSupport string

const (
	ZendeskSupportZendeskSupport ZendeskSupport = "zendesk-support"
)

func (e ZendeskSupport) ToPointer() *ZendeskSupport {
	return &e
}

func (e *ZendeskSupport) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zendesk-support":
		*e = ZendeskSupport(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ZendeskSupport: %v", v)
	}
}

type SourceZendeskSupport struct {
	// Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
	Credentials *SourceZendeskSupportAuthentication `json:"credentials,omitempty"`
	// Makes each stream read a single page of data.
	IgnorePagination *bool          `default:"false" json:"ignore_pagination"`
	sourceType       ZendeskSupport `const:"zendesk-support" json:"sourceType"`
	// The UTC date and time from which you'd like to replicate data, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
	// This is your unique Zendesk subdomain that can be found in your account URL. For example, in https://MY_SUBDOMAIN.zendesk.com/, MY_SUBDOMAIN is the value of your subdomain.
	Subdomain string `json:"subdomain"`
}

func (s SourceZendeskSupport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceZendeskSupport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceZendeskSupport) GetCredentials() *SourceZendeskSupportAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceZendeskSupport) GetIgnorePagination() *bool {
	if o == nil {
		return nil
	}
	return o.IgnorePagination
}

func (o *SourceZendeskSupport) GetSourceType() ZendeskSupport {
	return ZendeskSupportZendeskSupport
}

func (o *SourceZendeskSupport) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceZendeskSupport) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}
