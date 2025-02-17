// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceHubspotResourceModel) ToCreateSDKType() *shared.SourceHubspotCreateRequest {
	var credentials shared.SourceHubspotAuthentication
	var sourceHubspotOAuth *shared.SourceHubspotOAuth
	if r.Configuration.Credentials.OAuth != nil {
		clientID := r.Configuration.Credentials.OAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth.RefreshToken.ValueString()
		sourceHubspotOAuth = &shared.SourceHubspotOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceHubspotOAuth != nil {
		credentials = shared.SourceHubspotAuthentication{
			SourceHubspotOAuth: sourceHubspotOAuth,
		}
	}
	var sourceHubspotPrivateApp *shared.SourceHubspotPrivateApp
	if r.Configuration.Credentials.PrivateApp != nil {
		accessToken := r.Configuration.Credentials.PrivateApp.AccessToken.ValueString()
		sourceHubspotPrivateApp = &shared.SourceHubspotPrivateApp{
			AccessToken: accessToken,
		}
	}
	if sourceHubspotPrivateApp != nil {
		credentials = shared.SourceHubspotAuthentication{
			SourceHubspotPrivateApp: sourceHubspotPrivateApp,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceHubspot{
		Credentials: credentials,
		StartDate:   startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHubspotCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHubspotResourceModel) ToGetSDKType() *shared.SourceHubspotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHubspotResourceModel) ToUpdateSDKType() *shared.SourceHubspotPutRequest {
	var credentials shared.SourceHubspotUpdateAuthentication
	var sourceHubspotUpdateOAuth *shared.SourceHubspotUpdateOAuth
	if r.Configuration.Credentials.OAuth != nil {
		clientID := r.Configuration.Credentials.OAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth.RefreshToken.ValueString()
		sourceHubspotUpdateOAuth = &shared.SourceHubspotUpdateOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceHubspotUpdateOAuth != nil {
		credentials = shared.SourceHubspotUpdateAuthentication{
			SourceHubspotUpdateOAuth: sourceHubspotUpdateOAuth,
		}
	}
	var privateApp *shared.PrivateApp
	if r.Configuration.Credentials.PrivateApp != nil {
		accessToken := r.Configuration.Credentials.PrivateApp.AccessToken.ValueString()
		privateApp = &shared.PrivateApp{
			AccessToken: accessToken,
		}
	}
	if privateApp != nil {
		credentials = shared.SourceHubspotUpdateAuthentication{
			PrivateApp: privateApp,
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceHubspotUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHubspotPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHubspotResourceModel) ToDeleteSDKType() *shared.SourceHubspotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHubspotResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceHubspotResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
