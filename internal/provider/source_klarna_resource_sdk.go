// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceKlarnaResourceModel) ToCreateSDKType() *shared.SourceKlarnaCreateRequest {
	password := r.Configuration.Password.ValueString()
	playground := new(bool)
	if !r.Configuration.Playground.IsUnknown() && !r.Configuration.Playground.IsNull() {
		*playground = r.Configuration.Playground.ValueBool()
	} else {
		playground = nil
	}
	region := shared.SourceKlarnaRegion(r.Configuration.Region.ValueString())
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceKlarna{
		Password:   password,
		Playground: playground,
		Region:     region,
		Username:   username,
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
	out := shared.SourceKlarnaCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKlarnaResourceModel) ToGetSDKType() *shared.SourceKlarnaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKlarnaResourceModel) ToUpdateSDKType() *shared.SourceKlarnaPutRequest {
	password := r.Configuration.Password.ValueString()
	playground := new(bool)
	if !r.Configuration.Playground.IsUnknown() && !r.Configuration.Playground.IsNull() {
		*playground = r.Configuration.Playground.ValueBool()
	} else {
		playground = nil
	}
	region := shared.SourceKlarnaUpdateRegion(r.Configuration.Region.ValueString())
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceKlarnaUpdate{
		Password:   password,
		Playground: playground,
		Region:     region,
		Username:   username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceKlarnaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceKlarnaResourceModel) ToDeleteSDKType() *shared.SourceKlarnaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceKlarnaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceKlarnaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
