// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceOracleSchemasConnectionType string

const (
	SourceOracleSchemasConnectionTypeSid SourceOracleSchemasConnectionType = "sid"
)

func (e SourceOracleSchemasConnectionType) ToPointer() *SourceOracleSchemasConnectionType {
	return &e
}

func (e *SourceOracleSchemasConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sid":
		*e = SourceOracleSchemasConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleSchemasConnectionType: %v", v)
	}
}

// SourceOracleSystemIDSID - Use SID (Oracle System Identifier)
type SourceOracleSystemIDSID struct {
	connectionType *SourceOracleSchemasConnectionType `const:"sid" json:"connection_type"`
	Sid            string                             `json:"sid"`
}

func (s SourceOracleSystemIDSID) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleSystemIDSID) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleSystemIDSID) GetConnectionType() *SourceOracleSchemasConnectionType {
	return SourceOracleSchemasConnectionTypeSid.ToPointer()
}

func (o *SourceOracleSystemIDSID) GetSid() string {
	if o == nil {
		return ""
	}
	return o.Sid
}

type SourceOracleConnectionType string

const (
	SourceOracleConnectionTypeServiceName SourceOracleConnectionType = "service_name"
)

func (e SourceOracleConnectionType) ToPointer() *SourceOracleConnectionType {
	return &e
}

func (e *SourceOracleConnectionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_name":
		*e = SourceOracleConnectionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleConnectionType: %v", v)
	}
}

// SourceOracleServiceName - Use service name
type SourceOracleServiceName struct {
	connectionType *SourceOracleConnectionType `const:"service_name" json:"connection_type"`
	ServiceName    string                      `json:"service_name"`
}

func (s SourceOracleServiceName) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleServiceName) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleServiceName) GetConnectionType() *SourceOracleConnectionType {
	return SourceOracleConnectionTypeServiceName.ToPointer()
}

func (o *SourceOracleServiceName) GetServiceName() string {
	if o == nil {
		return ""
	}
	return o.ServiceName
}

type SourceOracleConnectByType string

const (
	SourceOracleConnectByTypeSourceOracleServiceName SourceOracleConnectByType = "source-oracle_Service name"
	SourceOracleConnectByTypeSourceOracleSystemIDSID SourceOracleConnectByType = "source-oracle_System ID (SID)"
)

type SourceOracleConnectBy struct {
	SourceOracleServiceName *SourceOracleServiceName
	SourceOracleSystemIDSID *SourceOracleSystemIDSID

	Type SourceOracleConnectByType
}

func CreateSourceOracleConnectBySourceOracleServiceName(sourceOracleServiceName SourceOracleServiceName) SourceOracleConnectBy {
	typ := SourceOracleConnectByTypeSourceOracleServiceName

	return SourceOracleConnectBy{
		SourceOracleServiceName: &sourceOracleServiceName,
		Type:                    typ,
	}
}

func CreateSourceOracleConnectBySourceOracleSystemIDSID(sourceOracleSystemIDSID SourceOracleSystemIDSID) SourceOracleConnectBy {
	typ := SourceOracleConnectByTypeSourceOracleSystemIDSID

	return SourceOracleConnectBy{
		SourceOracleSystemIDSID: &sourceOracleSystemIDSID,
		Type:                    typ,
	}
}

