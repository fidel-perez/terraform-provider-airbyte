// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMicrosoftTeamsResourceModel) ToCreateSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	var credentials *shared.SourceMicrosoftTeamsAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 *shared.SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20
		if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20 != nil {
			clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.RefreshToken.ValueString()
			tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.TenantID.ValueString()
			sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 = &shared.SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
				TenantID:     tenantID,
			}
		}
		if sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20 != nil {
			credentials = &shared.SourceMicrosoftTeamsAuthenticationMechanism{
				SourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20: sourceMicrosoftTeamsAuthenticateViaMicrosoftOAuth20,
			}
		}
		var sourceMicrosoftTeamsAuthenticateViaMicrosoft *shared.SourceMicrosoftTeamsAuthenticateViaMicrosoft
		if r.Configuration.Credentials.AuthenticateViaMicrosoft != nil {
			clientId1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.ClientSecret.ValueString()
			tenantId1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.TenantID.ValueString()
			sourceMicrosoftTeamsAuthenticateViaMicrosoft = &shared.SourceMicrosoftTeamsAuthenticateViaMicrosoft{
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
				TenantID:     tenantId1,
			}
		}
		if sourceMicrosoftTeamsAuthenticateViaMicrosoft != nil {
			credentials = &shared.SourceMicrosoftTeamsAuthenticationMechanism{
				SourceMicrosoftTeamsAuthenticateViaMicrosoft: sourceMicrosoftTeamsAuthenticateViaMicrosoft,
			}
		}
	}
	period := r.Configuration.Period.ValueString()
	configuration := shared.SourceMicrosoftTeams{
		Credentials: credentials,
		Period:      period,
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
	out := shared.SourceMicrosoftTeamsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftTeamsResourceModel) ToGetSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMicrosoftTeamsResourceModel) ToUpdateSDKType() *shared.SourceMicrosoftTeamsPutRequest {
	var credentials *shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var authenticateViaMicrosoftOAuth20 *shared.AuthenticateViaMicrosoftOAuth20
		if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20 != nil {
			clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.RefreshToken.ValueString()
			tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth20.TenantID.ValueString()
			authenticateViaMicrosoftOAuth20 = &shared.AuthenticateViaMicrosoftOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
				TenantID:     tenantID,
			}
		}
		if authenticateViaMicrosoftOAuth20 != nil {
			credentials = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism{
				AuthenticateViaMicrosoftOAuth20: authenticateViaMicrosoftOAuth20,
			}
		}
		var authenticateViaMicrosoft *shared.AuthenticateViaMicrosoft
		if r.Configuration.Credentials.AuthenticateViaMicrosoft != nil {
			clientId1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.ClientID.ValueString()
			clientSecret1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.ClientSecret.ValueString()
			tenantId1 := r.Configuration.Credentials.AuthenticateViaMicrosoft.TenantID.ValueString()
			authenticateViaMicrosoft = &shared.AuthenticateViaMicrosoft{
				ClientID:     clientId1,
				ClientSecret: clientSecret1,
				TenantID:     tenantId1,
			}
		}
		if authenticateViaMicrosoft != nil {
			credentials = &shared.SourceMicrosoftTeamsUpdateAuthenticationMechanism{
				AuthenticateViaMicrosoft: authenticateViaMicrosoft,
			}
		}
	}
	period := r.Configuration.Period.ValueString()
	configuration := shared.SourceMicrosoftTeamsUpdate{
		Credentials: credentials,
		Period:      period,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftTeamsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftTeamsResourceModel) ToDeleteSDKType() *shared.SourceMicrosoftTeamsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMicrosoftTeamsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMicrosoftTeamsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
