// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourcePipedriveResourceModel) ToCreateSDKType() *shared.SourcePipedriveCreateRequest {
	var authorization *shared.SourcePipedriveAPIKeyAuthentication
	if r.Configuration.Authorization != nil {
		apiToken := r.Configuration.Authorization.APIToken.ValueString()
		authorization = &shared.SourcePipedriveAPIKeyAuthentication{
			APIToken: apiToken,
		}
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourcePipedrive{
		Authorization:        authorization,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePipedriveCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePipedriveResourceModel) ToGetSDKType() *shared.SourcePipedriveCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePipedriveResourceModel) ToUpdateSDKType() *shared.SourcePipedrivePutRequest {
	var authorization *shared.SourcePipedriveUpdateAPIKeyAuthentication
	if r.Configuration.Authorization != nil {
		apiToken := r.Configuration.Authorization.APIToken.ValueString()
		authorization = &shared.SourcePipedriveUpdateAPIKeyAuthentication{
			APIToken: apiToken,
		}
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	configuration := shared.SourcePipedriveUpdate{
		Authorization:        authorization,
		ReplicationStartDate: replicationStartDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePipedrivePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePipedriveResourceModel) ToDeleteSDKType() *shared.SourcePipedriveCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePipedriveResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePipedriveResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
