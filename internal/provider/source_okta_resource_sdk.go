// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOktaResourceModel) ToCreateSDKType() *shared.SourceOktaCreateRequest {
	var credentials *shared.SourceOktaAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceOktaAuthorizationMethodOAuth20 *shared.SourceOktaAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20 != nil {
			clientID := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceOktaAuthorizationMethodOAuth20.RefreshToken.ValueString()
			sourceOktaAuthorizationMethodOAuth20 = &shared.SourceOktaAuthorizationMethodOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceOktaAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceOktaAuthorizationMethod{
				SourceOktaAuthorizationMethodOAuth20: sourceOktaAuthorizationMethodOAuth20,
			}
		}
		var sourceOktaAuthorizationMethodAPIToken *shared.SourceOktaAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceOktaAuthorizationMethodAPIToken.APIToken.ValueString()
			sourceOktaAuthorizationMethodAPIToken = &shared.SourceOktaAuthorizationMethodAPIToken{
				APIToken: apiToken,
			}
		}
		if sourceOktaAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceOktaAuthorizationMethod{
				SourceOktaAuthorizationMethodAPIToken: sourceOktaAuthorizationMethodAPIToken,
			}
		}
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	configuration := shared.SourceOkta{
		Credentials: credentials,
		Domain:      domain,
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
	out := shared.SourceOktaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOktaResourceModel) ToGetSDKType() *shared.SourceOktaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOktaResourceModel) ToUpdateSDKType() *shared.SourceOktaPutRequest {
	var credentials *shared.SourceOktaUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceOktaUpdateAuthorizationMethodOAuth20 *shared.SourceOktaUpdateAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20 != nil {
			clientID := r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
			sourceOktaUpdateAuthorizationMethodOAuth20 = &shared.SourceOktaUpdateAuthorizationMethodOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceOktaUpdateAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceOktaUpdateAuthorizationMethod{
				SourceOktaUpdateAuthorizationMethodOAuth20: sourceOktaUpdateAuthorizationMethodOAuth20,
			}
		}
		var sourceOktaUpdateAuthorizationMethodAPIToken *shared.SourceOktaUpdateAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceOktaUpdateAuthorizationMethodAPIToken.APIToken.ValueString()
			sourceOktaUpdateAuthorizationMethodAPIToken = &shared.SourceOktaUpdateAuthorizationMethodAPIToken{
				APIToken: apiToken,
			}
		}
		if sourceOktaUpdateAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceOktaUpdateAuthorizationMethod{
				SourceOktaUpdateAuthorizationMethodAPIToken: sourceOktaUpdateAuthorizationMethodAPIToken,
			}
		}
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	configuration := shared.SourceOktaUpdate{
		Credentials: credentials,
		Domain:      domain,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOktaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOktaResourceModel) ToDeleteSDKType() *shared.SourceOktaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOktaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOktaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
