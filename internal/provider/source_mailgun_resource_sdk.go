// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceMailgunResourceModel) ToCreateSDKType() *shared.SourceMailgunCreateRequest {
	domainRegion := new(string)
	if !r.Configuration.DomainRegion.IsUnknown() && !r.Configuration.DomainRegion.IsNull() {
		*domainRegion = r.Configuration.DomainRegion.ValueString()
	} else {
		domainRegion = nil
	}
	privateKey := r.Configuration.PrivateKey.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceMailgun{
		DomainRegion: domainRegion,
		PrivateKey:   privateKey,
		StartDate:    startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailgunCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailgunResourceModel) ToGetSDKType() *shared.SourceMailgunCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailgunResourceModel) ToUpdateSDKType() *shared.SourceMailgunPutRequest {
	domainRegion := new(string)
	if !r.Configuration.DomainRegion.IsUnknown() && !r.Configuration.DomainRegion.IsNull() {
		*domainRegion = r.Configuration.DomainRegion.ValueString()
	} else {
		domainRegion = nil
	}
	privateKey := r.Configuration.PrivateKey.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceMailgunUpdate{
		DomainRegion: domainRegion,
		PrivateKey:   privateKey,
		StartDate:    startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailgunPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailgunResourceModel) ToDeleteSDKType() *shared.SourceMailgunCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailgunResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMailgunResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
