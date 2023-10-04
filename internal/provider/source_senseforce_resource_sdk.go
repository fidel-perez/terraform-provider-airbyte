// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSenseforceResourceModel) ToCreateSDKType() *shared.SourceSenseforceCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	backendURL := r.Configuration.BackendURL.ValueString()
	datasetID := r.Configuration.DatasetID.ValueString()
	sliceRange := new(int64)
	if !r.Configuration.SliceRange.IsUnknown() && !r.Configuration.SliceRange.IsNull() {
		*sliceRange = r.Configuration.SliceRange.ValueInt64()
	} else {
		sliceRange = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSenseforce{
		AccessToken: accessToken,
		BackendURL:  backendURL,
		DatasetID:   datasetID,
		SliceRange:  sliceRange,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSenseforceCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSenseforceResourceModel) ToGetSDKType() *shared.SourceSenseforceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSenseforceResourceModel) ToUpdateSDKType() *shared.SourceSenseforcePutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	backendURL := r.Configuration.BackendURL.ValueString()
	datasetID := r.Configuration.DatasetID.ValueString()
	sliceRange := new(int64)
	if !r.Configuration.SliceRange.IsUnknown() && !r.Configuration.SliceRange.IsNull() {
		*sliceRange = r.Configuration.SliceRange.ValueInt64()
	} else {
		sliceRange = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSenseforceUpdate{
		AccessToken: accessToken,
		BackendURL:  backendURL,
		DatasetID:   datasetID,
		SliceRange:  sliceRange,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSenseforcePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSenseforceResourceModel) ToDeleteSDKType() *shared.SourceSenseforceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSenseforceResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSenseforceResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
