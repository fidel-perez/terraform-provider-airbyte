// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType string

const (
	SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthTypeService SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType = "Service"
)

func (e SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType) ToPointer() *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service":
		*e = SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication struct {
	authType SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType `const:"Service" json:"auth_type"`
	// The email of the user which has permissions to access the Google Workspace Admin APIs.
	Email string `json:"email"`
	// The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys">here</a>.
	ServiceAccountInfo string `json:"service_account_info"`
}

func (s SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) GetAuthType() SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthType {
	return SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthenticationAuthTypeService
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) GetServiceAccountInfo() string {
	if o == nil {
		return ""
	}
	return o.ServiceAccountInfo
}

type SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType string

const (
	SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthTypeClient SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType = "Client"
)

func (e SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType) ToPointer() *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth struct {
	// Access token for making authenticated requests. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	AccessToken *string                                                        `json:"access_token,omitempty"`
	authType    SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType `const:"Client" json:"auth_type"`
	// The client ID of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientID string `json:"client_id"`
	// The client secret of your Google Search Console developer application. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token. Read more <a href="https://developers.google.com/webmaster-tools/v1/how-tos/authorizing">here</a>.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) GetAuthType() SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthType {
	return SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuthAuthTypeClient
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceGoogleSearchConsoleUpdateAuthenticationTypeType string

const (
	SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth                           SourceGoogleSearchConsoleUpdateAuthenticationTypeType = "source-google-search-console-update_Authentication Type_OAuth"
	SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication SourceGoogleSearchConsoleUpdateAuthenticationTypeType = "source-google-search-console-update_Authentication Type_Service Account Key Authentication"
)

type SourceGoogleSearchConsoleUpdateAuthenticationType struct {
	SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth                           *SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth
	SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication *SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication

	Type SourceGoogleSearchConsoleUpdateAuthenticationTypeType
}

func CreateSourceGoogleSearchConsoleUpdateAuthenticationTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth(sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth) SourceGoogleSearchConsoleUpdateAuthenticationType {
	typ := SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth

	return SourceGoogleSearchConsoleUpdateAuthenticationType{
		SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth: &sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth,
		Type: typ,
	}
}

func CreateSourceGoogleSearchConsoleUpdateAuthenticationTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication(sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication) SourceGoogleSearchConsoleUpdateAuthenticationType {
	typ := SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication

	return SourceGoogleSearchConsoleUpdateAuthenticationType{
		SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication: &sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication,
		Type: typ,
	}
}

func (u *SourceGoogleSearchConsoleUpdateAuthenticationType) UnmarshalJSON(data []byte) error {

	sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication := new(SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication, "", true, true); err == nil {
		u.SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication = sourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication
		u.Type = SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication
		return nil
	}

	sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth := new(SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth)
	if err := utils.UnmarshalJSON(data, &sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth, "", true, true); err == nil {
		u.SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth = sourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth
		u.Type = SourceGoogleSearchConsoleUpdateAuthenticationTypeTypeSourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGoogleSearchConsoleUpdateAuthenticationType) MarshalJSON() ([]byte, error) {
	if u.SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth != nil {
		return utils.MarshalJSON(u.SourceGoogleSearchConsoleUpdateAuthenticationTypeOAuth, "", true)
	}

	if u.SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication != nil {
		return utils.MarshalJSON(u.SourceGoogleSearchConsoleUpdateAuthenticationTypeServiceAccountKeyAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums - An enumeration of dimensions.
type SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums string

const (
	SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnumsCountry SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums = "country"
	SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnumsDate    SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums = "date"
	SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnumsDevice  SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums = "device"
	SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnumsPage    SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums = "page"
	SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnumsQuery   SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums = "query"
)

func (e SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums) ToPointer() *SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "date":
		fallthrough
	case "device":
		fallthrough
	case "page":
		fallthrough
	case "query":
		*e = SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdateCustomReportConfig struct {
	// A list of dimensions (country, date, device, page, query)
	Dimensions []SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums `json:"dimensions"`
	// The name of the custom report, this name would be used as stream name
	Name string `json:"name"`
}

func (o *SourceGoogleSearchConsoleUpdateCustomReportConfig) GetDimensions() []SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums {
	if o == nil {
		return []SourceGoogleSearchConsoleUpdateCustomReportConfigValidEnums{}
	}
	return o.Dimensions
}

func (o *SourceGoogleSearchConsoleUpdateCustomReportConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// SourceGoogleSearchConsoleUpdateDataFreshness - If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
type SourceGoogleSearchConsoleUpdateDataFreshness string

const (
	SourceGoogleSearchConsoleUpdateDataFreshnessFinal SourceGoogleSearchConsoleUpdateDataFreshness = "final"
	SourceGoogleSearchConsoleUpdateDataFreshnessAll   SourceGoogleSearchConsoleUpdateDataFreshness = "all"
)

func (e SourceGoogleSearchConsoleUpdateDataFreshness) ToPointer() *SourceGoogleSearchConsoleUpdateDataFreshness {
	return &e
}

func (e *SourceGoogleSearchConsoleUpdateDataFreshness) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "final":
		fallthrough
	case "all":
		*e = SourceGoogleSearchConsoleUpdateDataFreshness(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSearchConsoleUpdateDataFreshness: %v", v)
	}
}

type SourceGoogleSearchConsoleUpdate struct {
	Authorization SourceGoogleSearchConsoleUpdateAuthenticationType `json:"authorization"`
	// (DEPRCATED) A JSON array describing the custom reports you want to sync from Google Search Console. See our <a href='https://docs.airbyte.com/integrations/sources/google-search-console'>documentation</a> for more information on formulating custom reports.
	CustomReports *string `json:"custom_reports,omitempty"`
	// You can add your Custom Analytics report by creating one.
	CustomReportsArray []SourceGoogleSearchConsoleUpdateCustomReportConfig `json:"custom_reports_array,omitempty"`
	// If set to 'final', the returned data will include only finalized, stable data. If set to 'all', fresh data will be included. When using Incremental sync mode, we do not recommend setting this parameter to 'all' as it may cause data loss. More information can be found in our <a href='https://docs.airbyte.com/integrations/source/google-search-console'>full documentation</a>.
	DataState *SourceGoogleSearchConsoleUpdateDataFreshness `default:"final" json:"data_state"`
	// UTC date in the format YYYY-MM-DD. Any data created after this date will not be replicated. Must be greater or equal to the start date field. Leaving this field blank will replicate all data from the start date onward.
	EndDate *types.Date `json:"end_date,omitempty"`
	// The URLs of the website property attached to your GSC account. Learn more about properties <a href="https://support.google.com/webmasters/answer/34592?hl=en">here</a>.
	SiteUrls []string `json:"site_urls"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate *types.Date `default:"2021-01-01" json:"start_date"`
}

func (s SourceGoogleSearchConsoleUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleSearchConsoleUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleSearchConsoleUpdate) GetAuthorization() SourceGoogleSearchConsoleUpdateAuthenticationType {
	if o == nil {
		return SourceGoogleSearchConsoleUpdateAuthenticationType{}
	}
	return o.Authorization
}

func (o *SourceGoogleSearchConsoleUpdate) GetCustomReports() *string {
	if o == nil {
		return nil
	}
	return o.CustomReports
}

func (o *SourceGoogleSearchConsoleUpdate) GetCustomReportsArray() []SourceGoogleSearchConsoleUpdateCustomReportConfig {
	if o == nil {
		return nil
	}
	return o.CustomReportsArray
}

func (o *SourceGoogleSearchConsoleUpdate) GetDataState() *SourceGoogleSearchConsoleUpdateDataFreshness {
	if o == nil {
		return nil
	}
	return o.DataState
}

func (o *SourceGoogleSearchConsoleUpdate) GetEndDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceGoogleSearchConsoleUpdate) GetSiteUrls() []string {
	if o == nil {
		return []string{}
	}
	return o.SiteUrls
}

func (o *SourceGoogleSearchConsoleUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
