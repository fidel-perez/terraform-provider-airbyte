// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

// DestinationBigqueryDenormalizedUpdateDatasetLocation - The location of the dataset. Warning: Changes made after creation will not be applied. The default "US" value is used if not set explicitly. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
type DestinationBigqueryDenormalizedUpdateDatasetLocation string

const (
	DestinationBigqueryDenormalizedUpdateDatasetLocationUs                     DestinationBigqueryDenormalizedUpdateDatasetLocation = "US"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEu                     DestinationBigqueryDenormalizedUpdateDatasetLocation = "EU"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaEast1              DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-east1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaEast2              DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-east2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaNortheast1         DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-northeast1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaNortheast2         DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-northeast2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaNortheast3         DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-northeast3"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaSouth1             DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-south1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaSouth2             DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-south2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaSoutheast1         DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-southeast1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAsiaSoutheast2         DestinationBigqueryDenormalizedUpdateDatasetLocation = "asia-southeast2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAustraliaSoutheast1    DestinationBigqueryDenormalizedUpdateDatasetLocation = "australia-southeast1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationAustraliaSoutheast2    DestinationBigqueryDenormalizedUpdateDatasetLocation = "australia-southeast2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeCentral1         DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-central1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeCentral2         DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-central2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeNorth1           DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-north1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeSouthwest1       DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-southwest1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest1            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest2            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest3            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west3"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest4            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west4"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest6            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west6"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest7            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west7"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest8            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west8"
	DestinationBigqueryDenormalizedUpdateDatasetLocationEuropeWest9            DestinationBigqueryDenormalizedUpdateDatasetLocation = "europe-west9"
	DestinationBigqueryDenormalizedUpdateDatasetLocationMeWest1                DestinationBigqueryDenormalizedUpdateDatasetLocation = "me-west1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationNorthamericaNortheast1 DestinationBigqueryDenormalizedUpdateDatasetLocation = "northamerica-northeast1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationNorthamericaNortheast2 DestinationBigqueryDenormalizedUpdateDatasetLocation = "northamerica-northeast2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationSouthamericaEast1      DestinationBigqueryDenormalizedUpdateDatasetLocation = "southamerica-east1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationSouthamericaWest1      DestinationBigqueryDenormalizedUpdateDatasetLocation = "southamerica-west1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsCentral1             DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-central1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsEast1                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-east1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsEast2                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-east2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsEast3                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-east3"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsEast4                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-east4"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsEast5                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-east5"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsWest1                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-west1"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsWest2                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-west2"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsWest3                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-west3"
	DestinationBigqueryDenormalizedUpdateDatasetLocationUsWest4                DestinationBigqueryDenormalizedUpdateDatasetLocation = "us-west4"
)

func (e DestinationBigqueryDenormalizedUpdateDatasetLocation) ToPointer() *DestinationBigqueryDenormalizedUpdateDatasetLocation {
	return &e
}

func (e *DestinationBigqueryDenormalizedUpdateDatasetLocation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		fallthrough
	case "asia-east1":
		fallthrough
	case "asia-east2":
		fallthrough
	case "asia-northeast1":
		fallthrough
	case "asia-northeast2":
		fallthrough
	case "asia-northeast3":
		fallthrough
	case "asia-south1":
		fallthrough
	case "asia-south2":
		fallthrough
	case "asia-southeast1":
		fallthrough
	case "asia-southeast2":
		fallthrough
	case "australia-southeast1":
		fallthrough
	case "australia-southeast2":
		fallthrough
	case "europe-central1":
		fallthrough
	case "europe-central2":
		fallthrough
	case "europe-north1":
		fallthrough
	case "europe-southwest1":
		fallthrough
	case "europe-west1":
		fallthrough
	case "europe-west2":
		fallthrough
	case "europe-west3":
		fallthrough
	case "europe-west4":
		fallthrough
	case "europe-west6":
		fallthrough
	case "europe-west7":
		fallthrough
	case "europe-west8":
		fallthrough
	case "europe-west9":
		fallthrough
	case "me-west1":
		fallthrough
	case "northamerica-northeast1":
		fallthrough
	case "northamerica-northeast2":
		fallthrough
	case "southamerica-east1":
		fallthrough
	case "southamerica-west1":
		fallthrough
	case "us-central1":
		fallthrough
	case "us-east1":
		fallthrough
	case "us-east2":
		fallthrough
	case "us-east3":
		fallthrough
	case "us-east4":
		fallthrough
	case "us-east5":
		fallthrough
	case "us-west1":
		fallthrough
	case "us-west2":
		fallthrough
	case "us-west3":
		fallthrough
	case "us-west4":
		*e = DestinationBigqueryDenormalizedUpdateDatasetLocation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDenormalizedUpdateDatasetLocation: %v", v)
	}
}

