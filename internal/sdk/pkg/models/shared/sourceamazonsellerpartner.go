// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

// SourceAmazonSellerPartnerAWSSellerPartnerAccountType - Type of the Account you're going to authorize the Airbyte application by
type SourceAmazonSellerPartnerAWSSellerPartnerAccountType string

const (
	SourceAmazonSellerPartnerAWSSellerPartnerAccountTypeSeller SourceAmazonSellerPartnerAWSSellerPartnerAccountType = "Seller"
	SourceAmazonSellerPartnerAWSSellerPartnerAccountTypeVendor SourceAmazonSellerPartnerAWSSellerPartnerAccountType = "Vendor"
)

func (e SourceAmazonSellerPartnerAWSSellerPartnerAccountType) ToPointer() *SourceAmazonSellerPartnerAWSSellerPartnerAccountType {
	return &e
}

func (e *SourceAmazonSellerPartnerAWSSellerPartnerAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Seller":
		fallthrough
	case "Vendor":
		*e = SourceAmazonSellerPartnerAWSSellerPartnerAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonSellerPartnerAWSSellerPartnerAccountType: %v", v)
	}
}

type SourceAmazonSellerPartnerAuthType string

const (
	SourceAmazonSellerPartnerAuthTypeOauth20 SourceAmazonSellerPartnerAuthType = "oauth2.0"
)

func (e SourceAmazonSellerPartnerAuthType) ToPointer() *SourceAmazonSellerPartnerAuthType {
	return &e
}

func (e *SourceAmazonSellerPartnerAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAmazonSellerPartnerAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonSellerPartnerAuthType: %v", v)
	}
}

// SourceAmazonSellerPartnerAWSEnvironment - Select the AWS Environment.
type SourceAmazonSellerPartnerAWSEnvironment string

const (
	SourceAmazonSellerPartnerAWSEnvironmentProduction SourceAmazonSellerPartnerAWSEnvironment = "PRODUCTION"
	SourceAmazonSellerPartnerAWSEnvironmentSandbox    SourceAmazonSellerPartnerAWSEnvironment = "SANDBOX"
)

func (e SourceAmazonSellerPartnerAWSEnvironment) ToPointer() *SourceAmazonSellerPartnerAWSEnvironment {
	return &e
}

func (e *SourceAmazonSellerPartnerAWSEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PRODUCTION":
		fallthrough
	case "SANDBOX":
		*e = SourceAmazonSellerPartnerAWSEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonSellerPartnerAWSEnvironment: %v", v)
	}
}

// SourceAmazonSellerPartnerAWSRegion - Select the AWS Region.
type SourceAmazonSellerPartnerAWSRegion string

const (
	SourceAmazonSellerPartnerAWSRegionAe SourceAmazonSellerPartnerAWSRegion = "AE"
	SourceAmazonSellerPartnerAWSRegionAu SourceAmazonSellerPartnerAWSRegion = "AU"
	SourceAmazonSellerPartnerAWSRegionBe SourceAmazonSellerPartnerAWSRegion = "BE"
	SourceAmazonSellerPartnerAWSRegionBr SourceAmazonSellerPartnerAWSRegion = "BR"
	SourceAmazonSellerPartnerAWSRegionCa SourceAmazonSellerPartnerAWSRegion = "CA"
	SourceAmazonSellerPartnerAWSRegionDe SourceAmazonSellerPartnerAWSRegion = "DE"
	SourceAmazonSellerPartnerAWSRegionEg SourceAmazonSellerPartnerAWSRegion = "EG"
	SourceAmazonSellerPartnerAWSRegionEs SourceAmazonSellerPartnerAWSRegion = "ES"
	SourceAmazonSellerPartnerAWSRegionFr SourceAmazonSellerPartnerAWSRegion = "FR"
	SourceAmazonSellerPartnerAWSRegionGb SourceAmazonSellerPartnerAWSRegion = "GB"
	SourceAmazonSellerPartnerAWSRegionIn SourceAmazonSellerPartnerAWSRegion = "IN"
	SourceAmazonSellerPartnerAWSRegionIt SourceAmazonSellerPartnerAWSRegion = "IT"
	SourceAmazonSellerPartnerAWSRegionJp SourceAmazonSellerPartnerAWSRegion = "JP"
	SourceAmazonSellerPartnerAWSRegionMx SourceAmazonSellerPartnerAWSRegion = "MX"
	SourceAmazonSellerPartnerAWSRegionNl SourceAmazonSellerPartnerAWSRegion = "NL"
	SourceAmazonSellerPartnerAWSRegionPl SourceAmazonSellerPartnerAWSRegion = "PL"
	SourceAmazonSellerPartnerAWSRegionSa SourceAmazonSellerPartnerAWSRegion = "SA"
	SourceAmazonSellerPartnerAWSRegionSe SourceAmazonSellerPartnerAWSRegion = "SE"
	SourceAmazonSellerPartnerAWSRegionSg SourceAmazonSellerPartnerAWSRegion = "SG"
	SourceAmazonSellerPartnerAWSRegionTr SourceAmazonSellerPartnerAWSRegion = "TR"
	SourceAmazonSellerPartnerAWSRegionUk SourceAmazonSellerPartnerAWSRegion = "UK"
	SourceAmazonSellerPartnerAWSRegionUs SourceAmazonSellerPartnerAWSRegion = "US"
)

