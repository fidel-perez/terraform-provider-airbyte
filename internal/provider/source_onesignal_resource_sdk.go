// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceOnesignalResourceModel) ToCreateSDKType() *shared.SourceOnesignalCreateRequest {
	var applications []shared.SourceOnesignalApplications = nil
	for _, applicationsItem := range r.Configuration.Applications {
		appAPIKey := applicationsItem.AppAPIKey.ValueString()
		appID := applicationsItem.AppID.ValueString()
		appName := new(string)
		if !applicationsItem.AppName.IsUnknown() && !applicationsItem.AppName.IsNull() {
			*appName = applicationsItem.AppName.ValueString()
		} else {
			appName = nil
		}
		applications = append(applications, shared.SourceOnesignalApplications{
			AppAPIKey: appAPIKey,
			AppID:     appID,
			AppName:   appName,
		})
	}
	outcomeNames := r.Configuration.OutcomeNames.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	userAuthKey := r.Configuration.UserAuthKey.ValueString()
	configuration := shared.SourceOnesignal{
		Applications: applications,
		OutcomeNames: outcomeNames,
		StartDate:    startDate,
		UserAuthKey:  userAuthKey,
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
	out := shared.SourceOnesignalCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOnesignalResourceModel) ToGetSDKType() *shared.SourceOnesignalCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOnesignalResourceModel) ToUpdateSDKType() *shared.SourceOnesignalPutRequest {
	var applications []shared.Applications = nil
	for _, applicationsItem := range r.Configuration.Applications {
		appAPIKey := applicationsItem.AppAPIKey.ValueString()
		appID := applicationsItem.AppID.ValueString()
		appName := new(string)
		if !applicationsItem.AppName.IsUnknown() && !applicationsItem.AppName.IsNull() {
			*appName = applicationsItem.AppName.ValueString()
		} else {
			appName = nil
		}
		applications = append(applications, shared.Applications{
			AppAPIKey: appAPIKey,
			AppID:     appID,
			AppName:   appName,
		})
	}
	outcomeNames := r.Configuration.OutcomeNames.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	userAuthKey := r.Configuration.UserAuthKey.ValueString()
	configuration := shared.SourceOnesignalUpdate{
		Applications: applications,
		OutcomeNames: outcomeNames,
		StartDate:    startDate,
		UserAuthKey:  userAuthKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOnesignalPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOnesignalResourceModel) ToDeleteSDKType() *shared.SourceOnesignalCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOnesignalResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOnesignalResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
