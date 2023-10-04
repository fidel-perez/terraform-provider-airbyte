// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType string

const (
	DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeJsonl DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType = "JSONL"
)

func (e DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType) ToPointer() *DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType {
	return &e
}

func (e *DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType: %v", v)
	}
}

// DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON - Output data format
type DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON struct {
	formatType DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType `const:"JSONL" json:"format_type"`
}

func (d DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON) GetFormatType() DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType {
	return DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatTypeJsonl
}

// DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening - Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
type DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening string

const (
	DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlatteningNoFlattening        DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening = "No flattening"
	DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlatteningRootLevelFlattening DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening = "Root level flattening"
)

func (e DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening) ToPointer() *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening {
	return &e
}

func (e *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No flattening":
		fallthrough
	case "Root level flattening":
		*e = DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening: %v", v)
	}
}

type DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType string

const (
	DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatTypeCsv DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType = "CSV"
)

func (e DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType) ToPointer() *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType {
	return &e
}

func (e *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CSV":
		*e = DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType: %v", v)
	}
}

// DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues - Output data format
type DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues struct {
	// Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
	Flattening *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening `default:"No flattening" json:"flattening"`
	formatType DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType               `const:"CSV" json:"format_type"`
}

func (d DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues) GetFlattening() *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening {
	if o == nil {
		return nil
	}
	return o.Flattening
}

func (o *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues) GetFormatType() DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType {
	return DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatTypeCsv
}

type DestinationAzureBlobStorageUpdateOutputFormatType string

const (
	DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues       DestinationAzureBlobStorageUpdateOutputFormatType = "destination-azure-blob-storage-update_Output Format_CSV: Comma-Separated Values"
	DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageUpdateOutputFormatType = "destination-azure-blob-storage-update_Output Format_JSON Lines: newline-delimited JSON"
)

type DestinationAzureBlobStorageUpdateOutputFormat struct {
	DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues       *DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues
	DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON *DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON

	Type DestinationAzureBlobStorageUpdateOutputFormatType
}

func CreateDestinationAzureBlobStorageUpdateOutputFormatDestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues(destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues) DestinationAzureBlobStorageUpdateOutputFormat {
	typ := DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues

	return DestinationAzureBlobStorageUpdateOutputFormat{
		DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues: &destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues,
		Type: typ,
	}
}

func CreateDestinationAzureBlobStorageUpdateOutputFormatDestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON(destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON) DestinationAzureBlobStorageUpdateOutputFormat {
	typ := DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON

	return DestinationAzureBlobStorageUpdateOutputFormat{
		DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON: &destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *DestinationAzureBlobStorageUpdateOutputFormat) UnmarshalJSON(data []byte) error {

	destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON := new(DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON)
	if err := utils.UnmarshalJSON(data, &destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON = destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON
		u.Type = DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON
		return nil
	}

	destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues := new(DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues)
	if err := utils.UnmarshalJSON(data, &destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues, "", true, true); err == nil {
		u.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues = destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues
		u.Type = DestinationAzureBlobStorageUpdateOutputFormatTypeDestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationAzureBlobStorageUpdateOutputFormat) MarshalJSON() ([]byte, error) {
	if u.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues != nil {
		return utils.MarshalJSON(u.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues, "", true)
	}

	if u.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationAzureBlobStorageUpdate struct {
	// The Azure blob storage account key.
	AzureBlobStorageAccountKey string `json:"azure_blob_storage_account_key"`
	// The account's name of the Azure Blob Storage.
	AzureBlobStorageAccountName string `json:"azure_blob_storage_account_name"`
	// The name of the Azure blob storage container. If not exists - will be created automatically. May be empty, then will be created automatically airbytecontainer+timestamp
	AzureBlobStorageContainerName *string `json:"azure_blob_storage_container_name,omitempty"`
	// This is Azure Blob Storage endpoint domain name. Leave default value (or leave it empty if run container from command line) to use Microsoft native from example.
	AzureBlobStorageEndpointDomainName *string `default:"blob.core.windows.net" json:"azure_blob_storage_endpoint_domain_name"`
	// The amount of megabytes to buffer for the output stream to Azure. This will impact memory footprint on workers, but may need adjustment for performance and appropriate block size in Azure.
	AzureBlobStorageOutputBufferSize *int64 `default:"5" json:"azure_blob_storage_output_buffer_size"`
	// The amount of megabytes after which the connector should spill the records in a new blob object. Make sure to configure size greater than individual records. Enter 0 if not applicable
	AzureBlobStorageSpillSize *int64 `default:"500" json:"azure_blob_storage_spill_size"`
	// Output data format
	Format DestinationAzureBlobStorageUpdateOutputFormat `json:"format"`
}

func (d DestinationAzureBlobStorageUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageAccountKey() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountKey
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageAccountName() string {
	if o == nil {
		return ""
	}
	return o.AzureBlobStorageAccountName
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageContainerName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageContainerName
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageEndpointDomainName() *string {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageEndpointDomainName
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageOutputBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageOutputBufferSize
}

func (o *DestinationAzureBlobStorageUpdate) GetAzureBlobStorageSpillSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AzureBlobStorageSpillSize
}

func (o *DestinationAzureBlobStorageUpdate) GetFormat() DestinationAzureBlobStorageUpdateOutputFormat {
	if o == nil {
		return DestinationAzureBlobStorageUpdateOutputFormat{}
	}
	return o.Format
}
