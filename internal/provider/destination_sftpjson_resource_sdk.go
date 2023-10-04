// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationSftpJSONResourceModel) ToCreateSDKType() *shared.DestinationSftpJSONCreateRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	host := r.Configuration.Host.ValueString()
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationSftpJSON{
		DestinationPath: destinationPath,
		Host:            host,
		Password:        password,
		Port:            port,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSftpJSONCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSftpJSONResourceModel) ToGetSDKType() *shared.DestinationSftpJSONCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSftpJSONResourceModel) ToUpdateSDKType() *shared.DestinationSftpJSONPutRequest {
	destinationPath := r.Configuration.DestinationPath.ValueString()
	host := r.Configuration.Host.ValueString()
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationSftpJSONUpdate{
		DestinationPath: destinationPath,
		Host:            host,
		Password:        password,
		Port:            port,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSftpJSONPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSftpJSONResourceModel) ToDeleteSDKType() *shared.DestinationSftpJSONCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSftpJSONResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationSftpJSONResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
