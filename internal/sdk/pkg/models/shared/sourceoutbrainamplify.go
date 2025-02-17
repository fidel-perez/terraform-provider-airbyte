// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest string

const (
	SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequestUsernamePassword SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest = "username_password"
)

func (e SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest) ToPointer() *SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest {
	return &e
}

func (e *SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "username_password":
		*e = SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest: %v", v)
	}
}

// SourceOutbrainAmplifyUsernamePassword - Credentials for making authenticated requests requires either username/password or access_token.
type SourceOutbrainAmplifyUsernamePassword struct {
	// Add Password for authentication.
	Password string                                                                         `json:"password"`
	type_    SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest `const:"username_password" json:"type"`
	// Add Username for authentication.
	Username string `json:"username"`
}

func (s SourceOutbrainAmplifyUsernamePassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOutbrainAmplifyUsernamePassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOutbrainAmplifyUsernamePassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceOutbrainAmplifyUsernamePassword) GetType() SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequest {
	return SourceOutbrainAmplifyBothUsernameAndPasswordIsRequiredForAuthenticationRequestUsernamePassword
}

func (o *SourceOutbrainAmplifyUsernamePassword) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests string

const (
	SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequestsAccessToken SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests = "access_token"
)

func (e SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests) ToPointer() *SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests {
	return &e
}

func (e *SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests: %v", v)
	}
}

// SourceOutbrainAmplifyAccessToken - Credentials for making authenticated requests requires either username/password or access_token.
type SourceOutbrainAmplifyAccessToken struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                              `json:"access_token"`
	type_       SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests `const:"access_token" json:"type"`
}

func (s SourceOutbrainAmplifyAccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOutbrainAmplifyAccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceOutbrainAmplifyAccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceOutbrainAmplifyAccessToken) GetType() SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequests {
	return SourceOutbrainAmplifyAccessTokenIsRequiredForAuthenticationRequestsAccessToken
}

type SourceOutbrainAmplifyAuthenticationMethodType string

const (
	SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyAccessToken      SourceOutbrainAmplifyAuthenticationMethodType = "source-outbrain-amplify_Access token"
	SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyUsernamePassword SourceOutbrainAmplifyAuthenticationMethodType = "source-outbrain-amplify_Username Password"
)

type SourceOutbrainAmplifyAuthenticationMethod struct {
	SourceOutbrainAmplifyAccessToken      *SourceOutbrainAmplifyAccessToken
	SourceOutbrainAmplifyUsernamePassword *SourceOutbrainAmplifyUsernamePassword

	Type SourceOutbrainAmplifyAuthenticationMethodType
}

func CreateSourceOutbrainAmplifyAuthenticationMethodSourceOutbrainAmplifyAccessToken(sourceOutbrainAmplifyAccessToken SourceOutbrainAmplifyAccessToken) SourceOutbrainAmplifyAuthenticationMethod {
	typ := SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyAccessToken

	return SourceOutbrainAmplifyAuthenticationMethod{
		SourceOutbrainAmplifyAccessToken: &sourceOutbrainAmplifyAccessToken,
		Type:                             typ,
	}
}

func CreateSourceOutbrainAmplifyAuthenticationMethodSourceOutbrainAmplifyUsernamePassword(sourceOutbrainAmplifyUsernamePassword SourceOutbrainAmplifyUsernamePassword) SourceOutbrainAmplifyAuthenticationMethod {
	typ := SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyUsernamePassword

	return SourceOutbrainAmplifyAuthenticationMethod{
		SourceOutbrainAmplifyUsernamePassword: &sourceOutbrainAmplifyUsernamePassword,
		Type:                                  typ,
	}
}

