// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType string

const (
	SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatTypeJsonl SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType = "JSONL"
)

func (e SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType) ToPointer() *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType {
	return &e
}

func (e *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType: %v", v)
	}
}

// SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON - Input data format
type SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON struct {
	formatType SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType `const:"JSONL" json:"format_type"`
}

func (s SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON) GetFormatType() SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatType {
	return SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSONFormatTypeJsonl
}

type SourceAzureBlobStorageInputFormatType string

const (
	SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON SourceAzureBlobStorageInputFormatType = "source-azure-blob-storage_Input Format_JSON Lines: newline-delimited JSON"
)

type SourceAzureBlobStorageInputFormat struct {
	SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON *SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON

	Type SourceAzureBlobStorageInputFormatType
}

func CreateSourceAzureBlobStorageInputFormatSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON(sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON) SourceAzureBlobStorageInputFormat {
	typ := SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON

	return SourceAzureBlobStorageInputFormat{
		SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON: &sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *SourceAzureBlobStorageInputFormat) UnmarshalJSON(data []byte) error {

	sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON := new(SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON)
	if err := utils.UnmarshalJSON(data, &sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON = sourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON
		u.Type = SourceAzureBlobStorageInputFormatTypeSourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceAzureBlobStorageInputFormat) MarshalJSON() ([]byte, error) {
	if u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.SourceAzureBlobStorageInputFormatJSONLinesNewlineDelimitedJSON, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceAzureBlobStorageAzureBlobStorage string

const (
	SourceAzureBlobStorageAzureBlobStorageAzureBlobStorage SourceAzureBlobStorageAzureBlobStorage = "azure-blob-storage"
)

func (e SourceAzureBlobStorageAzureBlobStorage) ToPointer() *SourceAzureBlobStorageAzureBlobStorage {
	return &e
}

func (e *SourceAzureBlobStorageAzureBlobStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure-blob-storage":
		*e = SourceAzureBlobStorageAzureBlobStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAzureBlobStorageAzureBlobStorage: %v", v)
	}
}

type SourceAzureBlobStorage struct {
	// The Azure blob storage account key.
	AzureBlobStorageAccountKey string `json:"azure_blob_storage_account_key"`
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The Azure blob storage prefix to be applied
	AzureBlobStorageBlobsPrefix *string `json:"azure_blob_storage_blobs_prefix,omitempty"`
	// The name of the Azure blob storage container.
	AzureBlobStorageContainerName string `json:"azure_blob_storage_container_name"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpoint *string `default:"blob.core.windows.net" json:"azure_blob_storage_endpoint"`
	// The Azure blob storage blobs to scan for inferring the schema, useful on large amounts of data with consistent structure
	AzureBlobStorageSchemaInferenceLimit *int64 `json:"azure_blob_storage_schema_inference_limit,omitempty"`
	// Input data format
	Format     SourceAzureBlobStorageInputFormat      `json:"format"`
	sourceType SourceAzureBlobStorageAzureBlobStorage `const:"azure-blob-storage" json:"sourceType"`
}

func (s SourceAzureBlobStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAzureBlobStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageAccountKey() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountKey
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountName
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageBlobsPrefix() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageBlobsPrefix
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageContainerName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageContainerName
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageEndpoint
}

func (o *SourceAzureBlobStorage) GetAzureBlobStorageSchemaInferenceLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageSchemaInferenceLimit
}

func (o *SourceAzureBlobStorage) GetFormat() SourceAzureBlobStorageInputFormat {
	if o == nil {
		return SourceAzureBlobStorageInputFormat{}
	}
	return o.Format
}

func (o *SourceAzureBlobStorage) GetSourceType() SourceAzureBlobStorageAzureBlobStorage {
	return SourceAzureBlobStorageAzureBlobStorageAzureBlobStorage
}
