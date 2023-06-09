// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZohoCrmResourceModel) ToCreateSDKType() *shared.SourceZohoCrmCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dcRegion := shared.SourceZohoCrmDataCenterLocation(r.Configuration.DcRegion.ValueString())
	edition := shared.SourceZohoCRMZohoCRMEdition(r.Configuration.Edition.ValueString())
	environment := shared.SourceZohoCrmEnvironment(r.Configuration.Environment.ValueString())
	refreshToken := r.Configuration.RefreshToken.ValueString()
	sourceType := shared.SourceZohoCrmZohoCrm(r.Configuration.SourceType.ValueString())
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceZohoCrm{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		DcRegion:      dcRegion,
		Edition:       edition,
		Environment:   environment,
		RefreshToken:  refreshToken,
		SourceType:    sourceType,
		StartDatetime: startDatetime,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZohoCrmCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZohoCrmResourceModel) ToGetSDKType() *shared.SourceZohoCrmCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZohoCrmResourceModel) ToUpdateSDKType() *shared.SourceZohoCrmPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dcRegion := shared.SourceZohoCrmUpdateDataCenterLocation(r.Configuration.DcRegion.ValueString())
	edition := shared.SourceZohoCRMUpdateZohoCRMEdition(r.Configuration.Edition.ValueString())
	environment := shared.SourceZohoCrmUpdateEnvironment(r.Configuration.Environment.ValueString())
	refreshToken := r.Configuration.RefreshToken.ValueString()
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceZohoCrmUpdate{
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		DcRegion:      dcRegion,
		Edition:       edition,
		Environment:   environment,
		RefreshToken:  refreshToken,
		StartDatetime: startDatetime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZohoCrmPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZohoCrmResourceModel) ToDeleteSDKType() *shared.SourceZohoCrmCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZohoCrmResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZohoCrmResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
