// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// DestinationAwsDatalakeUpdateCredentialsTitle - Name of the credentials
type DestinationAwsDatalakeUpdateCredentialsTitle string

const (
	DestinationAwsDatalakeUpdateCredentialsTitleIamUser DestinationAwsDatalakeUpdateCredentialsTitle = "IAM User"
)

func (e DestinationAwsDatalakeUpdateCredentialsTitle) ToPointer() *DestinationAwsDatalakeUpdateCredentialsTitle {
	return &e
}

func (e *DestinationAwsDatalakeUpdateCredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM User":
		*e = DestinationAwsDatalakeUpdateCredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeUpdateCredentialsTitle: %v", v)
	}
}

// IAMUser - Choose How to Authenticate to AWS.
type IAMUser struct {
	// AWS User Access Key Id
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// Secret Access Key
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// Name of the credentials
	credentialsTitle *DestinationAwsDatalakeUpdateCredentialsTitle `const:"IAM User" json:"credentials_title"`
}

func (i IAMUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IAMUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *IAMUser) GetAwsAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AwsAccessKeyID
}

func (o *IAMUser) GetAwsSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.AwsSecretAccessKey
}

func (o *IAMUser) GetCredentialsTitle() *DestinationAwsDatalakeUpdateCredentialsTitle {
	return DestinationAwsDatalakeUpdateCredentialsTitleIamUser.ToPointer()
}

// CredentialsTitle - Name of the credentials
type CredentialsTitle string

const (
	CredentialsTitleIamRole CredentialsTitle = "IAM Role"
)

func (e CredentialsTitle) ToPointer() *CredentialsTitle {
	return &e
}

func (e *CredentialsTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IAM Role":
		*e = CredentialsTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CredentialsTitle: %v", v)
	}
}

// IAMRole - Choose How to Authenticate to AWS.
type IAMRole struct {
	// Name of the credentials
	credentialsTitle *CredentialsTitle `const:"IAM Role" json:"credentials_title"`
	// Will assume this role to write data to s3
	RoleArn string `json:"role_arn"`
}

func (i IAMRole) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *IAMRole) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *IAMRole) GetCredentialsTitle() *CredentialsTitle {
	return CredentialsTitleIamRole.ToPointer()
}

func (o *IAMRole) GetRoleArn() string {
	if o == nil {
		return ""
	}
	return o.RoleArn
}

type AuthenticationModeType string

const (
	AuthenticationModeTypeIAMRole AuthenticationModeType = "IAM Role"
	AuthenticationModeTypeIAMUser AuthenticationModeType = "IAM User"
)

type AuthenticationMode struct {
	IAMRole *IAMRole
	IAMUser *IAMUser

	Type AuthenticationModeType
}

func CreateAuthenticationModeIAMRole(iamRole IAMRole) AuthenticationMode {
	typ := AuthenticationModeTypeIAMRole

	return AuthenticationMode{
		IAMRole: &iamRole,
		Type:    typ,
	}
}

func CreateAuthenticationModeIAMUser(iamUser IAMUser) AuthenticationMode {
	typ := AuthenticationModeTypeIAMUser

	return AuthenticationMode{
		IAMUser: &iamUser,
		Type:    typ,
	}
}

