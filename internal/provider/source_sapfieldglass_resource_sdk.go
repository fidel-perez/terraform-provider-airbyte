// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSapFieldglassResourceModel) ToCreateSDKType() *shared.SourceSapFieldglassCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceSapFieldglassSapFieldglass(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSapFieldglass{
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
	out := shared.SourceSapFieldglassCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSapFieldglassResourceModel) ToGetSDKType() *shared.SourceSapFieldglassCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSapFieldglassResourceModel) ToUpdateSDKType() *shared.SourceSapFieldglassPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceSapFieldglassUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSapFieldglassPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSapFieldglassResourceModel) ToDeleteSDKType() *shared.SourceSapFieldglassCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSapFieldglassResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSapFieldglassResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
