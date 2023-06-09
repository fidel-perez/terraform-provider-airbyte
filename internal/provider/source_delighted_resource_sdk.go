// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceDelightedResourceModel) ToCreateSDKType() *shared.SourceDelightedCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	since, _ := time.Parse(time.RFC3339Nano, r.Configuration.Since.ValueString())
	sourceType := shared.SourceDelightedDelighted(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceDelighted{
		APIKey:     apiKey,
		Since:      since,
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
	out := shared.SourceDelightedCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDelightedResourceModel) ToGetSDKType() *shared.SourceDelightedCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDelightedResourceModel) ToUpdateSDKType() *shared.SourceDelightedPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	since, _ := time.Parse(time.RFC3339Nano, r.Configuration.Since.ValueString())
	configuration := shared.SourceDelightedUpdate{
		APIKey: apiKey,
		Since:  since,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDelightedPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDelightedResourceModel) ToDeleteSDKType() *shared.SourceDelightedCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDelightedResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDelightedResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