type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialTypeHmacKey DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType = "HMAC_KEY"
)

func (e DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType) ToPointer() *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType {
	return &e
}

func (e *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HMAC_KEY":
		*e = DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType: %v", v)
	}
}

// DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey - An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey struct {
	credentialType DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType `const:"HMAC_KEY" json:"credential_type"`
	// HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.
	HmacKeyAccessID string `json:"hmac_key_access_id"`
	// The corresponding secret for the access ID. It is a 40-character base-64 encoded string.
	HmacKeySecret string `json:"hmac_key_secret"`
}

func (d DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) GetCredentialType() DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialType {
	return DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKeyCredentialTypeHmacKey
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) GetHmacKeyAccessID() string {
	if o == nil {
		return ""
	}
	return o.HmacKeyAccessID
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) GetHmacKeySecret() string {
	if o == nil {
		return ""
	}
	return o.HmacKeySecret
}

type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialType string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialType = "destination-bigquery-denormalized-update_Loading Method_GCS Staging_Credential_HMAC key"
)

type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential struct {
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey

	Type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialType
}

func CreateDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey(destinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey) DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential {
	typ := DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey

	return DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential{
		DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey: &destinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey,
		Type: typ,
	}
}

func (u *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential) UnmarshalJSON(data []byte) error {

	destinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey := new(DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey)
	if err := utils.UnmarshalJSON(data, &destinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey, "", true, true); err == nil {
		u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey = destinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey
		u.Type = DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey != nil {
		return utils.MarshalJSON(u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredentialHMACKey, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing - This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessingDeleteAllTmpFilesFromGcs DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing = "Delete all tmp files from GCS"
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessingKeepAllTmpFilesInGcs     DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing = "Keep all tmp files in GCS"
)

func (e DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing) ToPointer() *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing {
	return &e
}

func (e *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Delete all tmp files from GCS":
		fallthrough
	case "Keep all tmp files in GCS":
		*e = DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing: %v", v)
	}
}

type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethodGcsStaging DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod = "GCS Staging"
)

func (e DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod) ToPointer() *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod {
	return &e
}

func (e *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GCS Staging":
		*e = DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod: %v", v)
	}
}

// DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging - Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
type DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging struct {
	// An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>.
	Credential DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential `json:"credential"`
	// Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects
	FileBufferCount *int64 `default:"10" json:"file_buffer_count"`
	// The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.
	GcsBucketName string `json:"gcs_bucket_name"`
	// Directory under the GCS bucket where data will be written. Read more <a href="https://cloud.google.com/storage/docs/locations">here</a>.
	GcsBucketPath string `json:"gcs_bucket_path"`
	// This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.
	KeepFilesInGcsBucket *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing `default:"Delete all tmp files from GCS" json:"keep_files_in_gcs-bucket"`
	method               DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod                          `const:"GCS Staging" json:"method"`
}

