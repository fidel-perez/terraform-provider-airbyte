// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGitlabResourceModel) ToCreateSDKType() *shared.SourceGitlabCreateRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabAuthorizationMethod
	var sourceGitlabAuthorizationMethodOAuth20 *shared.SourceGitlabAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceGitlabAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceGitlabAuthorizationMethodOAuth20 = &shared.SourceGitlabAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabAuthorizationMethodOAuth20: sourceGitlabAuthorizationMethodOAuth20,
		}
	}
	var sourceGitlabAuthorizationMethodPrivateToken *shared.SourceGitlabAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceGitlabAuthorizationMethodPrivateToken.AccessToken.ValueString()
		sourceGitlabAuthorizationMethodPrivateToken = &shared.SourceGitlabAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
		}
	}
	if sourceGitlabAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceGitlabAuthorizationMethod{
			SourceGitlabAuthorizationMethodPrivateToken: sourceGitlabAuthorizationMethodPrivateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGitlab{
		APIURL:      apiURL,
		Credentials: credentials,
		Groups:      groups,
		Projects:    projects,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGitlabCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGitlabResourceModel) ToGetSDKType() *shared.SourceGitlabCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGitlabResourceModel) ToUpdateSDKType() *shared.SourceGitlabPutRequest {
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	var credentials shared.SourceGitlabUpdateAuthorizationMethod
	var sourceGitlabUpdateAuthorizationMethodOAuth20 *shared.SourceGitlabUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceGitlabUpdateAuthorizationMethodOAuth20 = &shared.SourceGitlabUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceGitlabUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			SourceGitlabUpdateAuthorizationMethodOAuth20: sourceGitlabUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourceGitlabUpdateAuthorizationMethodPrivateToken *shared.SourceGitlabUpdateAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceGitlabUpdateAuthorizationMethodPrivateToken.AccessToken.ValueString()
		sourceGitlabUpdateAuthorizationMethodPrivateToken = &shared.SourceGitlabUpdateAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
		}
	}
	if sourceGitlabUpdateAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceGitlabUpdateAuthorizationMethod{
			SourceGitlabUpdateAuthorizationMethodPrivateToken: sourceGitlabUpdateAuthorizationMethodPrivateToken,
		}
	}
	groups := new(string)
	if !r.Configuration.Groups.IsUnknown() && !r.Configuration.Groups.IsNull() {
		*groups = r.Configuration.Groups.ValueString()
	} else {
		groups = nil
	}
	projects := new(string)
	if !r.Configuration.Projects.IsUnknown() && !r.Configuration.Projects.IsNull() {
		*projects = r.Configuration.Projects.ValueString()
	} else {
		projects = nil
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceGitlabUpdate{
		APIURL:      apiURL,
		Credentials: credentials,
		Groups:      groups,
		Projects:    projects,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGitlabPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGitlabResourceModel) ToDeleteSDKType() *shared.SourceGitlabCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGitlabResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGitlabResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
