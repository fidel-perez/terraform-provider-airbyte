// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationDatabricksSchemasDataSourceDataSourceType string

const (
	DestinationDatabricksSchemasDataSourceDataSourceTypeAzureBlobStorage DestinationDatabricksSchemasDataSourceDataSourceType = "AZURE_BLOB_STORAGE"
)

func (e DestinationDatabricksSchemasDataSourceDataSourceType) ToPointer() *DestinationDatabricksSchemasDataSourceDataSourceType {
	return &e
}

func (e *DestinationDatabricksSchemasDataSourceDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AZURE_BLOB_STORAGE":
		*e = DestinationDatabricksSchemasDataSourceDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksSchemasDataSourceDataSourceType: %v", v)
	}
}

// DestinationDatabricksAzureBlobStorage - Storage on which the delta lake is built.
type DestinationDatabricksAzureBlobStorage struct {
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `default:"blob.core.windows.net" json:"azure_blob_storage_endpoint_domain_name"`
	// Shared access signature (SAS) token to grant limited access to objects in your storage account.
	AzureBlobStorageSasToken string                                               `json:"azure_blob_storage_sas_token"`
	dataSourceType           DestinationDatabricksSchemasDataSourceDataSourceType `const:"AZURE_BLOB_STORAGE" json:"data_source_type"`
}

func (d DestinationDatabricksAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksAzureBlobStorage) GetAzureBlobStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountName
}

func (o *DestinationDatabricksAzureBlobStorage) GetAzureBlobStorageContainerName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageContainerName
}

func (o *DestinationDatabricksAzureBlobStorage) GetAzureBlobStorageEndpointDomainName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageEndpointDomainName
}

func (o *DestinationDatabricksAzureBlobStorage) GetAzureBlobStorageSasToken() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageSasToken
}

func (o *DestinationDatabricksAzureBlobStorage) GetDataSourceType() DestinationDatabricksSchemasDataSourceDataSourceType {
	return DestinationDatabricksSchemasDataSourceDataSourceTypeAzureBlobStorage
}

type DestinationDatabricksSchemasDataSourceType string

const (
	DestinationDatabricksSchemasDataSourceTypeS3Storage DestinationDatabricksSchemasDataSourceType = "S3_STORAGE"
)

func (e DestinationDatabricksSchemasDataSourceType) ToPointer() *DestinationDatabricksSchemasDataSourceType {
	return &e
}

func (e *DestinationDatabricksSchemasDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3_STORAGE":
		*e = DestinationDatabricksSchemasDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksSchemasDataSourceType: %v", v)
	}
}

// DestinationDatabricksS3BucketRegion - The region of the S3 staging bucket to use if utilising a copy strategy.
type DestinationDatabricksS3BucketRegion string

