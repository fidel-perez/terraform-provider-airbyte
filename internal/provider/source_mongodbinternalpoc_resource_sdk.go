// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMongodbInternalPocResourceModel) ToCreateSDKType() *shared.SourceMongodbInternalPocCreateRequest {
	authSource := new(string)
	if !r.Configuration.AuthSource.IsUnknown() && !r.Configuration.AuthSource.IsNull() {
		*authSource = r.Configuration.AuthSource.ValueString()
	} else {
		authSource = nil
	}
	connectionString := new(string)
	if !r.Configuration.ConnectionString.IsUnknown() && !r.Configuration.ConnectionString.IsNull() {
		*connectionString = r.Configuration.ConnectionString.ValueString()
	} else {
		connectionString = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	replicaSet := new(string)
	if !r.Configuration.ReplicaSet.IsUnknown() && !r.Configuration.ReplicaSet.IsNull() {
		*replicaSet = r.Configuration.ReplicaSet.ValueString()
	} else {
		replicaSet = nil
	}
	user := new(string)
	if !r.Configuration.User.IsUnknown() && !r.Configuration.User.IsNull() {
		*user = r.Configuration.User.ValueString()
	} else {
		user = nil
	}
	configuration := shared.SourceMongodbInternalPoc{
		AuthSource:       authSource,
		ConnectionString: connectionString,
		Password:         password,
		ReplicaSet:       replicaSet,
		User:             user,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbInternalPocCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMongodbInternalPocResourceModel) ToGetSDKType() *shared.SourceMongodbInternalPocCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMongodbInternalPocResourceModel) ToUpdateSDKType() *shared.SourceMongodbInternalPocPutRequest {
	authSource := new(string)
	if !r.Configuration.AuthSource.IsUnknown() && !r.Configuration.AuthSource.IsNull() {
		*authSource = r.Configuration.AuthSource.ValueString()
	} else {
		authSource = nil
	}
	connectionString := new(string)
	if !r.Configuration.ConnectionString.IsUnknown() && !r.Configuration.ConnectionString.IsNull() {
		*connectionString = r.Configuration.ConnectionString.ValueString()
	} else {
		connectionString = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	replicaSet := new(string)
	if !r.Configuration.ReplicaSet.IsUnknown() && !r.Configuration.ReplicaSet.IsNull() {
		*replicaSet = r.Configuration.ReplicaSet.ValueString()
	} else {
		replicaSet = nil
	}
	user := new(string)
	if !r.Configuration.User.IsUnknown() && !r.Configuration.User.IsNull() {
		*user = r.Configuration.User.ValueString()
	} else {
		user = nil
	}
	configuration := shared.SourceMongodbInternalPocUpdate{
		AuthSource:       authSource,
		ConnectionString: connectionString,
		Password:         password,
		ReplicaSet:       replicaSet,
		User:             user,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMongodbInternalPocPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMongodbInternalPocResourceModel) ToDeleteSDKType() *shared.SourceMongodbInternalPocCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMongodbInternalPocResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMongodbInternalPocResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
