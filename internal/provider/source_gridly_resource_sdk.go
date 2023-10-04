// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGridlyResourceModel) ToCreateSDKType() *shared.SourceGridlyCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	gridID := r.Configuration.GridID.ValueString()
	configuration := shared.SourceGridly{
		APIKey: apiKey,
		GridID: gridID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGridlyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGridlyResourceModel) ToGetSDKType() *shared.SourceGridlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGridlyResourceModel) ToUpdateSDKType() *shared.SourceGridlyPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	gridID := r.Configuration.GridID.ValueString()
	configuration := shared.SourceGridlyUpdate{
		APIKey: apiKey,
		GridID: gridID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGridlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGridlyResourceModel) ToDeleteSDKType() *shared.SourceGridlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGridlyResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGridlyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