const (
	DestinationDatabricksS3BucketRegionUnknown      DestinationDatabricksS3BucketRegion = ""
	DestinationDatabricksS3BucketRegionUsEast1      DestinationDatabricksS3BucketRegion = "us-east-1"
	DestinationDatabricksS3BucketRegionUsEast2      DestinationDatabricksS3BucketRegion = "us-east-2"
	DestinationDatabricksS3BucketRegionUsWest1      DestinationDatabricksS3BucketRegion = "us-west-1"
	DestinationDatabricksS3BucketRegionUsWest2      DestinationDatabricksS3BucketRegion = "us-west-2"
	DestinationDatabricksS3BucketRegionAfSouth1     DestinationDatabricksS3BucketRegion = "af-south-1"
	DestinationDatabricksS3BucketRegionApEast1      DestinationDatabricksS3BucketRegion = "ap-east-1"
	DestinationDatabricksS3BucketRegionApSouth1     DestinationDatabricksS3BucketRegion = "ap-south-1"
	DestinationDatabricksS3BucketRegionApNortheast1 DestinationDatabricksS3BucketRegion = "ap-northeast-1"
	DestinationDatabricksS3BucketRegionApNortheast2 DestinationDatabricksS3BucketRegion = "ap-northeast-2"
	DestinationDatabricksS3BucketRegionApNortheast3 DestinationDatabricksS3BucketRegion = "ap-northeast-3"
	DestinationDatabricksS3BucketRegionApSoutheast1 DestinationDatabricksS3BucketRegion = "ap-southeast-1"
	DestinationDatabricksS3BucketRegionApSoutheast2 DestinationDatabricksS3BucketRegion = "ap-southeast-2"
	DestinationDatabricksS3BucketRegionCaCentral1   DestinationDatabricksS3BucketRegion = "ca-central-1"
	DestinationDatabricksS3BucketRegionCnNorth1     DestinationDatabricksS3BucketRegion = "cn-north-1"
	DestinationDatabricksS3BucketRegionCnNorthwest1 DestinationDatabricksS3BucketRegion = "cn-northwest-1"
	DestinationDatabricksS3BucketRegionEuCentral1   DestinationDatabricksS3BucketRegion = "eu-central-1"
	DestinationDatabricksS3BucketRegionEuNorth1     DestinationDatabricksS3BucketRegion = "eu-north-1"
	DestinationDatabricksS3BucketRegionEuSouth1     DestinationDatabricksS3BucketRegion = "eu-south-1"
	DestinationDatabricksS3BucketRegionEuWest1      DestinationDatabricksS3BucketRegion = "eu-west-1"
	DestinationDatabricksS3BucketRegionEuWest2      DestinationDatabricksS3BucketRegion = "eu-west-2"
	DestinationDatabricksS3BucketRegionEuWest3      DestinationDatabricksS3BucketRegion = "eu-west-3"
	DestinationDatabricksS3BucketRegionSaEast1      DestinationDatabricksS3BucketRegion = "sa-east-1"
	DestinationDatabricksS3BucketRegionMeSouth1     DestinationDatabricksS3BucketRegion = "me-south-1"
	DestinationDatabricksS3BucketRegionUsGovEast1   DestinationDatabricksS3BucketRegion = "us-gov-east-1"
	DestinationDatabricksS3BucketRegionUsGovWest1   DestinationDatabricksS3BucketRegion = "us-gov-west-1"
)

func (e DestinationDatabricksS3BucketRegion) ToPointer() *DestinationDatabricksS3BucketRegion {
	return &e
}

func (e *DestinationDatabricksS3BucketRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = DestinationDatabricksS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksS3BucketRegion: %v", v)
	}
}

