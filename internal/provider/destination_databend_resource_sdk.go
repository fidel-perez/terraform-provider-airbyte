// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDatabendResourceModel) ToCreateSDKType() *shared.DestinationDatabendCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	table := new(string)
	if !r.Configuration.Table.IsUnknown() && !r.Configuration.Table.IsNull() {
		*table = r.Configuration.Table.ValueString()
	} else {
		table = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationDatabend{
		Database: database,
		Host:     host,
		Password: password,
		Port:     port,
		Table:    table,
		Username: username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabendCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabendResourceModel) ToGetSDKType() *shared.DestinationDatabendCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabendResourceModel) ToUpdateSDKType() *shared.DestinationDatabendPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	table := new(string)
	if !r.Configuration.Table.IsUnknown() && !r.Configuration.Table.IsNull() {
		*table = r.Configuration.Table.ValueString()
	} else {
		table = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationDatabendUpdate{
		Database: database,
		Host:     host,
		Password: password,
		Port:     port,
		Table:    table,
		Username: username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabendPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabendResourceModel) ToDeleteSDKType() *shared.DestinationDatabendCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabendResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDatabendResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
