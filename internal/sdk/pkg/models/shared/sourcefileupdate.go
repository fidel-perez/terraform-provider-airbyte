// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// FileFormat - The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
type FileFormat string

const (
	FileFormatCsv         FileFormat = "csv"
	FileFormatJSON        FileFormat = "json"
	FileFormatJsonl       FileFormat = "jsonl"
	FileFormatExcel       FileFormat = "excel"
	FileFormatExcelBinary FileFormat = "excel_binary"
	FileFormatFeather     FileFormat = "feather"
	FileFormatParquet     FileFormat = "parquet"
	FileFormatYaml        FileFormat = "yaml"
)

func (e FileFormat) ToPointer() *FileFormat {
	return &e
}

func (e *FileFormat) UnmarshalJSON(data []byte) error {
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
		*e = FileFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FileFormat: %v", v)
	}
}

type SourceFileUpdateSchemasProviderStorageProvider7Storage string

const (
	SourceFileUpdateSchemasProviderStorageProvider7StorageSftp SourceFileUpdateSchemasProviderStorageProvider7Storage = "SFTP"
)

func (e SourceFileUpdateSchemasProviderStorageProvider7Storage) ToPointer() *SourceFileUpdateSchemasProviderStorageProvider7Storage {
	return &e
}

func (e *SourceFileUpdateSchemasProviderStorageProvider7Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SFTP":
		*e = SourceFileUpdateSchemasProviderStorageProvider7Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateSchemasProviderStorageProvider7Storage: %v", v)
	}
}

// SFTPSecureFileTransferProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SFTPSecureFileTransferProtocol struct {
	Host     string                                                 `json:"host"`
	Password *string                                                `json:"password,omitempty"`
	Port     *string                                                `default:"22" json:"port"`
	storage  SourceFileUpdateSchemasProviderStorageProvider7Storage `const:"SFTP" json:"storage"`
	User     string                                                 `json:"user"`
}

func (s SFTPSecureFileTransferProtocol) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SFTPSecureFileTransferProtocol) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SFTPSecureFileTransferProtocol) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SFTPSecureFileTransferProtocol) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SFTPSecureFileTransferProtocol) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SFTPSecureFileTransferProtocol) GetStorage() SourceFileUpdateSchemasProviderStorageProvider7Storage {
	return SourceFileUpdateSchemasProviderStorageProvider7StorageSftp
}

func (o *SFTPSecureFileTransferProtocol) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileUpdateSchemasProviderStorageProvider6Storage string

const (
	SourceFileUpdateSchemasProviderStorageProvider6StorageScp SourceFileUpdateSchemasProviderStorageProvider6Storage = "SCP"
)

func (e SourceFileUpdateSchemasProviderStorageProvider6Storage) ToPointer() *SourceFileUpdateSchemasProviderStorageProvider6Storage {
	return &e
}

func (e *SourceFileUpdateSchemasProviderStorageProvider6Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SCP":
		*e = SourceFileUpdateSchemasProviderStorageProvider6Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateSchemasProviderStorageProvider6Storage: %v", v)
	}
}

// SCPSecureCopyProtocol - The storage Provider or Location of the file(s) which should be replicated.
type SCPSecureCopyProtocol struct {
	Host     string                                                 `json:"host"`
	Password *string                                                `json:"password,omitempty"`
	Port     *string                                                `default:"22" json:"port"`
	storage  SourceFileUpdateSchemasProviderStorageProvider6Storage `const:"SCP" json:"storage"`
	User     string                                                 `json:"user"`
}

func (s SCPSecureCopyProtocol) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SCPSecureCopyProtocol) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SCPSecureCopyProtocol) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SCPSecureCopyProtocol) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SCPSecureCopyProtocol) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SCPSecureCopyProtocol) GetStorage() SourceFileUpdateSchemasProviderStorageProvider6Storage {
	return SourceFileUpdateSchemasProviderStorageProvider6StorageScp
}

func (o *SCPSecureCopyProtocol) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileUpdateSchemasProviderStorageProviderStorage string

