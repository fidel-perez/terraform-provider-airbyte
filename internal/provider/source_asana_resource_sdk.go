// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAsanaResourceModel) ToCreateSDKType() *shared.SourceAsanaCreateRequest {
	var credentials *shared.SourceAsanaAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceAsanaAuthenticateViaAsanaOauth *shared.SourceAsanaAuthenticateViaAsanaOauth
		if r.Configuration.Credentials.AuthenticateViaAsanaOauth != nil {
			clientID := r.Configuration.Credentials.AuthenticateViaAsanaOauth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaAsanaOauth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaAsanaOauth.RefreshToken.ValueString()
			sourceAsanaAuthenticateViaAsanaOauth = &shared.SourceAsanaAuthenticateViaAsanaOauth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceAsanaAuthenticateViaAsanaOauth != nil {
			credentials = &shared.SourceAsanaAuthenticationMechanism{
				SourceAsanaAuthenticateViaAsanaOauth: sourceAsanaAuthenticateViaAsanaOauth,
			}
		}
		var sourceAsanaAuthenticateWithPersonalAccessToken *shared.SourceAsanaAuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.AuthenticateWithPersonalAccessToken != nil {
			personalAccessToken := r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.PersonalAccessToken.ValueString()
			sourceAsanaAuthenticateWithPersonalAccessToken = &shared.SourceAsanaAuthenticateWithPersonalAccessToken{
				PersonalAccessToken: personalAccessToken,
			}
		}
		if sourceAsanaAuthenticateWithPersonalAccessToken != nil {
			credentials = &shared.SourceAsanaAuthenticationMechanism{
				SourceAsanaAuthenticateWithPersonalAccessToken: sourceAsanaAuthenticateWithPersonalAccessToken,
			}
		}
	}
	var organizationExportIds []interface{} = nil
	for _, organizationExportIdsItem := range r.Configuration.OrganizationExportIds {
		var organizationExportIdsTmp interface{}
		_ = json.Unmarshal([]byte(organizationExportIdsItem.ValueString()), &organizationExportIdsTmp)
		organizationExportIds = append(organizationExportIds, organizationExportIdsTmp)
	}
	testMode := new(bool)
	if !r.Configuration.TestMode.IsUnknown() && !r.Configuration.TestMode.IsNull() {
		*testMode = r.Configuration.TestMode.ValueBool()
	} else {
		testMode = nil
	}
	configuration := shared.SourceAsana{
		Credentials:           credentials,
		OrganizationExportIds: organizationExportIds,
		TestMode:              testMode,
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
	out := shared.SourceAsanaCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAsanaResourceModel) ToGetSDKType() *shared.SourceAsanaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAsanaResourceModel) ToUpdateSDKType() *shared.SourceAsanaPutRequest {
	var credentials *shared.AuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var authenticateViaAsanaOauth *shared.AuthenticateViaAsanaOauth
		if r.Configuration.Credentials.AuthenticateViaAsanaOauth != nil {
			clientID := r.Configuration.Credentials.AuthenticateViaAsanaOauth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.AuthenticateViaAsanaOauth.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.AuthenticateViaAsanaOauth.RefreshToken.ValueString()
			authenticateViaAsanaOauth = &shared.AuthenticateViaAsanaOauth{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if authenticateViaAsanaOauth != nil {
			credentials = &shared.AuthenticationMechanism{
				AuthenticateViaAsanaOauth: authenticateViaAsanaOauth,
			}
		}
		var authenticateWithPersonalAccessToken *shared.AuthenticateWithPersonalAccessToken
		if r.Configuration.Credentials.AuthenticateWithPersonalAccessToken != nil {
			personalAccessToken := r.Configuration.Credentials.AuthenticateWithPersonalAccessToken.PersonalAccessToken.ValueString()
			authenticateWithPersonalAccessToken = &shared.AuthenticateWithPersonalAccessToken{
				PersonalAccessToken: personalAccessToken,
			}
		}
		if authenticateWithPersonalAccessToken != nil {
			credentials = &shared.AuthenticationMechanism{
				AuthenticateWithPersonalAccessToken: authenticateWithPersonalAccessToken,
			}
		}
	}
	var organizationExportIds []interface{} = nil
	for _, organizationExportIdsItem := range r.Configuration.OrganizationExportIds {
		var organizationExportIdsTmp interface{}
		_ = json.Unmarshal([]byte(organizationExportIdsItem.ValueString()), &organizationExportIdsTmp)
		organizationExportIds = append(organizationExportIds, organizationExportIdsTmp)
	}
	testMode := new(bool)
	if !r.Configuration.TestMode.IsUnknown() && !r.Configuration.TestMode.IsNull() {
		*testMode = r.Configuration.TestMode.ValueBool()
	} else {
		testMode = nil
	}
	configuration := shared.SourceAsanaUpdate{
		Credentials:           credentials,
		OrganizationExportIds: organizationExportIds,
		TestMode:              testMode,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAsanaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAsanaResourceModel) ToDeleteSDKType() *shared.SourceAsanaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAsanaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAsanaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
