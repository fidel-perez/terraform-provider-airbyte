// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAwsCloudtrailResourceModel) ToCreateSDKType() *shared.SourceAwsCloudtrailCreateRequest {
	awsKeyID := r.Configuration.AwsKeyID.ValueString()
	awsRegionName := r.Configuration.AwsRegionName.ValueString()
	awsSecretKey := r.Configuration.AwsSecretKey.ValueString()
	sourceType := shared.SourceAwsCloudtrailAwsCloudtrail(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceAwsCloudtrail{
		AwsKeyID:      awsKeyID,
		AwsRegionName: awsRegionName,
		AwsSecretKey:  awsSecretKey,
		SourceType:    sourceType,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAwsCloudtrailCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAwsCloudtrailResourceModel) ToGetSDKType() *shared.SourceAwsCloudtrailCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAwsCloudtrailResourceModel) ToUpdateSDKType() *shared.SourceAwsCloudtrailPutRequest {
	awsKeyID := r.Configuration.AwsKeyID.ValueString()
	awsRegionName := r.Configuration.AwsRegionName.ValueString()
	awsSecretKey := r.Configuration.AwsSecretKey.ValueString()
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceAwsCloudtrailUpdate{
		AwsKeyID:      awsKeyID,
		AwsRegionName: awsRegionName,
		AwsSecretKey:  awsSecretKey,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAwsCloudtrailPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAwsCloudtrailResourceModel) ToDeleteSDKType() *shared.SourceAwsCloudtrailCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAwsCloudtrailResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAwsCloudtrailResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
