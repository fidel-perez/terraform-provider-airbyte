// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceInsightlyResourceModel) ToCreateSDKType() *shared.SourceInsightlyCreateRequest {
	sourceType := shared.SourceInsightlyInsightly(r.Configuration.SourceType.ValueString())
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	token := new(string)
	if !r.Configuration.Token.IsUnknown() && !r.Configuration.Token.IsNull() {
		*token = r.Configuration.Token.ValueString()
	} else {
		token = nil
	}
	configuration := shared.SourceInsightly{
		SourceType: sourceType,
		StartDate:  startDate,
		Token:      token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceInsightlyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceInsightlyResourceModel) ToGetSDKType() *shared.SourceInsightlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceInsightlyResourceModel) ToUpdateSDKType() *shared.SourceInsightlyPutRequest {
	startDate := new(string)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueString()
	} else {
		startDate = nil
	}
	token := new(string)
	if !r.Configuration.Token.IsUnknown() && !r.Configuration.Token.IsNull() {
		*token = r.Configuration.Token.ValueString()
	} else {
		token = nil
	}
	configuration := shared.SourceInsightlyUpdate{
		StartDate: startDate,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceInsightlyPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceInsightlyResourceModel) ToDeleteSDKType() *shared.SourceInsightlyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceInsightlyResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceInsightlyResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
