// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceLinkedinPagesResourceModel) ToCreateSDKType() *shared.SourceLinkedinPagesCreateRequest {
	var credentials *shared.SourceLinkedinPagesAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinPagesAuthenticationOAuth20 *shared.SourceLinkedinPagesAuthenticationOAuth20
		if r.Configuration.Credentials.SourceLinkedinPagesAuthenticationOAuth20 != nil {
			clientID := r.Configuration.Credentials.SourceLinkedinPagesAuthenticationOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceLinkedinPagesAuthenticationOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceLinkedinPagesAuthenticationOAuth20.RefreshToken.ValueString()
			sourceLinkedinPagesAuthenticationOAuth20 = &shared.SourceLinkedinPagesAuthenticationOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinPagesAuthenticationOAuth20 != nil {
			credentials = &shared.SourceLinkedinPagesAuthentication{
				SourceLinkedinPagesAuthenticationOAuth20: sourceLinkedinPagesAuthenticationOAuth20,
			}
		}
		var sourceLinkedinPagesAuthenticationAccessToken *shared.SourceLinkedinPagesAuthenticationAccessToken
		if r.Configuration.Credentials.SourceLinkedinPagesAuthenticationAccessToken != nil {
			accessToken := r.Configuration.Credentials.SourceLinkedinPagesAuthenticationAccessToken.AccessToken.ValueString()
			sourceLinkedinPagesAuthenticationAccessToken = &shared.SourceLinkedinPagesAuthenticationAccessToken{
				AccessToken: accessToken,
			}
		}
		if sourceLinkedinPagesAuthenticationAccessToken != nil {
			credentials = &shared.SourceLinkedinPagesAuthentication{
				SourceLinkedinPagesAuthenticationAccessToken: sourceLinkedinPagesAuthenticationAccessToken,
			}
		}
	}
	orgID := r.Configuration.OrgID.ValueString()
	configuration := shared.SourceLinkedinPages{
		Credentials: credentials,
		OrgID:       orgID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinPagesCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinPagesResourceModel) ToGetSDKType() *shared.SourceLinkedinPagesCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinPagesResourceModel) ToUpdateSDKType() *shared.SourceLinkedinPagesPutRequest {
	var credentials *shared.SourceLinkedinPagesUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceLinkedinPagesUpdateAuthenticationOAuth20 *shared.SourceLinkedinPagesUpdateAuthenticationOAuth20
		if r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationOAuth20 != nil {
			clientID := r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationOAuth20.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationOAuth20.RefreshToken.ValueString()
			sourceLinkedinPagesUpdateAuthenticationOAuth20 = &shared.SourceLinkedinPagesUpdateAuthenticationOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceLinkedinPagesUpdateAuthenticationOAuth20 != nil {
			credentials = &shared.SourceLinkedinPagesUpdateAuthentication{
				SourceLinkedinPagesUpdateAuthenticationOAuth20: sourceLinkedinPagesUpdateAuthenticationOAuth20,
			}
		}
		var sourceLinkedinPagesUpdateAuthenticationAccessToken *shared.SourceLinkedinPagesUpdateAuthenticationAccessToken
		if r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationAccessToken != nil {
			accessToken := r.Configuration.Credentials.SourceLinkedinPagesUpdateAuthenticationAccessToken.AccessToken.ValueString()
			sourceLinkedinPagesUpdateAuthenticationAccessToken = &shared.SourceLinkedinPagesUpdateAuthenticationAccessToken{
				AccessToken: accessToken,
			}
		}
		if sourceLinkedinPagesUpdateAuthenticationAccessToken != nil {
			credentials = &shared.SourceLinkedinPagesUpdateAuthentication{
				SourceLinkedinPagesUpdateAuthenticationAccessToken: sourceLinkedinPagesUpdateAuthenticationAccessToken,
			}
		}
	}
	orgID := r.Configuration.OrgID.ValueString()
	configuration := shared.SourceLinkedinPagesUpdate{
		Credentials: credentials,
		OrgID:       orgID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinkedinPagesPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinkedinPagesResourceModel) ToDeleteSDKType() *shared.SourceLinkedinPagesCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceLinkedinPagesResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceLinkedinPagesResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
