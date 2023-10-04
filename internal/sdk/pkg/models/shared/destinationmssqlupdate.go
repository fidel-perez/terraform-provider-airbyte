// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod string

const (
	DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethodEncryptedVerifyCertificate DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod = "encrypted_verify_certificate"
)

func (e DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod) ToPointer() *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod {
	return &e
}

func (e *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod: %v", v)
	}
}

// DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate struct {
	// Specifies the host name of the server. The value of this property must match the subject property of the certificate.
	HostNameInCertificate *string                                                             `json:"hostNameInCertificate,omitempty"`
	sslMethod             *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod `const:"encrypted_verify_certificate" json:"ssl_method"`
}

func (d DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate) GetHostNameInCertificate() *string {
	if o == nil {
		return nil
	}
	return o.HostNameInCertificate
}

func (o *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate) GetSslMethod() *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethod {
	return DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificateSSLMethodEncryptedVerifyCertificate.ToPointer()
}

type DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod string

const (
	DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethodEncryptedTrustServerCertificate DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod = "encrypted_trust_server_certificate"
)

func (e DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod) ToPointer() *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod {
	return &e
}

func (e *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_trust_server_certificate":
		*e = DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod: %v", v)
	}
}

// DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate - Use the certificate provided by the server without verification. (For testing purposes only!)
type DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate struct {
	sslMethod *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod `const:"encrypted_trust_server_certificate" json:"ssl_method"`
}

func (d DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate) GetSslMethod() *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethod {
	return DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificateSSLMethodEncryptedTrustServerCertificate.ToPointer()
}

type DestinationMssqlUpdateSSLMethodType string

const (
	DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate DestinationMssqlUpdateSSLMethodType = "destination-mssql-update_SSL Method_Encrypted (trust server certificate)"
	DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate      DestinationMssqlUpdateSSLMethodType = "destination-mssql-update_SSL Method_Encrypted (verify certificate)"
)

type DestinationMssqlUpdateSSLMethod struct {
	DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate *DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate
	DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate      *DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate

	Type DestinationMssqlUpdateSSLMethodType
}

func CreateDestinationMssqlUpdateSSLMethodDestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate(destinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate) DestinationMssqlUpdateSSLMethod {
	typ := DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate

	return DestinationMssqlUpdateSSLMethod{
		DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate: &destinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate,
		Type: typ,
	}
}

func CreateDestinationMssqlUpdateSSLMethodDestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate(destinationMssqlUpdateSSLMethodEncryptedVerifyCertificate DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate) DestinationMssqlUpdateSSLMethod {
	typ := DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate

	return DestinationMssqlUpdateSSLMethod{
		DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate: &destinationMssqlUpdateSSLMethodEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *DestinationMssqlUpdateSSLMethod) UnmarshalJSON(data []byte) error {

	destinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate := new(DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate)
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate = destinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate
		u.Type = DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate
		return nil
	}

	destinationMssqlUpdateSSLMethodEncryptedVerifyCertificate := new(DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate)
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSLMethodEncryptedVerifyCertificate, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate = destinationMssqlUpdateSSLMethodEncryptedVerifyCertificate
		u.Type = DestinationMssqlUpdateSSLMethodTypeDestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMssqlUpdateSSLMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSLMethodEncryptedTrustServerCertificate, "", true)
	}

	if u.DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSLMethodEncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelMethod() DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return DestinationMssqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth
}

func (o *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelMethod() DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth
}

func (o *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationMssqlUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMssqlUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationMssqlUpdateSSHTunnelMethodNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdateSSHTunnelMethodNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdateSSHTunnelMethodNoTunnel) GetTunnelMethod() DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return DestinationMssqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel
}

type DestinationMssqlUpdateSSHTunnelMethodType string

const (
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodNoTunnel               DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_SSH Tunnel Method_No Tunnel"
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication   DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_SSH Tunnel Method_SSH Key Authentication"
	DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication DestinationMssqlUpdateSSHTunnelMethodType = "destination-mssql-update_SSH Tunnel Method_Password Authentication"
)

type DestinationMssqlUpdateSSHTunnelMethod struct {
	DestinationMssqlUpdateSSHTunnelMethodNoTunnel               *DestinationMssqlUpdateSSHTunnelMethodNoTunnel
	DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication   *DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
	DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication *DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication

	Type DestinationMssqlUpdateSSHTunnelMethodType
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdateSSHTunnelMethodNoTunnel(destinationMssqlUpdateSSHTunnelMethodNoTunnel DestinationMssqlUpdateSSHTunnelMethodNoTunnel) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodNoTunnel

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdateSSHTunnelMethodNoTunnel: &destinationMssqlUpdateSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication(destinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication: &destinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMssqlUpdateSSHTunnelMethodDestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication(destinationMssqlUpdateSSHTunnelMethodPasswordAuthentication DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication) DestinationMssqlUpdateSSHTunnelMethod {
	typ := DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication

	return DestinationMssqlUpdateSSHTunnelMethod{
		DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication: &destinationMssqlUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMssqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationMssqlUpdateSSHTunnelMethodNoTunnel := new(DestinationMssqlUpdateSSHTunnelMethodNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSHTunnelMethodNoTunnel, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSHTunnelMethodNoTunnel = destinationMssqlUpdateSSHTunnelMethodNoTunnel
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	destinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication := new(DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication = destinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationMssqlUpdateSSHTunnelMethodPasswordAuthentication := new(DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMssqlUpdateSSHTunnelMethodPasswordAuthentication, "", true, true); err == nil {
		u.DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication = destinationMssqlUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationMssqlUpdateSSHTunnelMethodTypeDestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMssqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMssqlUpdateSSHTunnelMethodNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSHTunnelMethodNoTunnel, "", true)
	}

	if u.DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSHTunnelMethodSSHKeyAuthentication, "", true)
	}

	if u.DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMssqlUpdateSSHTunnelMethodPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationMssqlUpdate struct {
	// The name of the MSSQL database.
	Database string `json:"database"`
	// The host name of the MSSQL database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with this username.
	Password *string `json:"password,omitempty"`
	// The port of the MSSQL database.
	Port *int64 `default:"1433" json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// The encryption method which is used to communicate with the database.
	SslMethod *DestinationMssqlUpdateSSLMethod `json:"ssl_method,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMssqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (d DestinationMssqlUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMssqlUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMssqlUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationMssqlUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationMssqlUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationMssqlUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationMssqlUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationMssqlUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationMssqlUpdate) GetSslMethod() *DestinationMssqlUpdateSSLMethod {
	if o == nil {
		return nil
	}
	return o.SslMethod
}

func (o *DestinationMssqlUpdate) GetTunnelMethod() *DestinationMssqlUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationMssqlUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