func (u *AuthenticationMode) UnmarshalJSON(data []byte) error {

	iamRole := new(IAMRole)
	if err := utils.UnmarshalJSON(data, &iamRole, "", true, true); err == nil {
		u.IAMRole = iamRole
		u.Type = AuthenticationModeTypeIAMRole
		return nil
	}

	iamUser := new(IAMUser)
	if err := utils.UnmarshalJSON(data, &iamUser, "", true, true); err == nil {
		u.IAMUser = iamUser
		u.Type = AuthenticationModeTypeIAMUser
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AuthenticationMode) MarshalJSON() ([]byte, error) {
	if u.IAMRole != nil {
		return utils.MarshalJSON(u.IAMRole, "", true)
	}

	if u.IAMUser != nil {
		return utils.MarshalJSON(u.IAMUser, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationAwsDatalakeUpdateCompressionCodecOptional - The compression algorithm used to compress data.
type DestinationAwsDatalakeUpdateCompressionCodecOptional string

const (
	DestinationAwsDatalakeUpdateCompressionCodecOptionalUncompressed DestinationAwsDatalakeUpdateCompressionCodecOptional = "UNCOMPRESSED"
	DestinationAwsDatalakeUpdateCompressionCodecOptionalSnappy       DestinationAwsDatalakeUpdateCompressionCodecOptional = "SNAPPY"
	DestinationAwsDatalakeUpdateCompressionCodecOptionalGzip         DestinationAwsDatalakeUpdateCompressionCodecOptional = "GZIP"
	DestinationAwsDatalakeUpdateCompressionCodecOptionalZstd         DestinationAwsDatalakeUpdateCompressionCodecOptional = "ZSTD"
)

func (e DestinationAwsDatalakeUpdateCompressionCodecOptional) ToPointer() *DestinationAwsDatalakeUpdateCompressionCodecOptional {
	return &e
}

func (e *DestinationAwsDatalakeUpdateCompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "SNAPPY":
		fallthrough
	case "GZIP":
		fallthrough
	case "ZSTD":
		*e = DestinationAwsDatalakeUpdateCompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeUpdateCompressionCodecOptional: %v", v)
	}
}

type DestinationAwsDatalakeUpdateFormatTypeWildcard string

const (
	DestinationAwsDatalakeUpdateFormatTypeWildcardParquet DestinationAwsDatalakeUpdateFormatTypeWildcard = "Parquet"
)

func (e DestinationAwsDatalakeUpdateFormatTypeWildcard) ToPointer() *DestinationAwsDatalakeUpdateFormatTypeWildcard {
	return &e
}

func (e *DestinationAwsDatalakeUpdateFormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Parquet":
		*e = DestinationAwsDatalakeUpdateFormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAwsDatalakeUpdateFormatTypeWildcard: %v", v)
	}
}

// ParquetColumnarStorage - Format of the data output.
type ParquetColumnarStorage struct {
	// The compression algorithm used to compress data.
	CompressionCodec *DestinationAwsDatalakeUpdateCompressionCodecOptional `default:"SNAPPY" json:"compression_codec"`
	FormatType       *DestinationAwsDatalakeUpdateFormatTypeWildcard       `default:"Parquet" json:"format_type"`
}

func (p ParquetColumnarStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *ParquetColumnarStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ParquetColumnarStorage) GetCompressionCodec() *DestinationAwsDatalakeUpdateCompressionCodecOptional {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *ParquetColumnarStorage) GetFormatType() *DestinationAwsDatalakeUpdateFormatTypeWildcard {
	if o == nil {
		return nil
	}
	return o.FormatType
}

// CompressionCodecOptional - The compression algorithm used to compress data.
type CompressionCodecOptional string

const (
	CompressionCodecOptionalUncompressed CompressionCodecOptional = "UNCOMPRESSED"
	CompressionCodecOptionalGzip         CompressionCodecOptional = "GZIP"
)

func (e CompressionCodecOptional) ToPointer() *CompressionCodecOptional {
	return &e
}

func (e *CompressionCodecOptional) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNCOMPRESSED":
		fallthrough
	case "GZIP":
		*e = CompressionCodecOptional(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CompressionCodecOptional: %v", v)
	}
}

type FormatTypeWildcard string

const (
	FormatTypeWildcardJsonl FormatTypeWildcard = "JSONL"
)

func (e FormatTypeWildcard) ToPointer() *FormatTypeWildcard {
	return &e
}

func (e *FormatTypeWildcard) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = FormatTypeWildcard(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FormatTypeWildcard: %v", v)
	}
}

// JSONLinesNewlineDelimitedJSON - Format of the data output.
type JSONLinesNewlineDelimitedJSON struct {
	// The compression algorithm used to compress data.
	CompressionCodec *CompressionCodecOptional `default:"UNCOMPRESSED" json:"compression_codec"`
	FormatType       *FormatTypeWildcard       `default:"JSONL" json:"format_type"`
}

func (j JSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *JSONLinesNewlineDelimitedJSON) GetCompressionCodec() *CompressionCodecOptional {
	if o == nil {
		return nil
	}
	return o.CompressionCodec
}

func (o *JSONLinesNewlineDelimitedJSON) GetFormatType() *FormatTypeWildcard {
	if o == nil {
		return nil
	}
	return o.FormatType
}

type OutputFormatWildcardType string

const (
	OutputFormatWildcardTypeJSONLinesNewlineDelimitedJSON OutputFormatWildcardType = "JSON Lines: Newline-delimited JSON"
	OutputFormatWildcardTypeParquetColumnarStorage        OutputFormatWildcardType = "Parquet: Columnar Storage"
)

type OutputFormatWildcard struct {
	JSONLinesNewlineDelimitedJSON *JSONLinesNewlineDelimitedJSON
	ParquetColumnarStorage        *ParquetColumnarStorage

	Type OutputFormatWildcardType
}

func CreateOutputFormatWildcardJSONLinesNewlineDelimitedJSON(jsonLinesNewlineDelimitedJSON JSONLinesNewlineDelimitedJSON) OutputFormatWildcard {
	typ := OutputFormatWildcardTypeJSONLinesNewlineDelimitedJSON

	return OutputFormatWildcard{
		JSONLinesNewlineDelimitedJSON: &jsonLinesNewlineDelimitedJSON,
		Type:                          typ,
	}
}

