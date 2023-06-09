// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePunkAPIResourceModel) ToCreateSDKType() *shared.SourcePunkAPICreateRequest {
	brewedAfter := r.Configuration.BrewedAfter.ValueString()
	brewedBefore := r.Configuration.BrewedBefore.ValueString()
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	sourceType := shared.SourcePunkAPIPunkAPI(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePunkAPI{
		BrewedAfter:  brewedAfter,
		BrewedBefore: brewedBefore,
		ID:           id,
		SourceType:   sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePunkAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePunkAPIResourceModel) ToGetSDKType() *shared.SourcePunkAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePunkAPIResourceModel) ToUpdateSDKType() *shared.SourcePunkAPIPutRequest {
	brewedAfter := r.Configuration.BrewedAfter.ValueString()
	brewedBefore := r.Configuration.BrewedBefore.ValueString()
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	configuration := shared.SourcePunkAPIUpdate{
		BrewedAfter:  brewedAfter,
		BrewedBefore: brewedBefore,
		ID:           id,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePunkAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePunkAPIResourceModel) ToDeleteSDKType() *shared.SourcePunkAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePunkAPIResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePunkAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
