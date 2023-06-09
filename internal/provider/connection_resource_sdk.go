// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *ConnectionResourceModel) ToCreateSDKType() *shared.ConnectionCreateRequest {
	var configurations *shared.StreamConfigurations
	if r.Configurations != nil {
		streams := make([]shared.StreamConfiguration, 0)
		for _, streamsItem := range r.Configurations.Streams {
			cursorField := make([]string, 0)
			for _, cursorFieldItem := range streamsItem.CursorField {
				cursorField = append(cursorField, cursorFieldItem.ValueString())
			}
			name := streamsItem.Name.ValueString()
			primaryKey := make([][]string, 0)
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				primaryKeyTmp := make([]string, 0)
				for _, item := range primaryKeyItem {
					primaryKeyTmp = append(primaryKeyTmp, item.ValueString())
				}
				primaryKey = append(primaryKey, primaryKeyTmp)
			}
			syncMode := new(shared.ConnectionSyncModeEnum)
			if !streamsItem.SyncMode.IsUnknown() && !streamsItem.SyncMode.IsNull() {
				*syncMode = shared.ConnectionSyncModeEnum(streamsItem.SyncMode.ValueString())
			} else {
				syncMode = nil
			}
			streams = append(streams, shared.StreamConfiguration{
				CursorField: cursorField,
				Name:        name,
				PrimaryKey:  primaryKey,
				SyncMode:    syncMode,
			})
		}
		configurations = &shared.StreamConfigurations{
			Streams: streams,
		}
	}
	dataResidency := new(shared.GeographyEnum)
	if !r.DataResidency.IsUnknown() && !r.DataResidency.IsNull() {
		*dataResidency = shared.GeographyEnum(r.DataResidency.ValueString())
	} else {
		dataResidency = nil
	}
	destinationID := r.DestinationID.ValueString()
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	namespaceDefinition := new(shared.NamespaceDefinitionEnum)
	if !r.NamespaceDefinition.IsUnknown() && !r.NamespaceDefinition.IsNull() {
		*namespaceDefinition = shared.NamespaceDefinitionEnum(r.NamespaceDefinition.ValueString())
	} else {
		namespaceDefinition = nil
	}
	namespaceFormat := new(string)
	if !r.NamespaceFormat.IsUnknown() && !r.NamespaceFormat.IsNull() {
		*namespaceFormat = r.NamespaceFormat.ValueString()
	} else {
		namespaceFormat = nil
	}
	nonBreakingSchemaUpdatesBehavior := new(shared.NonBreakingSchemaUpdatesBehaviorEnum)
	if !r.NonBreakingSchemaUpdatesBehavior.IsUnknown() && !r.NonBreakingSchemaUpdatesBehavior.IsNull() {
		*nonBreakingSchemaUpdatesBehavior = shared.NonBreakingSchemaUpdatesBehaviorEnum(r.NonBreakingSchemaUpdatesBehavior.ValueString())
	} else {
		nonBreakingSchemaUpdatesBehavior = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var schedule *shared.ConnectionSchedule
	if r.Schedule != nil {
		scheduleType := shared.ScheduleTypeEnum(r.Schedule.ScheduleType.ValueString())
		schedule = &shared.ConnectionSchedule{
			ScheduleType: scheduleType,
		}
	}
	sourceID := r.SourceID.ValueString()
	status := new(shared.ConnectionStatusEnum)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.ConnectionStatusEnum(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.ConnectionCreateRequest{
		Configurations:                   configurations,
		DataResidency:                    dataResidency,
		DestinationID:                    destinationID,
		Name:                             name1,
		NamespaceDefinition:              namespaceDefinition,
		NamespaceFormat:                  namespaceFormat,
		NonBreakingSchemaUpdatesBehavior: nonBreakingSchemaUpdatesBehavior,
		Prefix:                           prefix,
		Schedule:                         schedule,
		SourceID:                         sourceID,
		Status:                           status,
	}
	return &out
}

func (r *ConnectionResourceModel) ToGetSDKType() *shared.ConnectionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *ConnectionResourceModel) ToUpdateSDKType() *shared.ConnectionPatchRequest {
	var configurations *shared.StreamConfigurations
	if r.Configurations != nil {
		streams := make([]shared.StreamConfiguration, 0)
		for _, streamsItem := range r.Configurations.Streams {
			cursorField := make([]string, 0)
			for _, cursorFieldItem := range streamsItem.CursorField {
				cursorField = append(cursorField, cursorFieldItem.ValueString())
			}
			name := streamsItem.Name.ValueString()
			primaryKey := make([][]string, 0)
			for _, primaryKeyItem := range streamsItem.PrimaryKey {
				primaryKeyTmp := make([]string, 0)
				for _, item := range primaryKeyItem {
					primaryKeyTmp = append(primaryKeyTmp, item.ValueString())
				}
				primaryKey = append(primaryKey, primaryKeyTmp)
			}
			syncMode := new(shared.ConnectionSyncModeEnum)
			if !streamsItem.SyncMode.IsUnknown() && !streamsItem.SyncMode.IsNull() {
				*syncMode = shared.ConnectionSyncModeEnum(streamsItem.SyncMode.ValueString())
			} else {
				syncMode = nil
			}
			streams = append(streams, shared.StreamConfiguration{
				CursorField: cursorField,
				Name:        name,
				PrimaryKey:  primaryKey,
				SyncMode:    syncMode,
			})
		}
		configurations = &shared.StreamConfigurations{
			Streams: streams,
		}
	}
	dataResidency := new(shared.GeographyEnumNoDefault)
	if !r.DataResidency.IsUnknown() && !r.DataResidency.IsNull() {
		*dataResidency = shared.GeographyEnumNoDefault(r.DataResidency.ValueString())
	} else {
		dataResidency = nil
	}
	name1 := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name1 = r.Name.ValueString()
	} else {
		name1 = nil
	}
	namespaceDefinition := new(shared.NamespaceDefinitionEnumNoDefault)
	if !r.NamespaceDefinition.IsUnknown() && !r.NamespaceDefinition.IsNull() {
		*namespaceDefinition = shared.NamespaceDefinitionEnumNoDefault(r.NamespaceDefinition.ValueString())
	} else {
		namespaceDefinition = nil
	}
	namespaceFormat := new(string)
	if !r.NamespaceFormat.IsUnknown() && !r.NamespaceFormat.IsNull() {
		*namespaceFormat = r.NamespaceFormat.ValueString()
	} else {
		namespaceFormat = nil
	}
	nonBreakingSchemaUpdatesBehavior := new(shared.NonBreakingSchemaUpdatesBehaviorEnumNoDefault)
	if !r.NonBreakingSchemaUpdatesBehavior.IsUnknown() && !r.NonBreakingSchemaUpdatesBehavior.IsNull() {
		*nonBreakingSchemaUpdatesBehavior = shared.NonBreakingSchemaUpdatesBehaviorEnumNoDefault(r.NonBreakingSchemaUpdatesBehavior.ValueString())
	} else {
		nonBreakingSchemaUpdatesBehavior = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var schedule *shared.ConnectionSchedule
	if r.Schedule != nil {
		scheduleType := shared.ScheduleTypeEnum(r.Schedule.ScheduleType.ValueString())
		schedule = &shared.ConnectionSchedule{
			ScheduleType: scheduleType,
		}
	}
	status := new(shared.ConnectionStatusEnum)
	if !r.Status.IsUnknown() && !r.Status.IsNull() {
		*status = shared.ConnectionStatusEnum(r.Status.ValueString())
	} else {
		status = nil
	}
	out := shared.ConnectionPatchRequest{
		Configurations:                   configurations,
		DataResidency:                    dataResidency,
		Name:                             name1,
		NamespaceDefinition:              namespaceDefinition,
		NamespaceFormat:                  namespaceFormat,
		NonBreakingSchemaUpdatesBehavior: nonBreakingSchemaUpdatesBehavior,
		Prefix:                           prefix,
		Schedule:                         schedule,
		Status:                           status,
	}
	return &out
}

