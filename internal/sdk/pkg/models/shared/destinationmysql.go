// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Mysql string

const (
	MysqlMysql Mysql = "mysql"
)

func (e Mysql) ToPointer() *Mysql {
	return &e
}

func (e *Mysql) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mysql":
		*e = Mysql(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mysql: %v", v)
	}
}

// DestinationMysqlSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMysqlSchemasTunnelMethodTunnelMethod string

const (
	DestinationMysqlSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationMysqlSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMysqlSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationMysqlSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationMysqlSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMysqlSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationMysqlPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationMysqlSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationMysqlPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMysqlPasswordAuthentication) GetTunnelMethod() DestinationMysqlSchemasTunnelMethodTunnelMethod {
	return DestinationMysqlSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationMysqlPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMysqlPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationMysqlPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationMysqlSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMysqlSchemasTunnelMethod string

const (
	DestinationMysqlSchemasTunnelMethodSSHKeyAuth DestinationMysqlSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMysqlSchemasTunnelMethod) ToPointer() *DestinationMysqlSchemasTunnelMethod {
	return &e
}

func (e *DestinationMysqlSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMysqlSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlSchemasTunnelMethod: %v", v)
	}
}

// DestinationMysqlSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationMysqlSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationMysqlSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationMysqlSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMysqlSSHKeyAuthentication) GetTunnelMethod() DestinationMysqlSchemasTunnelMethod {
	return DestinationMysqlSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationMysqlSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMysqlSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationMysqlTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMysqlTunnelMethod string

const (
	DestinationMysqlTunnelMethodNoTunnel DestinationMysqlTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMysqlTunnelMethod) ToPointer() *DestinationMysqlTunnelMethod {
	return &e
}

func (e *DestinationMysqlTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMysqlTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMysqlTunnelMethod: %v", v)
	}
}

// DestinationMysqlNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMysqlNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationMysqlTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationMysqlNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysqlNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysqlNoTunnel) GetTunnelMethod() DestinationMysqlTunnelMethod {
	return DestinationMysqlTunnelMethodNoTunnel
}

type DestinationMysqlSSHTunnelMethodType string

const (
	DestinationMysqlSSHTunnelMethodTypeDestinationMysqlNoTunnel               DestinationMysqlSSHTunnelMethodType = "destination-mysql_No Tunnel"
	DestinationMysqlSSHTunnelMethodTypeDestinationMysqlSSHKeyAuthentication   DestinationMysqlSSHTunnelMethodType = "destination-mysql_SSH Key Authentication"
	DestinationMysqlSSHTunnelMethodTypeDestinationMysqlPasswordAuthentication DestinationMysqlSSHTunnelMethodType = "destination-mysql_Password Authentication"
)

type DestinationMysqlSSHTunnelMethod struct {
	DestinationMysqlNoTunnel               *DestinationMysqlNoTunnel
	DestinationMysqlSSHKeyAuthentication   *DestinationMysqlSSHKeyAuthentication
	DestinationMysqlPasswordAuthentication *DestinationMysqlPasswordAuthentication

	Type DestinationMysqlSSHTunnelMethodType
}

func CreateDestinationMysqlSSHTunnelMethodDestinationMysqlNoTunnel(destinationMysqlNoTunnel DestinationMysqlNoTunnel) DestinationMysqlSSHTunnelMethod {
	typ := DestinationMysqlSSHTunnelMethodTypeDestinationMysqlNoTunnel

	return DestinationMysqlSSHTunnelMethod{
		DestinationMysqlNoTunnel: &destinationMysqlNoTunnel,
		Type:                     typ,
	}
}

func CreateDestinationMysqlSSHTunnelMethodDestinationMysqlSSHKeyAuthentication(destinationMysqlSSHKeyAuthentication DestinationMysqlSSHKeyAuthentication) DestinationMysqlSSHTunnelMethod {
	typ := DestinationMysqlSSHTunnelMethodTypeDestinationMysqlSSHKeyAuthentication

	return DestinationMysqlSSHTunnelMethod{
		DestinationMysqlSSHKeyAuthentication: &destinationMysqlSSHKeyAuthentication,
		Type:                                 typ,
	}
}

func CreateDestinationMysqlSSHTunnelMethodDestinationMysqlPasswordAuthentication(destinationMysqlPasswordAuthentication DestinationMysqlPasswordAuthentication) DestinationMysqlSSHTunnelMethod {
	typ := DestinationMysqlSSHTunnelMethodTypeDestinationMysqlPasswordAuthentication

	return DestinationMysqlSSHTunnelMethod{
		DestinationMysqlPasswordAuthentication: &destinationMysqlPasswordAuthentication,
		Type:                                   typ,
	}
}

func (u *DestinationMysqlSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationMysqlNoTunnel := new(DestinationMysqlNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationMysqlNoTunnel, "", true, true); err == nil {
		u.DestinationMysqlNoTunnel = destinationMysqlNoTunnel
		u.Type = DestinationMysqlSSHTunnelMethodTypeDestinationMysqlNoTunnel
		return nil
	}

	destinationMysqlSSHKeyAuthentication := new(DestinationMysqlSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMysqlSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationMysqlSSHKeyAuthentication = destinationMysqlSSHKeyAuthentication
		u.Type = DestinationMysqlSSHTunnelMethodTypeDestinationMysqlSSHKeyAuthentication
		return nil
	}

	destinationMysqlPasswordAuthentication := new(DestinationMysqlPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMysqlPasswordAuthentication, "", true, true); err == nil {
		u.DestinationMysqlPasswordAuthentication = destinationMysqlPasswordAuthentication
		u.Type = DestinationMysqlSSHTunnelMethodTypeDestinationMysqlPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMysqlSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMysqlNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationMysqlNoTunnel, "", true)
	}

	if u.DestinationMysqlSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMysqlSSHKeyAuthentication, "", true)
	}

	if u.DestinationMysqlPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMysqlPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationMysql struct {
	// Name of the database.
	Database        string `json:"database"`
	destinationType Mysql  `const:"mysql" json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"3306" json:"port"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMysqlSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationMysql) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMysql) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMysql) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationMysql) GetDestinationType() Mysql {
	return MysqlMysql
}

func (o *DestinationMysql) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationMysql) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationMysql) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationMysql) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationMysql) GetTunnelMethod() *DestinationMysqlSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationMysql) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
