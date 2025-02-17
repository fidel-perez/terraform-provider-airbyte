// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationElasticsearchResourceModel) ToCreateSDKType() *shared.DestinationElasticsearchCreateRequest {
	var authenticationMethod *shared.DestinationElasticsearchAuthenticationMethod
	if r.Configuration.AuthenticationMethod != nil {
		var destinationElasticsearchAPIKeySecret *shared.DestinationElasticsearchAPIKeySecret
		if r.Configuration.AuthenticationMethod.APIKeySecret != nil {
			apiKeyID := r.Configuration.AuthenticationMethod.APIKeySecret.APIKeyID.ValueString()
			apiKeySecret := r.Configuration.AuthenticationMethod.APIKeySecret.APIKeySecret.ValueString()
			destinationElasticsearchAPIKeySecret = &shared.DestinationElasticsearchAPIKeySecret{
				APIKeyID:     apiKeyID,
				APIKeySecret: apiKeySecret,
			}
		}
		if destinationElasticsearchAPIKeySecret != nil {
			authenticationMethod = &shared.DestinationElasticsearchAuthenticationMethod{
				DestinationElasticsearchAPIKeySecret: destinationElasticsearchAPIKeySecret,
			}
		}
		var destinationElasticsearchUsernamePassword *shared.DestinationElasticsearchUsernamePassword
		if r.Configuration.AuthenticationMethod.UsernamePassword != nil {
			password := r.Configuration.AuthenticationMethod.UsernamePassword.Password.ValueString()
			username := r.Configuration.AuthenticationMethod.UsernamePassword.Username.ValueString()
			destinationElasticsearchUsernamePassword = &shared.DestinationElasticsearchUsernamePassword{
				Password: password,
				Username: username,
			}
		}
		if destinationElasticsearchUsernamePassword != nil {
			authenticationMethod = &shared.DestinationElasticsearchAuthenticationMethod{
				DestinationElasticsearchUsernamePassword: destinationElasticsearchUsernamePassword,
			}
		}
	}
	caCertificate := new(string)
	if !r.Configuration.CaCertificate.IsUnknown() && !r.Configuration.CaCertificate.IsNull() {
		*caCertificate = r.Configuration.CaCertificate.ValueString()
	} else {
		caCertificate = nil
	}
	endpoint := r.Configuration.Endpoint.ValueString()
	upsert := new(bool)
	if !r.Configuration.Upsert.IsUnknown() && !r.Configuration.Upsert.IsNull() {
		*upsert = r.Configuration.Upsert.ValueBool()
	} else {
		upsert = nil
	}
	configuration := shared.DestinationElasticsearch{
		AuthenticationMethod: authenticationMethod,
		CaCertificate:        caCertificate,
		Endpoint:             endpoint,
		Upsert:               upsert,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationElasticsearchCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationElasticsearchResourceModel) ToGetSDKType() *shared.DestinationElasticsearchCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationElasticsearchResourceModel) ToUpdateSDKType() *shared.DestinationElasticsearchPutRequest {
	var authenticationMethod *shared.AuthenticationMethod
	if r.Configuration.AuthenticationMethod != nil {
		var apiKeySecret *shared.APIKeySecret
		if r.Configuration.AuthenticationMethod.APIKeySecret != nil {
			apiKeyID := r.Configuration.AuthenticationMethod.APIKeySecret.APIKeyID.ValueString()
			apiKeySecret1 := r.Configuration.AuthenticationMethod.APIKeySecret.APIKeySecret.ValueString()
			apiKeySecret = &shared.APIKeySecret{
				APIKeyID:     apiKeyID,
				APIKeySecret: apiKeySecret1,
			}
		}
		if apiKeySecret != nil {
			authenticationMethod = &shared.AuthenticationMethod{
				APIKeySecret: apiKeySecret,
			}
		}
		var usernamePassword *shared.UsernamePassword
		if r.Configuration.AuthenticationMethod.UsernamePassword != nil {
			password := r.Configuration.AuthenticationMethod.UsernamePassword.Password.ValueString()
			username := r.Configuration.AuthenticationMethod.UsernamePassword.Username.ValueString()
			usernamePassword = &shared.UsernamePassword{
				Password: password,
				Username: username,
			}
		}
		if usernamePassword != nil {
			authenticationMethod = &shared.AuthenticationMethod{
				UsernamePassword: usernamePassword,
			}
		}
	}
	caCertificate := new(string)
	if !r.Configuration.CaCertificate.IsUnknown() && !r.Configuration.CaCertificate.IsNull() {
		*caCertificate = r.Configuration.CaCertificate.ValueString()
	} else {
		caCertificate = nil
	}
	endpoint := r.Configuration.Endpoint.ValueString()
	upsert := new(bool)
	if !r.Configuration.Upsert.IsUnknown() && !r.Configuration.Upsert.IsNull() {
		*upsert = r.Configuration.Upsert.ValueBool()
	} else {
		upsert = nil
	}
	configuration := shared.DestinationElasticsearchUpdate{
		AuthenticationMethod: authenticationMethod,
		CaCertificate:        caCertificate,
		Endpoint:             endpoint,
		Upsert:               upsert,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationElasticsearchPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationElasticsearchResourceModel) ToDeleteSDKType() *shared.DestinationElasticsearchCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationElasticsearchResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationElasticsearchResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