func (u *SourceOracleConnectBy) UnmarshalJSON(data []byte) error {

	sourceOracleServiceName := new(SourceOracleServiceName)
	if err := utils.UnmarshalJSON(data, &sourceOracleServiceName, "", true, true); err == nil {
		u.SourceOracleServiceName = sourceOracleServiceName
		u.Type = SourceOracleConnectByTypeSourceOracleServiceName
		return nil
	}

	sourceOracleSystemIDSID := new(SourceOracleSystemIDSID)
	if err := utils.UnmarshalJSON(data, &sourceOracleSystemIDSID, "", true, true); err == nil {
		u.SourceOracleSystemIDSID = sourceOracleSystemIDSID
		u.Type = SourceOracleConnectByTypeSourceOracleSystemIDSID
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOracleConnectBy) MarshalJSON() ([]byte, error) {
	if u.SourceOracleServiceName != nil {
		return utils.MarshalJSON(u.SourceOracleServiceName, "", true)
	}

	if u.SourceOracleSystemIDSID != nil {
		return utils.MarshalJSON(u.SourceOracleSystemIDSID, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceOracleSchemasEncryptionMethod string

const (
	SourceOracleSchemasEncryptionMethodEncryptedVerifyCertificate SourceOracleSchemasEncryptionMethod = "encrypted_verify_certificate"
)

func (e SourceOracleSchemasEncryptionMethod) ToPointer() *SourceOracleSchemasEncryptionMethod {
	return &e
}

func (e *SourceOracleSchemasEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "encrypted_verify_certificate":
		*e = SourceOracleSchemasEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleSchemasEncryptionMethod: %v", v)
	}
}

// SourceOracleTLSEncryptedVerifyCertificate - Verify and use the certificate provided by the server.
type SourceOracleTLSEncryptedVerifyCertificate struct {
	encryptionMethod *SourceOracleSchemasEncryptionMethod `const:"encrypted_verify_certificate" json:"encryption_method"`
	// Privacy Enhanced Mail (PEM) files are concatenated certificate containers frequently used in certificate installations.
	SslCertificate string `json:"ssl_certificate"`
}

func (s SourceOracleTLSEncryptedVerifyCertificate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleTLSEncryptedVerifyCertificate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleTLSEncryptedVerifyCertificate) GetEncryptionMethod() *SourceOracleSchemasEncryptionMethod {
	return SourceOracleSchemasEncryptionMethodEncryptedVerifyCertificate.ToPointer()
}

func (o *SourceOracleTLSEncryptedVerifyCertificate) GetSslCertificate() string {
	if o == nil {
		return ""
	}
	return o.SslCertificate
}

// SourceOracleEncryptionAlgorithm - This parameter defines what encryption algorithm is used.
type SourceOracleEncryptionAlgorithm string

const (
	SourceOracleEncryptionAlgorithmAes256      SourceOracleEncryptionAlgorithm = "AES256"
	SourceOracleEncryptionAlgorithmRc456       SourceOracleEncryptionAlgorithm = "RC4_56"
	SourceOracleEncryptionAlgorithmThreeDes168 SourceOracleEncryptionAlgorithm = "3DES168"
)

func (e SourceOracleEncryptionAlgorithm) ToPointer() *SourceOracleEncryptionAlgorithm {
	return &e
}

func (e *SourceOracleEncryptionAlgorithm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AES256":
		fallthrough
	case "RC4_56":
		fallthrough
	case "3DES168":
		*e = SourceOracleEncryptionAlgorithm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleEncryptionAlgorithm: %v", v)
	}
}

type SourceOracleEncryptionMethod string

const (
	SourceOracleEncryptionMethodClientNne SourceOracleEncryptionMethod = "client_nne"
)

func (e SourceOracleEncryptionMethod) ToPointer() *SourceOracleEncryptionMethod {
	return &e
}

func (e *SourceOracleEncryptionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_nne":
		*e = SourceOracleEncryptionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleEncryptionMethod: %v", v)
	}
}

// SourceOracleNativeNetworkEncryptionNNE - The native network encryption gives you the ability to encrypt database connections, without the configuration overhead of TCP/IP and SSL/TLS and without the need to open and listen on different ports.
type SourceOracleNativeNetworkEncryptionNNE struct {
	// This parameter defines what encryption algorithm is used.
	EncryptionAlgorithm *SourceOracleEncryptionAlgorithm `default:"AES256" json:"encryption_algorithm"`
	encryptionMethod    *SourceOracleEncryptionMethod    `const:"client_nne" json:"encryption_method"`
}

func (s SourceOracleNativeNetworkEncryptionNNE) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleNativeNetworkEncryptionNNE) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleNativeNetworkEncryptionNNE) GetEncryptionAlgorithm() *SourceOracleEncryptionAlgorithm {
	if o == nil {
		return nil
	}
	return o.EncryptionAlgorithm
}

func (o *SourceOracleNativeNetworkEncryptionNNE) GetEncryptionMethod() *SourceOracleEncryptionMethod {
	return SourceOracleEncryptionMethodClientNne.ToPointer()
}

type SourceOracleEncryptionType string

const (
	SourceOracleEncryptionTypeSourceOracleNativeNetworkEncryptionNNE    SourceOracleEncryptionType = "source-oracle_Native Network Encryption (NNE)"
	SourceOracleEncryptionTypeSourceOracleTLSEncryptedVerifyCertificate SourceOracleEncryptionType = "source-oracle_TLS Encrypted (verify certificate)"
)

type SourceOracleEncryption struct {
	SourceOracleNativeNetworkEncryptionNNE    *SourceOracleNativeNetworkEncryptionNNE
	SourceOracleTLSEncryptedVerifyCertificate *SourceOracleTLSEncryptedVerifyCertificate

	Type SourceOracleEncryptionType
}

func CreateSourceOracleEncryptionSourceOracleNativeNetworkEncryptionNNE(sourceOracleNativeNetworkEncryptionNNE SourceOracleNativeNetworkEncryptionNNE) SourceOracleEncryption {
	typ := SourceOracleEncryptionTypeSourceOracleNativeNetworkEncryptionNNE

	return SourceOracleEncryption{
		SourceOracleNativeNetworkEncryptionNNE: &sourceOracleNativeNetworkEncryptionNNE,
		Type:                                   typ,
	}
}

