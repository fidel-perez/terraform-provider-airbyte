// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type SourceSmartsheetsSchemasAuthType string

const (
	SourceSmartsheetsSchemasAuthTypeAccessToken SourceSmartsheetsSchemasAuthType = "access_token"
)

func (e SourceSmartsheetsSchemasAuthType) ToPointer() *SourceSmartsheetsSchemasAuthType {
	return &e
}

func (e *SourceSmartsheetsSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceSmartsheetsSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsSchemasAuthType: %v", v)
	}
}

type SourceSmartsheetsAPIAccessToken struct {
	// The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.
	AccessToken string                            `json:"access_token"`
	authType    *SourceSmartsheetsSchemasAuthType `const:"access_token" json:"auth_type,omitempty"`
}

func (s SourceSmartsheetsAPIAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSmartsheetsAPIAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSmartsheetsAPIAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceSmartsheetsAPIAccessToken) GetAuthType() *SourceSmartsheetsSchemasAuthType {
	return SourceSmartsheetsSchemasAuthTypeAccessToken.ToPointer()
}

type SourceSmartsheetsAuthType string

const (
	SourceSmartsheetsAuthTypeOauth20 SourceSmartsheetsAuthType = "oauth2.0"
)

func (e SourceSmartsheetsAuthType) ToPointer() *SourceSmartsheetsAuthType {
	return &e
}

func (e *SourceSmartsheetsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceSmartsheetsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsAuthType: %v", v)
	}
}

type SourceSmartsheetsOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                     `json:"access_token"`
	authType    *SourceSmartsheetsAuthType `const:"oauth2.0" json:"auth_type,omitempty"`
	// The API ID of the SmartSheets developer application.
	ClientID string `json:"client_id"`
	// The API Secret the SmartSheets developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

func (s SourceSmartsheetsOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSmartsheetsOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceSmartsheetsOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceSmartsheetsOAuth20) GetAuthType() *SourceSmartsheetsAuthType {
	return SourceSmartsheetsAuthTypeOauth20.ToPointer()
}

func (o *SourceSmartsheetsOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSmartsheetsOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSmartsheetsOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceSmartsheetsOAuth20) GetTokenExpiryDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.TokenExpiryDate
}

type SourceSmartsheetsAuthorizationMethodType string

const (
	SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsOAuth20        SourceSmartsheetsAuthorizationMethodType = "source-smartsheets_OAuth2.0"
	SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAPIAccessToken SourceSmartsheetsAuthorizationMethodType = "source-smartsheets_API Access Token"
)

type SourceSmartsheetsAuthorizationMethod struct {
	SourceSmartsheetsOAuth20        *SourceSmartsheetsOAuth20
	SourceSmartsheetsAPIAccessToken *SourceSmartsheetsAPIAccessToken

	Type SourceSmartsheetsAuthorizationMethodType
}

func CreateSourceSmartsheetsAuthorizationMethodSourceSmartsheetsOAuth20(sourceSmartsheetsOAuth20 SourceSmartsheetsOAuth20) SourceSmartsheetsAuthorizationMethod {
	typ := SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsOAuth20

	return SourceSmartsheetsAuthorizationMethod{
		SourceSmartsheetsOAuth20: &sourceSmartsheetsOAuth20,
		Type:                     typ,
	}
}

func CreateSourceSmartsheetsAuthorizationMethodSourceSmartsheetsAPIAccessToken(sourceSmartsheetsAPIAccessToken SourceSmartsheetsAPIAccessToken) SourceSmartsheetsAuthorizationMethod {
	typ := SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAPIAccessToken

	return SourceSmartsheetsAuthorizationMethod{
		SourceSmartsheetsAPIAccessToken: &sourceSmartsheetsAPIAccessToken,
		Type:                            typ,
	}
}

