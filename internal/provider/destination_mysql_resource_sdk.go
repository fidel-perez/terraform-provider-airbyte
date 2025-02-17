// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationMysqlResourceModel) ToCreateSDKType() *shared.DestinationMysqlCreateRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var tunnelMethod *shared.DestinationMysqlSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMysqlNoTunnel *shared.DestinationMysqlNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMysqlNoTunnel = &shared.DestinationMysqlNoTunnel{}
		}
		if destinationMysqlNoTunnel != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlNoTunnel: destinationMysqlNoTunnel,
			}
		}
		var destinationMysqlSSHKeyAuthentication *shared.DestinationMysqlSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationMysqlSSHKeyAuthentication = &shared.DestinationMysqlSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMysqlSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlSSHKeyAuthentication: destinationMysqlSSHKeyAuthentication,
			}
		}
		var destinationMysqlPasswordAuthentication *shared.DestinationMysqlPasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMysqlPasswordAuthentication = &shared.DestinationMysqlPasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMysqlPasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlSSHTunnelMethod{
				DestinationMysqlPasswordAuthentication: destinationMysqlPasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMysql{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMysqlCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMysqlResourceModel) ToGetSDKType() *shared.DestinationMysqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMysqlResourceModel) ToUpdateSDKType() *shared.DestinationMysqlPutRequest {
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	password := new(string)
	if !r.Configuration.Password.IsUnknown() && !r.Configuration.Password.IsNull() {
		*password = r.Configuration.Password.ValueString()
	} else {
		password = nil
	}
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	var tunnelMethod *shared.DestinationMysqlUpdateSSHTunnelMethod
	if r.Configuration.TunnelMethod != nil {
		var destinationMysqlUpdateNoTunnel *shared.DestinationMysqlUpdateNoTunnel
		if r.Configuration.TunnelMethod.NoTunnel != nil {
			destinationMysqlUpdateNoTunnel = &shared.DestinationMysqlUpdateNoTunnel{}
		}
		if destinationMysqlUpdateNoTunnel != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdateNoTunnel: destinationMysqlUpdateNoTunnel,
			}
		}
		var destinationMysqlUpdateSSHKeyAuthentication *shared.DestinationMysqlUpdateSSHKeyAuthentication
		if r.Configuration.TunnelMethod.SSHKeyAuthentication != nil {
			sshKey := r.Configuration.TunnelMethod.SSHKeyAuthentication.SSHKey.ValueString()
			tunnelHost := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelHost.ValueString()
			tunnelPort := new(int64)
			if !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.IsNull() {
				*tunnelPort = r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort = nil
			}
			tunnelUser := r.Configuration.TunnelMethod.SSHKeyAuthentication.TunnelUser.ValueString()
			destinationMysqlUpdateSSHKeyAuthentication = &shared.DestinationMysqlUpdateSSHKeyAuthentication{
				SSHKey:     sshKey,
				TunnelHost: tunnelHost,
				TunnelPort: tunnelPort,
				TunnelUser: tunnelUser,
			}
		}
		if destinationMysqlUpdateSSHKeyAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdateSSHKeyAuthentication: destinationMysqlUpdateSSHKeyAuthentication,
			}
		}
		var destinationMysqlUpdatePasswordAuthentication *shared.DestinationMysqlUpdatePasswordAuthentication
		if r.Configuration.TunnelMethod.PasswordAuthentication != nil {
			tunnelHost1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelHost.ValueString()
			tunnelPort1 := new(int64)
			if !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsUnknown() && !r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.IsNull() {
				*tunnelPort1 = r.Configuration.TunnelMethod.PasswordAuthentication.TunnelPort.ValueInt64()
			} else {
				tunnelPort1 = nil
			}
			tunnelUser1 := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUser.ValueString()
			tunnelUserPassword := r.Configuration.TunnelMethod.PasswordAuthentication.TunnelUserPassword.ValueString()
			destinationMysqlUpdatePasswordAuthentication = &shared.DestinationMysqlUpdatePasswordAuthentication{
				TunnelHost:         tunnelHost1,
				TunnelPort:         tunnelPort1,
				TunnelUser:         tunnelUser1,
				TunnelUserPassword: tunnelUserPassword,
			}
		}
		if destinationMysqlUpdatePasswordAuthentication != nil {
			tunnelMethod = &shared.DestinationMysqlUpdateSSHTunnelMethod{
				DestinationMysqlUpdatePasswordAuthentication: destinationMysqlUpdatePasswordAuthentication,
			}
		}
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationMysqlUpdate{
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Password:      password,
		Port:          port,
		TunnelMethod:  tunnelMethod,
		Username:      username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationMysqlPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationMysqlResourceModel) ToDeleteSDKType() *shared.DestinationMysqlCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationMysqlResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationMysqlResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
