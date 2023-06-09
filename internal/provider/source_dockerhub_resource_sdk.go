// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDockerhubResourceModel) ToCreateSDKType() *shared.SourceDockerhubCreateRequest {
	dockerUsername := r.Configuration.DockerUsername.ValueString()
	sourceType := shared.SourceDockerhubDockerhub(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceDockerhub{
		DockerUsername: dockerUsername,
		SourceType:     sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDockerhubCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDockerhubResourceModel) ToGetSDKType() *shared.SourceDockerhubCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDockerhubResourceModel) ToUpdateSDKType() *shared.SourceDockerhubPutRequest {
	dockerUsername := r.Configuration.DockerUsername.ValueString()
	configuration := shared.SourceDockerhubUpdate{
		DockerUsername: dockerUsername,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDockerhubPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDockerhubResourceModel) ToDeleteSDKType() *shared.SourceDockerhubCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDockerhubResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDockerhubResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
