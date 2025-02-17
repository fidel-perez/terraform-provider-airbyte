// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSendgridResourceModel) ToCreateSDKType() *shared.SourceSendgridCreateRequest {
	apikey := r.Configuration.Apikey.ValueString()
	startTime := new(time.Time)
	if !r.Configuration.StartTime.IsUnknown() && !r.Configuration.StartTime.IsNull() {
		*startTime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartTime.ValueString())
	} else {
		startTime = nil
	}
	configuration := shared.SourceSendgrid{
		Apikey:    apikey,
		StartTime: startTime,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSendgridCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSendgridResourceModel) ToGetSDKType() *shared.SourceSendgridCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSendgridResourceModel) ToUpdateSDKType() *shared.SourceSendgridPutRequest {
	apikey := r.Configuration.Apikey.ValueString()
	startTime := new(time.Time)
	if !r.Configuration.StartTime.IsUnknown() && !r.Configuration.StartTime.IsNull() {
		*startTime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartTime.ValueString())
	} else {
		startTime = nil
	}
	configuration := shared.SourceSendgridUpdate{
		Apikey:    apikey,
		StartTime: startTime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSendgridPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSendgridResourceModel) ToDeleteSDKType() *shared.SourceSendgridCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSendgridResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSendgridResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
