// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTrustpilotResourceModel) ToCreateSDKType() *shared.SourceTrustpilotCreateRequest {
	var businessUnits []string = nil
	for _, businessUnitsItem := range r.Configuration.BusinessUnits {
		businessUnits = append(businessUnits, businessUnitsItem.ValueString())
	}
	var credentials shared.SourceTrustpilotAuthorizationMethod
	var sourceTrustpilotOAuth20 *shared.SourceTrustpilotOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceTrustpilotOAuth20 = &shared.SourceTrustpilotOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTrustpilotOAuth20 != nil {
		credentials = shared.SourceTrustpilotAuthorizationMethod{
			SourceTrustpilotOAuth20: sourceTrustpilotOAuth20,
		}
	}
	var sourceTrustpilotAPIKey *shared.SourceTrustpilotAPIKey
	if r.Configuration.Credentials.APIKey != nil {
		clientId1 := r.Configuration.Credentials.APIKey.ClientID.ValueString()
		sourceTrustpilotAPIKey = &shared.SourceTrustpilotAPIKey{
			ClientID: clientId1,
		}
	}
	if sourceTrustpilotAPIKey != nil {
		credentials = shared.SourceTrustpilotAuthorizationMethod{
			SourceTrustpilotAPIKey: sourceTrustpilotAPIKey,
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceTrustpilot{
		BusinessUnits: businessUnits,
		Credentials:   credentials,
		StartDate:     startDate,
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
	out := shared.SourceTrustpilotCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrustpilotResourceModel) ToGetSDKType() *shared.SourceTrustpilotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrustpilotResourceModel) ToUpdateSDKType() *shared.SourceTrustpilotPutRequest {
	var businessUnits []string = nil
	for _, businessUnitsItem := range r.Configuration.BusinessUnits {
		businessUnits = append(businessUnits, businessUnitsItem.ValueString())
	}
	var credentials shared.SourceTrustpilotUpdateAuthorizationMethod
	var sourceTrustpilotUpdateOAuth20 *shared.SourceTrustpilotUpdateOAuth20
	if r.Configuration.Credentials.OAuth20 != nil {
		accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
		clientID := r.Configuration.Credentials.OAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.OAuth20.TokenExpiryDate.ValueString())
		sourceTrustpilotUpdateOAuth20 = &shared.SourceTrustpilotUpdateOAuth20{
			AccessToken:     accessToken,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTrustpilotUpdateOAuth20 != nil {
		credentials = shared.SourceTrustpilotUpdateAuthorizationMethod{
			SourceTrustpilotUpdateOAuth20: sourceTrustpilotUpdateOAuth20,
		}
	}
	var sourceTrustpilotUpdateAPIKey *shared.SourceTrustpilotUpdateAPIKey
	if r.Configuration.Credentials.APIKey != nil {
		clientId1 := r.Configuration.Credentials.APIKey.ClientID.ValueString()
		sourceTrustpilotUpdateAPIKey = &shared.SourceTrustpilotUpdateAPIKey{
			ClientID: clientId1,
		}
	}
	if sourceTrustpilotUpdateAPIKey != nil {
		credentials = shared.SourceTrustpilotUpdateAuthorizationMethod{
			SourceTrustpilotUpdateAPIKey: sourceTrustpilotUpdateAPIKey,
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceTrustpilotUpdate{
		BusinessUnits: businessUnits,
		Credentials:   credentials,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrustpilotPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrustpilotResourceModel) ToDeleteSDKType() *shared.SourceTrustpilotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrustpilotResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTrustpilotResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