func CreateOutputFormatWildcardParquetColumnarStorage(parquetColumnarStorage ParquetColumnarStorage) OutputFormatWildcard {
	typ := OutputFormatWildcardTypeParquetColumnarStorage

	return OutputFormatWildcard{
		ParquetColumnarStorage: &parquetColumnarStorage,
		Type:                   typ,
	}
}

func (u *OutputFormatWildcard) UnmarshalJSON(data []byte) error {

	jsonLinesNewlineDelimitedJSON := new(JSONLinesNewlineDelimitedJSON)
	if err := utils.UnmarshalJSON(data, &jsonLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.JSONLinesNewlineDelimitedJSON = jsonLinesNewlineDelimitedJSON
		u.Type = OutputFormatWildcardTypeJSONLinesNewlineDelimitedJSON
		return nil
	}

	parquetColumnarStorage := new(ParquetColumnarStorage)
	if err := utils.UnmarshalJSON(data, &parquetColumnarStorage, "", true, true); err == nil {
		u.ParquetColumnarStorage = parquetColumnarStorage
		u.Type = OutputFormatWildcardTypeParquetColumnarStorage
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OutputFormatWildcard) MarshalJSON() ([]byte, error) {
	if u.JSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.JSONLinesNewlineDelimitedJSON, "", true)
	}

	if u.ParquetColumnarStorage != nil {
		return utils.MarshalJSON(u.ParquetColumnarStorage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// ChooseHowToPartitionData - Partition data by cursor fields when a cursor field is a date
type ChooseHowToPartitionData string

const (
	ChooseHowToPartitionDataNoPartitioning ChooseHowToPartitionData = "NO PARTITIONING"
	ChooseHowToPartitionDataDate           ChooseHowToPartitionData = "DATE"
	ChooseHowToPartitionDataYear           ChooseHowToPartitionData = "YEAR"
	ChooseHowToPartitionDataMonth          ChooseHowToPartitionData = "MONTH"
	ChooseHowToPartitionDataDay            ChooseHowToPartitionData = "DAY"
	ChooseHowToPartitionDataYearMonth      ChooseHowToPartitionData = "YEAR/MONTH"
	ChooseHowToPartitionDataYearMonthDay   ChooseHowToPartitionData = "YEAR/MONTH/DAY"
)

func (e ChooseHowToPartitionData) ToPointer() *ChooseHowToPartitionData {
	return &e
}

func (e *ChooseHowToPartitionData) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO PARTITIONING":
		fallthrough
	case "DATE":
		fallthrough
	case "YEAR":
		fallthrough
	case "MONTH":
		fallthrough
	case "DAY":
		fallthrough
	case "YEAR/MONTH":
		fallthrough
	case "YEAR/MONTH/DAY":
		*e = ChooseHowToPartitionData(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ChooseHowToPartitionData: %v", v)
	}
}

// S3BucketRegion - The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
type S3BucketRegion string

const (
	S3BucketRegionUnknown      S3BucketRegion = ""
	S3BucketRegionUsEast1      S3BucketRegion = "us-east-1"
	S3BucketRegionUsEast2      S3BucketRegion = "us-east-2"
	S3BucketRegionUsWest1      S3BucketRegion = "us-west-1"
	S3BucketRegionUsWest2      S3BucketRegion = "us-west-2"
	S3BucketRegionAfSouth1     S3BucketRegion = "af-south-1"
	S3BucketRegionApEast1      S3BucketRegion = "ap-east-1"
	S3BucketRegionApSouth1     S3BucketRegion = "ap-south-1"
	S3BucketRegionApNortheast1 S3BucketRegion = "ap-northeast-1"
	S3BucketRegionApNortheast2 S3BucketRegion = "ap-northeast-2"
	S3BucketRegionApNortheast3 S3BucketRegion = "ap-northeast-3"
	S3BucketRegionApSoutheast1 S3BucketRegion = "ap-southeast-1"
	S3BucketRegionApSoutheast2 S3BucketRegion = "ap-southeast-2"
	S3BucketRegionCaCentral1   S3BucketRegion = "ca-central-1"
	S3BucketRegionCnNorth1     S3BucketRegion = "cn-north-1"
	S3BucketRegionCnNorthwest1 S3BucketRegion = "cn-northwest-1"
	S3BucketRegionEuCentral1   S3BucketRegion = "eu-central-1"
	S3BucketRegionEuNorth1     S3BucketRegion = "eu-north-1"
	S3BucketRegionEuSouth1     S3BucketRegion = "eu-south-1"
	S3BucketRegionEuWest1      S3BucketRegion = "eu-west-1"
	S3BucketRegionEuWest2      S3BucketRegion = "eu-west-2"
	S3BucketRegionEuWest3      S3BucketRegion = "eu-west-3"
	S3BucketRegionSaEast1      S3BucketRegion = "sa-east-1"
	S3BucketRegionMeSouth1     S3BucketRegion = "me-south-1"
	S3BucketRegionUsGovEast1   S3BucketRegion = "us-gov-east-1"
	S3BucketRegionUsGovWest1   S3BucketRegion = "us-gov-west-1"
)

func (e S3BucketRegion) ToPointer() *S3BucketRegion {
	return &e
}

func (e *S3BucketRegion) UnmarshalJSON(data []byte) error {
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
		*e = S3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for S3BucketRegion: %v", v)
	}
}