func CreateSourceOracleEncryptionSourceOracleTLSEncryptedVerifyCertificate(sourceOracleTLSEncryptedVerifyCertificate SourceOracleTLSEncryptedVerifyCertificate) SourceOracleEncryption {
	typ := SourceOracleEncryptionTypeSourceOracleTLSEncryptedVerifyCertificate

	return SourceOracleEncryption{
		SourceOracleTLSEncryptedVerifyCertificate: &sourceOracleTLSEncryptedVerifyCertificate,
		Type: typ,
	}
}

func (u *SourceOracleEncryption) UnmarshalJSON(data []byte) error {

	sourceOracleNativeNetworkEncryptionNNE := new(SourceOracleNativeNetworkEncryptionNNE)
	if err := utils.UnmarshalJSON(data, &sourceOracleNativeNetworkEncryptionNNE, "", true, true); err == nil {
		u.SourceOracleNativeNetworkEncryptionNNE = sourceOracleNativeNetworkEncryptionNNE
		u.Type = SourceOracleEncryptionTypeSourceOracleNativeNetworkEncryptionNNE
		return nil
	}

	sourceOracleTLSEncryptedVerifyCertificate := new(SourceOracleTLSEncryptedVerifyCertificate)
	if err := utils.UnmarshalJSON(data, &sourceOracleTLSEncryptedVerifyCertificate, "", true, true); err == nil {
		u.SourceOracleTLSEncryptedVerifyCertificate = sourceOracleTLSEncryptedVerifyCertificate
		u.Type = SourceOracleEncryptionTypeSourceOracleTLSEncryptedVerifyCertificate
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOracleEncryption) MarshalJSON() ([]byte, error) {
	if u.SourceOracleNativeNetworkEncryptionNNE != nil {
		return utils.MarshalJSON(u.SourceOracleNativeNetworkEncryptionNNE, "", true)
	}

	if u.SourceOracleTLSEncryptedVerifyCertificate != nil {
		return utils.MarshalJSON(u.SourceOracleTLSEncryptedVerifyCertificate, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceOracleOracle string

const (
	SourceOracleOracleOracle SourceOracleOracle = "oracle"
)

func (e SourceOracleOracle) ToPointer() *SourceOracleOracle {
	return &e
}

func (e *SourceOracleOracle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oracle":
		*e = SourceOracleOracle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleOracle: %v", v)
	}
}

// SourceOracleSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceOracleSchemasTunnelMethodTunnelMethod string

const (
	SourceOracleSchemasTunnelMethodTunnelMethodSSHPasswordAuth SourceOracleSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceOracleSchemasTunnelMethodTunnelMethod) ToPointer() *SourceOracleSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *SourceOracleSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceOracleSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// SourceOraclePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceOraclePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod SourceOracleSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (s SourceOraclePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOraclePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOraclePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceOraclePasswordAuthentication) GetTunnelMethod() SourceOracleSchemasTunnelMethodTunnelMethod {
	return SourceOracleSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *SourceOraclePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceOraclePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *SourceOraclePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// SourceOracleSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceOracleSchemasTunnelMethod string

const (
	SourceOracleSchemasTunnelMethodSSHKeyAuth SourceOracleSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceOracleSchemasTunnelMethod) ToPointer() *SourceOracleSchemasTunnelMethod {
	return &e
}

func (e *SourceOracleSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceOracleSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleSchemasTunnelMethod: %v", v)
	}
}

// SourceOracleSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceOracleSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod SourceOracleSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (s SourceOracleSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *SourceOracleSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceOracleSSHKeyAuthentication) GetTunnelMethod() SourceOracleSchemasTunnelMethod {
	return SourceOracleSchemasTunnelMethodSSHKeyAuth
}

func (o *SourceOracleSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceOracleSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// SourceOracleTunnelMethod - No ssh tunnel needed to connect to database
type SourceOracleTunnelMethod string

const (
	SourceOracleTunnelMethodNoTunnel SourceOracleTunnelMethod = "NO_TUNNEL"
)

func (e SourceOracleTunnelMethod) ToPointer() *SourceOracleTunnelMethod {
	return &e
}

func (e *SourceOracleTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceOracleTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOracleTunnelMethod: %v", v)
	}
}

// SourceOracleNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceOracleNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod SourceOracleTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (s SourceOracleNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracleNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracleNoTunnel) GetTunnelMethod() SourceOracleTunnelMethod {
	return SourceOracleTunnelMethodNoTunnel
}

type SourceOracleSSHTunnelMethodType string

const (
	SourceOracleSSHTunnelMethodTypeSourceOracleNoTunnel               SourceOracleSSHTunnelMethodType = "source-oracle_No Tunnel"
	SourceOracleSSHTunnelMethodTypeSourceOracleSSHKeyAuthentication   SourceOracleSSHTunnelMethodType = "source-oracle_SSH Key Authentication"
	SourceOracleSSHTunnelMethodTypeSourceOraclePasswordAuthentication SourceOracleSSHTunnelMethodType = "source-oracle_Password Authentication"
)

