// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceMarketoResourceModel) ToCreateSDKType() *shared.SourceMarketoCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	domainURL := r.Configuration.DomainURL.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceMarketo{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		DomainURL:    domainURL,
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
	out := shared.SourceMarketoCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMarketoResourceModel) ToGetSDKType() *shared.SourceMarketoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMarketoResourceModel) ToUpdateSDKType() *shared.SourceMarketoPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	domainURL := r.Configuration.DomainURL.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceMarketoUpdate{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		DomainURL:    domainURL,
		StartDate:    startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMarketoPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMarketoResourceModel) ToDeleteSDKType() *shared.SourceMarketoCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMarketoResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMarketoResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