func (e SourceAmazonSellerPartnerAWSRegion) ToPointer() *SourceAmazonSellerPartnerAWSRegion {
	return &e
}

func (e *SourceAmazonSellerPartnerAWSRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AE":
		fallthrough
	case "AU":
		fallthrough
	case "BE":
		fallthrough
	case "BR":
		fallthrough
	case "CA":
		fallthrough
	case "DE":
		fallthrough
	case "EG":
		fallthrough
	case "ES":
		fallthrough
	case "FR":
		fallthrough
	case "GB":
		fallthrough
	case "IN":
		fallthrough
	case "IT":
		fallthrough
	case "JP":
		fallthrough
	case "MX":
		fallthrough
	case "NL":
		fallthrough
	case "PL":
		fallthrough
	case "SA":
		fallthrough
	case "SE":
		fallthrough
	case "SG":
		fallthrough
	case "TR":
		fallthrough
	case "UK":
		fallthrough
	case "US":
		*e = SourceAmazonSellerPartnerAWSRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonSellerPartnerAWSRegion: %v", v)
	}
}

type AmazonSellerPartner string

const (
	AmazonSellerPartnerAmazonSellerPartner AmazonSellerPartner = "amazon-seller-partner"
)

func (e AmazonSellerPartner) ToPointer() *AmazonSellerPartner {
	return &e
}

func (e *AmazonSellerPartner) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amazon-seller-partner":
		*e = AmazonSellerPartner(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AmazonSellerPartner: %v", v)
	}
}

type SourceAmazonSellerPartner struct {
	// Type of the Account you're going to authorize the Airbyte application by
	AccountType *SourceAmazonSellerPartnerAWSSellerPartnerAccountType `default:"Seller" json:"account_type"`
	// Additional information to configure report options. This varies by report type, not every report implement this kind of feature. Must be a valid json string.
	AdvancedStreamOptions *string                            `json:"advanced_stream_options,omitempty"`
	authType              *SourceAmazonSellerPartnerAuthType `const:"oauth2.0" json:"auth_type,omitempty"`
	// Select the AWS Environment.
	AwsEnvironment *SourceAmazonSellerPartnerAWSEnvironment `default:"PRODUCTION" json:"aws_environment"`
	// Your Login with Amazon Client ID.
	LwaAppID string `json:"lwa_app_id"`
	// Your Login with Amazon Client Secret.
	LwaClientSecret string `json:"lwa_client_secret"`
	// Will be used for stream slicing for initial full_refresh sync when no updated state is present for reports that support sliced incremental sync.
	PeriodInDays *int64 `default:"90" json:"period_in_days"`
	// The Refresh Token obtained via OAuth flow authorization.
	RefreshToken string `json:"refresh_token"`
	// Select the AWS Region.
	Region *SourceAmazonSellerPartnerAWSRegion `default:"US" json:"region"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
	ReplicationEndDate *time.Time `json:"replication_end_date,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	ReplicationStartDate time.Time `json:"replication_start_date"`
	// Additional information passed to reports. This varies by report type. Must be a valid json string.
	ReportOptions *string             `json:"report_options,omitempty"`
	sourceType    AmazonSellerPartner `const:"amazon-seller-partner" json:"sourceType"`
}

func (s SourceAmazonSellerPartner) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAmazonSellerPartner) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAmazonSellerPartner) GetAccountType() *SourceAmazonSellerPartnerAWSSellerPartnerAccountType {
	if o == nil {
		return nil
	}
	return o.AccountType
}

func (o *SourceAmazonSellerPartner) GetAdvancedStreamOptions() *string {
	if o == nil {
		return nil
	}
	return o.AdvancedStreamOptions
}

func (o *SourceAmazonSellerPartner) GetAuthType() *SourceAmazonSellerPartnerAuthType {
	return SourceAmazonSellerPartnerAuthTypeOauth20.ToPointer()
}

func (o *SourceAmazonSellerPartner) GetAwsEnvironment() *SourceAmazonSellerPartnerAWSEnvironment {
	if o == nil {
		return nil
	}
	return o.AwsEnvironment
}

func (o *SourceAmazonSellerPartner) GetLwaAppID() string {
	if o == nil {
		return ""
	}
	return o.LwaAppID
}

func (o *SourceAmazonSellerPartner) GetLwaClientSecret() string {
	if o == nil {
		return ""
	}
	return o.LwaClientSecret
}

func (o *SourceAmazonSellerPartner) GetPeriodInDays() *int64 {
	if o == nil {
		return nil
	}
	return o.PeriodInDays
}

func (o *SourceAmazonSellerPartner) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceAmazonSellerPartner) GetRegion() *SourceAmazonSellerPartnerAWSRegion {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *SourceAmazonSellerPartner) GetReplicationEndDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ReplicationEndDate
}

func (o *SourceAmazonSellerPartner) GetReplicationStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ReplicationStartDate
}

func (o *SourceAmazonSellerPartner) GetReportOptions() *string {
	if o == nil {
		return nil
	}
	return o.ReportOptions
}

func (o *SourceAmazonSellerPartner) GetSourceType() AmazonSellerPartner {
	return AmazonSellerPartnerAmazonSellerPartner
}