type SourceOracleSSHTunnelMethod struct {
	SourceOracleNoTunnel               *SourceOracleNoTunnel
	SourceOracleSSHKeyAuthentication   *SourceOracleSSHKeyAuthentication
	SourceOraclePasswordAuthentication *SourceOraclePasswordAuthentication

	Type SourceOracleSSHTunnelMethodType
}

func CreateSourceOracleSSHTunnelMethodSourceOracleNoTunnel(sourceOracleNoTunnel SourceOracleNoTunnel) SourceOracleSSHTunnelMethod {
	typ := SourceOracleSSHTunnelMethodTypeSourceOracleNoTunnel

	return SourceOracleSSHTunnelMethod{
		SourceOracleNoTunnel: &sourceOracleNoTunnel,
		Type:                 typ,
	}
}

func CreateSourceOracleSSHTunnelMethodSourceOracleSSHKeyAuthentication(sourceOracleSSHKeyAuthentication SourceOracleSSHKeyAuthentication) SourceOracleSSHTunnelMethod {
	typ := SourceOracleSSHTunnelMethodTypeSourceOracleSSHKeyAuthentication

	return SourceOracleSSHTunnelMethod{
		SourceOracleSSHKeyAuthentication: &sourceOracleSSHKeyAuthentication,
		Type:                             typ,
	}
}

func CreateSourceOracleSSHTunnelMethodSourceOraclePasswordAuthentication(sourceOraclePasswordAuthentication SourceOraclePasswordAuthentication) SourceOracleSSHTunnelMethod {
	typ := SourceOracleSSHTunnelMethodTypeSourceOraclePasswordAuthentication

	return SourceOracleSSHTunnelMethod{
		SourceOraclePasswordAuthentication: &sourceOraclePasswordAuthentication,
		Type:                               typ,
	}
}

func (u *SourceOracleSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	sourceOracleNoTunnel := new(SourceOracleNoTunnel)
	if err := utils.UnmarshalJSON(data, &sourceOracleNoTunnel, "", true, true); err == nil {
		u.SourceOracleNoTunnel = sourceOracleNoTunnel
		u.Type = SourceOracleSSHTunnelMethodTypeSourceOracleNoTunnel
		return nil
	}

	sourceOracleSSHKeyAuthentication := new(SourceOracleSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceOracleSSHKeyAuthentication, "", true, true); err == nil {
		u.SourceOracleSSHKeyAuthentication = sourceOracleSSHKeyAuthentication
		u.Type = SourceOracleSSHTunnelMethodTypeSourceOracleSSHKeyAuthentication
		return nil
	}

	sourceOraclePasswordAuthentication := new(SourceOraclePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceOraclePasswordAuthentication, "", true, true); err == nil {
		u.SourceOraclePasswordAuthentication = sourceOraclePasswordAuthentication
		u.Type = SourceOracleSSHTunnelMethodTypeSourceOraclePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOracleSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceOracleNoTunnel != nil {
		return utils.MarshalJSON(u.SourceOracleNoTunnel, "", true)
	}

	if u.SourceOracleSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceOracleSSHKeyAuthentication, "", true)
	}

	if u.SourceOraclePasswordAuthentication != nil {
		return utils.MarshalJSON(u.SourceOraclePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceOracle struct {
	// Connect data that will be used for DB connection
	ConnectionData *SourceOracleConnectBy `json:"connection_data,omitempty"`
	// The encryption method with is used when communicating with the database.
	Encryption SourceOracleEncryption `json:"encryption"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	// Oracle Corporations recommends the following port numbers:
	// 1521 - Default listening port for client connections to the listener.
	// 2484 - Recommended and officially registered listening port for client connections to the listener using TCP/IP with SSL
	Port *int64 `default:"1521" json:"port"`
	// The list of schemas to sync from. Defaults to user. Case sensitive.
	Schemas    []string           `json:"schemas,omitempty"`
	sourceType SourceOracleOracle `const:"oracle" json:"sourceType"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceOracleSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (s SourceOracle) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOracle) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOracle) GetConnectionData() *SourceOracleConnectBy {
	if o == nil {
		return nil
	}
	return o.ConnectionData
}

func (o *SourceOracle) GetEncryption() SourceOracleEncryption {
	if o == nil {
		return SourceOracleEncryption{}
	}
	return o.Encryption
}

func (o *SourceOracle) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceOracle) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceOracle) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceOracle) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceOracle) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *SourceOracle) GetSourceType() SourceOracleOracle {
	return SourceOracleOracleOracle
}

func (o *SourceOracle) GetTunnelMethod() *SourceOracleSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *SourceOracle) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
