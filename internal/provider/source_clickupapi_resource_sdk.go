// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceClickupAPIResourceModel) ToCreateSDKType() *shared.SourceClickupAPICreateRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	folderID := new(string)
	if !r.Configuration.FolderID.IsUnknown() && !r.Configuration.FolderID.IsNull() {
		*folderID = r.Configuration.FolderID.ValueString()
	} else {
		folderID = nil
	}
	includeClosedTasks := new(bool)
	if !r.Configuration.IncludeClosedTasks.IsUnknown() && !r.Configuration.IncludeClosedTasks.IsNull() {
		*includeClosedTasks = r.Configuration.IncludeClosedTasks.ValueBool()
	} else {
		includeClosedTasks = nil
	}
	listID := new(string)
	if !r.Configuration.ListID.IsUnknown() && !r.Configuration.ListID.IsNull() {
		*listID = r.Configuration.ListID.ValueString()
	} else {
		listID = nil
	}
	sourceType := shared.SourceClickupAPIClickupAPI(r.Configuration.SourceType.ValueString())
	spaceID := new(string)
	if !r.Configuration.SpaceID.IsUnknown() && !r.Configuration.SpaceID.IsNull() {
		*spaceID = r.Configuration.SpaceID.ValueString()
	} else {
		spaceID = nil
	}
	teamID := new(string)
	if !r.Configuration.TeamID.IsUnknown() && !r.Configuration.TeamID.IsNull() {
		*teamID = r.Configuration.TeamID.ValueString()
	} else {
		teamID = nil
	}
	configuration := shared.SourceClickupAPI{
		APIToken:           apiToken,
		FolderID:           folderID,
		IncludeClosedTasks: includeClosedTasks,
		ListID:             listID,
		SourceType:         sourceType,
		SpaceID:            spaceID,
		TeamID:             teamID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceClickupAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceClickupAPIResourceModel) ToGetSDKType() *shared.SourceClickupAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceClickupAPIResourceModel) ToUpdateSDKType() *shared.SourceClickupAPIPutRequest {
	apiToken := r.Configuration.APIToken.ValueString()
	folderID := new(string)
	if !r.Configuration.FolderID.IsUnknown() && !r.Configuration.FolderID.IsNull() {
		*folderID = r.Configuration.FolderID.ValueString()
	} else {
		folderID = nil
	}
	includeClosedTasks := new(bool)
	if !r.Configuration.IncludeClosedTasks.IsUnknown() && !r.Configuration.IncludeClosedTasks.IsNull() {
		*includeClosedTasks = r.Configuration.IncludeClosedTasks.ValueBool()
	} else {
		includeClosedTasks = nil
	}
	listID := new(string)
	if !r.Configuration.ListID.IsUnknown() && !r.Configuration.ListID.IsNull() {
		*listID = r.Configuration.ListID.ValueString()
	} else {
		listID = nil
	}
	spaceID := new(string)
	if !r.Configuration.SpaceID.IsUnknown() && !r.Configuration.SpaceID.IsNull() {
		*spaceID = r.Configuration.SpaceID.ValueString()
	} else {
		spaceID = nil
	}
	teamID := new(string)
	if !r.Configuration.TeamID.IsUnknown() && !r.Configuration.TeamID.IsNull() {
		*teamID = r.Configuration.TeamID.ValueString()
	} else {
		teamID = nil
	}
	configuration := shared.SourceClickupAPIUpdate{
		APIToken:           apiToken,
		FolderID:           folderID,
		IncludeClosedTasks: includeClosedTasks,
		ListID:             listID,
		SpaceID:            spaceID,
		TeamID:             teamID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceClickupAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceClickupAPIResourceModel) ToDeleteSDKType() *shared.SourceClickupAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceClickupAPIResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceClickupAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
