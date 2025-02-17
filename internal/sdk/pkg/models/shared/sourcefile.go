// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// SourceFileFileFormat - The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
type SourceFileFileFormat string

const (
	SourceFileFileFormatCsv         SourceFileFileFormat = "csv"
	SourceFileFileFormatJSON        SourceFileFileFormat = "json"
	SourceFileFileFormatJsonl       SourceFileFileFormat = "jsonl"
	SourceFileFileFormatExcel       SourceFileFileFormat = "excel"
	SourceFileFileFormatExcelBinary SourceFileFileFormat = "excel_binary"
	SourceFileFileFormatFeather     SourceFileFileFormat = "feather"
	SourceFileFileFormatParquet     SourceFileFileFormat = "parquet"
	SourceFileFileFormatYaml        SourceFileFileFormat = "yaml"
)

func (e SourceFileFileFormat) ToPointer() *SourceFileFileFormat {
	return &e
}

func (e *SourceFileFileFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "csv":
		fallthrough
	case "json":
		fallthrough
	case "jsonl":
		fallthrough
	case "excel":
		fallthrough
	case "excel_binary":
		fallthrough
	case "feather":
		fallthrough
	case "parquet":
		fallthrough
	case "yaml":
		*e = SourceFileFileFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileFileFormat: %v", v)
	}
}

type SourceFileSchemasProviderStorageProvider7Storage string

const (
	SourceFileSchemasProviderStorageProvider7StorageSftp SourceFileSchemasProviderStorageProvider7Storage = "SFTP"
)

func (e SourceFileSchemasProviderStorageProvider7Storage) ToPointer() *SourceFileSchemasProviderStorageProvider7Storage {
	return &e
}

func (e *SourceFileSchemasProviderStorageProvider7Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SFTP":
		*e = SourceFileSchemasProviderStorageProvider7Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasProviderStorageProvider7Storage: %v", v)
	}
}

// SourceFileSFTPSecureFileTransferProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSFTPSecureFileTransferProtocol struct {
	Host     string                                           `json:"host"`
	Password *string                                          `json:"password,omitempty"`
	Port     *string                                          `default:"22" json:"port"`
	storage  SourceFileSchemasProviderStorageProvider7Storage `const:"SFTP" json:"storage"`
	User     string                                           `json:"user"`
}

func (s SourceFileSFTPSecureFileTransferProtocol) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileSFTPSecureFileTransferProtocol) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileSFTPSecureFileTransferProtocol) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceFileSFTPSecureFileTransferProtocol) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceFileSFTPSecureFileTransferProtocol) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceFileSFTPSecureFileTransferProtocol) GetStorage() SourceFileSchemasProviderStorageProvider7Storage {
	return SourceFileSchemasProviderStorageProvider7StorageSftp
}

func (o *SourceFileSFTPSecureFileTransferProtocol) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileSchemasProviderStorageProvider6Storage string

const (
	SourceFileSchemasProviderStorageProvider6StorageScp SourceFileSchemasProviderStorageProvider6Storage = "SCP"
)

func (e SourceFileSchemasProviderStorageProvider6Storage) ToPointer() *SourceFileSchemasProviderStorageProvider6Storage {
	return &e
}

func (e *SourceFileSchemasProviderStorageProvider6Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SCP":
		*e = SourceFileSchemasProviderStorageProvider6Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasProviderStorageProvider6Storage: %v", v)
	}
}

// SourceFileSCPSecureCopyProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSCPSecureCopyProtocol struct {
	Host     string                                           `json:"host"`
	Password *string                                          `json:"password,omitempty"`
	Port     *string                                          `default:"22" json:"port"`
	storage  SourceFileSchemasProviderStorageProvider6Storage `const:"SCP" json:"storage"`
	User     string                                           `json:"user"`
}

func (s SourceFileSCPSecureCopyProtocol) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileSCPSecureCopyProtocol) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileSCPSecureCopyProtocol) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceFileSCPSecureCopyProtocol) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceFileSCPSecureCopyProtocol) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceFileSCPSecureCopyProtocol) GetStorage() SourceFileSchemasProviderStorageProvider6Storage {
	return SourceFileSchemasProviderStorageProvider6StorageScp
}