type DestinationAwsDatalakeUpdate struct {
	// target aws account id
	AwsAccountID *string `json:"aws_account_id,omitempty"`
	// The name of the S3 bucket. Read more <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-bucket-overview.html">here</a>.
	BucketName string `json:"bucket_name"`
	// S3 prefix
	BucketPrefix *string `json:"bucket_prefix,omitempty"`
	// Choose How to Authenticate to AWS.
	Credentials AuthenticationMode `json:"credentials"`
	// Format of the data output.
	Format *OutputFormatWildcard `json:"format,omitempty"`
	// Cast float/double as decimal(38,18). This can help achieve higher accuracy and represent numbers correctly as received from the source.
	GlueCatalogFloatAsDecimal *bool `default:"false" json:"glue_catalog_float_as_decimal"`
	// Add a default tag key to databases created by this destination
	LakeformationDatabaseDefaultTagKey *string `json:"lakeformation_database_default_tag_key,omitempty"`
	// Add default values for the `Tag Key` to databases created by this destination. Comma separate for multiple values.
	LakeformationDatabaseDefaultTagValues *string `json:"lakeformation_database_default_tag_values,omitempty"`
	// The default database this destination will use to create tables in per stream. Can be changed per connection by customizing the namespace.
	LakeformationDatabaseName string `json:"lakeformation_database_name"`
	// Whether to create tables as LF governed tables.
	LakeformationGovernedTables *bool `default:"false" json:"lakeformation_governed_tables"`
	// Partition data by cursor fields when a cursor field is a date
	Partitioning *ChooseHowToPartitionData `default:"NO PARTITIONING" json:"partitioning"`
	// The region of the S3 bucket. See <a href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions">here</a> for all region codes.
	Region *S3BucketRegion `default:"" json:"region"`
}

func (d DestinationAwsDatalakeUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAwsDatalakeUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAwsDatalakeUpdate) GetAwsAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AwsAccountID
}

func (o *DestinationAwsDatalakeUpdate) GetBucketName() string {
	if o == nil {
		return ""
	}
	return o.BucketName
}

func (o *DestinationAwsDatalakeUpdate) GetBucketPrefix() *string {
	if o == nil {
		return nil
	}
	return o.BucketPrefix
}

func (o *DestinationAwsDatalakeUpdate) GetCredentials() AuthenticationMode {
	if o == nil {
		return AuthenticationMode{}
	}
	return o.Credentials
}

func (o *DestinationAwsDatalakeUpdate) GetFormat() *OutputFormatWildcard {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *DestinationAwsDatalakeUpdate) GetGlueCatalogFloatAsDecimal() *bool {
	if o == nil {
		return nil
	}
	return o.GlueCatalogFloatAsDecimal
}

func (o *DestinationAwsDatalakeUpdate) GetLakeformationDatabaseDefaultTagKey() *string {
	if o == nil {
		return nil
	}
	return o.LakeformationDatabaseDefaultTagKey
}

func (o *DestinationAwsDatalakeUpdate) GetLakeformationDatabaseDefaultTagValues() *string {
	if o == nil {
		return nil
	}
	return o.LakeformationDatabaseDefaultTagValues
}

func (o *DestinationAwsDatalakeUpdate) GetLakeformationDatabaseName() string {
	if o == nil {
		return ""
	}
	return o.LakeformationDatabaseName
}

func (o *DestinationAwsDatalakeUpdate) GetLakeformationGovernedTables() *bool {
	if o == nil {
		return nil
	}
	return o.LakeformationGovernedTables
}

func (o *DestinationAwsDatalakeUpdate) GetPartitioning() *ChooseHowToPartitionData {
	if o == nil {
		return nil
	}
	return o.Partitioning
}

func (o *DestinationAwsDatalakeUpdate) GetRegion() *S3BucketRegion {
	if o == nil {
		return nil
	}
	return o.Region
}
