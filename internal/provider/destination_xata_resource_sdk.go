// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationXataResourceModel) ToCreateSDKType() *shared.DestinationXataCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dbURL := r.Configuration.DbURL.ValueString()
	configuration := shared.DestinationXata{
		APIKey: apiKey,
		DbURL:  dbURL,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationXataCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationXataResourceModel) ToGetSDKType() *shared.DestinationXataCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationXataResourceModel) ToUpdateSDKType() *shared.DestinationXataPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	dbURL := r.Configuration.DbURL.ValueString()
	configuration := shared.DestinationXataUpdate{
		APIKey: apiKey,
		DbURL:  dbURL,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationXataPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationXataResourceModel) ToDeleteSDKType() *shared.DestinationXataCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationXataResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationXataResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
