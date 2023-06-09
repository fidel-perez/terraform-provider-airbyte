// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceEmailoctopusResourceModel) ToCreateSDKType() *shared.SourceEmailoctopusCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceEmailoctopusEmailoctopus(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceEmailoctopus{
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
	out := shared.SourceEmailoctopusCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceEmailoctopusResourceModel) ToGetSDKType() *shared.SourceEmailoctopusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceEmailoctopusResourceModel) ToUpdateSDKType() *shared.SourceEmailoctopusPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceEmailoctopusUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceEmailoctopusPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceEmailoctopusResourceModel) ToDeleteSDKType() *shared.SourceEmailoctopusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceEmailoctopusResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceEmailoctopusResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