const (
	SourceFileUpdateSchemasProviderStorageProviderStorageSSH SourceFileUpdateSchemasProviderStorageProviderStorage = "SSH"
)

func (e SourceFileUpdateSchemasProviderStorageProviderStorage) ToPointer() *SourceFileUpdateSchemasProviderStorageProviderStorage {
	return &e
}

func (e *SourceFileUpdateSchemasProviderStorageProviderStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH":
		*e = SourceFileUpdateSchemasProviderStorageProviderStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateSchemasProviderStorageProviderStorage: %v", v)
	}
}

// SSHSecureShell - The storage Provider or Location of the file(s) which should be replicated.
type SSHSecureShell struct {
	Host     string                                                `json:"host"`
	Password *string                                               `json:"password,omitempty"`
	Port     *string                                               `default:"22" json:"port"`
	storage  SourceFileUpdateSchemasProviderStorageProviderStorage `const:"SSH" json:"storage"`
	User     string                                                `json:"user"`
}

func (s SSHSecureShell) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SSHSecureShell) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SSHSecureShell) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SSHSecureShell) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SSHSecureShell) GetPort() *string {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SSHSecureShell) GetStorage() SourceFileUpdateSchemasProviderStorageProviderStorage {
	return SourceFileUpdateSchemasProviderStorageProviderStorageSSH
}

func (o *SSHSecureShell) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

type SourceFileUpdateSchemasProviderStorage string

const (
	SourceFileUpdateSchemasProviderStorageAzBlob SourceFileUpdateSchemasProviderStorage = "AzBlob"
)

func (e SourceFileUpdateSchemasProviderStorage) ToPointer() *SourceFileUpdateSchemasProviderStorage {
	return &e
}

func (e *SourceFileUpdateSchemasProviderStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AzBlob":
		*e = SourceFileUpdateSchemasProviderStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateSchemasProviderStorage: %v", v)
	}
}

// AzBlobAzureBlobStorage - The storage Provider or Location of the file(s) which should be replicated.
type AzBlobAzureBlobStorage struct {
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a SAS (Shared Access Signature) token. If accessing publicly available data, this field is not necessary.
	SasToken *string `json:"sas_token,omitempty"`
	// To access Azure Blob Storage, this connector would need credentials with the proper permissions. One option is a storage account shared key (aka account key or access key). If accessing publicly available data, this field is not necessary.
	SharedKey *string                                `json:"shared_key,omitempty"`
	storage   SourceFileUpdateSchemasProviderStorage `const:"AzBlob" json:"storage"`
	// The globally unique name of the storage account that the desired blob sits within. See <a href="https://docs.microsoft.com/en-us/azure/storage/common/storage-account-overview" target="_blank">here</a> for more details.
	StorageAccount string `json:"storage_account"`
}

func (a AzBlobAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AzBlobAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AzBlobAzureBlobStorage) GetSasToken() *string {
	if o == nil {
		return nil
	}
	return o.SasToken
}

func (o *AzBlobAzureBlobStorage) GetSharedKey() *string {
	if o == nil {
		return nil
	}
	return o.SharedKey
}

func (o *AzBlobAzureBlobStorage) GetStorage() SourceFileUpdateSchemasProviderStorage {
	return SourceFileUpdateSchemasProviderStorageAzBlob
}

func (o *AzBlobAzureBlobStorage) GetStorageAccount() string {
	if o == nil {
		return ""
	}
	return o.StorageAccount
}

type SourceFileUpdateSchemasStorage string

const (
	SourceFileUpdateSchemasStorageS3 SourceFileUpdateSchemasStorage = "S3"
)

func (e SourceFileUpdateSchemasStorage) ToPointer() *SourceFileUpdateSchemasStorage {
	return &e
}

func (e *SourceFileUpdateSchemasStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3":
		*e = SourceFileUpdateSchemasStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateSchemasStorage: %v", v)
	}
}

