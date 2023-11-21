// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGoogleDirectoryResourceModel) ToCreateSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	var credentials *shared.SourceGoogleDirectoryGoogleCredentials
	if r.Configuration.Credentials != nil {
		var sourceGoogleDirectorySignInViaGoogleOAuth *shared.SourceGoogleDirectorySignInViaGoogleOAuth
		if r.Configuration.Credentials.SignInViaGoogleOAuth != nil {
			clientID := r.Configuration.Credentials.SignInViaGoogleOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SignInViaGoogleOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SignInViaGoogleOAuth.RefreshToken.ValueString()
			sourceGoogleDirectorySignInViaGoogleOAuth = &shared.SourceGoogleDirectorySignInViaGoogleOAuth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceGoogleDirectorySignInViaGoogleOAuth != nil {
			credentials = &shared.SourceGoogleDirectoryGoogleCredentials{
				SourceGoogleDirectorySignInViaGoogleOAuth: sourceGoogleDirectorySignInViaGoogleOAuth,
			}
		}
		var sourceGoogleDirectoryServiceAccountKey *shared.SourceGoogleDirectoryServiceAccountKey
		if r.Configuration.Credentials.ServiceAccountKey != nil {
			credentialsJSON := r.Configuration.Credentials.ServiceAccountKey.CredentialsJSON.ValueString()
			email := r.Configuration.Credentials.ServiceAccountKey.Email.ValueString()
			sourceGoogleDirectoryServiceAccountKey = &shared.SourceGoogleDirectoryServiceAccountKey{
				CredentialsJSON: credentialsJSON,
				Email:           email,
			}
		}
		if sourceGoogleDirectoryServiceAccountKey != nil {
			credentials = &shared.SourceGoogleDirectoryGoogleCredentials{
				SourceGoogleDirectoryServiceAccountKey: sourceGoogleDirectoryServiceAccountKey,
			}
		}
	}
	configuration := shared.SourceGoogleDirectory{
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
	out := shared.SourceGoogleDirectoryCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleDirectoryResourceModel) ToGetSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleDirectoryResourceModel) ToUpdateSDKType() *shared.SourceGoogleDirectoryPutRequest {
	var credentials *shared.SourceGoogleDirectoryUpdateGoogleCredentials
	if r.Configuration.Credentials != nil {
		var signInViaGoogleOAuth *shared.SignInViaGoogleOAuth
		if r.Configuration.Credentials.SignInViaGoogleOAuth != nil {
			clientID := r.Configuration.Credentials.SignInViaGoogleOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SignInViaGoogleOAuth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SignInViaGoogleOAuth.RefreshToken.ValueString()
			signInViaGoogleOAuth = &shared.SignInViaGoogleOAuth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if signInViaGoogleOAuth != nil {
			credentials = &shared.SourceGoogleDirectoryUpdateGoogleCredentials{
				SignInViaGoogleOAuth: signInViaGoogleOAuth,
			}
		}
		var serviceAccountKey *shared.ServiceAccountKey
		if r.Configuration.Credentials.ServiceAccountKey != nil {
			credentialsJSON := r.Configuration.Credentials.ServiceAccountKey.CredentialsJSON.ValueString()
			email := r.Configuration.Credentials.ServiceAccountKey.Email.ValueString()
			serviceAccountKey = &shared.ServiceAccountKey{
				CredentialsJSON: credentialsJSON,
				Email:           email,
			}
		}
		if serviceAccountKey != nil {
			credentials = &shared.SourceGoogleDirectoryUpdateGoogleCredentials{
				ServiceAccountKey: serviceAccountKey,
			}
		}
	}
	configuration := shared.SourceGoogleDirectoryUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleDirectoryPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleDirectoryResourceModel) ToDeleteSDKType() *shared.SourceGoogleDirectoryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGoogleDirectoryResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGoogleDirectoryResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
