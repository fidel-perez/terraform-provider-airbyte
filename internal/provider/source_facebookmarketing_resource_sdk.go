// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFacebookMarketingResourceModel) ToCreateSDKType() *shared.SourceFacebookMarketingCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var customInsights []shared.SourceFacebookMarketingInsightConfig = nil
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.SourceFacebookMarketingValidActionBreakdowns = nil
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.SourceFacebookMarketingActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.SourceFacebookMarketingActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.SourceFacebookMarketingValidBreakdowns = nil
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingValidEnums = nil
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1 := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	configuration := shared.SourceFacebookMarketing{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToGetSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) ToUpdateSDKType() *shared.SourceFacebookMarketingPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	clientID := new(string)
	if !r.Configuration.ClientID.IsUnknown() && !r.Configuration.ClientID.IsNull() {
		*clientID = r.Configuration.ClientID.ValueString()
	} else {
		clientID = nil
	}
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	var customInsights []shared.InsightConfig = nil
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		var actionBreakdowns []shared.ValidActionBreakdowns = nil
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.ValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		actionReportTime := new(shared.ActionReportTime)
		if !customInsightsItem.ActionReportTime.IsUnknown() && !customInsightsItem.ActionReportTime.IsNull() {
			*actionReportTime = shared.ActionReportTime(customInsightsItem.ActionReportTime.ValueString())
		} else {
			actionReportTime = nil
		}
		var breakdowns []shared.ValidBreakdowns = nil
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.ValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		var fields []shared.SourceFacebookMarketingUpdateValidEnums = nil
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingUpdateValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.Level)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.Level(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.InsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			ActionReportTime:       actionReportTime,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1 := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	configuration := shared.SourceFacebookMarketingUpdate{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		ClientID:                   clientID,
		ClientSecret:               clientSecret,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToDeleteSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