// SourceFileUpdateS3AmazonWebServices - The storage Provider or Location of the file(s) which should be replicated.
type SourceFileUpdateS3AmazonWebServices struct {
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsAccessKeyID *string `json:"aws_access_key_id,omitempty"`
	// In order to access private Buckets stored on AWS S3, this connector would need credentials with the proper permissions. If accessing publicly available data, this field is not necessary.
	AwsSecretAccessKey *string                        `json:"aws_secret_access_key,omitempty"`
	storage            SourceFileUpdateSchemasStorage `const:"S3" json:"storage"`
}

func (s SourceFileUpdateS3AmazonWebServices) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileUpdateS3AmazonWebServices) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileUpdateS3AmazonWebServices) GetAwsAccessKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccessKeyID
}

func (o *SourceFileUpdateS3AmazonWebServices) GetAwsSecretAccessKey() *string {
	if o == nil {
		return nil
	}
	return o.AwsSecretAccessKey
}

func (o *SourceFileUpdateS3AmazonWebServices) GetStorage() SourceFileUpdateSchemasStorage {
	return SourceFileUpdateSchemasStorageS3
}

type SourceFileUpdateStorage string

const (
	SourceFileUpdateStorageGcs SourceFileUpdateStorage = "GCS"
)

func (e SourceFileUpdateStorage) ToPointer() *SourceFileUpdateStorage {
	return &e
}

func (e *SourceFileUpdateStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS":
		*e = SourceFileUpdateStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFileUpdateStorage: %v", v)
	}
}

// GCSGoogleCloudStorage - The storage Provider or Location of the file(s) which should be replicated.
type GCSGoogleCloudStorage struct {
	// In order to access private Buckets stored on Google Cloud, this connector would need a service account json credentials with the proper permissions as described <a href="https://cloud.google.com/iam/docs/service-accounts" target="_blank">here</a>. Please generate the credentials.json file and copy/paste its content to this field (expecting JSON formats). If accessing publicly available data, this field is not necessary.
	ServiceAccountJSON *string                 `json:"service_account_json,omitempty"`
	storage            SourceFileUpdateStorage `const:"GCS" json:"storage"`
}

func (g GCSGoogleCloudStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GCSGoogleCloudStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GCSGoogleCloudStorage) GetServiceAccountJSON() *string {
	if o == nil {
		return nil
	}
	return o.ServiceAccountJSON
}

func (o *GCSGoogleCloudStorage) GetStorage() SourceFileUpdateStorage {
	return SourceFileUpdateStorageGcs
}

type Storage string

const (
	StorageHTTPS Storage = "HTTPS"
)

func (e Storage) ToPointer() *Storage {
	return &e
}

func (e *Storage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HTTPS":
		*e = Storage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Storage: %v", v)
	}
}

// HTTPSPublicWeb - The storage Provider or Location of the file(s) which should be replicated.
type HTTPSPublicWeb struct {
	storage Storage `const:"HTTPS" json:"storage"`
	// Add User-Agent to request
	UserAgent *bool `default:"false" json:"user_agent"`
}

func (h HTTPSPublicWeb) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HTTPSPublicWeb) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *HTTPSPublicWeb) GetStorage() Storage {
	return StorageHTTPS
}

func (o *HTTPSPublicWeb) GetUserAgent() *bool {
	if o == nil {
		return nil
	}
	return o.UserAgent
}

type StorageProviderType string

const (
	StorageProviderTypeHTTPSPublicWeb                      StorageProviderType = "HTTPS: Public Web"
	StorageProviderTypeGCSGoogleCloudStorage               StorageProviderType = "GCS: Google Cloud Storage"
	StorageProviderTypeSourceFileUpdateS3AmazonWebServices StorageProviderType = "source-file-update_S3: Amazon Web Services"
	StorageProviderTypeAzBlobAzureBlobStorage              StorageProviderType = "AzBlob: Azure Blob Storage"
	StorageProviderTypeSSHSecureShell                      StorageProviderType = "SSH: Secure Shell"
	StorageProviderTypeSCPSecureCopyProtocol               StorageProviderType = "SCP: Secure copy protocol"
	StorageProviderTypeSFTPSecureFileTransferProtocol      StorageProviderType = "SFTP: Secure File Transfer Protocol"
)

