// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceXeroResourceModel) ToCreateSDKType() *shared.SourceXeroCreateRequest {
	accessToken := r.Configuration.Authentication.AccessToken.ValueString()
	clientID := r.Configuration.Authentication.ClientID.ValueString()
	clientSecret := r.Configuration.Authentication.ClientSecret.ValueString()
	refreshToken := r.Configuration.Authentication.RefreshToken.ValueString()
	tokenExpiryDate := r.Configuration.Authentication.TokenExpiryDate.ValueString()
	authentication := shared.SourceXeroAuthenticateViaXeroOAuth{
		AccessToken:     accessToken,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		RefreshToken:    refreshToken,
		TokenExpiryDate: tokenExpiryDate,
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	tenantID := r.Configuration.TenantID.ValueString()
	configuration := shared.SourceXero{
		Authentication: authentication,
		StartDate:      startDate,
		TenantID:       tenantID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceXeroCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceXeroResourceModel) ToGetSDKType() *shared.SourceXeroCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceXeroResourceModel) ToUpdateSDKType() *shared.SourceXeroPutRequest {
	accessToken := r.Configuration.Authentication.AccessToken.ValueString()
	clientID := r.Configuration.Authentication.ClientID.ValueString()
	clientSecret := r.Configuration.Authentication.ClientSecret.ValueString()
	refreshToken := r.Configuration.Authentication.RefreshToken.ValueString()
	tokenExpiryDate := r.Configuration.Authentication.TokenExpiryDate.ValueString()
	authentication := shared.SourceXeroUpdateAuthenticateViaXeroOAuth{
		AccessToken:     accessToken,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		RefreshToken:    refreshToken,
		TokenExpiryDate: tokenExpiryDate,
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	tenantID := r.Configuration.TenantID.ValueString()
	configuration := shared.SourceXeroUpdate{
		Authentication: authentication,
		StartDate:      startDate,
		TenantID:       tenantID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceXeroPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceXeroResourceModel) ToDeleteSDKType() *shared.SourceXeroCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceXeroResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceXeroResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
