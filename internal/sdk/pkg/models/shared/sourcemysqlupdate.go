// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod string

const (
	SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethodStandard SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod = "STANDARD"
)

func (e SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod) ToPointer() *SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod {
	return &e
}

func (e *SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STANDARD":
		*e = SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod: %v", v)
	}
}

// SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor - Incrementally detects new inserts and updates using the <a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor">cursor column</a> chosen when configuring a connection (e.g. created_at, updated_at).
type SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor struct {
	method SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod `const:"STANDARD" json:"method"`
}

func (s SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor) GetMethod() SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethod {
	return SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursorMethodStandard
}

type SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod string

const (
	SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethodCdc SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod = "CDC"
)

func (e SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod) ToPointer() *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod {
	return &e
}

func (e *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CDC":
		*e = SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod: %v", v)
	}
}

// SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC - <i>Recommended</i> - Incrementally reads new inserts, updates, and deletes using the MySQL <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">binary log</a>. This must be enabled on your database.
type SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC struct {
	// The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
	InitialWaitingSeconds *int64                                                          `default:"300" json:"initial_waiting_seconds"`
	method                SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod `const:"CDC" json:"method"`
	// Enter the configured MySQL server timezone. This should only be done if the configured timezone in your MySQL instance does not conform to IANNA standard.
	ServerTimeZone *string `json:"server_time_zone,omitempty"`
}

func (s SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) GetInitialWaitingSeconds() *int64 {
	if o == nil {
		return nil
	}
	return o.InitialWaitingSeconds
}

func (o *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) GetMethod() SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethod {
	return SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDCMethodCdc
}

func (o *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) GetServerTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.ServerTimeZone
}

type SourceMysqlUpdateUpdateMethodType string

const (
	SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC     SourceMysqlUpdateUpdateMethodType = "source-mysql-update_Update Method_Read Changes using Binary Log (CDC)"
	SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor SourceMysqlUpdateUpdateMethodType = "source-mysql-update_Update Method_Scan Changes with User Defined Cursor"
)

type SourceMysqlUpdateUpdateMethod struct {
	SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC     *SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC
	SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor *SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor

	Type SourceMysqlUpdateUpdateMethodType
}

func CreateSourceMysqlUpdateUpdateMethodSourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC(sourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC) SourceMysqlUpdateUpdateMethod {
	typ := SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC

	return SourceMysqlUpdateUpdateMethod{
		SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC: &sourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC,
		Type: typ,
	}
}

func CreateSourceMysqlUpdateUpdateMethodSourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor(sourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor) SourceMysqlUpdateUpdateMethod {
	typ := SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor

	return SourceMysqlUpdateUpdateMethod{
		SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor: &sourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor,
		Type: typ,
	}
}

func (u *SourceMysqlUpdateUpdateMethod) UnmarshalJSON(data []byte) error {

	sourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor := new(SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor, "", true, true); err == nil {
		u.SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor = sourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor
		u.Type = SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor
		return nil
	}

	sourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC := new(SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC, "", true, true); err == nil {
		u.SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC = sourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC
		u.Type = SourceMysqlUpdateUpdateMethodTypeSourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMysqlUpdateUpdateMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateUpdateMethodReadChangesUsingBinaryLogCDC, "", true)
	}

	if u.SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateUpdateMethodScanChangesWithUserDefinedCursor, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMysqlUpdateSSLModesVerifyIdentityMode string

const (
	SourceMysqlUpdateSSLModesVerifyIdentityModeVerifyIdentity SourceMysqlUpdateSSLModesVerifyIdentityMode = "verify_identity"
)

func (e SourceMysqlUpdateSSLModesVerifyIdentityMode) ToPointer() *SourceMysqlUpdateSSLModesVerifyIdentityMode {
	return &e
}

func (e *SourceMysqlUpdateSSLModesVerifyIdentityMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify_identity":
		*e = SourceMysqlUpdateSSLModesVerifyIdentityMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSLModesVerifyIdentityMode: %v", v)
	}
}

