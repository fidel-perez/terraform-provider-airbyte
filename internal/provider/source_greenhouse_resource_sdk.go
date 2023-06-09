// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceGreenhouseResourceModel) ToCreateSDKType() *shared.SourceGreenhouseCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceGreenhouseGreenhouse(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceGreenhouse{
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
	out := shared.SourceGreenhouseCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGreenhouseResourceModel) ToGetSDKType() *shared.SourceGreenhouseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGreenhouseResourceModel) ToUpdateSDKType() *shared.SourceGreenhousePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceGreenhouseUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGreenhousePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGreenhouseResourceModel) ToDeleteSDKType() *shared.SourceGreenhouseCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGreenhouseResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGreenhouseResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
