// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePexelsAPIResourceModel) ToCreateSDKType() *shared.SourcePexelsAPICreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	color := new(string)
	if !r.Configuration.Color.IsUnknown() && !r.Configuration.Color.IsNull() {
		*color = r.Configuration.Color.ValueString()
	} else {
		color = nil
	}
	locale := new(string)
	if !r.Configuration.Locale.IsUnknown() && !r.Configuration.Locale.IsNull() {
		*locale = r.Configuration.Locale.ValueString()
	} else {
		locale = nil
	}
	orientation := new(string)
	if !r.Configuration.Orientation.IsUnknown() && !r.Configuration.Orientation.IsNull() {
		*orientation = r.Configuration.Orientation.ValueString()
	} else {
		orientation = nil
	}
	query := r.Configuration.Query.ValueString()
	size := new(string)
	if !r.Configuration.Size.IsUnknown() && !r.Configuration.Size.IsNull() {
		*size = r.Configuration.Size.ValueString()
	} else {
		size = nil
	}
	sourceType := shared.SourcePexelsAPIPexelsAPI(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePexelsAPI{
		APIKey:      apiKey,
		Color:       color,
		Locale:      locale,
		Orientation: orientation,
		Query:       query,
		Size:        size,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePexelsAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePexelsAPIResourceModel) ToGetSDKType() *shared.SourcePexelsAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePexelsAPIResourceModel) ToUpdateSDKType() *shared.SourcePexelsAPIPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	color := new(string)
	if !r.Configuration.Color.IsUnknown() && !r.Configuration.Color.IsNull() {
		*color = r.Configuration.Color.ValueString()
	} else {
		color = nil
	}
	locale := new(string)
	if !r.Configuration.Locale.IsUnknown() && !r.Configuration.Locale.IsNull() {
		*locale = r.Configuration.Locale.ValueString()
	} else {
		locale = nil
	}
	orientation := new(string)
	if !r.Configuration.Orientation.IsUnknown() && !r.Configuration.Orientation.IsNull() {
		*orientation = r.Configuration.Orientation.ValueString()
	} else {
		orientation = nil
	}
	query := r.Configuration.Query.ValueString()
	size := new(string)
	if !r.Configuration.Size.IsUnknown() && !r.Configuration.Size.IsNull() {
		*size = r.Configuration.Size.ValueString()
	} else {
		size = nil
	}
	configuration := shared.SourcePexelsAPIUpdate{
		APIKey:      apiKey,
		Color:       color,
		Locale:      locale,
		Orientation: orientation,
		Query:       query,
		Size:        size,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePexelsAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePexelsAPIResourceModel) ToDeleteSDKType() *shared.SourcePexelsAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePexelsAPIResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePexelsAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