func (o *SourceFileSCPSecureCopyProtocol) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileSchemasProviderStorageProvider5Storage string

const (
	SourceFileSchemasProviderStorageProvider5StorageSSH SourceFileSchemasProviderStorageProvider5Storage = "SSH"
)

func (e SourceFileSchemasProviderStorageProvider5Storage) ToPointer() *SourceFileSchemasProviderStorageProvider5Storage {
	return &e
}

func (e *SourceFileSchemasProviderStorageProvider5Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH":
		*e = SourceFileSchemasProviderStorageProvider5Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasProviderStorageProvider5Storage: %v", v)
	}
}

// SourceFileSSHSecureShell - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileSSHSecureShell struct {
	Host     string                                           `json:"host"`
	Password *string                                          `json:"password,omitempty"`
	Port     *string                                          `default:"22" json:"port"`
	storage  SourceFileSchemasProviderStorageProvider5Storage `const:"SSH" json:"storage"`
	User     string                                           `json:"user"`
}

func (s SourceFileSSHSecureShell) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileSSHSecureShell) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileSSHSecureShell) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceFileSSHSecureShell) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceFileSSHSecureShell) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceFileSSHSecureShell) GetStorage() SourceFileSchemasProviderStorageProvider5Storage {
	return SourceFileSchemasProviderStorageProvider5StorageSSH
}

func (o *SourceFileSSHSecureShell) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileSchemasProviderStorageProviderStorage string

const (
	SourceFileSchemasProviderStorageProviderStorageAzBlob SourceFileSchemasProviderStorageProviderStorage = "AzBlob"
)

func (e SourceFileSchemasProviderStorageProviderStorage) ToPointer() *SourceFileSchemasProviderStorageProviderStorage {
	return &e
}

func (e *SourceFileSchemasProviderStorageProviderStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AzBlob":
		*e = SourceFileSchemasProviderStorageProviderStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasProviderStorageProviderStorage: %v", v)
	}
}

// SourceFileAzBlobAzureBlobStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileAzBlobAzureBlobStorage struct {
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
	SasToken *string `json:"sas_token,omitempty"`
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
	SharedKey *string                                         `json:"shared_key,omitempty"`
	storage   SourceFileSchemasProviderStorageProviderStorage `const:"AzBlob" json:"storage"`
	// The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.
	StorageAccount string `json:"storage_account"`
}

func (s SourceFileAzBlobAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileAzBlobAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileAzBlobAzureBlobStorage) GetSasToken() *string {
	if o == nil {
		return nil
	}
	return o.SasToken
}

func (o *SourceFileAzBlobAzureBlobStorage) GetSharedKey() *string {
	if o == nil {
		return nil
	}
	return o.SharedKey
}

func (o *SourceFileAzBlobAzureBlobStorage) GetStorage() SourceFileSchemasProviderStorageProviderStorage {
	return SourceFileSchemasProviderStorageProviderStorageAzBlob
}

func (o *SourceFileAzBlobAzureBlobStorage) GetStorageAccount() string {
	if o == nil {
		return ""
	}
	return o.StorageAccount
}

type SourceFileSchemasProviderStorage string

const (
	SourceFileSchemasProviderStorageS3 SourceFileSchemasProviderStorage = "S3"
)

func (e SourceFileSchemasProviderStorage) ToPointer() *SourceFileSchemasProviderStorage {
	return &e
}

func (e *SourceFileSchemasProviderStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = SourceFileSchemasProviderStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasProviderStorage: %v", v)
	}
}

// SourceFileS3AmazonWebServices - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileS3AmazonWebServices struct {
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsSecretAccessKey *string                          `json:"aws_secret_access_key,omitempty"`
	storage            SourceFileSchemasProviderStorage `const:"S3" json:"storage"`
}

func (s SourceFileS3AmazonWebServices) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileS3AmazonWebServices) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileS3AmazonWebServices) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *SourceFileS3AmazonWebServices) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *SourceFileS3AmazonWebServices) GetStorage() SourceFileSchemasProviderStorage {
	return SourceFileSchemasProviderStorageS3
}

