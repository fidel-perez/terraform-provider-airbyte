// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceBigqueryResourceModel) ToCreateSDKType() *shared.SourceBigqueryCreateRequest {
	credentialsJSON := r.Configuration.CredentialsJSON.ValueString()
	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.SourceBigquery{
		CredentialsJSON: credentialsJSON,
		DatasetID:       datasetID,
		ProjectID:       projectID,
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
	out := shared.SourceBigqueryCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigqueryResourceModel) ToGetSDKType() *shared.SourceBigqueryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBigqueryResourceModel) ToUpdateSDKType() *shared.SourceBigqueryPutRequest {
	credentialsJSON := r.Configuration.CredentialsJSON.ValueString()
	datasetID := new(string)
	if !r.Configuration.DatasetID.IsUnknown() && !r.Configuration.DatasetID.IsNull() {
		*datasetID = r.Configuration.DatasetID.ValueString()
	} else {
		datasetID = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.SourceBigqueryUpdate{
		CredentialsJSON: credentialsJSON,
		DatasetID:       datasetID,
		ProjectID:       projectID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBigqueryPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigqueryResourceModel) ToDeleteSDKType() *shared.SourceBigqueryCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceBigqueryResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceBigqueryResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
