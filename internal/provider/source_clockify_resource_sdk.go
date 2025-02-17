// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceClockifyResourceModel) ToCreateSDKType() *shared.SourceClockifyCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	workspaceID := r.Configuration.WorkspaceID.ValueString()
	configuration := shared.SourceClockify{
		APIKey:      apiKey,
		APIURL:      apiURL,
		WorkspaceID: workspaceID,
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
	workspaceId1 := r.WorkspaceID.ValueString()
	out := shared.SourceClockifyCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceId1,
	}
	return &out
}

func (r *SourceClockifyResourceModel) ToGetSDKType() *shared.SourceClockifyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceClockifyResourceModel) ToUpdateSDKType() *shared.SourceClockifyPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	apiURL := new(string)
	if !r.Configuration.APIURL.IsUnknown() && !r.Configuration.APIURL.IsNull() {
		*apiURL = r.Configuration.APIURL.ValueString()
	} else {
		apiURL = nil
	}
	workspaceID := r.Configuration.WorkspaceID.ValueString()
	configuration := shared.SourceClockifyUpdate{
		APIKey:      apiKey,
		APIURL:      apiURL,
		WorkspaceID: workspaceID,
	}
	name := r.Name.ValueString()
	workspaceId1 := r.WorkspaceID.ValueString()
	out := shared.SourceClockifyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceId1,
	}
	return &out
}

func (r *SourceClockifyResourceModel) ToDeleteSDKType() *shared.SourceClockifyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceClockifyResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceClockifyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