type SourceFileSchemasStorage string

const (
	SourceFileSchemasStorageGcs SourceFileSchemasStorage = "GCS"
)

func (e SourceFileSchemasStorage) ToPointer() *SourceFileSchemasStorage {
	return &e
}

func (e *SourceFileSchemasStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS":
		*e = SourceFileSchemasStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileSchemasStorage: %v", v)
	}
}

// SourceFileGCSGoogleCloudStorage - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileGCSGoogleCloudStorage struct {
	// In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
	ServiceAccountJSON *string                  `json:"service_account_json,omitempty"`
	storage            SourceFileSchemasStorage `const:"GCS" json:"storage"`
}

func (s SourceFileGCSGoogleCloudStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileGCSGoogleCloudStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileGCSGoogleCloudStorage) GetServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.ServiceAccountJSON
}

func (o *SourceFileGCSGoogleCloudStorage) GetStorage() SourceFileSchemasStorage {
	return SourceFileSchemasStorageGcs
}

type SourceFileStorage string

const (
	SourceFileStorageHTTPS SourceFileStorage = "HTTPS"
)

func (e SourceFileStorage) ToPointer() *SourceFileStorage {
	return &e
}

func (e *SourceFileStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTPS":
		*e = SourceFileStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileStorage: %v", v)
	}
}

// SourceFileHTTPSPublicWeb - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileHTTPSPublicWeb struct {
	storage SourceFileStorage `const:"HTTPS" json:"storage"`
	// Add User-Agent to request
	UserAgent *bool `default:"false" json:"user_agent"`
}

func (s SourceFileHTTPSPublicWeb) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileHTTPSPublicWeb) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileHTTPSPublicWeb) GetStorage() SourceFileStorage {
	return SourceFileStorageHTTPS
}

func (o *SourceFileHTTPSPublicWeb) GetUserAgent() *bool {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

type SourceFileStorageProviderType string

const (
	SourceFileStorageProviderTypeSourceFileHTTPSPublicWeb                 SourceFileStorageProviderType = "source-file_HTTPS: Public Web"
	SourceFileStorageProviderTypeSourceFileGCSGoogleCloudStorage          SourceFileStorageProviderType = "source-file_GCS: Google Cloud Storage"
	SourceFileStorageProviderTypeSourceFileS3AmazonWebServices            SourceFileStorageProviderType = "source-file_S3: Amazon Web Services"
	SourceFileStorageProviderTypeSourceFileAzBlobAzureBlobStorage         SourceFileStorageProviderType = "source-file_AzBlob: Azure Blob Storage"
	SourceFileStorageProviderTypeSourceFileSSHSecureShell                 SourceFileStorageProviderType = "source-file_SSH: Secure Shell"
	SourceFileStorageProviderTypeSourceFileSCPSecureCopyProtocol          SourceFileStorageProviderType = "source-file_SCP: Secure copy protocol"
	SourceFileStorageProviderTypeSourceFileSFTPSecureFileTransferProtocol SourceFileStorageProviderType = "source-file_SFTP: Secure File Transfer Protocol"
)

type SourceFileStorageProvider struct {
	SourceFileHTTPSPublicWeb                 *SourceFileHTTPSPublicWeb
	SourceFileGCSGoogleCloudStorage          *SourceFileGCSGoogleCloudStorage
	SourceFileS3AmazonWebServices            *SourceFileS3AmazonWebServices
	SourceFileAzBlobAzureBlobStorage         *SourceFileAzBlobAzureBlobStorage
	SourceFileSSHSecureShell                 *SourceFileSSHSecureShell
	SourceFileSCPSecureCopyProtocol          *SourceFileSCPSecureCopyProtocol
	SourceFileSFTPSecureFileTransferProtocol *SourceFileSFTPSecureFileTransferProtocol

	Type SourceFileStorageProviderType
}

func CreateSourceFileStorageProviderSourceFileHTTPSPublicWeb(sourceFileHTTPSPublicWeb SourceFileHTTPSPublicWeb) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileHTTPSPublicWeb

	return SourceFileStorageProvider{
		SourceFileHTTPSPublicWeb: &sourceFileHTTPSPublicWeb,
		Type:                     typ,
	}
}

