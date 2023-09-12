// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleSheetsResourceModel) ToCreateSDKType() *shared.SourceGoogleSheetsCreateRequest {
	var credentials shared.SourceGoogleSheetsAuthentication
	var sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth *shared.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth != nil {
		authType := shared.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuthAuthType(r.Configuration.Credentials.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth.AuthType.ValueString())
		clientID := r.Configuration.Credentials.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth = &shared.SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth{
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleSheetsAuthentication{
			SourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth: sourceGoogleSheetsAuthenticationAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication *shared.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication
	if r.Configuration.Credentials.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication != nil {
		authType1 := shared.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication.AuthType.ValueString())
		serviceAccountInfo := r.Configuration.Credentials.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication = &shared.SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication{
			AuthType:           authType1,
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleSheetsAuthentication{
			SourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication: sourceGoogleSheetsAuthenticationServiceAccountKeyAuthentication,
		}
	}
	namesConversion := new(bool)
	if !r.Configuration.NamesConversion.IsUnknown() && !r.Configuration.NamesConversion.IsNull() {
		*namesConversion = r.Configuration.NamesConversion.ValueBool()
	} else {
		namesConversion = nil
	}
	sourceType := shared.SourceGoogleSheetsGoogleSheets(r.Configuration.SourceType.ValueString())
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.SourceGoogleSheets{
		Credentials:     credentials,
		NamesConversion: namesConversion,
		SourceType:      sourceType,
		SpreadsheetID:   spreadsheetID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleSheetsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleSheetsResourceModel) ToGetSDKType() *shared.SourceGoogleSheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleSheetsResourceModel) ToUpdateSDKType() *shared.SourceGoogleSheetsPutRequest {
	var credentials shared.SourceGoogleSheetsUpdateAuthentication
	var sourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth *shared.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth != nil {
		authType := shared.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuthAuthType(r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth.AuthType.ValueString())
		clientID := r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth = &shared.SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth{
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleSheetsUpdateAuthentication{
			SourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth: sourceGoogleSheetsUpdateAuthenticationAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication *shared.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication
	if r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication != nil {
		authType1 := shared.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthenticationAuthType(r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication.AuthType.ValueString())
		serviceAccountInfo := r.Configuration.Credentials.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication = &shared.SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication{
			AuthType:           authType1,
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleSheetsUpdateAuthentication{
			SourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication: sourceGoogleSheetsUpdateAuthenticationServiceAccountKeyAuthentication,
		}
	}
	namesConversion := new(bool)
	if !r.Configuration.NamesConversion.IsUnknown() && !r.Configuration.NamesConversion.IsNull() {
		*namesConversion = r.Configuration.NamesConversion.ValueBool()
	} else {
		namesConversion = nil
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.SourceGoogleSheetsUpdate{
		Credentials:     credentials,
		NamesConversion: namesConversion,
		SpreadsheetID:   spreadsheetID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleSheetsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleSheetsResourceModel) ToDeleteSDKType() *shared.SourceGoogleSheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleSheetsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGoogleSheetsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
