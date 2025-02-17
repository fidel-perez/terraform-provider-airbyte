// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationWeaviateIndexing struct {
	AdditionalHeaders []Header                          `tfsdk:"additional_headers"`
	Auth              DestinationWeaviateAuthentication `tfsdk:"auth"`
	BatchSize         types.Int64                       `tfsdk:"batch_size"`
	DefaultVectorizer types.String                      `tfsdk:"default_vectorizer"`
	Host              types.String                      `tfsdk:"host"`
	TextField         types.String                      `tfsdk:"text_field"`
}
