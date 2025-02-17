// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceAmazonSellerPartnerResourceModel) ToCreateSDKType() *shared.SourceAmazonSellerPartnerCreateRequest {
	accountType := new(shared.SourceAmazonSellerPartnerAWSSellerPartnerAccountType)
	if !r.Configuration.AccountType.IsUnknown() && !r.Configuration.AccountType.IsNull() {
		*accountType = shared.SourceAmazonSellerPartnerAWSSellerPartnerAccountType(r.Configuration.AccountType.ValueString())
	} else {
		accountType = nil
	}
	advancedStreamOptions := new(string)
	if !r.Configuration.AdvancedStreamOptions.IsUnknown() && !r.Configuration.AdvancedStreamOptions.IsNull() {
		*advancedStreamOptions = r.Configuration.AdvancedStreamOptions.ValueString()
	} else {
		advancedStreamOptions = nil
	}
	awsEnvironment := new(shared.SourceAmazonSellerPartnerAWSEnvironment)
	if !r.Configuration.AwsEnvironment.IsUnknown() && !r.Configuration.AwsEnvironment.IsNull() {
		*awsEnvironment = shared.SourceAmazonSellerPartnerAWSEnvironment(r.Configuration.AwsEnvironment.ValueString())
	} else {
		awsEnvironment = nil
	}
	lwaAppID := r.Configuration.LwaAppID.ValueString()
	lwaClientSecret := r.Configuration.LwaClientSecret.ValueString()
	periodInDays := new(int64)
	if !r.Configuration.PeriodInDays.IsUnknown() && !r.Configuration.PeriodInDays.IsNull() {
		*periodInDays = r.Configuration.PeriodInDays.ValueInt64()
	} else {
		periodInDays = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	region := new(shared.SourceAmazonSellerPartnerAWSRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceAmazonSellerPartnerAWSRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	reportOptions := new(string)
	if !r.Configuration.ReportOptions.IsUnknown() && !r.Configuration.ReportOptions.IsNull() {
		*reportOptions = r.Configuration.ReportOptions.ValueString()
	} else {
		reportOptions = nil
	}
	configuration := shared.SourceAmazonSellerPartner{
		AccountType:           accountType,
		AdvancedStreamOptions: advancedStreamOptions,
		AwsEnvironment:        awsEnvironment,
		LwaAppID:              lwaAppID,
		LwaClientSecret:       lwaClientSecret,
		PeriodInDays:          periodInDays,
		RefreshToken:          refreshToken,
		Region:                region,
		ReplicationEndDate:    replicationEndDate,
		ReplicationStartDate:  replicationStartDate,
		ReportOptions:         reportOptions,
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
	out := shared.SourceAmazonSellerPartnerCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmazonSellerPartnerResourceModel) ToGetSDKType() *shared.SourceAmazonSellerPartnerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmazonSellerPartnerResourceModel) ToUpdateSDKType() *shared.SourceAmazonSellerPartnerPutRequest {
	accountType := new(shared.AWSSellerPartnerAccountType)
	if !r.Configuration.AccountType.IsUnknown() && !r.Configuration.AccountType.IsNull() {
		*accountType = shared.AWSSellerPartnerAccountType(r.Configuration.AccountType.ValueString())
	} else {
		accountType = nil
	}
	advancedStreamOptions := new(string)
	if !r.Configuration.AdvancedStreamOptions.IsUnknown() && !r.Configuration.AdvancedStreamOptions.IsNull() {
		*advancedStreamOptions = r.Configuration.AdvancedStreamOptions.ValueString()
	} else {
		advancedStreamOptions = nil
	}
	awsEnvironment := new(shared.AWSEnvironment)
	if !r.Configuration.AwsEnvironment.IsUnknown() && !r.Configuration.AwsEnvironment.IsNull() {
		*awsEnvironment = shared.AWSEnvironment(r.Configuration.AwsEnvironment.ValueString())
	} else {
		awsEnvironment = nil
	}
	lwaAppID := r.Configuration.LwaAppID.ValueString()
	lwaClientSecret := r.Configuration.LwaClientSecret.ValueString()
	periodInDays := new(int64)
	if !r.Configuration.PeriodInDays.IsUnknown() && !r.Configuration.PeriodInDays.IsNull() {
		*periodInDays = r.Configuration.PeriodInDays.ValueInt64()
	} else {
		periodInDays = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	region := new(shared.AWSRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.AWSRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	replicationEndDate := new(time.Time)
	if !r.Configuration.ReplicationEndDate.IsUnknown() && !r.Configuration.ReplicationEndDate.IsNull() {
		*replicationEndDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.ReplicationEndDate.ValueString())
	} else {
		replicationEndDate = nil
	}
	replicationStartDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.ReplicationStartDate.ValueString())
	reportOptions := new(string)
	if !r.Configuration.ReportOptions.IsUnknown() && !r.Configuration.ReportOptions.IsNull() {
		*reportOptions = r.Configuration.ReportOptions.ValueString()
	} else {
		reportOptions = nil
	}
	configuration := shared.SourceAmazonSellerPartnerUpdate{
		AccountType:           accountType,
		AdvancedStreamOptions: advancedStreamOptions,
		AwsEnvironment:        awsEnvironment,
		LwaAppID:              lwaAppID,
		LwaClientSecret:       lwaClientSecret,
		PeriodInDays:          periodInDays,
		RefreshToken:          refreshToken,
		Region:                region,
		ReplicationEndDate:    replicationEndDate,
		ReplicationStartDate:  replicationStartDate,
		ReportOptions:         reportOptions,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAmazonSellerPartnerPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAmazonSellerPartnerResourceModel) ToDeleteSDKType() *shared.SourceAmazonSellerPartnerCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAmazonSellerPartnerResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceAmazonSellerPartnerResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