// SourceMysqlUpdateSSLModesVerifyIdentity - Always connect with SSL. Verify both CA and Hostname.
type SourceMysqlUpdateSSLModesVerifyIdentity struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate (this is not a required field, but if you want to use it, you will need to add the <b>Client key</b> as well)
	ClientCertificate *string `json:"client_certificate,omitempty"`
	// Client key (this is not a required field, but if you want to use it, you will need to add the <b>Client certificate</b> as well)
	ClientKey *string `json:"client_key,omitempty"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                                     `json:"client_key_password,omitempty"`
	mode              SourceMysqlUpdateSSLModesVerifyIdentityMode `const:"verify_identity" json:"mode"`
}

func (s SourceMysqlUpdateSSLModesVerifyIdentity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSLModesVerifyIdentity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSLModesVerifyIdentity) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *SourceMysqlUpdateSSLModesVerifyIdentity) GetClientCertificate() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *SourceMysqlUpdateSSLModesVerifyIdentity) GetClientKey() *string {
	if o == nil {
		return nil
	}
	return o.ClientKey
}

func (o *SourceMysqlUpdateSSLModesVerifyIdentity) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *SourceMysqlUpdateSSLModesVerifyIdentity) GetMode() SourceMysqlUpdateSSLModesVerifyIdentityMode {
	return SourceMysqlUpdateSSLModesVerifyIdentityModeVerifyIdentity
}

type SourceMysqlUpdateSSLModesVerifyCAMode string

const (
	SourceMysqlUpdateSSLModesVerifyCAModeVerifyCa SourceMysqlUpdateSSLModesVerifyCAMode = "verify_ca"
)

func (e SourceMysqlUpdateSSLModesVerifyCAMode) ToPointer() *SourceMysqlUpdateSSLModesVerifyCAMode {
	return &e
}

func (e *SourceMysqlUpdateSSLModesVerifyCAMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "verify_ca":
		*e = SourceMysqlUpdateSSLModesVerifyCAMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSLModesVerifyCAMode: %v", v)
	}
}

// SourceMysqlUpdateSSLModesVerifyCA - Always connect with SSL. Verifies CA, but allows connection even if Hostname does not match.
type SourceMysqlUpdateSSLModesVerifyCA struct {
	// CA certificate
	CaCertificate string `json:"ca_certificate"`
	// Client certificate (this is not a required field, but if you want to use it, you will need to add the <b>Client key</b> as well)
	ClientCertificate *string `json:"client_certificate,omitempty"`
	// Client key (this is not a required field, but if you want to use it, you will need to add the <b>Client certificate</b> as well)
	ClientKey *string `json:"client_key,omitempty"`
	// Password for keystorage. This field is optional. If you do not add it - the password will be generated automatically.
	ClientKeyPassword *string                               `json:"client_key_password,omitempty"`
	mode              SourceMysqlUpdateSSLModesVerifyCAMode `const:"verify_ca" json:"mode"`
}

func (s SourceMysqlUpdateSSLModesVerifyCA) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSLModesVerifyCA) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSLModesVerifyCA) GetCaCertificate() string {
	if o == nil {
		return ""
	}
	return o.CaCertificate
}

func (o *SourceMysqlUpdateSSLModesVerifyCA) GetClientCertificate() *string {
	if o == nil {
		return nil
	}
	return o.ClientCertificate
}

func (o *SourceMysqlUpdateSSLModesVerifyCA) GetClientKey() *string {
	if o == nil {
		return nil
	}
	return o.ClientKey
}

func (o *SourceMysqlUpdateSSLModesVerifyCA) GetClientKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.ClientKeyPassword
}

func (o *SourceMysqlUpdateSSLModesVerifyCA) GetMode() SourceMysqlUpdateSSLModesVerifyCAMode {
	return SourceMysqlUpdateSSLModesVerifyCAModeVerifyCa
}

type SourceMysqlUpdateSSLModesRequiredMode string

const (
	SourceMysqlUpdateSSLModesRequiredModeRequired SourceMysqlUpdateSSLModesRequiredMode = "required"
)

func (e SourceMysqlUpdateSSLModesRequiredMode) ToPointer() *SourceMysqlUpdateSSLModesRequiredMode {
	return &e
}

func (e *SourceMysqlUpdateSSLModesRequiredMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "required":
		*e = SourceMysqlUpdateSSLModesRequiredMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSLModesRequiredMode: %v", v)
	}
}

// SourceMysqlUpdateSSLModesRequired - Always connect with SSL. If the MySQL server doesn’t support SSL, the connection will not be established. Certificate Authority (CA) and Hostname are not verified.
type SourceMysqlUpdateSSLModesRequired struct {
	mode SourceMysqlUpdateSSLModesRequiredMode `const:"required" json:"mode"`
}

func (s SourceMysqlUpdateSSLModesRequired) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSLModesRequired) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSLModesRequired) GetMode() SourceMysqlUpdateSSLModesRequiredMode {
	return SourceMysqlUpdateSSLModesRequiredModeRequired
}

type SourceMysqlUpdateSSLModesPreferredMode string

const (
	SourceMysqlUpdateSSLModesPreferredModePreferred SourceMysqlUpdateSSLModesPreferredMode = "preferred"
)

func (e SourceMysqlUpdateSSLModesPreferredMode) ToPointer() *SourceMysqlUpdateSSLModesPreferredMode {
	return &e
}

func (e *SourceMysqlUpdateSSLModesPreferredMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preferred":
		*e = SourceMysqlUpdateSSLModesPreferredMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSLModesPreferredMode: %v", v)
	}
}

// SourceMysqlUpdateSSLModesPreferred - Automatically attempt SSL connection. If the MySQL server does not support SSL, continue with a regular connection.
type SourceMysqlUpdateSSLModesPreferred struct {
	mode SourceMysqlUpdateSSLModesPreferredMode `const:"preferred" json:"mode"`
}

func (s SourceMysqlUpdateSSLModesPreferred) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSLModesPreferred) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSLModesPreferred) GetMode() SourceMysqlUpdateSSLModesPreferredMode {
	return SourceMysqlUpdateSSLModesPreferredModePreferred
}

type SourceMysqlUpdateSSLModesType string

const (
	SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesPreferred      SourceMysqlUpdateSSLModesType = "source-mysql-update_SSL modes_preferred"
	SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesRequired       SourceMysqlUpdateSSLModesType = "source-mysql-update_SSL modes_required"
	SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyCA       SourceMysqlUpdateSSLModesType = "source-mysql-update_SSL modes_Verify CA"
	SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyIdentity SourceMysqlUpdateSSLModesType = "source-mysql-update_SSL modes_Verify Identity"
)

type SourceMysqlUpdateSSLModes struct {
	SourceMysqlUpdateSSLModesPreferred      *SourceMysqlUpdateSSLModesPreferred
	SourceMysqlUpdateSSLModesRequired       *SourceMysqlUpdateSSLModesRequired
	SourceMysqlUpdateSSLModesVerifyCA       *SourceMysqlUpdateSSLModesVerifyCA
	SourceMysqlUpdateSSLModesVerifyIdentity *SourceMysqlUpdateSSLModesVerifyIdentity

	Type SourceMysqlUpdateSSLModesType
}

func CreateSourceMysqlUpdateSSLModesSourceMysqlUpdateSSLModesPreferred(sourceMysqlUpdateSSLModesPreferred SourceMysqlUpdateSSLModesPreferred) SourceMysqlUpdateSSLModes {
	typ := SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesPreferred

	return SourceMysqlUpdateSSLModes{
		SourceMysqlUpdateSSLModesPreferred: &sourceMysqlUpdateSSLModesPreferred,
		Type:                               typ,
	}
}

func CreateSourceMysqlUpdateSSLModesSourceMysqlUpdateSSLModesRequired(sourceMysqlUpdateSSLModesRequired SourceMysqlUpdateSSLModesRequired) SourceMysqlUpdateSSLModes {
	typ := SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesRequired

	return SourceMysqlUpdateSSLModes{
		SourceMysqlUpdateSSLModesRequired: &sourceMysqlUpdateSSLModesRequired,
		Type:                              typ,
	}
}

func CreateSourceMysqlUpdateSSLModesSourceMysqlUpdateSSLModesVerifyCA(sourceMysqlUpdateSSLModesVerifyCA SourceMysqlUpdateSSLModesVerifyCA) SourceMysqlUpdateSSLModes {
	typ := SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyCA

	return SourceMysqlUpdateSSLModes{
		SourceMysqlUpdateSSLModesVerifyCA: &sourceMysqlUpdateSSLModesVerifyCA,
		Type:                              typ,
	}
}

func CreateSourceMysqlUpdateSSLModesSourceMysqlUpdateSSLModesVerifyIdentity(sourceMysqlUpdateSSLModesVerifyIdentity SourceMysqlUpdateSSLModesVerifyIdentity) SourceMysqlUpdateSSLModes {
	typ := SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyIdentity

	return SourceMysqlUpdateSSLModes{
		SourceMysqlUpdateSSLModesVerifyIdentity: &sourceMysqlUpdateSSLModesVerifyIdentity,
		Type:                                    typ,
	}
}

func (u *SourceMysqlUpdateSSLModes) UnmarshalJSON(data []byte) error {

	sourceMysqlUpdateSSLModesPreferred := new(SourceMysqlUpdateSSLModesPreferred)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSLModesPreferred, "", true, true); err == nil {
		u.SourceMysqlUpdateSSLModesPreferred = sourceMysqlUpdateSSLModesPreferred
		u.Type = SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesPreferred
		return nil
	}

	sourceMysqlUpdateSSLModesRequired := new(SourceMysqlUpdateSSLModesRequired)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSLModesRequired, "", true, true); err == nil {
		u.SourceMysqlUpdateSSLModesRequired = sourceMysqlUpdateSSLModesRequired
		u.Type = SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesRequired
		return nil
	}

	sourceMysqlUpdateSSLModesVerifyCA := new(SourceMysqlUpdateSSLModesVerifyCA)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSLModesVerifyCA, "", true, true); err == nil {
		u.SourceMysqlUpdateSSLModesVerifyCA = sourceMysqlUpdateSSLModesVerifyCA
		u.Type = SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyCA
		return nil
	}

	sourceMysqlUpdateSSLModesVerifyIdentity := new(SourceMysqlUpdateSSLModesVerifyIdentity)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSLModesVerifyIdentity, "", true, true); err == nil {
		u.SourceMysqlUpdateSSLModesVerifyIdentity = sourceMysqlUpdateSSLModesVerifyIdentity
		u.Type = SourceMysqlUpdateSSLModesTypeSourceMysqlUpdateSSLModesVerifyIdentity
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMysqlUpdateSSLModes) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlUpdateSSLModesPreferred != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSLModesPreferred, "", true)
	}

	if u.SourceMysqlUpdateSSLModesRequired != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSLModesRequired, "", true)
	}

	if u.SourceMysqlUpdateSSLModesVerifyCA != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSLModesVerifyCA, "", true)
	}

	if u.SourceMysqlUpdateSSLModesVerifyIdentity != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSLModesVerifyIdentity, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod string

const (
	SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) ToPointer() *SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (s SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelMethod() SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethod {
	return SourceMysqlUpdateSSHTunnelMethodPasswordAuthenticationTunnelMethodSSHPasswordAuth
}

func (o *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod string

const (
	SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod = "SSH_KEY_AUTH"
)

func (e SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) ToPointer() *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return &e
}

func (e *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod: %v", v)
	}
}

// SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (s SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelMethod() SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethod {
	return SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthenticationTunnelMethodSSHKeyAuth
}

func (o *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod - No ssh tunnel needed to connect to database
type SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod string

const (
	SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod = "NO_TUNNEL"
)

func (e SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) ToPointer() *SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return &e
}

func (e *SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod: %v", v)
	}
}

// SourceMysqlUpdateSSHTunnelMethodNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type SourceMysqlUpdateSSHTunnelMethodNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (s SourceMysqlUpdateSSHTunnelMethodNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdateSSHTunnelMethodNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdateSSHTunnelMethodNoTunnel) GetTunnelMethod() SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethod {
	return SourceMysqlUpdateSSHTunnelMethodNoTunnelTunnelMethodNoTunnel
}

type SourceMysqlUpdateSSHTunnelMethodType string

const (
	SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodNoTunnel               SourceMysqlUpdateSSHTunnelMethodType = "source-mysql-update_SSH Tunnel Method_No Tunnel"
	SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication   SourceMysqlUpdateSSHTunnelMethodType = "source-mysql-update_SSH Tunnel Method_SSH Key Authentication"
	SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodPasswordAuthentication SourceMysqlUpdateSSHTunnelMethodType = "source-mysql-update_SSH Tunnel Method_Password Authentication"
)

type SourceMysqlUpdateSSHTunnelMethod struct {
	SourceMysqlUpdateSSHTunnelMethodNoTunnel               *SourceMysqlUpdateSSHTunnelMethodNoTunnel
	SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication   *SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
	SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication *SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication

	Type SourceMysqlUpdateSSHTunnelMethodType
}

func CreateSourceMysqlUpdateSSHTunnelMethodSourceMysqlUpdateSSHTunnelMethodNoTunnel(sourceMysqlUpdateSSHTunnelMethodNoTunnel SourceMysqlUpdateSSHTunnelMethodNoTunnel) SourceMysqlUpdateSSHTunnelMethod {
	typ := SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodNoTunnel

	return SourceMysqlUpdateSSHTunnelMethod{
		SourceMysqlUpdateSSHTunnelMethodNoTunnel: &sourceMysqlUpdateSSHTunnelMethodNoTunnel,
		Type:                                     typ,
	}
}

func CreateSourceMysqlUpdateSSHTunnelMethodSourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication(sourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication) SourceMysqlUpdateSSHTunnelMethod {
	typ := SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication

	return SourceMysqlUpdateSSHTunnelMethod{
		SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication: &sourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateSourceMysqlUpdateSSHTunnelMethodSourceMysqlUpdateSSHTunnelMethodPasswordAuthentication(sourceMysqlUpdateSSHTunnelMethodPasswordAuthentication SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication) SourceMysqlUpdateSSHTunnelMethod {
	typ := SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodPasswordAuthentication

	return SourceMysqlUpdateSSHTunnelMethod{
		SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication: &sourceMysqlUpdateSSHTunnelMethodPasswordAuthentication,
		Type: typ,
	}
}

func (u *SourceMysqlUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	sourceMysqlUpdateSSHTunnelMethodNoTunnel := new(SourceMysqlUpdateSSHTunnelMethodNoTunnel)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSHTunnelMethodNoTunnel, "", true, true); err == nil {
		u.SourceMysqlUpdateSSHTunnelMethodNoTunnel = sourceMysqlUpdateSSHTunnelMethodNoTunnel
		u.Type = SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodNoTunnel
		return nil
	}

	sourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication := new(SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication, "", true, true); err == nil {
		u.SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication = sourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
		u.Type = SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication
		return nil
	}

	sourceMysqlUpdateSSHTunnelMethodPasswordAuthentication := new(SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceMysqlUpdateSSHTunnelMethodPasswordAuthentication, "", true, true); err == nil {
		u.SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication = sourceMysqlUpdateSSHTunnelMethodPasswordAuthentication
		u.Type = SourceMysqlUpdateSSHTunnelMethodTypeSourceMysqlUpdateSSHTunnelMethodPasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMysqlUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.SourceMysqlUpdateSSHTunnelMethodNoTunnel != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSHTunnelMethodNoTunnel, "", true)
	}

	if u.SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSHTunnelMethodSSHKeyAuthentication, "", true)
	}

	if u.SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication != nil {
		return utils.MarshalJSON(u.SourceMysqlUpdateSSHTunnelMethodPasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMysqlUpdate struct {
	// The database name.
	Database string `json:"database"`
	// The host name of the database.
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3). For more information read about <a href="https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-jdbc-url-format.html">JDBC URL parameters</a>.
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The password associated with the username.
	Password *string `json:"password,omitempty"`
	// The port to connect to.
	Port *int64 `default:"3306" json:"port"`
	// Configures how data is extracted from the database.
	ReplicationMethod SourceMysqlUpdateUpdateMethod `json:"replication_method"`
	// SSL connection modes. Read more <a href="https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-reference-using-ssl.html"> in the docs</a>.
	SslMode *SourceMysqlUpdateSSLModes `json:"ssl_mode,omitempty"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *SourceMysqlUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The username which is used to access the database.
	Username string `json:"username"`
}

func (s SourceMysqlUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMysqlUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMysqlUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceMysqlUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceMysqlUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceMysqlUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceMysqlUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceMysqlUpdate) GetReplicationMethod() SourceMysqlUpdateUpdateMethod {
	if o == nil {
		return SourceMysqlUpdateUpdateMethod{}
	}
	return o.ReplicationMethod
}

func (o *SourceMysqlUpdate) GetSslMode() *SourceMysqlUpdateSSLModes {
	if o == nil {
		return nil
	}
	return o.SslMode
}

func (o *SourceMysqlUpdate) GetTunnelMethod() *SourceMysqlUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *SourceMysqlUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
