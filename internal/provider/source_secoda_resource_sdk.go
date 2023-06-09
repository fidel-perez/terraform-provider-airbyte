// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSecodaResourceModel) ToCreateSDKType() *shared.SourceSecodaCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceSecodaSecoda(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSecoda{
		APIKey:     apiKey,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSecodaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSecodaResourceModel) ToGetSDKType() *shared.SourceSecodaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSecodaResourceModel) ToUpdateSDKType() *shared.SourceSecodaPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceSecodaUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSecodaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSecodaResourceModel) ToDeleteSDKType() *shared.SourceSecodaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSecodaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSecodaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
