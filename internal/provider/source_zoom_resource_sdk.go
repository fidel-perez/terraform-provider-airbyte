// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceZoomResourceModel) ToCreateSDKType() *shared.SourceZoomCreateRequest {
	jwtToken := r.Configuration.JwtToken.ValueString()
	configuration := shared.SourceZoom{
		JwtToken: jwtToken,
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
	out := shared.SourceZoomCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZoomResourceModel) ToGetSDKType() *shared.SourceZoomCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZoomResourceModel) ToUpdateSDKType() *shared.SourceZoomPutRequest {
	jwtToken := r.Configuration.JwtToken.ValueString()
	configuration := shared.SourceZoomUpdate{
		JwtToken: jwtToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZoomPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZoomResourceModel) ToDeleteSDKType() *shared.SourceZoomCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZoomResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZoomResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