func (r *ConnectionResourceModel) ToDeleteSDKType() *shared.ConnectionCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *ConnectionResourceModel) RefreshFromGetResponse(resp *shared.ConnectionResponse) {
	if r.Configurations == nil {
		r.Configurations = &StreamConfigurations{}
	}
	r.Configurations.Streams = nil
	for _, streamsItem := range resp.Configurations.Streams {
		var streams1 StreamConfiguration
		streams1.CursorField = nil
		for _, v := range streamsItem.CursorField {
			streams1.CursorField = append(streams1.CursorField, types.StringValue(v))
		}
		streams1.Name = types.StringValue(streamsItem.Name)
		streams1.PrimaryKey = nil
		for _, primaryKeyItem := range streamsItem.PrimaryKey {
			var primaryKey1 []types.String
			primaryKey1 = nil
			for _, v := range primaryKeyItem {
				primaryKey1 = append(primaryKey1, types.StringValue(v))
			}
			streams1.PrimaryKey = append(streams1.PrimaryKey, primaryKey1)
		}
		if streamsItem.SyncMode != nil {
			streams1.SyncMode = types.StringValue(string(*streamsItem.SyncMode))
		} else {
			streams1.SyncMode = types.StringNull()
		}
		r.Configurations.Streams = append(r.Configurations.Streams, streams1)
	}
	r.ConnectionID = types.StringValue(resp.ConnectionID)
	r.DataResidency = types.StringValue(string(resp.DataResidency))
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.Name = types.StringValue(resp.Name)
	if resp.NamespaceDefinition != nil {
		r.NamespaceDefinition = types.StringValue(string(*resp.NamespaceDefinition))
	} else {
		r.NamespaceDefinition = types.StringNull()
	}
	if resp.NamespaceFormat != nil {
		r.NamespaceFormat = types.StringValue(*resp.NamespaceFormat)
	} else {
		r.NamespaceFormat = types.StringNull()
	}
	if resp.NonBreakingSchemaUpdatesBehavior != nil {
		r.NonBreakingSchemaUpdatesBehavior = types.StringValue(string(*resp.NonBreakingSchemaUpdatesBehavior))
	} else {
		r.NonBreakingSchemaUpdatesBehavior = types.StringNull()
	}
	if resp.Prefix != nil {
		r.Prefix = types.StringValue(*resp.Prefix)
	} else {
		r.Prefix = types.StringNull()
	}
	if r.Schedule == nil {
		r.Schedule = &ConnectionSchedule{}
	}
	if resp.Schedule.BasicTiming != nil {
		r.Schedule.BasicTiming = types.StringValue(*resp.Schedule.BasicTiming)
	} else {
		r.Schedule.BasicTiming = types.StringNull()
	}
	r.Schedule.ScheduleType = types.StringValue(string(resp.Schedule.ScheduleType))
	r.SourceID = types.StringValue(resp.SourceID)
	r.Status = types.StringValue(string(resp.Status))
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *ConnectionResourceModel) RefreshFromCreateResponse(resp *shared.ConnectionResponse) {
	r.RefreshFromGetResponse(resp)
}

func (r *ConnectionResourceModel) RefreshFromUpdateResponse(resp *shared.ConnectionResponse) {
	r.RefreshFromGetResponse(resp)
}
