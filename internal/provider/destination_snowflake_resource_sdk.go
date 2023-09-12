// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationSnowflakeResourceModel) ToCreateSDKType() *shared.DestinationSnowflakeCreateRequest {
	var credentials *shared.DestinationSnowflakeAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var destinationSnowflakeAuthorizationMethodOAuth20 *shared.DestinationSnowflakeAuthorizationMethodOAuth20
		if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AccessToken.ValueString()
			authType := new(shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.IsNull() {
				*authType = shared.DestinationSnowflakeAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodOAuth20.RefreshToken.ValueString()
			destinationSnowflakeAuthorizationMethodOAuth20 = &shared.DestinationSnowflakeAuthorizationMethodOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if destinationSnowflakeAuthorizationMethodOAuth20 != nil {
			credentials = &shared.DestinationSnowflakeAuthorizationMethod{
				DestinationSnowflakeAuthorizationMethodOAuth20: destinationSnowflakeAuthorizationMethodOAuth20,
			}
		}
		var destinationSnowflakeAuthorizationMethodKeyPairAuthentication *shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication
		if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
			authType1 := new(shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.IsNull() {
				*authType1 = shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthenticationAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			privateKey := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKey.ValueString()
			privateKeyPassword := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsNull() {
				*privateKeyPassword = r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.ValueString()
			} else {
				privateKeyPassword = nil
			}
			destinationSnowflakeAuthorizationMethodKeyPairAuthentication = &shared.DestinationSnowflakeAuthorizationMethodKeyPairAuthentication{
				AuthType:           authType1,
				PrivateKey:         privateKey,
				PrivateKeyPassword: privateKeyPassword,
			}
		}
		if destinationSnowflakeAuthorizationMethodKeyPairAuthentication != nil {
			credentials = &shared.DestinationSnowflakeAuthorizationMethod{
				DestinationSnowflakeAuthorizationMethodKeyPairAuthentication: destinationSnowflakeAuthorizationMethodKeyPairAuthentication,
			}
		}
		var destinationSnowflakeAuthorizationMethodUsernameAndPassword *shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword
		if r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
			authType2 := new(shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.IsNull() {
				*authType2 = shared.DestinationSnowflakeAuthorizationMethodUsernameAndPasswordAuthType(r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
			} else {
				authType2 = nil
			}
			password := r.Configuration.Credentials.DestinationSnowflakeAuthorizationMethodUsernameAndPassword.Password.ValueString()
			destinationSnowflakeAuthorizationMethodUsernameAndPassword = &shared.DestinationSnowflakeAuthorizationMethodUsernameAndPassword{
				AuthType: authType2,
				Password: password,
			}
		}
		if destinationSnowflakeAuthorizationMethodUsernameAndPassword != nil {
			credentials = &shared.DestinationSnowflakeAuthorizationMethod{
				DestinationSnowflakeAuthorizationMethodUsernameAndPassword: destinationSnowflakeAuthorizationMethodUsernameAndPassword,
			}
		}
	}
	database := r.Configuration.Database.ValueString()
	destinationType := shared.DestinationSnowflakeSnowflake(r.Configuration.DestinationType.ValueString())
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	rawDataSchema := new(string)
	if !r.Configuration.RawDataSchema.IsUnknown() && !r.Configuration.RawDataSchema.IsNull() {
		*rawDataSchema = r.Configuration.RawDataSchema.ValueString()
	} else {
		rawDataSchema = nil
	}
	role := r.Configuration.Role.ValueString()
	schema := r.Configuration.Schema.ValueString()
	username := r.Configuration.Username.ValueString()
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.DestinationSnowflake{
		Credentials:     credentials,
		Database:        database,
		DestinationType: destinationType,
		Host:            host,
		JdbcURLParams:   jdbcURLParams,
		RawDataSchema:   rawDataSchema,
		Role:            role,
		Schema:          schema,
		Username:        username,
		Warehouse:       warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSnowflakeCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSnowflakeResourceModel) ToGetSDKType() *shared.DestinationSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSnowflakeResourceModel) ToUpdateSDKType() *shared.DestinationSnowflakePutRequest {
	var credentials *shared.DestinationSnowflakeUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var destinationSnowflakeUpdateAuthorizationMethodOAuth20 *shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20
		if r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
			authType := new(shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20AuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
				*authType = shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			refreshToken := r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
			destinationSnowflakeUpdateAuthorizationMethodOAuth20 = &shared.DestinationSnowflakeUpdateAuthorizationMethodOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if destinationSnowflakeUpdateAuthorizationMethodOAuth20 != nil {
			credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
				DestinationSnowflakeUpdateAuthorizationMethodOAuth20: destinationSnowflakeUpdateAuthorizationMethodOAuth20,
			}
		}
		var destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication *shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication
		if r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication != nil {
			authType1 := new(shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthenticationAuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.AuthType.IsNull() {
				*authType1 = shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthenticationAuthType(r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			privateKey := r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.PrivateKey.ValueString()
			privateKeyPassword := new(string)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.IsNull() {
				*privateKeyPassword = r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication.PrivateKeyPassword.ValueString()
			} else {
				privateKeyPassword = nil
			}
			destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication = &shared.DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication{
				AuthType:           authType1,
				PrivateKey:         privateKey,
				PrivateKeyPassword: privateKeyPassword,
			}
		}
		if destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication != nil {
			credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
				DestinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication: destinationSnowflakeUpdateAuthorizationMethodKeyPairAuthentication,
			}
		}
		var destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword *shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword
		if r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword != nil {
			authType2 := new(shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPasswordAuthType)
			if !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword.AuthType.IsUnknown() && !r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword.AuthType.IsNull() {
				*authType2 = shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPasswordAuthType(r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
			} else {
				authType2 = nil
			}
			password := r.Configuration.Credentials.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword.Password.ValueString()
			destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword = &shared.DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword{
				AuthType: authType2,
				Password: password,
			}
		}
		if destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword != nil {
			credentials = &shared.DestinationSnowflakeUpdateAuthorizationMethod{
				DestinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword: destinationSnowflakeUpdateAuthorizationMethodUsernameAndPassword,
			}
		}
	}
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	rawDataSchema := new(string)
	if !r.Configuration.RawDataSchema.IsUnknown() && !r.Configuration.RawDataSchema.IsNull() {
		*rawDataSchema = r.Configuration.RawDataSchema.ValueString()
	} else {
		rawDataSchema = nil
	}
	role := r.Configuration.Role.ValueString()
	schema := r.Configuration.Schema.ValueString()
	username := r.Configuration.Username.ValueString()
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.DestinationSnowflakeUpdate{
		Credentials:   credentials,
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		RawDataSchema: rawDataSchema,
		Role:          role,
		Schema:        schema,
		Username:      username,
		Warehouse:     warehouse,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSnowflakePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSnowflakeResourceModel) ToDeleteSDKType() *shared.DestinationSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationSnowflakeResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationSnowflakeResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