func (d DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetCredential() DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential {
	if o == nil {
		return DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingCredential{}
	}
	return o.Credential
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetFileBufferCount() *int64 {
	if o == nil {
		return nil
	}
	return o.FileBufferCount
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetGcsBucketName() string {
	if o == nil {
		return ""
	}
	return o.GcsBucketName
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetGcsBucketPath() string {
	if o == nil {
		return ""
	}
	return o.GcsBucketPath
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetKeepFilesInGcsBucket() *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingGCSTmpFilesAfterwardProcessing {
	if o == nil {
		return nil
	}
	return o.KeepFilesInGcsBucket
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) GetMethod() DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethod {
	return DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStagingMethodGcsStaging
}

type DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethodStandard DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod = "Standard"
)

func (e DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod) ToPointer() *DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod {
	return &e
}

func (e *DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		*e = DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod: %v", v)
	}
}

// DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts - Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
type DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts struct {
	method DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod `const:"Standard" json:"method"`
}

func (d DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts) GetMethod() DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethod {
	return DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInsertsMethodStandard
}

type DestinationBigqueryDenormalizedUpdateLoadingMethodType string

const (
	DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts DestinationBigqueryDenormalizedUpdateLoadingMethodType = "destination-bigquery-denormalized-update_Loading Method_Standard Inserts"
	DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging      DestinationBigqueryDenormalizedUpdateLoadingMethodType = "destination-bigquery-denormalized-update_Loading Method_GCS Staging"
)

type DestinationBigqueryDenormalizedUpdateLoadingMethod struct {
	DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts *DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts
	DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging      *DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging

	Type DestinationBigqueryDenormalizedUpdateLoadingMethodType
}

func CreateDestinationBigqueryDenormalizedUpdateLoadingMethodDestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts(destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts) DestinationBigqueryDenormalizedUpdateLoadingMethod {
	typ := DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts

	return DestinationBigqueryDenormalizedUpdateLoadingMethod{
		DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts: &destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts,
		Type: typ,
	}
}

func CreateDestinationBigqueryDenormalizedUpdateLoadingMethodDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging(destinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging) DestinationBigqueryDenormalizedUpdateLoadingMethod {
	typ := DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging

	return DestinationBigqueryDenormalizedUpdateLoadingMethod{
		DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging: &destinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging,
		Type: typ,
	}
}

func (u *DestinationBigqueryDenormalizedUpdateLoadingMethod) UnmarshalJSON(data []byte) error {

	destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts := new(DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts)
	if err := utils.UnmarshalJSON(data, &destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts, "", true, true); err == nil {
		u.DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts = destinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts
		u.Type = DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts
		return nil
	}

	destinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging := new(DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging)
	if err := utils.UnmarshalJSON(data, &destinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging, "", true, true); err == nil {
		u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging = destinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging
		u.Type = DestinationBigqueryDenormalizedUpdateLoadingMethodTypeDestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationBigqueryDenormalizedUpdateLoadingMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts != nil {
		return utils.MarshalJSON(u.DestinationBigqueryDenormalizedUpdateLoadingMethodStandardInserts, "", true)
	}

	if u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging != nil {
		return utils.MarshalJSON(u.DestinationBigqueryDenormalizedUpdateLoadingMethodGCSStaging, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationBigqueryDenormalizedUpdate struct {
	// Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more <a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html">here</a>.
	BigQueryClientBufferSizeMb *int64 `default:"15" json:"big_query_client_buffer_size_mb"`
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
	CredentialsJSON *string `json:"credentials_json,omitempty"`
	// The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more <a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset">here</a>.
	DatasetID string `json:"dataset_id"`
	// The location of the dataset. Warning: Changes made after creation will not be applied. The default "US" value is used if not set explicitly. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
	DatasetLocation *DestinationBigqueryDenormalizedUpdateDatasetLocation `default:"US" json:"dataset_location"`
	// Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>.
	LoadingMethod *DestinationBigqueryDenormalizedUpdateLoadingMethod `json:"loading_method,omitempty"`
	// The GCP project ID for the project containing the target BigQuery dataset. Read more <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">here</a>.
	ProjectID string `json:"project_id"`
}

func (d DestinationBigqueryDenormalizedUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationBigqueryDenormalizedUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationBigqueryDenormalizedUpdate) GetBigQueryClientBufferSizeMb() *int64 {
	if o == nil {
		return nil
	}
	return o.BigQueryClientBufferSizeMb
}

func (o *DestinationBigqueryDenormalizedUpdate) GetCredentialsJSON() *string {
	if o == nil {
		return nil
	}
	return o.CredentialsJSON
}

func (o *DestinationBigqueryDenormalizedUpdate) GetDatasetID() string {
	if o == nil {
		return ""
	}
	return o.DatasetID
}

func (o *DestinationBigqueryDenormalizedUpdate) GetDatasetLocation() *DestinationBigqueryDenormalizedUpdateDatasetLocation {
	if o == nil {
		return nil
	}
	return o.DatasetLocation
}

func (o *DestinationBigqueryDenormalizedUpdate) GetLoadingMethod() *DestinationBigqueryDenormalizedUpdateLoadingMethod {
	if o == nil {
		return nil
	}
	return o.LoadingMethod
}

func (o *DestinationBigqueryDenormalizedUpdate) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}
