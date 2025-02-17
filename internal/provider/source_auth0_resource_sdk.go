// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAuth0ResourceModel) ToCreateSDKType() *shared.SourceAuth0CreateRequest {
	baseURL := r.Configuration.BaseURL.ValueString()
	var credentials shared.SourceAuth0AuthenticationMethod
	var sourceAuth0OAuth2ConfidentialApplication *shared.SourceAuth0OAuth2ConfidentialApplication
	if r.Configuration.Credentials.OAuth2ConfidentialApplication != nil {
		audience := r.Configuration.Credentials.OAuth2ConfidentialApplication.Audience.ValueString()
		clientID := r.Configuration.Credentials.OAuth2ConfidentialApplication.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth2ConfidentialApplication.ClientSecret.ValueString()
		sourceAuth0OAuth2ConfidentialApplication = &shared.SourceAuth0OAuth2ConfidentialApplication{
			Audience:     audience,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if sourceAuth0OAuth2ConfidentialApplication != nil {
		credentials = shared.SourceAuth0AuthenticationMethod{
			SourceAuth0OAuth2ConfidentialApplication: sourceAuth0OAuth2ConfidentialApplication,
		}
	}
	var sourceAuth0OAuth2AccessToken *shared.SourceAuth0OAuth2AccessToken
	if r.Configuration.Credentials.OAuth2AccessToken != nil {
		accessToken := r.Configuration.Credentials.OAuth2AccessToken.AccessToken.ValueString()
		sourceAuth0OAuth2AccessToken = &shared.SourceAuth0OAuth2AccessToken{
			AccessToken: accessToken,
		}
	}
	if sourceAuth0OAuth2AccessToken != nil {
		credentials = shared.SourceAuth0AuthenticationMethod{
			SourceAuth0OAuth2AccessToken: sourceAuth0OAuth2AccessToken,
		}
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	configuration := shared.SourceAuth0{
		BaseURL:     baseURL,
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
	out := shared.SourceAuth0CreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAuth0ResourceModel) ToGetSDKType() *shared.SourceAuth0CreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAuth0ResourceModel) ToUpdateSDKType() *shared.SourceAuth0PutRequest {
	baseURL := r.Configuration.BaseURL.ValueString()
	var credentials shared.SourceAuth0UpdateAuthenticationMethod
	var oAuth2ConfidentialApplication *shared.OAuth2ConfidentialApplication
	if r.Configuration.Credentials.OAuth2ConfidentialApplication != nil {
		audience := r.Configuration.Credentials.OAuth2ConfidentialApplication.Audience.ValueString()
		clientID := r.Configuration.Credentials.OAuth2ConfidentialApplication.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth2ConfidentialApplication.ClientSecret.ValueString()
		oAuth2ConfidentialApplication = &shared.OAuth2ConfidentialApplication{
			Audience:     audience,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if oAuth2ConfidentialApplication != nil {
		credentials = shared.SourceAuth0UpdateAuthenticationMethod{
			OAuth2ConfidentialApplication: oAuth2ConfidentialApplication,
		}
	}
	var oAuth2AccessToken *shared.OAuth2AccessToken
	if r.Configuration.Credentials.OAuth2AccessToken != nil {
		accessToken := r.Configuration.Credentials.OAuth2AccessToken.AccessToken.ValueString()
		oAuth2AccessToken = &shared.OAuth2AccessToken{
			AccessToken: accessToken,
		}
	}
	if oAuth2AccessToken != nil {
		credentials = shared.SourceAuth0UpdateAuthenticationMethod{
			OAuth2AccessToken: oAuth2AccessToken,
		}
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	configuration := shared.SourceAuth0Update{
		BaseURL:     baseURL,
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAuth0PutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAuth0ResourceModel) ToDeleteSDKType() *shared.SourceAuth0CreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAuth0ResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAuth0ResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
