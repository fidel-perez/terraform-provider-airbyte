// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOutreachResourceModel) ToCreateSDKType() *shared.SourceOutreachCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	redirectURI := r.Configuration.RedirectURI.ValueString()
	refreshToken := r.Configuration.RefreshToken.ValueString()
	sourceType := shared.SourceOutreachOutreach(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceOutreach{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		RefreshToken: refreshToken,
		SourceType:   sourceType,
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
	out := shared.SourceOutreachCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOutreachResourceModel) ToGetSDKType() *shared.SourceOutreachCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOutreachResourceModel) ToUpdateSDKType() *shared.SourceOutreachPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	redirectURI := r.Configuration.RedirectURI.ValueString()
	refreshToken := r.Configuration.RefreshToken.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceOutreachUpdate{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		RefreshToken: refreshToken,
		StartDate:    startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOutreachPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOutreachResourceModel) ToDeleteSDKType() *shared.SourceOutreachCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOutreachResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOutreachResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
