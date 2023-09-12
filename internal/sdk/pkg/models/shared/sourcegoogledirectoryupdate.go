// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle - Authentication Scenario
type SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle string

const (
	SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitleServiceAccounts SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle = "Service accounts"
)

func (e SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle) ToPointer() *SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle {
	return &e
}

func (e *SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service accounts":
		*e = SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle: %v", v)
	}
}

// SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey - For these scenario user should obtain service account's credentials from the Google API Console and provide delegated email.
type SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey struct {
	// The contents of the JSON service account key. See the <a href="https://developers.google.com/admin-sdk/directory/v1/guides/delegation">docs</a> for more information on how to generate this key.
	CredentialsJSON string `json:"credentials_json"`
	// Authentication Scenario
	CredentialsTitle *SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKeyCredentialsTitle `json:"credentials_title,omitempty"`
	// The email of the user, which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
}

// SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle - Authentication Scenario
type SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle string

const (
	SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitleWebServerApp SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle = "Web server app"
)

func (e SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle) ToPointer() *SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle {
	return &e
}

func (e *SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Web server app":
		*e = SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle: %v", v)
	}
}

// SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth - For these scenario user only needs to give permission to read Google Directory data.
type SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth struct {
	// The Client ID of the developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of the developer application.
	ClientSecret string `json:"client_secret"`
	// Authentication Scenario
	CredentialsTitle *SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuthCredentialsTitle `json:"credentials_title,omitempty"`
	// The Token for obtaining a new access token.
	RefreshToken string `json:"refresh_token"`
}

type SourceGoogleDirectoryUpdateGoogleCredentialsType string

const (
	SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth SourceGoogleDirectoryUpdateGoogleCredentialsType = "source-google-directory-update_Google Credentials_Sign in via Google (OAuth)"
	SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey    SourceGoogleDirectoryUpdateGoogleCredentialsType = "source-google-directory-update_Google Credentials_Service Account Key"
)

type SourceGoogleDirectoryUpdateGoogleCredentials struct {
	SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth *SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth
	SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey    *SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey

	Type SourceGoogleDirectoryUpdateGoogleCredentialsType
}

func CreateSourceGoogleDirectoryUpdateGoogleCredentialsSourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth(sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth) SourceGoogleDirectoryUpdateGoogleCredentials {
	typ := SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth

	return SourceGoogleDirectoryUpdateGoogleCredentials{
		SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth: &sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth,
		Type: typ,
	}
}

func CreateSourceGoogleDirectoryUpdateGoogleCredentialsSourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey(sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey) SourceGoogleDirectoryUpdateGoogleCredentials {
	typ := SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey

	return SourceGoogleDirectoryUpdateGoogleCredentials{
		SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey: &sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey,
		Type: typ,
	}
}

func (u *SourceGoogleDirectoryUpdateGoogleCredentials) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey := new(SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey); err == nil {
		u.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey = sourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey
		u.Type = SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey
		return nil
	}

	sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth := new(SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth); err == nil {
		u.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth = sourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth
		u.Type = SourceGoogleDirectoryUpdateGoogleCredentialsTypeSourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleDirectoryUpdateGoogleCredentials) MarshalJSON() ([]byte, error) {
	if u.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey != nil {
		return json.Marshal(u.SourceGoogleDirectoryUpdateGoogleCredentialsServiceAccountKey)
	}

	if u.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth != nil {
		return json.Marshal(u.SourceGoogleDirectoryUpdateGoogleCredentialsSignInViaGoogleOAuth)
	}

	return nil, nil
}

type SourceGoogleDirectoryUpdate struct {
	// Google APIs use the OAuth 2.0 protocol for authentication and authorization. The Source supports <a href="https://developers.google.com/identity/protocols/oauth2#webserver" target="_blank">Web server application</a> and <a href="https://developers.google.com/identity/protocols/oauth2#serviceaccount" target="_blank">Service accounts</a> scenarios.
	Credentials *SourceGoogleDirectoryUpdateGoogleCredentials `json:"credentials,omitempty"`
}
