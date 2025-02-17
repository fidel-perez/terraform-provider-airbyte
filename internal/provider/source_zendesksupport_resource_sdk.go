// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskSupportResourceModel) ToCreateSDKType() *shared.SourceZendeskSupportCreateRequest {
	var credentials *shared.SourceZendeskSupportAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskSupportOAuth20 *shared.SourceZendeskSupportOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.OAuth20.AdditionalProperties.ValueString()), &additionalProperties)
			}
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceZendeskSupportOAuth20 = &shared.SourceZendeskSupportOAuth20{
				AdditionalProperties: additionalProperties,
				AccessToken:          accessToken,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
			}
		}
		if sourceZendeskSupportOAuth20 != nil {
			credentials = &shared.SourceZendeskSupportAuthentication{
				SourceZendeskSupportOAuth20: sourceZendeskSupportOAuth20,
			}
		}
		var sourceZendeskSupportAPIToken *shared.SourceZendeskSupportAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.APIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.APIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.APIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskSupportAPIToken = &shared.SourceZendeskSupportAPIToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
				Email:                email,
			}
		}
		if sourceZendeskSupportAPIToken != nil {
			credentials = &shared.SourceZendeskSupportAuthentication{
				SourceZendeskSupportAPIToken: sourceZendeskSupportAPIToken,
			}
		}
	}
	ignorePagination := new(bool)
	if !r.Configuration.IgnorePagination.IsUnknown() && !r.Configuration.IgnorePagination.IsNull() {
		*ignorePagination = r.Configuration.IgnorePagination.ValueBool()
	} else {
		ignorePagination = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSupport{
		Credentials:      credentials,
		IgnorePagination: ignorePagination,
		StartDate:        startDate,
		Subdomain:        subdomain,
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
	out := shared.SourceZendeskSupportCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSupportResourceModel) ToGetSDKType() *shared.SourceZendeskSupportCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSupportResourceModel) ToUpdateSDKType() *shared.SourceZendeskSupportPutRequest {
	var credentials *shared.SourceZendeskSupportUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskSupportUpdateOAuth20 *shared.SourceZendeskSupportUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			var additionalProperties interface{}
			if !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.OAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.OAuth20.AdditionalProperties.ValueString()), &additionalProperties)
			}
			accessToken := r.Configuration.Credentials.OAuth20.AccessToken.ValueString()
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			sourceZendeskSupportUpdateOAuth20 = &shared.SourceZendeskSupportUpdateOAuth20{
				AdditionalProperties: additionalProperties,
				AccessToken:          accessToken,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
			}
		}
		if sourceZendeskSupportUpdateOAuth20 != nil {
			credentials = &shared.SourceZendeskSupportUpdateAuthentication{
				SourceZendeskSupportUpdateOAuth20: sourceZendeskSupportUpdateOAuth20,
			}
		}
		var sourceZendeskSupportUpdateAPIToken *shared.SourceZendeskSupportUpdateAPIToken
		if r.Configuration.Credentials.APIToken != nil {
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.APIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.APIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.APIToken.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			apiToken := r.Configuration.Credentials.APIToken.APIToken.ValueString()
			email := r.Configuration.Credentials.APIToken.Email.ValueString()
			sourceZendeskSupportUpdateAPIToken = &shared.SourceZendeskSupportUpdateAPIToken{
				AdditionalProperties: additionalProperties1,
				APIToken:             apiToken,
				Email:                email,
			}
		}
		if sourceZendeskSupportUpdateAPIToken != nil {
			credentials = &shared.SourceZendeskSupportUpdateAuthentication{
				SourceZendeskSupportUpdateAPIToken: sourceZendeskSupportUpdateAPIToken,
			}
		}
	}
	ignorePagination := new(bool)
	if !r.Configuration.IgnorePagination.IsUnknown() && !r.Configuration.IgnorePagination.IsNull() {
		*ignorePagination = r.Configuration.IgnorePagination.ValueBool()
	} else {
		ignorePagination = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskSupportUpdate{
		Credentials:      credentials,
		IgnorePagination: ignorePagination,
		StartDate:        startDate,
		Subdomain:        subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskSupportPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskSupportResourceModel) ToDeleteSDKType() *shared.SourceZendeskSupportCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskSupportResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskSupportResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
