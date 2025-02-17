// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// SourceGoogleDirectorySchemasCredentialsTitle - Authentication Scenario
type SourceGoogleDirectorySchemasCredentialsTitle string

const (
	SourceGoogleDirectorySchemasCredentialsTitleServiceAccounts SourceGoogleDirectorySchemasCredentialsTitle = "Service accounts"
)

func (e SourceGoogleDirectorySchemasCredentialsTitle) ToPointer() *SourceGoogleDirectorySchemasCredentialsTitle {
	return &e
}

func (e *SourceGoogleDirectorySchemasCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service accounts":
		*e = SourceGoogleDirectorySchemasCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleDirectorySchemasCredentialsTitle: %v", v)
	}
}

// SourceGoogleDirectoryServiceAccountKey - For these scenario user should obtain service account's credentials from the Google API Console and provide delegated email.
type SourceGoogleDirectoryServiceAccountKey struct {
	// The contents of the JSON service account key. See the <a href="https://developers.google.com/admin-sdk/directory/v1/guides/delegation">docs</a> for more information on how to generate this key.
	CredentialsJSON string `json:"credentials_json"`
	// Authentication Scenario
	credentialsTitle *SourceGoogleDirectorySchemasCredentialsTitle `const:"Service accounts" json:"credentials_title,omitempty"`
	// The email of the user, which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
}

func (s SourceGoogleDirectoryServiceAccountKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleDirectoryServiceAccountKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleDirectoryServiceAccountKey) GetCredentialsJSON() string {
	if o == nil {
		return ""
	}
	return o.CredentialsJSON
}

func (o *SourceGoogleDirectoryServiceAccountKey) GetCredentialsTitle() *SourceGoogleDirectorySchemasCredentialsTitle {
	return SourceGoogleDirectorySchemasCredentialsTitleServiceAccounts.ToPointer()
}

func (o *SourceGoogleDirectoryServiceAccountKey) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

// SourceGoogleDirectoryCredentialsTitle - Authentication Scenario
type SourceGoogleDirectoryCredentialsTitle string

const (
	SourceGoogleDirectoryCredentialsTitleWebServerApp SourceGoogleDirectoryCredentialsTitle = "Web server app"
)

func (e SourceGoogleDirectoryCredentialsTitle) ToPointer() *SourceGoogleDirectoryCredentialsTitle {
	return &e
}

func (e *SourceGoogleDirectoryCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Web server app":
		*e = SourceGoogleDirectoryCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleDirectoryCredentialsTitle: %v", v)
	}
}

// SourceGoogleDirectorySignInViaGoogleOAuth - For these scenario user only needs to give permission to read Google Directory data.
type SourceGoogleDirectorySignInViaGoogleOAuth struct {
	// The Client ID of the developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of the developer application.
	ClientSecret string `json:"client_secret"`
	// Authentication Scenario
	credentialsTitle *SourceGoogleDirectoryCredentialsTitle `const:"Web server app" json:"credentials_title,omitempty"`
	// The Token for obtaining a new access token.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceGoogleDirectorySignInViaGoogleOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleDirectorySignInViaGoogleOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleDirectorySignInViaGoogleOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceGoogleDirectorySignInViaGoogleOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceGoogleDirectorySignInViaGoogleOAuth) GetCredentialsTitle() *SourceGoogleDirectoryCredentialsTitle {
	return SourceGoogleDirectoryCredentialsTitleWebServerApp.ToPointer()
}

