// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceHarvestResourceModel) ToCreateSDKType() *shared.SourceHarvestCreateRequest {
	accountID := r.Configuration.AccountID.ValueString()
	var credentials *shared.SourceHarvestAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth *shared.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth
		if r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			clientID := r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth.RefreshToken.ValueString()
			sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth = &shared.SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth{
				AdditionalProperties: additionalProperties,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
			}
		}
		if sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
			credentials = &shared.SourceHarvestAuthenticationMechanism{
				SourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth: sourceHarvestAuthenticationMechanismAuthenticateViaHarvestOAuth,
			}
		}
		var sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken *shared.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken.APIToken.ValueString()
			sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken = &shared.SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
			}
		}
		if sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
			credentials = &shared.SourceHarvestAuthenticationMechanism{
				SourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken: sourceHarvestAuthenticationMechanismAuthenticateWithPersonalAccessToken,
			}
		}
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourceHarvest{
		AccountID:            accountID,
		Credentials:          credentials,
		ReplicationEndDate:   replicationEndDate,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHarvestCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHarvestResourceModel) ToGetSDKType() *shared.SourceHarvestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHarvestResourceModel) ToUpdateSDKType() *shared.SourceHarvestPutRequest {
	accountID := r.Configuration.AccountID.ValueString()
	var credentials *shared.SourceHarvestUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth *shared.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth
		if r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.AdditionalProperties.ValueString()), &additionalProperties)
			}
			clientID := r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth.RefreshToken.ValueString()
			sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth = &shared.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth{
				AdditionalProperties: additionalProperties,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				RefreshToken:         refreshToken,
			}
		}
		if sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
			credentials = &shared.SourceHarvestUpdateAuthenticationMechanism{
				SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth: sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth,
			}
		}
		var sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken *shared.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.APIToken.ValueString()
			sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken = &shared.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
			}
		}
		if sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
			credentials = &shared.SourceHarvestUpdateAuthenticationMechanism{
				SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken: sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken,
			}
		}
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourceHarvestUpdate{
		AccountID:            accountID,
		Credentials:          credentials,
		ReplicationEndDate:   replicationEndDate,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceHarvestPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceHarvestResourceModel) ToDeleteSDKType() *shared.SourceHarvestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceHarvestResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceHarvestResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
