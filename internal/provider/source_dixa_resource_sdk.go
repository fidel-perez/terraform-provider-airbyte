// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDixaResourceModel) ToCreateSDKType() *shared.SourceDixaCreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceDixa{
		APIToken:  apiToken,
		BatchSize: batchSize,
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
	out := shared.SourceDixaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDixaResourceModel) ToGetSDKType() *shared.SourceDixaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDixaResourceModel) ToUpdateSDKType() *shared.SourceDixaPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	batchSize := new(int64)
	if !r.Configuration.BatchSize.IsUnknown() && !r.Configuration.BatchSize.IsNull() {
		*batchSize = r.Configuration.BatchSize.ValueInt64()
	} else {
		batchSize = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceDixaUpdate{
		APIToken:  apiToken,
		BatchSize: batchSize,
		StartDate: startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDixaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDixaResourceModel) ToDeleteSDKType() *shared.SourceDixaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDixaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDixaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
