// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTrelloResourceModel) ToCreateSDKType() *shared.SourceTrelloCreateRequest {
	var boardIds []string = nil
	for _, boardIdsItem := range r.Configuration.BoardIds {
		boardIds = append(boardIds, boardIdsItem.ValueString())
	}
	key := r.Configuration.Key.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceTrello{
		BoardIds:  boardIds,
		Key:       key,
		StartDate: startDate,
		Token:     token,
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
	out := shared.SourceTrelloCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrelloResourceModel) ToGetSDKType() *shared.SourceTrelloCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrelloResourceModel) ToUpdateSDKType() *shared.SourceTrelloPutRequest {
	var boardIds []string = nil
	for _, boardIdsItem := range r.Configuration.BoardIds {
		boardIds = append(boardIds, boardIdsItem.ValueString())
	}
	key := r.Configuration.Key.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceTrelloUpdate{
		BoardIds:  boardIds,
		Key:       key,
		StartDate: startDate,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrelloPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrelloResourceModel) ToDeleteSDKType() *shared.SourceTrelloCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrelloResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTrelloResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
