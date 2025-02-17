// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceExchangeRatesResourceModel) ToCreateSDKType() *shared.SourceExchangeRatesCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	base := new(string)
	if !r.Configuration.Base.IsUnknown() && !r.Configuration.Base.IsNull() {
		*base = r.Configuration.Base.ValueString()
	} else {
		base = nil
	}
	ignoreWeekends := new(bool)
	if !r.Configuration.IgnoreWeekends.IsUnknown() && !r.Configuration.IgnoreWeekends.IsNull() {
		*ignoreWeekends = r.Configuration.IgnoreWeekends.ValueBool()
	} else {
		ignoreWeekends = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceExchangeRates{
		AccessKey:      accessKey,
		Base:           base,
		IgnoreWeekends: ignoreWeekends,
		StartDate:      startDate,
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
	out := shared.SourceExchangeRatesCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceExchangeRatesResourceModel) ToGetSDKType() *shared.SourceExchangeRatesCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceExchangeRatesResourceModel) ToUpdateSDKType() *shared.SourceExchangeRatesPutRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	base := new(string)
	if !r.Configuration.Base.IsUnknown() && !r.Configuration.Base.IsNull() {
		*base = r.Configuration.Base.ValueString()
	} else {
		base = nil
	}
	ignoreWeekends := new(bool)
	if !r.Configuration.IgnoreWeekends.IsUnknown() && !r.Configuration.IgnoreWeekends.IsNull() {
		*ignoreWeekends = r.Configuration.IgnoreWeekends.ValueBool()
	} else {
		ignoreWeekends = nil
	}
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	configuration := shared.SourceExchangeRatesUpdate{
		AccessKey:      accessKey,
		Base:           base,
		IgnoreWeekends: ignoreWeekends,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceExchangeRatesPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceExchangeRatesResourceModel) ToDeleteSDKType() *shared.SourceExchangeRatesCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceExchangeRatesResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceExchangeRatesResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