func (u *SourceSmartsheetsAuthorizationMethod) UnmarshalJSON(data []byte) error {

	sourceSmartsheetsAPIAccessToken := new(SourceSmartsheetsAPIAccessToken)
	if err := utils.UnmarshalJSON(data, &sourceSmartsheetsAPIAccessToken, "", true, true); err == nil {
		u.SourceSmartsheetsAPIAccessToken = sourceSmartsheetsAPIAccessToken
		u.Type = SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsAPIAccessToken
		return nil
	}

	sourceSmartsheetsOAuth20 := new(SourceSmartsheetsOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceSmartsheetsOAuth20, "", true, true); err == nil {
		u.SourceSmartsheetsOAuth20 = sourceSmartsheetsOAuth20
		u.Type = SourceSmartsheetsAuthorizationMethodTypeSourceSmartsheetsOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSmartsheetsAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceSmartsheetsOAuth20 != nil {
		return utils.MarshalJSON(u.SourceSmartsheetsOAuth20, "", true)
	}

	if u.SourceSmartsheetsAPIAccessToken != nil {
		return utils.MarshalJSON(u.SourceSmartsheetsAPIAccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceSmartsheetsValidenums string

const (
	SourceSmartsheetsValidenumsSheetcreatedAt   SourceSmartsheetsValidenums = "sheetcreatedAt"
	SourceSmartsheetsValidenumsSheetid          SourceSmartsheetsValidenums = "sheetid"
	SourceSmartsheetsValidenumsSheetmodifiedAt  SourceSmartsheetsValidenums = "sheetmodifiedAt"
	SourceSmartsheetsValidenumsSheetname        SourceSmartsheetsValidenums = "sheetname"
	SourceSmartsheetsValidenumsSheetpermalink   SourceSmartsheetsValidenums = "sheetpermalink"
	SourceSmartsheetsValidenumsSheetversion     SourceSmartsheetsValidenums = "sheetversion"
	SourceSmartsheetsValidenumsSheetaccessLevel SourceSmartsheetsValidenums = "sheetaccess_level"
	SourceSmartsheetsValidenumsRowID            SourceSmartsheetsValidenums = "row_id"
	SourceSmartsheetsValidenumsRowAccessLevel   SourceSmartsheetsValidenums = "row_access_level"
	SourceSmartsheetsValidenumsRowCreatedAt     SourceSmartsheetsValidenums = "row_created_at"
	SourceSmartsheetsValidenumsRowCreatedBy     SourceSmartsheetsValidenums = "row_created_by"
	SourceSmartsheetsValidenumsRowExpanded      SourceSmartsheetsValidenums = "row_expanded"
	SourceSmartsheetsValidenumsRowModifiedBy    SourceSmartsheetsValidenums = "row_modified_by"
	SourceSmartsheetsValidenumsRowParentID      SourceSmartsheetsValidenums = "row_parent_id"
	SourceSmartsheetsValidenumsRowPermalink     SourceSmartsheetsValidenums = "row_permalink"
	SourceSmartsheetsValidenumsRowNumber        SourceSmartsheetsValidenums = "row_number"
	SourceSmartsheetsValidenumsRowVersion       SourceSmartsheetsValidenums = "row_version"
)

func (e SourceSmartsheetsValidenums) ToPointer() *SourceSmartsheetsValidenums {
	return &e
}

func (e *SourceSmartsheetsValidenums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sheetcreatedAt":
		fallthrough
	case "sheetid":
		fallthrough
	case "sheetmodifiedAt":
		fallthrough
	case "sheetname":
		fallthrough
	case "sheetpermalink":
		fallthrough
	case "sheetversion":
		fallthrough
	case "sheetaccess_level":
		fallthrough
	case "row_id":
		fallthrough
	case "row_access_level":
		fallthrough
	case "row_created_at":
		fallthrough
	case "row_created_by":
		fallthrough
	case "row_expanded":
		fallthrough
	case "row_modified_by":
		fallthrough
	case "row_parent_id":
		fallthrough
	case "row_permalink":
		fallthrough
	case "row_number":
		fallthrough
	case "row_version":
		*e = SourceSmartsheetsValidenums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsValidenums: %v", v)
	}
}

type Smartsheets string

const (
	SmartsheetsSmartsheets Smartsheets = "smartsheets"
)

func (e Smartsheets) ToPointer() *Smartsheets {
	return &e
}

func (e *Smartsheets) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smartsheets":
		*e = Smartsheets(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Smartsheets: %v", v)
	}
}

type SourceSmartsheets struct {
	Credentials SourceSmartsheetsAuthorizationMethod `json:"credentials"`
	// A List of available columns which metadata can be pulled from.
	MetadataFields []SourceSmartsheetsValidenums `json:"metadata_fields,omitempty"`
	sourceType     Smartsheets                   `const:"smartsheets" json:"sourceType"`
	// The spreadsheet ID. Find it by opening the spreadsheet then navigating to File > Properties
	SpreadsheetID string `json:"spreadsheet_id"`
	// Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: `2000-01-01T13:00:00`
	StartDatetime *time.Time `default:"2020-01-01T00:00:00+00:00" json:"start_datetime"`
}

func (s SourceSmartsheets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSmartsheets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSmartsheets) GetCredentials() SourceSmartsheetsAuthorizationMethod {
	if o == nil {
		return SourceSmartsheetsAuthorizationMethod{}
	}
	return o.Credentials
}

func (o *SourceSmartsheets) GetMetadataFields() []SourceSmartsheetsValidenums {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *SourceSmartsheets) GetSourceType() Smartsheets {
	return SmartsheetsSmartsheets
}

func (o *SourceSmartsheets) GetSpreadsheetID() string {
	if o == nil {
		return ""
	}
	return o.SpreadsheetID
}

func (o *SourceSmartsheets) GetStartDatetime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDatetime
}
