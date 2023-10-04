// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceTempoResourceModel) ToCreateSDKType() *shared.SourceTempoCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	configuration := shared.SourceTempo{
		APIToken: apiToken,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTempoCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTempoResourceModel) ToGetSDKType() *shared.SourceTempoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTempoResourceModel) ToUpdateSDKType() *shared.SourceTempoPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	configuration := shared.SourceTempoUpdate{
		APIToken: apiToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTempoPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTempoResourceModel) ToDeleteSDKType() *shared.SourceTempoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTempoResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTempoResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
