// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePinterestResourceModel) ToCreateSDKType() *shared.SourcePinterestCreateRequest {
	var credentials *shared.SourcePinterestAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourcePinterestOAuth20 *shared.SourcePinterestOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			sourcePinterestOAuth20 = &shared.SourcePinterestOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourcePinterestOAuth20 != nil {
			credentials = &shared.SourcePinterestAuthorizationMethod{
				SourcePinterestOAuth20: sourcePinterestOAuth20,
			}
		}
	}
	var customReports []shared.SourcePinterestReportConfig = nil
	for _, customReportsItem := range r.Configuration.CustomReports {
		var attributionTypes []shared.SourcePinterestValidEnums = nil
		for _, attributionTypesItem := range customReportsItem.AttributionTypes {
			attributionTypes = append(attributionTypes, shared.SourcePinterestValidEnums(attributionTypesItem.ValueString()))
		}
		clickWindowDays := new(shared.SourcePinterestClickWindowDays)
		if !customReportsItem.ClickWindowDays.IsUnknown() && !customReportsItem.ClickWindowDays.IsNull() {
			*clickWindowDays = shared.SourcePinterestClickWindowDays(customReportsItem.ClickWindowDays.ValueInt64())
		} else {
			clickWindowDays = nil
		}
		var columns []shared.SourcePinterestSchemasValidEnums = nil
		for _, columnsItem := range customReportsItem.Columns {
			columns = append(columns, shared.SourcePinterestSchemasValidEnums(columnsItem.ValueString()))
		}
		conversionReportTime := new(shared.SourcePinterestConversionReportTime)
		if !customReportsItem.ConversionReportTime.IsUnknown() && !customReportsItem.ConversionReportTime.IsNull() {
			*conversionReportTime = shared.SourcePinterestConversionReportTime(customReportsItem.ConversionReportTime.ValueString())
		} else {
			conversionReportTime = nil
		}
		engagementWindowDays := new(shared.SourcePinterestEngagementWindowDays)
		if !customReportsItem.EngagementWindowDays.IsUnknown() && !customReportsItem.EngagementWindowDays.IsNull() {
			*engagementWindowDays = shared.SourcePinterestEngagementWindowDays(customReportsItem.EngagementWindowDays.ValueInt64())
		} else {
			engagementWindowDays = nil
		}
		granularity := new(shared.SourcePinterestGranularity)
		if !customReportsItem.Granularity.IsUnknown() && !customReportsItem.Granularity.IsNull() {
			*granularity = shared.SourcePinterestGranularity(customReportsItem.Granularity.ValueString())
		} else {
			granularity = nil
		}
		level := new(shared.SourcePinterestLevel)
		if !customReportsItem.Level.IsUnknown() && !customReportsItem.Level.IsNull() {
			*level = shared.SourcePinterestLevel(customReportsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customReportsItem.Name.ValueString()
		startDate := new(customTypes.Date)
		if !customReportsItem.StartDate.IsUnknown() && !customReportsItem.StartDate.IsNull() {
			startDate = customTypes.MustNewDateFromString(customReportsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		viewWindowDays := new(shared.SourcePinterestViewWindowDays)
		if !customReportsItem.ViewWindowDays.IsUnknown() && !customReportsItem.ViewWindowDays.IsNull() {
			*viewWindowDays = shared.SourcePinterestViewWindowDays(customReportsItem.ViewWindowDays.ValueInt64())
		} else {
			viewWindowDays = nil
		}
		customReports = append(customReports, shared.SourcePinterestReportConfig{
			AttributionTypes:     attributionTypes,
			ClickWindowDays:      clickWindowDays,
			Columns:              columns,
			ConversionReportTime: conversionReportTime,
			EngagementWindowDays: engagementWindowDays,
			Granularity:          granularity,
			Level:                level,
			Name:                 name,
			StartDate:            startDate,
			ViewWindowDays:       viewWindowDays,
		})
	}
	startDate1 := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate1 = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	var status []shared.SourcePinterestStatus = nil
	for _, statusItem := range r.Configuration.Status {
		status = append(status, shared.SourcePinterestStatus(statusItem.ValueString()))
	}
	configuration := shared.SourcePinterest{
		Credentials:   credentials,
		CustomReports: customReports,
		StartDate:     startDate1,
		Status:        status,
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
	out := shared.SourcePinterestCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePinterestResourceModel) ToGetSDKType() *shared.SourcePinterestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePinterestResourceModel) ToUpdateSDKType() *shared.SourcePinterestPutRequest {
	var credentials *shared.SourcePinterestUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourcePinterestUpdateOAuth20 *shared.SourcePinterestUpdateOAuth20
		if r.Configuration.Credentials.OAuth20 != nil {
			clientID := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.OAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.OAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.OAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.OAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.OAuth20.RefreshToken.ValueString()
			sourcePinterestUpdateOAuth20 = &shared.SourcePinterestUpdateOAuth20{
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourcePinterestUpdateOAuth20 != nil {
			credentials = &shared.SourcePinterestUpdateAuthorizationMethod{
				SourcePinterestUpdateOAuth20: sourcePinterestUpdateOAuth20,
			}
		}
	}
	var customReports []shared.ReportConfig = nil
	for _, customReportsItem := range r.Configuration.CustomReports {
		var attributionTypes []shared.SourcePinterestUpdateValidEnums = nil
		for _, attributionTypesItem := range customReportsItem.AttributionTypes {
			attributionTypes = append(attributionTypes, shared.SourcePinterestUpdateValidEnums(attributionTypesItem.ValueString()))
		}
		clickWindowDays := new(shared.ClickWindowDays)
		if !customReportsItem.ClickWindowDays.IsUnknown() && !customReportsItem.ClickWindowDays.IsNull() {
			*clickWindowDays = shared.ClickWindowDays(customReportsItem.ClickWindowDays.ValueInt64())
		} else {
			clickWindowDays = nil
		}
		var columns []shared.SourcePinterestUpdateSchemasValidEnums = nil
		for _, columnsItem := range customReportsItem.Columns {
			columns = append(columns, shared.SourcePinterestUpdateSchemasValidEnums(columnsItem.ValueString()))
		}
		conversionReportTime := new(shared.ConversionReportTime)
		if !customReportsItem.ConversionReportTime.IsUnknown() && !customReportsItem.ConversionReportTime.IsNull() {
			*conversionReportTime = shared.ConversionReportTime(customReportsItem.ConversionReportTime.ValueString())
		} else {
			conversionReportTime = nil
		}
		engagementWindowDays := new(shared.EngagementWindowDays)
		if !customReportsItem.EngagementWindowDays.IsUnknown() && !customReportsItem.EngagementWindowDays.IsNull() {
			*engagementWindowDays = shared.EngagementWindowDays(customReportsItem.EngagementWindowDays.ValueInt64())
		} else {
			engagementWindowDays = nil
		}
		granularity := new(shared.Granularity)
		if !customReportsItem.Granularity.IsUnknown() && !customReportsItem.Granularity.IsNull() {
			*granularity = shared.Granularity(customReportsItem.Granularity.ValueString())
		} else {
			granularity = nil
		}
		level := new(shared.SourcePinterestUpdateLevel)
		if !customReportsItem.Level.IsUnknown() && !customReportsItem.Level.IsNull() {
			*level = shared.SourcePinterestUpdateLevel(customReportsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customReportsItem.Name.ValueString()
		startDate := new(customTypes.Date)
		if !customReportsItem.StartDate.IsUnknown() && !customReportsItem.StartDate.IsNull() {
			startDate = customTypes.MustNewDateFromString(customReportsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		viewWindowDays := new(shared.ViewWindowDays)
		if !customReportsItem.ViewWindowDays.IsUnknown() && !customReportsItem.ViewWindowDays.IsNull() {
			*viewWindowDays = shared.ViewWindowDays(customReportsItem.ViewWindowDays.ValueInt64())
		} else {
			viewWindowDays = nil
		}
		customReports = append(customReports, shared.ReportConfig{
			AttributionTypes:     attributionTypes,
			ClickWindowDays:      clickWindowDays,
			Columns:              columns,
			ConversionReportTime: conversionReportTime,
			EngagementWindowDays: engagementWindowDays,
			Granularity:          granularity,
			Level:                level,
			Name:                 name,
			StartDate:            startDate,
			ViewWindowDays:       viewWindowDays,
		})
	}
	startDate1 := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate1 = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate1 = nil
	}
	var status []shared.Status = nil
	for _, statusItem := range r.Configuration.Status {
		status = append(status, shared.Status(statusItem.ValueString()))
	}
	configuration := shared.SourcePinterestUpdate{
		Credentials:   credentials,
		CustomReports: customReports,
		StartDate:     startDate1,
		Status:        status,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePinterestPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePinterestResourceModel) ToDeleteSDKType() *shared.SourcePinterestCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePinterestResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePinterestResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
