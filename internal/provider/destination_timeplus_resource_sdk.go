// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationTimeplusResourceModel) ToCreateSDKType() *shared.DestinationTimeplusCreateRequest {
	apikey := r.Configuration.Apikey.ValueString()
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	configuration := shared.DestinationTimeplus{
		Apikey:   apikey,
		Endpoint: endpoint,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationTimeplusCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationTimeplusResourceModel) ToGetSDKType() *shared.DestinationTimeplusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationTimeplusResourceModel) ToUpdateSDKType() *shared.DestinationTimeplusPutRequest {
	apikey := r.Configuration.Apikey.ValueString()
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	configuration := shared.DestinationTimeplusUpdate{
		Apikey:   apikey,
		Endpoint: endpoint,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationTimeplusPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationTimeplusResourceModel) ToDeleteSDKType() *shared.DestinationTimeplusCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationTimeplusResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationTimeplusResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