func CreateSourceFileStorageProviderSourceFileGCSGoogleCloudStorage(sourceFileGCSGoogleCloudStorage SourceFileGCSGoogleCloudStorage) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileGCSGoogleCloudStorage

	return SourceFileStorageProvider{
		SourceFileGCSGoogleCloudStorage: &sourceFileGCSGoogleCloudStorage,
		Type:                            typ,
	}
}

func CreateSourceFileStorageProviderSourceFileS3AmazonWebServices(sourceFileS3AmazonWebServices SourceFileS3AmazonWebServices) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileS3AmazonWebServices

	return SourceFileStorageProvider{
		SourceFileS3AmazonWebServices: &sourceFileS3AmazonWebServices,
		Type:                          typ,
	}
}

func CreateSourceFileStorageProviderSourceFileAzBlobAzureBlobStorage(sourceFileAzBlobAzureBlobStorage SourceFileAzBlobAzureBlobStorage) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileAzBlobAzureBlobStorage

	return SourceFileStorageProvider{
		SourceFileAzBlobAzureBlobStorage: &sourceFileAzBlobAzureBlobStorage,
		Type:                             typ,
	}
}

func CreateSourceFileStorageProviderSourceFileSSHSecureShell(sourceFileSSHSecureShell SourceFileSSHSecureShell) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileSSHSecureShell

	return SourceFileStorageProvider{
		SourceFileSSHSecureShell: &sourceFileSSHSecureShell,
		Type:                     typ,
	}
}

func CreateSourceFileStorageProviderSourceFileSCPSecureCopyProtocol(sourceFileSCPSecureCopyProtocol SourceFileSCPSecureCopyProtocol) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileSCPSecureCopyProtocol

	return SourceFileStorageProvider{
		SourceFileSCPSecureCopyProtocol: &sourceFileSCPSecureCopyProtocol,
		Type:                            typ,
	}
}

func CreateSourceFileStorageProviderSourceFileSFTPSecureFileTransferProtocol(sourceFileSFTPSecureFileTransferProtocol SourceFileSFTPSecureFileTransferProtocol) SourceFileStorageProvider {
	typ := SourceFileStorageProviderTypeSourceFileSFTPSecureFileTransferProtocol

	return SourceFileStorageProvider{
		SourceFileSFTPSecureFileTransferProtocol: &sourceFileSFTPSecureFileTransferProtocol,
		Type:                                     typ,
	}
}

