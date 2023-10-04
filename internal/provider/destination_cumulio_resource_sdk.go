// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationCumulioResourceModel) ToCreateSDKType() *shared.DestinationCumulioCreateRequest {
	apiHost := new(string)
	if !r.Configuration.APIHost.IsUnknown() && !r.Configuration.APIHost.IsNull() {
		*apiHost = r.Configuration.APIHost.ValueString()
	} else {
		apiHost = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	apiToken := r.Configuration.APIToken.ValueString()
	configuration := shared.DestinationCumulio{
		APIHost:  apiHost,
		APIKey:   apiKey,
		APIToken: apiToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationCumulioCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationCumulioResourceModel) ToGetSDKType() *shared.DestinationCumulioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationCumulioResourceModel) ToUpdateSDKType() *shared.DestinationCumulioPutRequest {
	apiHost := new(string)
	if !r.Configuration.APIHost.IsUnknown() && !r.Configuration.APIHost.IsNull() {
		*apiHost = r.Configuration.APIHost.ValueString()
	} else {
		apiHost = nil
	}
	apiKey := r.Configuration.APIKey.ValueString()
	apiToken := r.Configuration.APIToken.ValueString()
	configuration := shared.DestinationCumulioUpdate{
		APIHost:  apiHost,
		APIKey:   apiKey,
		APIToken: apiToken,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationCumulioPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationCumulioResourceModel) ToDeleteSDKType() *shared.DestinationCumulioCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationCumulioResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationCumulioResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
