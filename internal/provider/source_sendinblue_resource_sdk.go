// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSendinblueResourceModel) ToCreateSDKType() *shared.SourceSendinblueCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceSendinblue{
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
	out := shared.SourceSendinblueCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSendinblueResourceModel) ToGetSDKType() *shared.SourceSendinblueCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSendinblueResourceModel) ToUpdateSDKType() *shared.SourceSendinbluePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceSendinblueUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSendinbluePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSendinblueResourceModel) ToDeleteSDKType() *shared.SourceSendinblueCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSendinblueResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSendinblueResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
