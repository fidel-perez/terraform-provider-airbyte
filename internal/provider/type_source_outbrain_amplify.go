// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceOutbrainAmplify struct {
	Credentials          SourceOutbrainAmplifyAuthenticationMethod `tfsdk:"credentials"`
	EndDate              types.String                              `tfsdk:"end_date"`
	GeoLocationBreakdown types.String                              `tfsdk:"geo_location_breakdown"`
	ReportGranularity    types.String                              `tfsdk:"report_granularity"`
	StartDate            types.String                              `tfsdk:"start_date"`
}
