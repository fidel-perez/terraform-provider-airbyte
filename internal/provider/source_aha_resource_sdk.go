// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAhaResourceModel) ToCreateSDKType() *shared.SourceAhaCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceAha{
		APIKey: apiKey,
		URL:    url,
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
	out := shared.SourceAhaCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAhaResourceModel) ToGetSDKType() *shared.SourceAhaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAhaResourceModel) ToUpdateSDKType() *shared.SourceAhaPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourceAhaUpdate{
		APIKey: apiKey,
		URL:    url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAhaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAhaResourceModel) ToDeleteSDKType() *shared.SourceAhaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAhaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAhaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