type StorageProvider struct {
	HTTPSPublicWeb                      *HTTPSPublicWeb
	GCSGoogleCloudStorage               *GCSGoogleCloudStorage
	SourceFileUpdateS3AmazonWebServices *SourceFileUpdateS3AmazonWebServices
	AzBlobAzureBlobStorage              *AzBlobAzureBlobStorage
	SSHSecureShell                      *SSHSecureShell
	SCPSecureCopyProtocol               *SCPSecureCopyProtocol
	SFTPSecureFileTransferProtocol      *SFTPSecureFileTransferProtocol

	Type StorageProviderType
}

func CreateStorageProviderHTTPSPublicWeb(httpsPublicWeb HTTPSPublicWeb) StorageProvider {
	typ := StorageProviderTypeHTTPSPublicWeb

	return StorageProvider{
		HTTPSPublicWeb: &httpsPublicWeb,
		Type:           typ,
	}
}

func CreateStorageProviderGCSGoogleCloudStorage(gcsGoogleCloudStorage GCSGoogleCloudStorage) StorageProvider {
	typ := StorageProviderTypeGCSGoogleCloudStorage

	return StorageProvider{
		GCSGoogleCloudStorage: &gcsGoogleCloudStorage,
		Type:                  typ,
	}
}

func CreateStorageProviderSourceFileUpdateS3AmazonWebServices(sourceFileUpdateS3AmazonWebServices SourceFileUpdateS3AmazonWebServices) StorageProvider {
	typ := StorageProviderTypeSourceFileUpdateS3AmazonWebServices

	return StorageProvider{
		SourceFileUpdateS3AmazonWebServices: &sourceFileUpdateS3AmazonWebServices,
		Type:                                typ,
	}
}

func CreateStorageProviderAzBlobAzureBlobStorage(azBlobAzureBlobStorage AzBlobAzureBlobStorage) StorageProvider {
	typ := StorageProviderTypeAzBlobAzureBlobStorage

	return StorageProvider{
		AzBlobAzureBlobStorage: &azBlobAzureBlobStorage,
		Type:                   typ,
	}
}

func CreateStorageProviderSSHSecureShell(sshSecureShell SSHSecureShell) StorageProvider {
	typ := StorageProviderTypeSSHSecureShell

	return StorageProvider{
		SSHSecureShell: &sshSecureShell,
		Type:           typ,
	}
}

func CreateStorageProviderSCPSecureCopyProtocol(scpSecureCopyProtocol SCPSecureCopyProtocol) StorageProvider {
	typ := StorageProviderTypeSCPSecureCopyProtocol

	return StorageProvider{
		SCPSecureCopyProtocol: &scpSecureCopyProtocol,
		Type:                  typ,
	}
}

func CreateStorageProviderSFTPSecureFileTransferProtocol(sftpSecureFileTransferProtocol SFTPSecureFileTransferProtocol) StorageProvider {
	typ := StorageProviderTypeSFTPSecureFileTransferProtocol

	return StorageProvider{
		SFTPSecureFileTransferProtocol: &sftpSecureFileTransferProtocol,
		Type:                           typ,
	}
}

