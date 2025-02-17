// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationAzureBlobStorageUpdateFormatType string

const (
	DestinationAzureBlobStorageUpdateFormatTypeJsonl DestinationAzureBlobStorageUpdateFormatType = "JSONL"
)

func (e DestinationAzureBlobStorageUpdateFormatType) ToPointer() *DestinationAzureBlobStorageUpdateFormatType {
	return &e
}

func (e *DestinationAzureBlobStorageUpdateFormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JSONL":
		*e = DestinationAzureBlobStorageUpdateFormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationAzureBlobStorageUpdateFormatType: %v", v)
	}
}

// DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON - Output data format
type DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON struct {
	formatType DestinationAzureBlobStorageUpdateFormatType `const:"JSONL" json:"format_type"`
}

func (d DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON) GetFormatType() DestinationAzureBlobStorageUpdateFormatType {
	return DestinationAzureBlobStorageUpdateFormatTypeJsonl
}

// NormalizationFlattening - Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
type NormalizationFlattening string

const (
	NormalizationFlatteningNoFlattening        NormalizationFlattening = "No flattening"
	NormalizationFlatteningRootLevelFlattening NormalizationFlattening = "Root level flattening"
)

func (e NormalizationFlattening) ToPointer() *NormalizationFlattening {
	return &e
}

func (e *NormalizationFlattening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "No flattening":
		fallthrough
	case "Root level flattening":
		*e = NormalizationFlattening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NormalizationFlattening: %v", v)
	}
}

type FormatType string

const (
	FormatTypeCsv FormatType = "CSV"
)

func (e FormatType) ToPointer() *FormatType {
	return &e
}

func (e *FormatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CSV":
		*e = FormatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FormatType: %v", v)
	}
}

// CSVCommaSeparatedValues - Output data format
type CSVCommaSeparatedValues struct {
	// Whether the input json data should be normalized (flattened) in the output CSV. Please refer to docs for details.
	Flattening *NormalizationFlattening `default:"No flattening" json:"flattening"`
	formatType FormatType               `const:"CSV" json:"format_type"`
}

func (c CSVCommaSeparatedValues) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CSVCommaSeparatedValues) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *CSVCommaSeparatedValues) GetFlattening() *NormalizationFlattening {
	if o == nil {
		return nil
	}
	return o.Flattening
}

func (o *CSVCommaSeparatedValues) GetFormatType() FormatType {
	return FormatTypeCsv
}

type OutputFormatType string

const (
	OutputFormatTypeCSVCommaSeparatedValues                                        OutputFormatType = "CSV: Comma-Separated Values"
	OutputFormatTypeDestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON OutputFormatType = "destination-azure-blob-storage-update_JSON Lines: newline-delimited JSON"
)

type OutputFormat struct {
	CSVCommaSeparatedValues                                        *CSVCommaSeparatedValues
	DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON *DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON

	Type OutputFormatType
}

func CreateOutputFormatCSVCommaSeparatedValues(csvCommaSeparatedValues CSVCommaSeparatedValues) OutputFormat {
	typ := OutputFormatTypeCSVCommaSeparatedValues

	return OutputFormat{
		CSVCommaSeparatedValues: &csvCommaSeparatedValues,
		Type:                    typ,
	}
}

func CreateOutputFormatDestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON(destinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON) OutputFormat {
	typ := OutputFormatTypeDestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON

	return OutputFormat{
		DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON: &destinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON,
		Type: typ,
	}
}

func (u *OutputFormat) UnmarshalJSON(data []byte) error {

	destinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON := new(DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON)
	if err := utils.UnmarshalJSON(data, &destinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON, "", true, true); err == nil {
		u.DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON = destinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON
		u.Type = OutputFormatTypeDestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON
		return nil
	}

	csvCommaSeparatedValues := new(CSVCommaSeparatedValues)
	if err := utils.UnmarshalJSON(data, &csvCommaSeparatedValues, "", true, true); err == nil {
		u.CSVCommaSeparatedValues = csvCommaSeparatedValues
		u.Type = OutputFormatTypeCSVCommaSeparatedValues
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u OutputFormat) MarshalJSON() ([]byte, error) {
	if u.CSVCommaSeparatedValues != nil {
		return utils.MarshalJSON(u.CSVCommaSeparatedValues, "", true)
	}

	if u.DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON != nil {
		return utils.MarshalJSON(u.DestinationAzureBlobStorageUpdateJSONLinesNewlineDelimitedJSON, "", true)
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
	Format OutputFormat `json:"format"`
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

func (o *DestinationAzureBlobStorageUpdate) GetFormat() OutputFormat {
	if o == nil {
		return OutputFormat{}
	}
	return o.Format
}
