// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMailchimpResourceModel) ToCreateSDKType() *shared.SourceMailchimpCreateRequest {
	campaignID := new(string)
	if !r.Configuration.CampaignID.IsUnknown() && !r.Configuration.CampaignID.IsNull() {
		*campaignID = r.Configuration.CampaignID.ValueString()
	} else {
		campaignID = nil
	}
	var credentials *shared.SourceMailchimpAuthentication
	if r.Configuration.Credentials != nil {
		var sourceMailchimpOAuth20 *shared.SourceMailchimpOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
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
			sourceMailchimpOAuth20 = &shared.SourceMailchimpOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceMailchimpOAuth20 != nil {
			credentials = &shared.SourceMailchimpAuthentication{
				SourceMailchimpOAuth20: sourceMailchimpOAuth20,
			}
		}
		var sourceMailchimpAPIKey *shared.SourceMailchimpAPIKey
		if r.Configuration.Credentials.APIKey != nil {
			apikey := r.Configuration.Credentials.APIKey.Apikey.ValueString()
			sourceMailchimpAPIKey = &shared.SourceMailchimpAPIKey{
				Apikey: apikey,
			}
		}
		if sourceMailchimpAPIKey != nil {
			credentials = &shared.SourceMailchimpAuthentication{
				SourceMailchimpAPIKey: sourceMailchimpAPIKey,
			}
		}
	}
	configuration := shared.SourceMailchimp{
		CampaignID:  campaignID,
		Credentials: credentials,
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
	out := shared.SourceMailchimpCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailchimpResourceModel) ToGetSDKType() *shared.SourceMailchimpCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailchimpResourceModel) ToUpdateSDKType() *shared.SourceMailchimpPutRequest {
	campaignID := new(string)
	if !r.Configuration.CampaignID.IsUnknown() && !r.Configuration.CampaignID.IsNull() {
		*campaignID = r.Configuration.CampaignID.ValueString()
	} else {
		campaignID = nil
	}
	var credentials *shared.SourceMailchimpUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceMailchimpUpdateOAuth20 *shared.SourceMailchimpUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
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
			sourceMailchimpUpdateOAuth20 = &shared.SourceMailchimpUpdateOAuth20{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			}
		}
		if sourceMailchimpUpdateOAuth20 != nil {
			credentials = &shared.SourceMailchimpUpdateAuthentication{
				SourceMailchimpUpdateOAuth20: sourceMailchimpUpdateOAuth20,
			}
		}
		var apiKey *shared.APIKey
		if r.Configuration.Credentials.APIKey != nil {
			apikey := r.Configuration.Credentials.APIKey.Apikey.ValueString()
			apiKey = &shared.APIKey{
				Apikey: apikey,
			}
		}
		if apiKey != nil {
			credentials = &shared.SourceMailchimpUpdateAuthentication{
				APIKey: apiKey,
			}
		}
	}
	configuration := shared.SourceMailchimpUpdate{
		CampaignID:  campaignID,
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailchimpPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailchimpResourceModel) ToDeleteSDKType() *shared.SourceMailchimpCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailchimpResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMailchimpResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
