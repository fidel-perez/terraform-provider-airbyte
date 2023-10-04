// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationPostgresPostgres string

const (
	DestinationPostgresPostgresPostgres DestinationPostgresPostgres = "postgres"
)

func (e DestinationPostgresPostgres) ToPointer() *DestinationPostgresPostgres {
	return &e
}

func (e *DestinationPostgresPostgres) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postgres":
		*e = DestinationPostgresPostgres(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresPostgres: %v", v)
	}
}

type DestinationPostgresSSLModesVerifyFullMode string

const (
	DestinationPostgresSSLModesVerifyFullModeVerifyFull DestinationPostgresSSLModesVerifyFullMode = "verify-full"
)

func (e DestinationPostgresSSLModesVerifyFullMode) ToPointer() *DestinationPostgresSSLModesVerifyFullMode {
	return &e
}

func (e *DestinationPostgresSSLModesVerifyFullMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-full":
		*e = DestinationPostgresSSLModesVerifyFullMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyFullMode: %v", v)
	}
}

// DestinationPostgresSSLModesVerifyFull - Verify-full SSL mode.
type DestinationPostgresSSLModesVerifyFull struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate
	ClientCertificate string `json:"client_certificate"`
	// Client key
	ClientKey string `json:"client_key"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                    `json:"client_key_password,omitempty"`
	mode              *DestinationPostgresSSLModesVerifyFullMode `const:"verify-full" json:"mode"`
}

func (d DestinationPostgresSSLModesVerifyFull) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesVerifyFull) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesVerifyFull) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *DestinationPostgresSSLModesVerifyFull) GetClientCertificate() string {
	if o == nil {
		return ""
	}
	return o.ClientCertificate
}

func (o *DestinationPostgresSSLModesVerifyFull) GetClientKey() string {
	if o == nil {
		return ""
	}
	return o.ClientKey
}

func (o *DestinationPostgresSSLModesVerifyFull) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *DestinationPostgresSSLModesVerifyFull) GetMode() *DestinationPostgresSSLModesVerifyFullMode {
	return DestinationPostgresSSLModesVerifyFullModeVerifyFull.ToPointer()
}

type DestinationPostgresSSLModesVerifyCaMode string

const (
	DestinationPostgresSSLModesVerifyCaModeVerifyCa DestinationPostgresSSLModesVerifyCaMode = "verify-ca"
)

func (e DestinationPostgresSSLModesVerifyCaMode) ToPointer() *DestinationPostgresSSLModesVerifyCaMode {
	return &e
}

func (e *DestinationPostgresSSLModesVerifyCaMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify-ca":
		*e = DestinationPostgresSSLModesVerifyCaMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesVerifyCaMode: %v", v)
	}
}

// DestinationPostgresSSLModesVerifyCa - Verify-ca SSL mode.
type DestinationPostgresSSLModesVerifyCa struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                  `json:"client_key_password,omitempty"`
	mode              *DestinationPostgresSSLModesVerifyCaMode `const:"verify-ca" json:"mode"`
}

func (d DestinationPostgresSSLModesVerifyCa) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesVerifyCa) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesVerifyCa) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *DestinationPostgresSSLModesVerifyCa) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *DestinationPostgresSSLModesVerifyCa) GetMode() *DestinationPostgresSSLModesVerifyCaMode {
	return DestinationPostgresSSLModesVerifyCaModeVerifyCa.ToPointer()
}

type DestinationPostgresSSLModesRequireMode string

const (
	DestinationPostgresSSLModesRequireModeRequire DestinationPostgresSSLModesRequireMode = "require"
)

func (e DestinationPostgresSSLModesRequireMode) ToPointer() *DestinationPostgresSSLModesRequireMode {
	return &e
}

func (e *DestinationPostgresSSLModesRequireMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "require":
		*e = DestinationPostgresSSLModesRequireMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesRequireMode: %v", v)
	}
}

// DestinationPostgresSSLModesRequire - Require SSL mode.
type DestinationPostgresSSLModesRequire struct {
	mode *DestinationPostgresSSLModesRequireMode `const:"require" json:"mode"`
}

func (d DestinationPostgresSSLModesRequire) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesRequire) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesRequire) GetMode() *DestinationPostgresSSLModesRequireMode {
	return DestinationPostgresSSLModesRequireModeRequire.ToPointer()
}

type DestinationPostgresSSLModesPreferMode string

const (
	DestinationPostgresSSLModesPreferModePrefer DestinationPostgresSSLModesPreferMode = "prefer"
)

func (e DestinationPostgresSSLModesPreferMode) ToPointer() *DestinationPostgresSSLModesPreferMode {
	return &e
}

func (e *DestinationPostgresSSLModesPreferMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "prefer":
		*e = DestinationPostgresSSLModesPreferMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesPreferMode: %v", v)
	}
}

// DestinationPostgresSSLModesPrefer - Prefer SSL mode.
type DestinationPostgresSSLModesPrefer struct {
	mode *DestinationPostgresSSLModesPreferMode `const:"prefer" json:"mode"`
}

func (d DestinationPostgresSSLModesPrefer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesPrefer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesPrefer) GetMode() *DestinationPostgresSSLModesPreferMode {
	return DestinationPostgresSSLModesPreferModePrefer.ToPointer()
}

type DestinationPostgresSSLModesAllowMode string

const (
	DestinationPostgresSSLModesAllowModeAllow DestinationPostgresSSLModesAllowMode = "allow"
)

func (e DestinationPostgresSSLModesAllowMode) ToPointer() *DestinationPostgresSSLModesAllowMode {
	return &e
}

func (e *DestinationPostgresSSLModesAllowMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "allow":
		*e = DestinationPostgresSSLModesAllowMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesAllowMode: %v", v)
	}
}

// DestinationPostgresSSLModesAllow - Allow SSL mode.
type DestinationPostgresSSLModesAllow struct {
	mode *DestinationPostgresSSLModesAllowMode `const:"allow" json:"mode"`
}

func (d DestinationPostgresSSLModesAllow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesAllow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesAllow) GetMode() *DestinationPostgresSSLModesAllowMode {
	return DestinationPostgresSSLModesAllowModeAllow.ToPointer()
}

type DestinationPostgresSSLModesDisableMode string

const (
	DestinationPostgresSSLModesDisableModeDisable DestinationPostgresSSLModesDisableMode = "disable"
)

func (e DestinationPostgresSSLModesDisableMode) ToPointer() *DestinationPostgresSSLModesDisableMode {
	return &e
}

func (e *DestinationPostgresSSLModesDisableMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disable":
		*e = DestinationPostgresSSLModesDisableMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSLModesDisableMode: %v", v)
	}
}

// DestinationPostgresSSLModesDisable - Disable SSL.
type DestinationPostgresSSLModesDisable struct {
	mode *DestinationPostgresSSLModesDisableMode `const:"disable" json:"mode"`
}

func (d DestinationPostgresSSLModesDisable) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSLModesDisable) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSLModesDisable) GetMode() *DestinationPostgresSSLModesDisableMode {
	return DestinationPostgresSSLModesDisableModeDisable.ToPointer()
}

type DestinationPostgresSSLModesType string

const (
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_disable"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow      DestinationPostgresSSLModesType = "destination-postgres_SSL modes_allow"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer     DestinationPostgresSSLModesType = "destination-postgres_SSL modes_prefer"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire    DestinationPostgresSSLModesType = "destination-postgres_SSL modes_require"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa   DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-ca"
	DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesType = "destination-postgres_SSL modes_verify-full"
)

type DestinationPostgresSSLModes struct {
	DestinationPostgresSSLModesDisable    *DestinationPostgresSSLModesDisable
	DestinationPostgresSSLModesAllow      *DestinationPostgresSSLModesAllow
	DestinationPostgresSSLModesPrefer     *DestinationPostgresSSLModesPrefer
	DestinationPostgresSSLModesRequire    *DestinationPostgresSSLModesRequire
	DestinationPostgresSSLModesVerifyCa   *DestinationPostgresSSLModesVerifyCa
	DestinationPostgresSSLModesVerifyFull *DestinationPostgresSSLModesVerifyFull

	Type DestinationPostgresSSLModesType
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesDisable(destinationPostgresSSLModesDisable DestinationPostgresSSLModesDisable) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesDisable: &destinationPostgresSSLModesDisable,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesAllow(destinationPostgresSSLModesAllow DestinationPostgresSSLModesAllow) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesAllow: &destinationPostgresSSLModesAllow,
		Type:                             typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesPrefer(destinationPostgresSSLModesPrefer DestinationPostgresSSLModesPrefer) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesPrefer: &destinationPostgresSSLModesPrefer,
		Type:                              typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesRequire(destinationPostgresSSLModesRequire DestinationPostgresSSLModesRequire) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesRequire: &destinationPostgresSSLModesRequire,
		Type:                               typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyCa(destinationPostgresSSLModesVerifyCa DestinationPostgresSSLModesVerifyCa) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyCa: &destinationPostgresSSLModesVerifyCa,
		Type:                                typ,
	}
}

func CreateDestinationPostgresSSLModesDestinationPostgresSSLModesVerifyFull(destinationPostgresSSLModesVerifyFull DestinationPostgresSSLModesVerifyFull) DestinationPostgresSSLModes {
	typ := DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull

	return DestinationPostgresSSLModes{
		DestinationPostgresSSLModesVerifyFull: &destinationPostgresSSLModesVerifyFull,
		Type:                                  typ,
	}
}

func (u *DestinationPostgresSSLModes) UnmarshalJSON(data []byte) error {

	destinationPostgresSSLModesDisable := new(DestinationPostgresSSLModesDisable)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesDisable, "", true, true); err == nil {
		u.DestinationPostgresSSLModesDisable = destinationPostgresSSLModesDisable
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesDisable
		return nil
	}

	destinationPostgresSSLModesAllow := new(DestinationPostgresSSLModesAllow)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesAllow, "", true, true); err == nil {
		u.DestinationPostgresSSLModesAllow = destinationPostgresSSLModesAllow
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesAllow
		return nil
	}

	destinationPostgresSSLModesPrefer := new(DestinationPostgresSSLModesPrefer)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesPrefer, "", true, true); err == nil {
		u.DestinationPostgresSSLModesPrefer = destinationPostgresSSLModesPrefer
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesPrefer
		return nil
	}

	destinationPostgresSSLModesRequire := new(DestinationPostgresSSLModesRequire)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesRequire, "", true, true); err == nil {
		u.DestinationPostgresSSLModesRequire = destinationPostgresSSLModesRequire
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesRequire
		return nil
	}

	destinationPostgresSSLModesVerifyCa := new(DestinationPostgresSSLModesVerifyCa)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesVerifyCa, "", true, true); err == nil {
		u.DestinationPostgresSSLModesVerifyCa = destinationPostgresSSLModesVerifyCa
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyCa
		return nil
	}

	destinationPostgresSSLModesVerifyFull := new(DestinationPostgresSSLModesVerifyFull)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSLModesVerifyFull, "", true, true); err == nil {
		u.DestinationPostgresSSLModesVerifyFull = destinationPostgresSSLModesVerifyFull
		u.Type = DestinationPostgresSSLModesTypeDestinationPostgresSSLModesVerifyFull
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSLModes) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSLModesDisable != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesDisable, "", true)
	}

	if u.DestinationPostgresSSLModesAllow != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesAllow, "", true)
	}

	if u.DestinationPostgresSSLModesPrefer != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesPrefer, "", true)
	}

	if u.DestinationPostgresSSLModesRequire != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesRequire, "", true)
	}

	if u.DestinationPostgresSSLModesVerifyCa != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesVerifyCa, "", true)
	}

	if u.DestinationPostgresSSLModesVerifyFull != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSLModesVerifyFull, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationPostgresSSHTunnelMethodPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSHTunnelMethodPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSHTunnelMethodPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationPostgresSSHTunnelMethodPasswordAuthentication) GetTunnelMethod() DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return DestinationPostgresSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth
}

func (o *DestinationPostgresSSHTunnelMethodPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationPostgresSSHTunnelMethodPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationPostgresSSHTunnelMethodPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) GetTunnelMethod() DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return DestinationPostgresSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth
}

func (o *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod string

const (
	DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodNoTunnel DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// DestinationPostgresSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationPostgresSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationPostgresSSHTunnelMethodNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgresSSHTunnelMethodNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgresSSHTunnelMethodNoTunnel) GetTunnelMethod() DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethod {
	return DestinationPostgresSSHTunnelMethodNoTunnelTunnelMethodNoTunnel
}

type DestinationPostgresSSHTunnelMethodType string

const (
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel               DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_No Tunnel"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication   DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_SSH Key Authentication"
	DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodType = "destination-postgres_SSH Tunnel Method_Password Authentication"
)

type DestinationPostgresSSHTunnelMethod struct {
	DestinationPostgresSSHTunnelMethodNoTunnel               *DestinationPostgresSSHTunnelMethodNoTunnel
	DestinationPostgresSSHTunnelMethodSSHKeyAuthentication   *DestinationPostgresSSHTunnelMethodSSHKeyAuthentication
	DestinationPostgresSSHTunnelMethodPasswordAuthentication *DestinationPostgresSSHTunnelMethodPasswordAuthentication

	Type DestinationPostgresSSHTunnelMethodType
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodNoTunnel(destinationPostgresSSHTunnelMethodNoTunnel DestinationPostgresSSHTunnelMethodNoTunnel) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodNoTunnel: &destinationPostgresSSHTunnelMethodNoTunnel,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodSSHKeyAuthentication(destinationPostgresSSHTunnelMethodSSHKeyAuthentication DestinationPostgresSSHTunnelMethodSSHKeyAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodSSHKeyAuthentication: &destinationPostgresSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationPostgresSSHTunnelMethodDestinationPostgresSSHTunnelMethodPasswordAuthentication(destinationPostgresSSHTunnelMethodPasswordAuthentication DestinationPostgresSSHTunnelMethodPasswordAuthentication) DestinationPostgresSSHTunnelMethod {
	typ := DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication

	return DestinationPostgresSSHTunnelMethod{
		DestinationPostgresSSHTunnelMethodPasswordAuthentication: &destinationPostgresSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationPostgresSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationPostgresSSHTunnelMethodNoTunnel := new(DestinationPostgresSSHTunnelMethodNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSHTunnelMethodNoTunnel, "", true, true); err == nil {
		u.DestinationPostgresSSHTunnelMethodNoTunnel = destinationPostgresSSHTunnelMethodNoTunnel
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodNoTunnel
		return nil
	}

	destinationPostgresSSHTunnelMethodSSHKeyAuthentication := new(DestinationPostgresSSHTunnelMethodSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSHTunnelMethodSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication = destinationPostgresSSHTunnelMethodSSHKeyAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	destinationPostgresSSHTunnelMethodPasswordAuthentication := new(DestinationPostgresSSHTunnelMethodPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationPostgresSSHTunnelMethodPasswordAuthentication, "", true, true); err == nil {
		u.DestinationPostgresSSHTunnelMethodPasswordAuthentication = destinationPostgresSSHTunnelMethodPasswordAuthentication
		u.Type = DestinationPostgresSSHTunnelMethodTypeDestinationPostgresSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPostgresSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationPostgresSSHTunnelMethodNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSHTunnelMethodNoTunnel, "", true)
	}

	if u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSHTunnelMethodSSHKeyAuthentication, "", true)
	}

	if u.DestinationPostgresSSHTunnelMethodPasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationPostgresSSHTunnelMethodPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationPostgres struct {
	// Name of the database.
	Database        string                      `json:"database"`
	destinationType DestinationPostgresPostgres `const:"postgres" json:"destinationType"`
	// Hostname of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password *string `json:"password,omitempty"`
	// Port of the database.
	Port *int64 `default:"5432" json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// SSL connection modes.
	//  <b>disable</b> - Chose this mode to disable encryption of communication between Airbyte and destination database
	//  <b>allow</b> - Chose this mode to enable encryption only when required by the source database
	//  <b>prefer</b> - Chose this mode to allow unencrypted connection only if the source database does not support encryption
	//  <b>require</b> - Chose this mode to always require encryption. If the source database server does not support encryption, connection will fail
	//   <b>verify-ca</b> - Chose this mode to always require encryption and to verify that the source database server has a valid SSL certificate
	//   <b>verify-full</b> - This is the most secure mode. Chose this mode to always require encryption and to verify the identity of the source database server
	//  See more information - <a href="https://jdbc.postgresql.org/documentation/head/ssl-client.html"> in the docs</a>.
	SslMode *DestinationPostgresSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationPostgresSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationPostgres) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPostgres) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPostgres) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationPostgres) GetDestinationType() DestinationPostgresPostgres {
	return DestinationPostgresPostgresPostgres
}

func (o *DestinationPostgres) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationPostgres) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationPostgres) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *DestinationPostgres) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationPostgres) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationPostgres) GetSslMode() *DestinationPostgresSSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *DestinationPostgres) GetTunnelMethod() *DestinationPostgresSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationPostgres) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
