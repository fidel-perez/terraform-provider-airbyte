// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGcsResourceModel) ToCreateSDKType() *shared.SourceGcsCreateRequest {
	bucket := r.Configuration.Bucket.ValueString()
	serviceAccount := r.Configuration.ServiceAccount.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGCSSourceGCSStreamConfig = nil
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGcsFormat
		var sourceGcsCSVFormat *shared.SourceGcsCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = nil
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGcsCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGcsFromCSV *shared.SourceGcsFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGcsFromCSV = &shared.SourceGcsFromCSV{}
				}
				if sourceGcsFromCSV != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsFromCSV: sourceGcsFromCSV,
					}
				}
				var sourceGcsAutogenerated *shared.SourceGcsAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGcsAutogenerated = &shared.SourceGcsAutogenerated{}
				}
				if sourceGcsAutogenerated != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsAutogenerated: sourceGcsAutogenerated,
					}
				}
				var sourceGcsUserProvided *shared.SourceGcsUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = nil
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGcsUserProvided = &shared.SourceGcsUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGcsUserProvided != nil {
					headerDefinition = &shared.SourceGcsCSVHeaderDefinition{
						SourceGcsUserProvided: sourceGcsUserProvided,
					}
				}
			}
			inferenceType := new(shared.SourceGcsInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceGcsInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = nil
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = nil
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGcsCSVFormat = &shared.SourceGcsCSVFormat{
				Delimiter:            delimiter,
				DoubleQuote:          doubleQuote,
				Encoding:             encoding,
				EscapeChar:           escapeChar,
				FalseValues:          falseValues,
				HeaderDefinition:     headerDefinition,
				InferenceType:        inferenceType,
				NullValues:           nullValues,
				QuoteChar:            quoteChar,
				SkipRowsAfterHeader:  skipRowsAfterHeader,
				SkipRowsBeforeHeader: skipRowsBeforeHeader,
				StringsCanBeNull:     stringsCanBeNull,
				TrueValues:           trueValues,
			}
		}
		if sourceGcsCSVFormat != nil {
			format = shared.SourceGcsFormat{
				SourceGcsCSVFormat: sourceGcsCSVFormat,
			}
		}
		var globs []string = nil
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGcsValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGcsValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGCSSourceGCSStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			LegacyPrefix:              legacyPrefix,
			Name:                      name,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceGcs{
		Bucket:         bucket,
		ServiceAccount: serviceAccount,
		StartDate:      startDate,
		Streams:        streams,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGcsCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGcsResourceModel) ToGetSDKType() *shared.SourceGcsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGcsResourceModel) ToUpdateSDKType() *shared.SourceGcsPutRequest {
	bucket := r.Configuration.Bucket.ValueString()
	serviceAccount := r.Configuration.ServiceAccount.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGCSStreamConfig = nil
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGcsUpdateFormat
		var sourceGcsUpdateCSVFormat *shared.SourceGcsUpdateCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = nil
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGcsUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGcsUpdateFromCSV *shared.SourceGcsUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGcsUpdateFromCSV = &shared.SourceGcsUpdateFromCSV{}
				}
				if sourceGcsUpdateFromCSV != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateFromCSV: sourceGcsUpdateFromCSV,
					}
				}
				var sourceGcsUpdateAutogenerated *shared.SourceGcsUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGcsUpdateAutogenerated = &shared.SourceGcsUpdateAutogenerated{}
				}
				if sourceGcsUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateAutogenerated: sourceGcsUpdateAutogenerated,
					}
				}
				var sourceGcsUpdateUserProvided *shared.SourceGcsUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = nil
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGcsUpdateUserProvided = &shared.SourceGcsUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGcsUpdateUserProvided != nil {
					headerDefinition = &shared.SourceGcsUpdateCSVHeaderDefinition{
						SourceGcsUpdateUserProvided: sourceGcsUpdateUserProvided,
					}
				}
			}
			inferenceType := new(shared.SourceGcsUpdateInferenceType)
			if !streamsItem.Format.CSVFormat.InferenceType.IsUnknown() && !streamsItem.Format.CSVFormat.InferenceType.IsNull() {
				*inferenceType = shared.SourceGcsUpdateInferenceType(streamsItem.Format.CSVFormat.InferenceType.ValueString())
			} else {
				inferenceType = nil
			}
			var nullValues []string = nil
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = nil
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGcsUpdateCSVFormat = &shared.SourceGcsUpdateCSVFormat{
				Delimiter:            delimiter,
				DoubleQuote:          doubleQuote,
				Encoding:             encoding,
				EscapeChar:           escapeChar,
				FalseValues:          falseValues,
				HeaderDefinition:     headerDefinition,
				InferenceType:        inferenceType,
				NullValues:           nullValues,
				QuoteChar:            quoteChar,
				SkipRowsAfterHeader:  skipRowsAfterHeader,
				SkipRowsBeforeHeader: skipRowsBeforeHeader,
				StringsCanBeNull:     stringsCanBeNull,
				TrueValues:           trueValues,
			}
		}
		if sourceGcsUpdateCSVFormat != nil {
			format = shared.SourceGcsUpdateFormat{
				SourceGcsUpdateCSVFormat: sourceGcsUpdateCSVFormat,
			}
		}
		var globs []string = nil
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		legacyPrefix := new(string)
		if !streamsItem.LegacyPrefix.IsUnknown() && !streamsItem.LegacyPrefix.IsNull() {
			*legacyPrefix = streamsItem.LegacyPrefix.ValueString()
		} else {
			legacyPrefix = nil
		}
		name := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGcsUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGcsUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGCSStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			LegacyPrefix:              legacyPrefix,
			Name:                      name,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceGcsUpdate{
		Bucket:         bucket,
		ServiceAccount: serviceAccount,
		StartDate:      startDate,
		Streams:        streams,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGcsPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGcsResourceModel) ToDeleteSDKType() *shared.SourceGcsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceGcsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceGcsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
