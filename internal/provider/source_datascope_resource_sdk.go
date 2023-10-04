// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDatascopeResourceModel) ToCreateSDKType() *shared.SourceDatascopeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceDatascope{
		APIKey:    apiKey,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDatascopeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDatascopeResourceModel) ToGetSDKType() *shared.SourceDatascopeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDatascopeResourceModel) ToUpdateSDKType() *shared.SourceDatascopePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceDatascopeUpdate{
		APIKey:    apiKey,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDatascopePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDatascopeResourceModel) ToDeleteSDKType() *shared.SourceDatascopeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDatascopeResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDatascopeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