func (u *StorageProvider) UnmarshalJSON(data []byte) error {

	httpsPublicWeb := new(HTTPSPublicWeb)
	if err := utils.UnmarshalJSON(data, &httpsPublicWeb, "", true, true); err == nil {
		u.HTTPSPublicWeb = httpsPublicWeb
		u.Type = StorageProviderTypeHTTPSPublicWeb
		return nil
	}

	gcsGoogleCloudStorage := new(GCSGoogleCloudStorage)
	if err := utils.UnmarshalJSON(data, &gcsGoogleCloudStorage, "", true, true); err == nil {
		u.GCSGoogleCloudStorage = gcsGoogleCloudStorage
		u.Type = StorageProviderTypeGCSGoogleCloudStorage
		return nil
	}

	sourceFileUpdateS3AmazonWebServices := new(SourceFileUpdateS3AmazonWebServices)
	if err := utils.UnmarshalJSON(data, &sourceFileUpdateS3AmazonWebServices, "", true, true); err == nil {
		u.SourceFileUpdateS3AmazonWebServices = sourceFileUpdateS3AmazonWebServices
		u.Type = StorageProviderTypeSourceFileUpdateS3AmazonWebServices
		return nil
	}

	azBlobAzureBlobStorage := new(AzBlobAzureBlobStorage)
	if err := utils.UnmarshalJSON(data, &azBlobAzureBlobStorage, "", true, true); err == nil {
		u.AzBlobAzureBlobStorage = azBlobAzureBlobStorage
		u.Type = StorageProviderTypeAzBlobAzureBlobStorage
		return nil
	}

	sshSecureShell := new(SSHSecureShell)
	if err := utils.UnmarshalJSON(data, &sshSecureShell, "", true, true); err == nil {
		u.SSHSecureShell = sshSecureShell
		u.Type = StorageProviderTypeSSHSecureShell
		return nil
	}

	scpSecureCopyProtocol := new(SCPSecureCopyProtocol)
	if err := utils.UnmarshalJSON(data, &scpSecureCopyProtocol, "", true, true); err == nil {
		u.SCPSecureCopyProtocol = scpSecureCopyProtocol
		u.Type = StorageProviderTypeSCPSecureCopyProtocol
		return nil
	}

	sftpSecureFileTransferProtocol := new(SFTPSecureFileTransferProtocol)
	if err := utils.UnmarshalJSON(data, &sftpSecureFileTransferProtocol, "", true, true); err == nil {
		u.SFTPSecureFileTransferProtocol = sftpSecureFileTransferProtocol
		u.Type = StorageProviderTypeSFTPSecureFileTransferProtocol
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u StorageProvider) MarshalJSON() ([]byte, error) {
	if u.HTTPSPublicWeb != nil {
		return utils.MarshalJSON(u.HTTPSPublicWeb, "", true)
	}

	if u.GCSGoogleCloudStorage != nil {
		return utils.MarshalJSON(u.GCSGoogleCloudStorage, "", true)
	}

	if u.SourceFileUpdateS3AmazonWebServices != nil {
		return utils.MarshalJSON(u.SourceFileUpdateS3AmazonWebServices, "", true)
	}

	if u.AzBlobAzureBlobStorage != nil {
		return utils.MarshalJSON(u.AzBlobAzureBlobStorage, "", true)
	}

	if u.SSHSecureShell != nil {
		return utils.MarshalJSON(u.SSHSecureShell, "", true)
	}

	if u.SCPSecureCopyProtocol != nil {
		return utils.MarshalJSON(u.SCPSecureCopyProtocol, "", true)
	}

	if u.SFTPSecureFileTransferProtocol != nil {
		return utils.MarshalJSON(u.SFTPSecureFileTransferProtocol, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceFileUpdate struct {
	// The Name of the final table to replicate this file into (should include letters, numbers dash and underscores only).
	DatasetName string `json:"dataset_name"`
	// The Format of the file which should be replicated (Warning: some formats may be experimental, please refer to the docs).
	Format *FileFormat `default:"csv" json:"format"`
	// The storage Provider or Location of the file(s) which should be replicated.
	Provider StorageProvider `json:"provider"`
	// This should be a string in JSON format. It depends on the chosen file format to provide additional options and tune its behavior.
	ReaderOptions *string `json:"reader_options,omitempty"`
	// The URL path to access the file which should be replicated.
	URL string `json:"url"`
}

func (s SourceFileUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFileUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFileUpdate) GetDatasetName() string {
	if o == nil {
		return ""
	}
	return o.DatasetName
}

func (o *SourceFileUpdate) GetFormat() *FileFormat {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *SourceFileUpdate) GetProvider() StorageProvider {
	if o == nil {
		return StorageProvider{}
	}
	return o.Provider
}

func (o *SourceFileUpdate) GetReaderOptions() *string {
	if o == nil {
		return nil
	}
	return o.ReaderOptions
}

func (o *SourceFileUpdate) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
