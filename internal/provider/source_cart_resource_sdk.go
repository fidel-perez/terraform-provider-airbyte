// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceCartResourceModel) ToCreateSDKType() *shared.SourceCartCreateRequest {
	var credentials *shared.SourceCartAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceCartCentralAPIRouter *shared.SourceCartCentralAPIRouter
		if r.Configuration.Credentials.CentralAPIRouter != nil {
			siteID := r.Configuration.Credentials.CentralAPIRouter.SiteID.ValueString()
			userName := r.Configuration.Credentials.CentralAPIRouter.UserName.ValueString()
			userSecret := r.Configuration.Credentials.CentralAPIRouter.UserSecret.ValueString()
			sourceCartCentralAPIRouter = &shared.SourceCartCentralAPIRouter{
				SiteID:     siteID,
				UserName:   userName,
				UserSecret: userSecret,
			}
		}
		if sourceCartCentralAPIRouter != nil {
			credentials = &shared.SourceCartAuthorizationMethod{
				SourceCartCentralAPIRouter: sourceCartCentralAPIRouter,
			}
		}
		var sourceCartSingleStoreAccessToken *shared.SourceCartSingleStoreAccessToken
		if r.Configuration.Credentials.SingleStoreAccessToken != nil {
			accessToken := r.Configuration.Credentials.SingleStoreAccessToken.AccessToken.ValueString()
			storeName := r.Configuration.Credentials.SingleStoreAccessToken.StoreName.ValueString()
			sourceCartSingleStoreAccessToken = &shared.SourceCartSingleStoreAccessToken{
				AccessToken: accessToken,
				StoreName:   storeName,
			}
		}
		if sourceCartSingleStoreAccessToken != nil {
			credentials = &shared.SourceCartAuthorizationMethod{
				SourceCartSingleStoreAccessToken: sourceCartSingleStoreAccessToken,
			}
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceCart{
		Credentials: credentials,
		StartDate:   startDate,
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
	out := shared.SourceCartCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceCartResourceModel) ToGetSDKType() *shared.SourceCartCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceCartResourceModel) ToUpdateSDKType() *shared.SourceCartPutRequest {
	var credentials *shared.SourceCartUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var centralAPIRouter *shared.CentralAPIRouter
		if r.Configuration.Credentials.CentralAPIRouter != nil {
			siteID := r.Configuration.Credentials.CentralAPIRouter.SiteID.ValueString()
			userName := r.Configuration.Credentials.CentralAPIRouter.UserName.ValueString()
			userSecret := r.Configuration.Credentials.CentralAPIRouter.UserSecret.ValueString()
			centralAPIRouter = &shared.CentralAPIRouter{
				SiteID:     siteID,
				UserName:   userName,
				UserSecret: userSecret,
			}
		}
		if centralAPIRouter != nil {
			credentials = &shared.SourceCartUpdateAuthorizationMethod{
				CentralAPIRouter: centralAPIRouter,
			}
		}
		var singleStoreAccessToken *shared.SingleStoreAccessToken
		if r.Configuration.Credentials.SingleStoreAccessToken != nil {
			accessToken := r.Configuration.Credentials.SingleStoreAccessToken.AccessToken.ValueString()
			storeName := r.Configuration.Credentials.SingleStoreAccessToken.StoreName.ValueString()
			singleStoreAccessToken = &shared.SingleStoreAccessToken{
				AccessToken: accessToken,
				StoreName:   storeName,
			}
		}
		if singleStoreAccessToken != nil {
			credentials = &shared.SourceCartUpdateAuthorizationMethod{
				SingleStoreAccessToken: singleStoreAccessToken,
			}
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceCartUpdate{
		Credentials: credentials,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceCartPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceCartResourceModel) ToDeleteSDKType() *shared.SourceCartCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceCartResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceCartResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
