// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceInstatusResourceModel) ToCreateSDKType() *shared.SourceInstatusCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceInstatus{
		APIKey: apiKey,
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
	out := shared.SourceInstatusCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceInstatusResourceModel) ToGetSDKType() *shared.SourceInstatusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceInstatusResourceModel) ToUpdateSDKType() *shared.SourceInstatusPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceInstatusUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceInstatusPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceInstatusResourceModel) ToDeleteSDKType() *shared.SourceInstatusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceInstatusResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceInstatusResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
