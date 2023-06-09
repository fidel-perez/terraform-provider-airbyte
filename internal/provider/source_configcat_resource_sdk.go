// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceConfigcatResourceModel) ToCreateSDKType() *shared.SourceConfigcatCreateRequest {
	password := r.Configuration.Password.ValueString()
	sourceType := shared.SourceConfigcatConfigcat(r.Configuration.SourceType.ValueString())
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceConfigcat{
		Password:   password,
		SourceType: sourceType,
		Username:   username,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceConfigcatCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConfigcatResourceModel) ToGetSDKType() *shared.SourceConfigcatCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceConfigcatResourceModel) ToUpdateSDKType() *shared.SourceConfigcatPutRequest {
	password := r.Configuration.Password.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceConfigcatUpdate{
		Password: password,
		Username: username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceConfigcatPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceConfigcatResourceModel) ToDeleteSDKType() *shared.SourceConfigcatCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceConfigcatResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceConfigcatResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
