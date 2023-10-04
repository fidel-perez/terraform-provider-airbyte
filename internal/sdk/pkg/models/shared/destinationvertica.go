// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationVerticaVertica string

const (
	DestinationVerticaVerticaVertica DestinationVerticaVertica = "vertica"
)

func (e DestinationVerticaVertica) ToPointer() *DestinationVerticaVertica {
	return &e
}

func (e *DestinationVerticaVertica) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "vertica":
		*e = DestinationVerticaVertica(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaVertica: %v", v)
	}
}

// DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationVerticaSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationVerticaSSHTunnelMethodPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaSSHTunnelMethodPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaSSHTunnelMethodPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaSSHTunnelMethodPasswordAuthentication) GetTunnelMethod() DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return DestinationVerticaSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth
}

func (o *DestinationVerticaSSHTunnelMethodPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaSSHTunnelMethodPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationVerticaSSHTunnelMethodPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationVerticaSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) GetTunnelMethod() DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return DestinationVerticaSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth
}

func (o *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationVerticaSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationVerticaSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationVerticaSSHTunnelMethodNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVerticaSSHTunnelMethodNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVerticaSSHTunnelMethodNoTunnel) GetTunnelMethod() DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethod {
	return DestinationVerticaSSHTunnelMethodNoTunnelTunnelMethodNoTunnel
}

type DestinationVerticaSSHTunnelMethodType string

const (
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodNoTunnel               DestinationVerticaSSHTunnelMethodType = "destination-vertica_SSH Tunnel Method_No Tunnel"
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodSSHKeyAuthentication   DestinationVerticaSSHTunnelMethodType = "destination-vertica_SSH Tunnel Method_SSH Key Authentication"
	DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodPasswordAuthentication DestinationVerticaSSHTunnelMethodType = "destination-vertica_SSH Tunnel Method_Password Authentication"
)

type DestinationVerticaSSHTunnelMethod struct {
	DestinationVerticaSSHTunnelMethodNoTunnel               *DestinationVerticaSSHTunnelMethodNoTunnel
	DestinationVerticaSSHTunnelMethodSSHKeyAuthentication   *DestinationVerticaSSHTunnelMethodSSHKeyAuthentication
	DestinationVerticaSSHTunnelMethodPasswordAuthentication *DestinationVerticaSSHTunnelMethodPasswordAuthentication

	Type DestinationVerticaSSHTunnelMethodType
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaSSHTunnelMethodNoTunnel(destinationVerticaSSHTunnelMethodNoTunnel DestinationVerticaSSHTunnelMethodNoTunnel) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodNoTunnel

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaSSHTunnelMethodNoTunnel: &destinationVerticaSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaSSHTunnelMethodSSHKeyAuthentication(destinationVerticaSSHTunnelMethodSSHKeyAuthentication DestinationVerticaSSHTunnelMethodSSHKeyAuthentication) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodSSHKeyAuthentication

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaSSHTunnelMethodSSHKeyAuthentication: &destinationVerticaSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationVerticaSSHTunnelMethodDestinationVerticaSSHTunnelMethodPasswordAuthentication(destinationVerticaSSHTunnelMethodPasswordAuthentication DestinationVerticaSSHTunnelMethodPasswordAuthentication) DestinationVerticaSSHTunnelMethod {
	typ := DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodPasswordAuthentication

	return DestinationVerticaSSHTunnelMethod{
		DestinationVerticaSSHTunnelMethodPasswordAuthentication: &destinationVerticaSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationVerticaSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationVerticaSSHTunnelMethodNoTunnel := new(DestinationVerticaSSHTunnelMethodNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationVerticaSSHTunnelMethodNoTunnel, "", true, true); err == nil {
		u.DestinationVerticaSSHTunnelMethodNoTunnel = destinationVerticaSSHTunnelMethodNoTunnel
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodNoTunnel
		return nil
	}

	destinationVerticaSSHTunnelMethodSSHKeyAuthentication := new(DestinationVerticaSSHTunnelMethodSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationVerticaSSHTunnelMethodSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationVerticaSSHTunnelMethodSSHKeyAuthentication = destinationVerticaSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationVerticaSSHTunnelMethodPasswordAuthentication := new(DestinationVerticaSSHTunnelMethodPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationVerticaSSHTunnelMethodPasswordAuthentication, "", true, true); err == nil {
		u.DestinationVerticaSSHTunnelMethodPasswordAuthentication = destinationVerticaSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationVerticaSSHTunnelMethodTypeDestinationVerticaSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationVerticaSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationVerticaSSHTunnelMethodNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationVerticaSSHTunnelMethodNoTunnel, "", true)
	}

	if u.DestinationVerticaSSHTunnelMethodSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationVerticaSSHTunnelMethodSSHKeyAuthentication, "", true)
	}

	if u.DestinationVerticaSSHTunnelMethodPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationVerticaSSHTunnelMethodPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationVertica struct {
	// Name of the database.
	Database        string                    `json:"database"`
	destinationType DestinationVerticaVertica `const:"vertica" json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5433" json:"port"`
	// Schema for vertica destination
	Schema string `json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationVerticaSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationVertica) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationVertica) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationVertica) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationVertica) GetDestinationType() DestinationVerticaVertica {
	return DestinationVerticaVerticaVertica
}

func (o *DestinationVertica) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationVertica) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationVertica) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationVertica) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationVertica) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationVertica) GetTunnelMethod() *DestinationVerticaSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationVertica) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