func (u *SourceOutbrainAmplifyAuthenticationMethod) UnmarshalJSON(data []byte) error {

	sourceOutbrainAmplifyAccessToken := new(SourceOutbrainAmplifyAccessToken)
	if err := utils.UnmarshalJSON(data, &sourceOutbrainAmplifyAccessToken, "", true, true); err == nil {
		u.SourceOutbrainAmplifyAccessToken = sourceOutbrainAmplifyAccessToken
		u.Type = SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyAccessToken
		return nil
	}

	sourceOutbrainAmplifyUsernamePassword := new(SourceOutbrainAmplifyUsernamePassword)
	if err := utils.UnmarshalJSON(data, &sourceOutbrainAmplifyUsernamePassword, "", true, true); err == nil {
		u.SourceOutbrainAmplifyUsernamePassword = sourceOutbrainAmplifyUsernamePassword
		u.Type = SourceOutbrainAmplifyAuthenticationMethodTypeSourceOutbrainAmplifyUsernamePassword
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceOutbrainAmplifyAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceOutbrainAmplifyAccessToken != nil {
		return utils.MarshalJSON(u.SourceOutbrainAmplifyAccessToken, "", true)
	}

	if u.SourceOutbrainAmplifyUsernamePassword != nil {
		return utils.MarshalJSON(u.SourceOutbrainAmplifyUsernamePassword, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SourceOutbrainAmplifyGranularityForGeoLocationRegion - The granularity used for geo location data in reports.
type SourceOutbrainAmplifyGranularityForGeoLocationRegion string

const (
	SourceOutbrainAmplifyGranularityForGeoLocationRegionCountry   SourceOutbrainAmplifyGranularityForGeoLocationRegion = "country"
	SourceOutbrainAmplifyGranularityForGeoLocationRegionRegion    SourceOutbrainAmplifyGranularityForGeoLocationRegion = "region"
	SourceOutbrainAmplifyGranularityForGeoLocationRegionSubregion SourceOutbrainAmplifyGranularityForGeoLocationRegion = "subregion"
)

func (e SourceOutbrainAmplifyGranularityForGeoLocationRegion) ToPointer() *SourceOutbrainAmplifyGranularityForGeoLocationRegion {
	return &e
}

func (e *SourceOutbrainAmplifyGranularityForGeoLocationRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "country":
		fallthrough
	case "region":
		fallthrough
	case "subregion":
		*e = SourceOutbrainAmplifyGranularityForGeoLocationRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOutbrainAmplifyGranularityForGeoLocationRegion: %v", v)
	}
}

// SourceOutbrainAmplifyGranularityForPeriodicReports - The granularity used for periodic data in reports. See <a href="https://amplifyv01.docs.apiary.io/#reference/performance-reporting/periodic/retrieve-performance-statistics-for-all-marketer-campaigns-by-periodic-breakdown">the docs</a>.
type SourceOutbrainAmplifyGranularityForPeriodicReports string

const (
	SourceOutbrainAmplifyGranularityForPeriodicReportsDaily   SourceOutbrainAmplifyGranularityForPeriodicReports = "daily"
	SourceOutbrainAmplifyGranularityForPeriodicReportsWeekly  SourceOutbrainAmplifyGranularityForPeriodicReports = "weekly"
	SourceOutbrainAmplifyGranularityForPeriodicReportsMonthly SourceOutbrainAmplifyGranularityForPeriodicReports = "monthly"
)

func (e SourceOutbrainAmplifyGranularityForPeriodicReports) ToPointer() *SourceOutbrainAmplifyGranularityForPeriodicReports {
	return &e
}

func (e *SourceOutbrainAmplifyGranularityForPeriodicReports) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "weekly":
		fallthrough
	case "monthly":
		*e = SourceOutbrainAmplifyGranularityForPeriodicReports(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOutbrainAmplifyGranularityForPeriodicReports: %v", v)
	}
}

type OutbrainAmplify string

const (
	OutbrainAmplifyOutbrainAmplify OutbrainAmplify = "outbrain-amplify"
)

func (e OutbrainAmplify) ToPointer() *OutbrainAmplify {
	return &e
}

func (e *OutbrainAmplify) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "outbrain-amplify":
		*e = OutbrainAmplify(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OutbrainAmplify: %v", v)
	}
}

type SourceOutbrainAmplify struct {
	// Credentials for making authenticated requests requires either username/password or access_token.
	Credentials SourceOutbrainAmplifyAuthenticationMethod `json:"credentials"`
	// Date in the format YYYY-MM-DD.
	EndDate *string `json:"end_date,omitempty"`
	// The granularity used for geo location data in reports.
	GeoLocationBreakdown *SourceOutbrainAmplifyGranularityForGeoLocationRegion `json:"geo_location_breakdown,omitempty"`
	// The granularity used for periodic data in reports. See <a href="https://amplifyv01.docs.apiary.io/#reference/performance-reporting/periodic/retrieve-performance-statistics-for-all-marketer-campaigns-by-periodic-breakdown">the docs</a>.
	ReportGranularity *SourceOutbrainAmplifyGranularityForPeriodicReports `json:"report_granularity,omitempty"`
	sourceType        OutbrainAmplify                                     `const:"outbrain-amplify" json:"sourceType"`
	// Date in the format YYYY-MM-DD eg. 2017-01-25. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}

func (s SourceOutbrainAmplify) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceOutbrainAmplify) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceOutbrainAmplify) GetCredentials() SourceOutbrainAmplifyAuthenticationMethod {
	if o == nil {
		return SourceOutbrainAmplifyAuthenticationMethod{}
	}
	return o.Credentials
}

func (o *SourceOutbrainAmplify) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceOutbrainAmplify) GetGeoLocationBreakdown() *SourceOutbrainAmplifyGranularityForGeoLocationRegion {
	if o == nil {
		return nil
	}
	return o.GeoLocationBreakdown
}

func (o *SourceOutbrainAmplify) GetReportGranularity() *SourceOutbrainAmplifyGranularityForPeriodicReports {
	if o == nil {
		return nil
	}
	return o.ReportGranularity
}

func (o *SourceOutbrainAmplify) GetSourceType() OutbrainAmplify {
	return OutbrainAmplifyOutbrainAmplify
}

func (o *SourceOutbrainAmplify) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