// DestinationDatabricksAmazonS3 - Storage on which the delta lake is built.
type DestinationDatabricksAmazonS3 struct {
	dataSourceType DestinationDatabricksSchemasDataSourceType `const:"S3_STORAGE" json:"data_source_type"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string `json:"file_name_pattern,omitempty"`
	// The Access Key Id granting allow one to access the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket.
	S3AccessKeyID string `json:"s3_access_key_id"`
	// The name of the S3 bucket to use for intermittent staging of the data.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written.
	S3BucketPath string `json:"s3_bucket_path"`
	// The region of the S3 staging bucket to use if utilising a copy strategy.
	S3BucketRegion *DestinationDatabricksS3BucketRegion `default:"" json:"s3_bucket_region"`
	// The corresponding secret to the above access key id.
	S3SecretAccessKey string `json:"s3_secret_access_key"`
}

func (d DestinationDatabricksAmazonS3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksAmazonS3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksAmazonS3) GetDataSourceType() DestinationDatabricksSchemasDataSourceType {
	return DestinationDatabricksSchemasDataSourceTypeS3Storage
}

func (o *DestinationDatabricksAmazonS3) GetFileNamePattern() *string {
	if o == nil {
		return nil
	}
	return o.FileNamePattern
}

func (o *DestinationDatabricksAmazonS3) GetS3AccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.S3AccessKeyID
}

func (o *DestinationDatabricksAmazonS3) GetS3BucketName() string {
	if o == nil {
		return ""
	}
	return o.S3BucketName
}

func (o *DestinationDatabricksAmazonS3) GetS3BucketPath() string {
	if o == nil {
		return ""
	}
	return o.S3BucketPath
}

func (o *DestinationDatabricksAmazonS3) GetS3BucketRegion() *DestinationDatabricksS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *DestinationDatabricksAmazonS3) GetS3SecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.S3SecretAccessKey
}

type DestinationDatabricksDataSourceType string

const (
	DestinationDatabricksDataSourceTypeManagedTablesStorage DestinationDatabricksDataSourceType = "MANAGED_TABLES_STORAGE"
)

func (e DestinationDatabricksDataSourceType) ToPointer() *DestinationDatabricksDataSourceType {
	return &e
}

func (e *DestinationDatabricksDataSourceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MANAGED_TABLES_STORAGE":
		*e = DestinationDatabricksDataSourceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDatabricksDataSourceType: %v", v)
	}
}

// DestinationDatabricksRecommendedManagedTables - Storage on which the delta lake is built.
type DestinationDatabricksRecommendedManagedTables struct {
	dataSourceType DestinationDatabricksDataSourceType `const:"MANAGED_TABLES_STORAGE" json:"data_source_type"`
}

func (d DestinationDatabricksRecommendedManagedTables) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricksRecommendedManagedTables) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricksRecommendedManagedTables) GetDataSourceType() DestinationDatabricksDataSourceType {
	return DestinationDatabricksDataSourceTypeManagedTablesStorage
}

type DestinationDatabricksDataSourceUnionType string

const (
	DestinationDatabricksDataSourceUnionTypeDestinationDatabricksRecommendedManagedTables DestinationDatabricksDataSourceUnionType = "destination-databricks_[Recommended] Managed tables"
	DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAmazonS3                 DestinationDatabricksDataSourceUnionType = "destination-databricks_Amazon S3"
	DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAzureBlobStorage         DestinationDatabricksDataSourceUnionType = "destination-databricks_Azure Blob Storage"
)

type DestinationDatabricksDataSource struct {
	DestinationDatabricksRecommendedManagedTables *DestinationDatabricksRecommendedManagedTables
	DestinationDatabricksAmazonS3                 *DestinationDatabricksAmazonS3
	DestinationDatabricksAzureBlobStorage         *DestinationDatabricksAzureBlobStorage

	Type DestinationDatabricksDataSourceUnionType
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksRecommendedManagedTables(destinationDatabricksRecommendedManagedTables DestinationDatabricksRecommendedManagedTables) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceUnionTypeDestinationDatabricksRecommendedManagedTables

	return DestinationDatabricksDataSource{
		DestinationDatabricksRecommendedManagedTables: &destinationDatabricksRecommendedManagedTables,
		Type: typ,
	}
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksAmazonS3(destinationDatabricksAmazonS3 DestinationDatabricksAmazonS3) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAmazonS3

	return DestinationDatabricksDataSource{
		DestinationDatabricksAmazonS3: &destinationDatabricksAmazonS3,
		Type:                          typ,
	}
}

func CreateDestinationDatabricksDataSourceDestinationDatabricksAzureBlobStorage(destinationDatabricksAzureBlobStorage DestinationDatabricksAzureBlobStorage) DestinationDatabricksDataSource {
	typ := DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAzureBlobStorage

	return DestinationDatabricksDataSource{
		DestinationDatabricksAzureBlobStorage: &destinationDatabricksAzureBlobStorage,
		Type:                                  typ,
	}
}

func (u *DestinationDatabricksDataSource) UnmarshalJSON(data []byte) error {

	destinationDatabricksRecommendedManagedTables := new(DestinationDatabricksRecommendedManagedTables)
	if err := utils.UnmarshalJSON(data, &destinationDatabricksRecommendedManagedTables, "", true, true); err == nil {
		u.DestinationDatabricksRecommendedManagedTables = destinationDatabricksRecommendedManagedTables
		u.Type = DestinationDatabricksDataSourceUnionTypeDestinationDatabricksRecommendedManagedTables
		return nil
	}

	destinationDatabricksAzureBlobStorage := new(DestinationDatabricksAzureBlobStorage)
	if err := utils.UnmarshalJSON(data, &destinationDatabricksAzureBlobStorage, "", true, true); err == nil {
		u.DestinationDatabricksAzureBlobStorage = destinationDatabricksAzureBlobStorage
		u.Type = DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAzureBlobStorage
		return nil
	}

	destinationDatabricksAmazonS3 := new(DestinationDatabricksAmazonS3)
	if err := utils.UnmarshalJSON(data, &destinationDatabricksAmazonS3, "", true, true); err == nil {
		u.DestinationDatabricksAmazonS3 = destinationDatabricksAmazonS3
		u.Type = DestinationDatabricksDataSourceUnionTypeDestinationDatabricksAmazonS3
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationDatabricksDataSource) MarshalJSON() ([]byte, error) {
	if u.DestinationDatabricksRecommendedManagedTables != nil {
		return utils.MarshalJSON(u.DestinationDatabricksRecommendedManagedTables, "", true)
	}

	if u.DestinationDatabricksAmazonS3 != nil {
		return utils.MarshalJSON(u.DestinationDatabricksAmazonS3, "", true)
	}

	if u.DestinationDatabricksAzureBlobStorage != nil {
		return utils.MarshalJSON(u.DestinationDatabricksAzureBlobStorage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type Databricks string

const (
	DatabricksDatabricks Databricks = "databricks"
)

func (e Databricks) ToPointer() *Databricks {
	return &e
}

func (e *Databricks) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "databricks":
		*e = Databricks(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Databricks: %v", v)
	}
}

type DestinationDatabricks struct {
	// You must agree to the Databricks JDBC Driver <a href="https://databricks.com/jdbc-odbc-driver-license">Terms & Conditions</a> to use this connector.
	AcceptTerms *bool `default:"false" json:"accept_terms"`
	// Storage on which the delta lake is built.
	DataSource DestinationDatabricksDataSource `json:"data_source"`
	// The name of the catalog. If not specified otherwise, the "hive_metastore" will be used.
	Database *string `json:"database,omitempty"`
	// Databricks Cluster HTTP Path.
	DatabricksHTTPPath string `json:"databricks_http_path"`
	// Databricks Personal Access Token for making authenticated requests.
	DatabricksPersonalAccessToken string `json:"databricks_personal_access_token"`
	// Databricks Cluster Port.
	DatabricksPort *string `default:"443" json:"databricks_port"`
	// Databricks Cluster Server Hostname.
	DatabricksServerHostname string     `json:"databricks_server_hostname"`
	destinationType          Databricks `const:"databricks" json:"destinationType"`
	// Support schema evolution for all streams. If "false", the connector might fail when a stream's schema changes.
	EnableSchemaEvolution *bool `default:"false" json:"enable_schema_evolution"`
	// Default to 'true'. Switch it to 'false' for debugging purpose.
	PurgeStagingData *bool `default:"true" json:"purge_staging_data"`
	// The default schema tables are written. If not specified otherwise, the "default" will be used.
	Schema *string `default:"default" json:"schema"`
}

func (d DestinationDatabricks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDatabricks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDatabricks) GetAcceptTerms() *bool {
	if o == nil {
		return nil
	}
	return o.AcceptTerms
}

func (o *DestinationDatabricks) GetDataSource() DestinationDatabricksDataSource {
	if o == nil {
		return DestinationDatabricksDataSource{}
	}
	return o.DataSource
}

func (o *DestinationDatabricks) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *DestinationDatabricks) GetDatabricksHTTPPath() string {
	if o == nil {
		return ""
	}
	return o.DatabricksHTTPPath
}

func (o *DestinationDatabricks) GetDatabricksPersonalAccessToken() string {
	if o == nil {
		return ""
	}
	return o.DatabricksPersonalAccessToken
}

func (o *DestinationDatabricks) GetDatabricksPort() *string {
	if o == nil {
		return nil
	}
	return o.DatabricksPort
}

func (o *DestinationDatabricks) GetDatabricksServerHostname() string {
	if o == nil {
		return ""
	}
	return o.DatabricksServerHostname
}

func (o *DestinationDatabricks) GetDestinationType() Databricks {
	return DatabricksDatabricks
}

func (o *DestinationDatabricks) GetEnableSchemaEvolution() *bool {
	if o == nil {
		return nil
	}
	return o.EnableSchemaEvolution
}

func (o *DestinationDatabricks) GetPurgeStagingData() *bool {
	if o == nil {
		return nil
	}
	return o.PurgeStagingData
}

func (o *DestinationDatabricks) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
