// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceApifyDataset struct {
	Clean     types.Bool   `tfsdk:"clean"`
	DatasetID types.String `tfsdk:"dataset_id"`
	Token     types.String `tfsdk:"token"`
}
