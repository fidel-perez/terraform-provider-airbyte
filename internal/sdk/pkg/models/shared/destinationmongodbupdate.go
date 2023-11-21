// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationMongodbUpdateAuthorization string

const (
	DestinationMongodbUpdateAuthorizationLoginPassword DestinationMongodbUpdateAuthorization = "login/password"
)

func (e DestinationMongodbUpdateAuthorization) ToPointer() *DestinationMongodbUpdateAuthorization {
	return &e
}

func (e *DestinationMongodbUpdateAuthorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "login/password":
		*e = DestinationMongodbUpdateAuthorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateAuthorization: %v", v)
	}
}

// LoginPassword - Login/Password.
type LoginPassword struct {
	authorization DestinationMongodbUpdateAuthorization `const:"login/password" json:"authorization"`
	// Password associated with the username.
	Password string `json:"password"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (l LoginPassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LoginPassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LoginPassword) GetAuthorization() DestinationMongodbUpdateAuthorization {
	return DestinationMongodbUpdateAuthorizationLoginPassword
}

func (o *LoginPassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *LoginPassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type Authorization string

const (
	AuthorizationNone Authorization = "none"
)

func (e Authorization) ToPointer() *Authorization {
	return &e
}

func (e *Authorization) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = Authorization(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Authorization: %v", v)
	}
}

// None - None.
type None struct {
	authorization Authorization `const:"none" json:"authorization"`
}

func (n None) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *None) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *None) GetAuthorization() Authorization {
	return AuthorizationNone
}

type AuthorizationTypeType string

const (
	AuthorizationTypeTypeNone          AuthorizationTypeType = "None"
	AuthorizationTypeTypeLoginPassword AuthorizationTypeType = "Login/Password"
)

type AuthorizationType struct {
	None          *None
	LoginPassword *LoginPassword

	Type AuthorizationTypeType
}

func CreateAuthorizationTypeNone(none None) AuthorizationType {
	typ := AuthorizationTypeTypeNone

	return AuthorizationType{
		None: &none,
		Type: typ,
	}
}

func CreateAuthorizationTypeLoginPassword(loginPassword LoginPassword) AuthorizationType {
	typ := AuthorizationTypeTypeLoginPassword

	return AuthorizationType{
		LoginPassword: &loginPassword,
		Type:          typ,
	}
}

func (u *AuthorizationType) UnmarshalJSON(data []byte) error {

	none := new(None)
	if err := utils.UnmarshalJSON(data, &none, "", true, true); err == nil {
		u.None = none
		u.Type = AuthorizationTypeTypeNone
		return nil
	}

	loginPassword := new(LoginPassword)
	if err := utils.UnmarshalJSON(data, &loginPassword, "", true, true); err == nil {
		u.LoginPassword = loginPassword
		u.Type = AuthorizationTypeTypeLoginPassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthorizationType) MarshalJSON() ([]byte, error) {
	if u.None != nil {
		return utils.MarshalJSON(u.None, "", true)
	}

	if u.LoginPassword != nil {
		return utils.MarshalJSON(u.LoginPassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationMongodbUpdateSchemasInstance string

const (
	DestinationMongodbUpdateSchemasInstanceAtlas DestinationMongodbUpdateSchemasInstance = "atlas"
)

func (e DestinationMongodbUpdateSchemasInstance) ToPointer() *DestinationMongodbUpdateSchemasInstance {
	return &e
}

func (e *DestinationMongodbUpdateSchemasInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "atlas":
		*e = DestinationMongodbUpdateSchemasInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSchemasInstance: %v", v)
	}
}

// MongoDBAtlas - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type MongoDBAtlas struct {
	// URL of a cluster to connect to.
	ClusterURL string                                   `json:"cluster_url"`
	Instance   *DestinationMongodbUpdateSchemasInstance `default:"atlas" json:"instance"`
}

func (m MongoDBAtlas) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MongoDBAtlas) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *MongoDBAtlas) GetClusterURL() string {
	if o == nil {
		return ""
	}
	return o.ClusterURL
}

func (o *MongoDBAtlas) GetInstance() *DestinationMongodbUpdateSchemasInstance {
	if o == nil {
		return nil
	}
	return o.Instance
}

type DestinationMongodbUpdateInstance string

const (
	DestinationMongodbUpdateInstanceReplica DestinationMongodbUpdateInstance = "replica"
)

func (e DestinationMongodbUpdateInstance) ToPointer() *DestinationMongodbUpdateInstance {
	return &e
}

func (e *DestinationMongodbUpdateInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "replica":
		*e = DestinationMongodbUpdateInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateInstance: %v", v)
	}
}

// ReplicaSet - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type ReplicaSet struct {
	Instance *DestinationMongodbUpdateInstance `default:"replica" json:"instance"`
	// A replica set name.
	ReplicaSet *string `json:"replica_set,omitempty"`
	// The members of a replica set. Please specify `host`:`port` of each member seperated by comma.
	ServerAddresses string `json:"server_addresses"`
}

func (r ReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ReplicaSet) GetInstance() *DestinationMongodbUpdateInstance {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ReplicaSet) GetReplicaSet() *string {
	if o == nil {
		return nil
	}
	return o.ReplicaSet
}

func (o *ReplicaSet) GetServerAddresses() string {
	if o == nil {
		return ""
	}
	return o.ServerAddresses
}

type Instance string

const (
	InstanceStandalone Instance = "standalone"
)

func (e Instance) ToPointer() *Instance {
	return &e
}

func (e *Instance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standalone":
		*e = Instance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Instance: %v", v)
	}
}

// StandaloneMongoDbInstance - MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type StandaloneMongoDbInstance struct {
	// The Host of a Mongo database to be replicated.
	Host     string    `json:"host"`
	Instance *Instance `default:"standalone" json:"instance"`
	// The Port of a Mongo database to be replicated.
	Port *int64 `default:"27017" json:"port"`
}

func (s StandaloneMongoDbInstance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *StandaloneMongoDbInstance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *StandaloneMongoDbInstance) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *StandaloneMongoDbInstance) GetInstance() *Instance {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *StandaloneMongoDbInstance) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type MongoDbInstanceTypeType string

const (
	MongoDbInstanceTypeTypeStandaloneMongoDbInstance MongoDbInstanceTypeType = "Standalone MongoDb Instance"
	MongoDbInstanceTypeTypeReplicaSet                MongoDbInstanceTypeType = "Replica Set"
	MongoDbInstanceTypeTypeMongoDBAtlas              MongoDbInstanceTypeType = "MongoDB Atlas"
)

type MongoDbInstanceType struct {
	StandaloneMongoDbInstance *StandaloneMongoDbInstance
	ReplicaSet                *ReplicaSet
	MongoDBAtlas              *MongoDBAtlas

	Type MongoDbInstanceTypeType
}

func CreateMongoDbInstanceTypeStandaloneMongoDbInstance(standaloneMongoDbInstance StandaloneMongoDbInstance) MongoDbInstanceType {
	typ := MongoDbInstanceTypeTypeStandaloneMongoDbInstance

	return MongoDbInstanceType{
		StandaloneMongoDbInstance: &standaloneMongoDbInstance,
		Type:                      typ,
	}
}

func CreateMongoDbInstanceTypeReplicaSet(replicaSet ReplicaSet) MongoDbInstanceType {
	typ := MongoDbInstanceTypeTypeReplicaSet

	return MongoDbInstanceType{
		ReplicaSet: &replicaSet,
		Type:       typ,
	}
}

func CreateMongoDbInstanceTypeMongoDBAtlas(mongoDBAtlas MongoDBAtlas) MongoDbInstanceType {
	typ := MongoDbInstanceTypeTypeMongoDBAtlas

	return MongoDbInstanceType{
		MongoDBAtlas: &mongoDBAtlas,
		Type:         typ,
	}
}

func (u *MongoDbInstanceType) UnmarshalJSON(data []byte) error {

	mongoDBAtlas := new(MongoDBAtlas)
	if err := utils.UnmarshalJSON(data, &mongoDBAtlas, "", true, true); err == nil {
		u.MongoDBAtlas = mongoDBAtlas
		u.Type = MongoDbInstanceTypeTypeMongoDBAtlas
		return nil
	}

	standaloneMongoDbInstance := new(StandaloneMongoDbInstance)
	if err := utils.UnmarshalJSON(data, &standaloneMongoDbInstance, "", true, true); err == nil {
		u.StandaloneMongoDbInstance = standaloneMongoDbInstance
		u.Type = MongoDbInstanceTypeTypeStandaloneMongoDbInstance
		return nil
	}

	replicaSet := new(ReplicaSet)
	if err := utils.UnmarshalJSON(data, &replicaSet, "", true, true); err == nil {
		u.ReplicaSet = replicaSet
		u.Type = MongoDbInstanceTypeTypeReplicaSet
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u MongoDbInstanceType) MarshalJSON() ([]byte, error) {
	if u.StandaloneMongoDbInstance != nil {
		return utils.MarshalJSON(u.StandaloneMongoDbInstance, "", true)
	}

	if u.ReplicaSet != nil {
		return utils.MarshalJSON(u.ReplicaSet, "", true)
	}

	if u.MongoDBAtlas != nil {
		return utils.MarshalJSON(u.MongoDBAtlas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationMongodbUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdatePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationMongodbUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMongodbUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMongodbUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMongodbUpdatePasswordAuthentication) GetTunnelMethod() DestinationMongodbUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationMongodbUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationMongodbUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMongodbUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationMongodbUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationMongodbUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationMongodbUpdateSchemasTunnelMethod string

const (
	DestinationMongodbUpdateSchemasTunnelMethodSSHKeyAuth DestinationMongodbUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationMongodbUpdateSchemasTunnelMethod) ToPointer() *DestinationMongodbUpdateSchemasTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationMongodbUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateSchemasTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdateSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationMongodbUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationMongodbUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMongodbUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMongodbUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationMongodbUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationMongodbUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationMongodbUpdateSchemasTunnelMethod {
	return DestinationMongodbUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationMongodbUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationMongodbUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationMongodbUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationMongodbUpdateTunnelMethod string

const (
	DestinationMongodbUpdateTunnelMethodNoTunnel DestinationMongodbUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationMongodbUpdateTunnelMethod) ToPointer() *DestinationMongodbUpdateTunnelMethod {
	return &e
}

func (e *DestinationMongodbUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationMongodbUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationMongodbUpdateTunnelMethod: %v", v)
	}
}

// DestinationMongodbUpdateNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationMongodbUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationMongodbUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationMongodbUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationMongodbUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationMongodbUpdateNoTunnel) GetTunnelMethod() DestinationMongodbUpdateTunnelMethod {
	return DestinationMongodbUpdateTunnelMethodNoTunnel
}

type DestinationMongodbUpdateSSHTunnelMethodType string

const (
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateNoTunnel               DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_No Tunnel"
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHKeyAuthentication   DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_SSH Key Authentication"
	DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdatePasswordAuthentication DestinationMongodbUpdateSSHTunnelMethodType = "destination-mongodb-update_Password Authentication"
)

type DestinationMongodbUpdateSSHTunnelMethod struct {
	DestinationMongodbUpdateNoTunnel               *DestinationMongodbUpdateNoTunnel
	DestinationMongodbUpdateSSHKeyAuthentication   *DestinationMongodbUpdateSSHKeyAuthentication
	DestinationMongodbUpdatePasswordAuthentication *DestinationMongodbUpdatePasswordAuthentication

	Type DestinationMongodbUpdateSSHTunnelMethodType
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdateNoTunnel(destinationMongodbUpdateNoTunnel DestinationMongodbUpdateNoTunnel) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateNoTunnel

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdateNoTunnel: &destinationMongodbUpdateNoTunnel,
		Type:                             typ,
	}
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdateSSHKeyAuthentication(destinationMongodbUpdateSSHKeyAuthentication DestinationMongodbUpdateSSHKeyAuthentication) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHKeyAuthentication

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdateSSHKeyAuthentication: &destinationMongodbUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationMongodbUpdateSSHTunnelMethodDestinationMongodbUpdatePasswordAuthentication(destinationMongodbUpdatePasswordAuthentication DestinationMongodbUpdatePasswordAuthentication) DestinationMongodbUpdateSSHTunnelMethod {
	typ := DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdatePasswordAuthentication

	return DestinationMongodbUpdateSSHTunnelMethod{
		DestinationMongodbUpdatePasswordAuthentication: &destinationMongodbUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationMongodbUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationMongodbUpdateNoTunnel := new(DestinationMongodbUpdateNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationMongodbUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationMongodbUpdateNoTunnel = destinationMongodbUpdateNoTunnel
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateNoTunnel
		return nil
	}

	destinationMongodbUpdateSSHKeyAuthentication := new(DestinationMongodbUpdateSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMongodbUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationMongodbUpdateSSHKeyAuthentication = destinationMongodbUpdateSSHKeyAuthentication
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdateSSHKeyAuthentication
		return nil
	}

	destinationMongodbUpdatePasswordAuthentication := new(DestinationMongodbUpdatePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationMongodbUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationMongodbUpdatePasswordAuthentication = destinationMongodbUpdatePasswordAuthentication
		u.Type = DestinationMongodbUpdateSSHTunnelMethodTypeDestinationMongodbUpdatePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationMongodbUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationMongodbUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationMongodbUpdateNoTunnel, "", true)
	}

	if u.DestinationMongodbUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMongodbUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationMongodbUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationMongodbUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationMongodbUpdate struct {
	// Authorization type.
	AuthType AuthorizationType `json:"auth_type"`
	// Name of the database.
	Database string `json:"database"`
	// MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
	InstanceType *MongoDbInstanceType `json:"instance_type,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationMongodbUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
}

func (o *DestinationMongodbUpdate) GetAuthType() AuthorizationType {
	if o == nil {
		return AuthorizationType{}
	}
	return o.AuthType
}

func (o *DestinationMongodbUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationMongodbUpdate) GetInstanceType() *MongoDbInstanceType {
	if o == nil {
		return nil
	}
	return o.InstanceType
}

func (o *DestinationMongodbUpdate) GetTunnelMethod() *DestinationMongodbUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}
