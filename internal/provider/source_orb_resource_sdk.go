// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOrbResourceModel) ToCreateSDKType() *shared.SourceOrbCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = nil
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = nil
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrb{
		APIKey:                       apiKey,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
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
	out := shared.SourceOrbCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbResourceModel) ToGetSDKType() *shared.SourceOrbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbResourceModel) ToUpdateSDKType() *shared.SourceOrbPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	lookbackWindowDays := new(int64)
	if !r.Configuration.LookbackWindowDays.IsUnknown() && !r.Configuration.LookbackWindowDays.IsNull() {
		*lookbackWindowDays = r.Configuration.LookbackWindowDays.ValueInt64()
	} else {
		lookbackWindowDays = nil
	}
	var numericEventPropertiesKeys []string = nil
	for _, numericEventPropertiesKeysItem := range r.Configuration.NumericEventPropertiesKeys {
		numericEventPropertiesKeys = append(numericEventPropertiesKeys, numericEventPropertiesKeysItem.ValueString())
	}
	planID := new(string)
	if !r.Configuration.PlanID.IsUnknown() && !r.Configuration.PlanID.IsNull() {
		*planID = r.Configuration.PlanID.ValueString()
	} else {
		planID = nil
	}
	startDate := r.Configuration.StartDate.ValueString()
	var stringEventPropertiesKeys []string = nil
	for _, stringEventPropertiesKeysItem := range r.Configuration.StringEventPropertiesKeys {
		stringEventPropertiesKeys = append(stringEventPropertiesKeys, stringEventPropertiesKeysItem.ValueString())
	}
	subscriptionUsageGroupingKey := new(string)
	if !r.Configuration.SubscriptionUsageGroupingKey.IsUnknown() && !r.Configuration.SubscriptionUsageGroupingKey.IsNull() {
		*subscriptionUsageGroupingKey = r.Configuration.SubscriptionUsageGroupingKey.ValueString()
	} else {
		subscriptionUsageGroupingKey = nil
	}
	configuration := shared.SourceOrbUpdate{
		APIKey:                       apiKey,
		LookbackWindowDays:           lookbackWindowDays,
		NumericEventPropertiesKeys:   numericEventPropertiesKeys,
		PlanID:                       planID,
		StartDate:                    startDate,
		StringEventPropertiesKeys:    stringEventPropertiesKeys,
		SubscriptionUsageGroupingKey: subscriptionUsageGroupingKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOrbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOrbResourceModel) ToDeleteSDKType() *shared.SourceOrbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOrbResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOrbResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