func (u *SourceFileStorageProvider) UnmarshalJSON(data []byte) error {

	sourceFileHTTPSPublicWeb := new(SourceFileHTTPSPublicWeb)
	if err := utils.UnmarshalJSON(data, &sourceFileHTTPSPublicWeb, "", true, true); err == nil {
		u.SourceFileHTTPSPublicWeb = sourceFileHTTPSPublicWeb
		u.Type = SourceFileStorageProviderTypeSourceFileHTTPSPublicWeb
		return nil
	}

	sourceFileGCSGoogleCloudStorage := new(SourceFileGCSGoogleCloudStorage)
	if err := utils.UnmarshalJSON(data, &sourceFileGCSGoogleCloudStorage, "", true, true); err == nil {
		u.SourceFileGCSGoogleCloudStorage = sourceFileGCSGoogleCloudStorage
		u.Type = SourceFileStorageProviderTypeSourceFileGCSGoogleCloudStorage
		return nil
	}

	sourceFileS3AmazonWebServices := new(SourceFileS3AmazonWebServices)
	if err := utils.UnmarshalJSON(data, &sourceFileS3AmazonWebServices, "", true, true); err == nil {
		u.SourceFileS3AmazonWebServices = sourceFileS3AmazonWebServices
		u.Type = SourceFileStorageProviderTypeSourceFileS3AmazonWebServices
		return nil
	}

	sourceFileAzBlobAzureBlobStorage := new(SourceFileAzBlobAzureBlobStorage)
	if err := utils.UnmarshalJSON(data, &sourceFileAzBlobAzureBlobStorage, "", true, true); err == nil {
		u.SourceFileAzBlobAzureBlobStorage = sourceFileAzBlobAzureBlobStorage
		u.Type = SourceFileStorageProviderTypeSourceFileAzBlobAzureBlobStorage
		return nil
	}

	sourceFileSSHSecureShell := new(SourceFileSSHSecureShell)
	if err := utils.UnmarshalJSON(data, &sourceFileSSHSecureShell, "", true, true); err == nil {
		u.SourceFileSSHSecureShell = sourceFileSSHSecureShell
		u.Type = SourceFileStorageProviderTypeSourceFileSSHSecureShell
		return nil
	}

	sourceFileSCPSecureCopyProtocol := new(SourceFileSCPSecureCopyProtocol)
	if err := utils.UnmarshalJSON(data, &sourceFileSCPSecureCopyProtocol, "", true, true); err == nil {
		u.SourceFileSCPSecureCopyProtocol = sourceFileSCPSecureCopyProtocol
		u.Type = SourceFileStorageProviderTypeSourceFileSCPSecureCopyProtocol
		return nil
	}

	sourceFileSFTPSecureFileTransferProtocol := new(SourceFileSFTPSecureFileTransferProtocol)
	if err := utils.UnmarshalJSON(data, &sourceFileSFTPSecureFileTransferProtocol, "", true, true); err == nil {
		u.SourceFileSFTPSecureFileTransferProtocol = sourceFileSFTPSecureFileTransferProtocol
		u.Type = SourceFileStorageProviderTypeSourceFileSFTPSecureFileTransferProtocol
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceFileStorageProvider) MarshalJSON() ([]byte, error) {
	if u.SourceFileHTTPSPublicWeb != nil {
		return utils.MarshalJSON(u.SourceFileHTTPSPublicWeb, "", true)
	}

	if u.SourceFileGCSGoogleCloudStorage != nil {
		return utils.MarshalJSON(u.SourceFileGCSGoogleCloudStorage, "", true)
	}

	if u.SourceFileS3AmazonWebServices != nil {
		return utils.MarshalJSON(u.SourceFileS3AmazonWebServices, "", true)
	}

	if u.SourceFileAzBlobAzureBlobStorage != nil {
		return utils.MarshalJSON(u.SourceFileAzBlobAzureBlobStorage, "", true)
	}

	if u.SourceFileSSHSecureShell != nil {
		return utils.MarshalJSON(u.SourceFileSSHSecureShell, "", true)
	}

	if u.SourceFileSCPSecureCopyProtocol != nil {
		return utils.MarshalJSON(u.SourceFileSCPSecureCopyProtocol, "", true)
	}

	if u.SourceFileSFTPSecureFileTransferProtocol != nil {
		return utils.MarshalJSON(u.SourceFileSFTPSecureFileTransferProtocol, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type File string

const (
	FileFile File = "file"
)

func (e File) ToPointer() *File {
	return &e
}

func (e *File) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "file":
		*e = File(v)
		return nil
	default:
		return fmt.Errorf("invalid value for File: %v", v)
	}
}

type SourceFile struct {
	// The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
	DatasetName string `json:"dataset_name"`
	// The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
	Format *SourceFileFileFormat `default:"csv" json:"format"`
	// The storage Provider or Location of the file(s) which should be replicated.
	Provider SourceFileStorageProvider `json:"provider"`
	// This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.
	ReaderOptions *string `json:"reader_options,omitempty"`
	sourceType    File    `const:"file" json:"sourceType"`
	// The URL path to access the file which should be replicated.
	URL string `json:"url"`
}

func (s SourceFile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFile) GetDatasetName() string {
	if o == nil {
		return ""
	}
	return o.DatasetName
}

func (o *SourceFile) GetFormat() *SourceFileFileFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *SourceFile) GetProvider() SourceFileStorageProvider {
	if o == nil {
		return SourceFileStorageProvider{}
	}
	return o.Provider
}

func (o *SourceFile) GetReaderOptions() *string {
	if o == nil {
		return nil
	}
	return o.ReaderOptions
}

func (o *SourceFile) GetSourceType() File {
	return FileFile
}

func (o *SourceFile) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
