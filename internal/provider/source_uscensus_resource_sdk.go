// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceUsCensusResourceModel) ToCreateSDKType() *shared.SourceUsCensusCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	queryParams := new(string)
	if !r.Configuration.QueryParams.IsUnknown() && !r.Configuration.QueryParams.IsNull() {
		*queryParams = r.Configuration.QueryParams.ValueString()
	} else {
		queryParams = nil
	}
	queryPath := r.Configuration.QueryPath.ValueString()
	configuration := shared.SourceUsCensus{
		APIKey:      apiKey,
		QueryParams: queryParams,
		QueryPath:   queryPath,
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
	out := shared.SourceUsCensusCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceUsCensusResourceModel) ToGetSDKType() *shared.SourceUsCensusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceUsCensusResourceModel) ToUpdateSDKType() *shared.SourceUsCensusPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	queryParams := new(string)
	if !r.Configuration.QueryParams.IsUnknown() && !r.Configuration.QueryParams.IsNull() {
		*queryParams = r.Configuration.QueryParams.ValueString()
	} else {
		queryParams = nil
	}
	queryPath := r.Configuration.QueryPath.ValueString()
	configuration := shared.SourceUsCensusUpdate{
		APIKey:      apiKey,
		QueryParams: queryParams,
		QueryPath:   queryPath,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceUsCensusPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceUsCensusResourceModel) ToDeleteSDKType() *shared.SourceUsCensusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceUsCensusResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceUsCensusResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