func (o *SourceGoogleDirectorySignInViaGoogleOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceGoogleDirectoryGoogleCredentialsType string

const (
	SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectorySignInViaGoogleOAuth SourceGoogleDirectoryGoogleCredentialsType = "source-google-directory_Sign in via Google (OAuth)"
	SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectoryServiceAccountKey    SourceGoogleDirectoryGoogleCredentialsType = "source-google-directory_Service Account Key"
)

type SourceGoogleDirectoryGoogleCredentials struct {
	SourceGoogleDirectorySignInViaGoogleOAuth *SourceGoogleDirectorySignInViaGoogleOAuth
	SourceGoogleDirectoryServiceAccountKey    *SourceGoogleDirectoryServiceAccountKey

	Type SourceGoogleDirectoryGoogleCredentialsType
}

func CreateSourceGoogleDirectoryGoogleCredentialsSourceGoogleDirectorySignInViaGoogleOAuth(sourceGoogleDirectorySignInViaGoogleOAuth SourceGoogleDirectorySignInViaGoogleOAuth) SourceGoogleDirectoryGoogleCredentials {
	typ := SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectorySignInViaGoogleOAuth

	return SourceGoogleDirectoryGoogleCredentials{
		SourceGoogleDirectorySignInViaGoogleOAuth: &sourceGoogleDirectorySignInViaGoogleOAuth,
		Type: typ,
	}
}

func CreateSourceGoogleDirectoryGoogleCredentialsSourceGoogleDirectoryServiceAccountKey(sourceGoogleDirectoryServiceAccountKey SourceGoogleDirectoryServiceAccountKey) SourceGoogleDirectoryGoogleCredentials {
	typ := SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectoryServiceAccountKey

	return SourceGoogleDirectoryGoogleCredentials{
		SourceGoogleDirectoryServiceAccountKey: &sourceGoogleDirectoryServiceAccountKey,
		Type:                                   typ,
	}
}

func (u *SourceGoogleDirectoryGoogleCredentials) UnmarshalJSON(data []byte) error {

	sourceGoogleDirectoryServiceAccountKey := new(SourceGoogleDirectoryServiceAccountKey)
	if err := utils.UnmarshalJSON(data, &sourceGoogleDirectoryServiceAccountKey, "", true, true); err == nil {
		u.SourceGoogleDirectoryServiceAccountKey = sourceGoogleDirectoryServiceAccountKey
		u.Type = SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectoryServiceAccountKey
		return nil
	}

	sourceGoogleDirectorySignInViaGoogleOAuth := new(SourceGoogleDirectorySignInViaGoogleOAuth)
	if err := utils.UnmarshalJSON(data, &sourceGoogleDirectorySignInViaGoogleOAuth, "", true, true); err == nil {
		u.SourceGoogleDirectorySignInViaGoogleOAuth = sourceGoogleDirectorySignInViaGoogleOAuth
		u.Type = SourceGoogleDirectoryGoogleCredentialsTypeSourceGoogleDirectorySignInViaGoogleOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleDirectoryGoogleCredentials) MarshalJSON() ([]byte, error) {
	if u.SourceGoogleDirectorySignInViaGoogleOAuth != nil {
		return utils.MarshalJSON(u.SourceGoogleDirectorySignInViaGoogleOAuth, "", true)
	}

	if u.SourceGoogleDirectoryServiceAccountKey != nil {
		return utils.MarshalJSON(u.SourceGoogleDirectoryServiceAccountKey, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GoogleDirectory string

const (
	GoogleDirectoryGoogleDirectory GoogleDirectory = "google-directory"
)

func (e GoogleDirectory) ToPointer() *GoogleDirectory {
	return &e
}

func (e *GoogleDirectory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-directory":
		*e = GoogleDirectory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GoogleDirectory: %v", v)
	}
}

type SourceGoogleDirectory struct {
	// Google APIs use the OAuth 2.0 protocol for authentication and authorization. The Source supports <a href="https://developers.google.com/identity/protocols/oauth2#webserver" target="_blank">Web server application</a> and <a href="https://developers.google.com/identity/protocols/oauth2#serviceaccount" target="_blank">Service accounts</a> scenarios.
	Credentials *SourceGoogleDirectoryGoogleCredentials `json:"credentials,omitempty"`
	sourceType  GoogleDirectory                         `const:"google-directory" json:"sourceType"`
}

func (s SourceGoogleDirectory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleDirectory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleDirectory) GetCredentials() *SourceGoogleDirectoryGoogleCredentials {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceGoogleDirectory) GetSourceType() GoogleDirectory {
	return GoogleDirectoryGoogleDirectory
}
