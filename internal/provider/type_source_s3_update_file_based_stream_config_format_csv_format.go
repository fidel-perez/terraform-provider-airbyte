// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceS3UpdateFileBasedStreamConfigFormatCSVFormat struct {
	Delimiter            types.String                                                           `tfsdk:"delimiter"`
	DoubleQuote          types.Bool                                                             `tfsdk:"double_quote"`
	Encoding             types.String                                                           `tfsdk:"encoding"`
	EscapeChar           types.String                                                           `tfsdk:"escape_char"`
	FalseValues          []types.String                                                         `tfsdk:"false_values"`
	HeaderDefinition     *SourceS3UpdateFileBasedStreamConfigFormatCSVFormatCSVHeaderDefinition `tfsdk:"header_definition"`
	InferenceType        types.String                                                           `tfsdk:"inference_type"`
	NullValues           []types.String                                                         `tfsdk:"null_values"`
	QuoteChar            types.String                                                           `tfsdk:"quote_char"`
	SkipRowsAfterHeader  types.Int64                                                            `tfsdk:"skip_rows_after_header"`
	SkipRowsBeforeHeader types.Int64                                                            `tfsdk:"skip_rows_before_header"`
	StringsCanBeNull     types.Bool                                                             `tfsdk:"strings_can_be_null"`
	TrueValues           []types.String                                                         `tfsdk:"true_values"`
}
