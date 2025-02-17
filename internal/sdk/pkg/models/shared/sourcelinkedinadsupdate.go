// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// PivotCategory - Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives.
type PivotCategory string

const (
	PivotCategoryCompany                     PivotCategory = "COMPANY"
	PivotCategoryAccount                     PivotCategory = "ACCOUNT"
	PivotCategoryShare                       PivotCategory = "SHARE"
	PivotCategoryCampaign                    PivotCategory = "CAMPAIGN"
	PivotCategoryCreative                    PivotCategory = "CREATIVE"
	PivotCategoryCampaignGroup               PivotCategory = "CAMPAIGN_GROUP"
	PivotCategoryConversion                  PivotCategory = "CONVERSION"
	PivotCategoryConversationNode            PivotCategory = "CONVERSATION_NODE"
	PivotCategoryConversationNodeOptionIndex PivotCategory = "CONVERSATION_NODE_OPTION_INDEX"
	PivotCategoryServingLocation             PivotCategory = "SERVING_LOCATION"
	PivotCategoryCardIndex                   PivotCategory = "CARD_INDEX"
	PivotCategoryMemberCompanySize           PivotCategory = "MEMBER_COMPANY_SIZE"
	PivotCategoryMemberIndustry              PivotCategory = "MEMBER_INDUSTRY"
	PivotCategoryMemberSeniority             PivotCategory = "MEMBER_SENIORITY"
	PivotCategoryMemberJobTitle              PivotCategory = "MEMBER_JOB_TITLE "
	PivotCategoryMemberJobFunction           PivotCategory = "MEMBER_JOB_FUNCTION "
	PivotCategoryMemberCountryV2             PivotCategory = "MEMBER_COUNTRY_V2 "
	PivotCategoryMemberRegionV2              PivotCategory = "MEMBER_REGION_V2"
	PivotCategoryMemberCompany               PivotCategory = "MEMBER_COMPANY"
	PivotCategoryPlacementName               PivotCategory = "PLACEMENT_NAME"
	PivotCategoryImpressionDeviceType        PivotCategory = "IMPRESSION_DEVICE_TYPE"
)

func (e PivotCategory) ToPointer() *PivotCategory {
	return &e
}

func (e *PivotCategory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPANY":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "SHARE":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CREATIVE":
		fallthrough
	case "CAMPAIGN_GROUP":
		fallthrough
	case "CONVERSION":
		fallthrough
	case "CONVERSATION_NODE":
		fallthrough
	case "CONVERSATION_NODE_OPTION_INDEX":
		fallthrough
	case "SERVING_LOCATION":
		fallthrough
	case "CARD_INDEX":
		fallthrough
	case "MEMBER_COMPANY_SIZE":
		fallthrough
	case "MEMBER_INDUSTRY":
		fallthrough
	case "MEMBER_SENIORITY":
		fallthrough
	case "MEMBER_JOB_TITLE ":
		fallthrough
	case "MEMBER_JOB_FUNCTION ":
		fallthrough
	case "MEMBER_COUNTRY_V2 ":
		fallthrough
	case "MEMBER_REGION_V2":
		fallthrough
	case "MEMBER_COMPANY":
		fallthrough
	case "PLACEMENT_NAME":
		fallthrough
	case "IMPRESSION_DEVICE_TYPE":
		*e = PivotCategory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PivotCategory: %v", v)
	}
}

// TimeGranularity - Choose how to group the data in your report by time. The options are:<br>- 'ALL': A single result summarizing the entire time range.<br>- 'DAILY': Group results by each day.<br>- 'MONTHLY': Group results by each month.<br>- 'YEARLY': Group results by each year.<br>Selecting a time grouping helps you analyze trends and patterns over different time periods.
type TimeGranularity string

const (
	TimeGranularityAll     TimeGranularity = "ALL"
	TimeGranularityDaily   TimeGranularity = "DAILY"
	TimeGranularityMonthly TimeGranularity = "MONTHLY"
	TimeGranularityYearly  TimeGranularity = "YEARLY"
)

func (e TimeGranularity) ToPointer() *TimeGranularity {
	return &e
}

func (e *TimeGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "DAILY":
		fallthrough
	case "MONTHLY":
		fallthrough
	case "YEARLY":
		*e = TimeGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TimeGranularity: %v", v)
	}
}

// AdAnalyticsReportConfiguration - Config for custom ad Analytics Report
type AdAnalyticsReportConfiguration struct {
	// The name for the custom report.
	Name string `json:"name"`
	// Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives.
	PivotBy PivotCategory `json:"pivot_by"`
	// Choose how to group the data in your report by time. The options are:<br>- 'ALL': A single result summarizing the entire time range.<br>- 'DAILY': Group results by each day.<br>- 'MONTHLY': Group results by each month.<br>- 'YEARLY': Group results by each year.<br>Selecting a time grouping helps you analyze trends and patterns over different time periods.
	TimeGranularity TimeGranularity `json:"time_granularity"`
}

func (o *AdAnalyticsReportConfiguration) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AdAnalyticsReportConfiguration) GetPivotBy() PivotCategory {
	if o == nil {
		return PivotCategory("")
	}
	return o.PivotBy
}

func (o *AdAnalyticsReportConfiguration) GetTimeGranularity() TimeGranularity {
	if o == nil {
		return TimeGranularity("")
	}
	return o.TimeGranularity
}

type SourceLinkedinAdsUpdateSchemasAuthMethod string

const (
	SourceLinkedinAdsUpdateSchemasAuthMethodAccessToken SourceLinkedinAdsUpdateSchemasAuthMethod = "access_token"
)

func (e SourceLinkedinAdsUpdateSchemasAuthMethod) ToPointer() *SourceLinkedinAdsUpdateSchemasAuthMethod {
	return &e
}

func (e *SourceLinkedinAdsUpdateSchemasAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinAdsUpdateSchemasAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateSchemasAuthMethod: %v", v)
	}
}

type AccessToken struct {
	// The access token generated for your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	AccessToken string                                    `json:"access_token"`
	authMethod  *SourceLinkedinAdsUpdateSchemasAuthMethod `const:"access_token" json:"auth_method,omitempty"`
}

func (a AccessToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccessToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *AccessToken) GetAuthMethod() *SourceLinkedinAdsUpdateSchemasAuthMethod {
	return SourceLinkedinAdsUpdateSchemasAuthMethodAccessToken.ToPointer()
}

type SourceLinkedinAdsUpdateAuthMethod string

const (
	SourceLinkedinAdsUpdateAuthMethodOAuth20 SourceLinkedinAdsUpdateAuthMethod = "oAuth2.0"
)

func (e SourceLinkedinAdsUpdateAuthMethod) ToPointer() *SourceLinkedinAdsUpdateAuthMethod {
	return &e
}

func (e *SourceLinkedinAdsUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinAdsUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateAuthMethod: %v", v)
	}
}

type SourceLinkedinAdsUpdateOAuth20 struct {
	authMethod *SourceLinkedinAdsUpdateAuthMethod `const:"oAuth2.0" json:"auth_method,omitempty"`
	// The client ID of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	ClientID string `json:"client_id"`
	// The client secret of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	RefreshToken string `json:"refresh_token"`
}

func (s SourceLinkedinAdsUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinAdsUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinAdsUpdateOAuth20) GetAuthMethod() *SourceLinkedinAdsUpdateAuthMethod {
	return SourceLinkedinAdsUpdateAuthMethodOAuth20.ToPointer()
}

func (o *SourceLinkedinAdsUpdateOAuth20) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceLinkedinAdsUpdateOAuth20) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceLinkedinAdsUpdateOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type SourceLinkedinAdsUpdateAuthenticationType string

const (
	SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateOAuth20 SourceLinkedinAdsUpdateAuthenticationType = "source-linkedin-ads-update_OAuth2.0"
	SourceLinkedinAdsUpdateAuthenticationTypeAccessToken                    SourceLinkedinAdsUpdateAuthenticationType = "Access Token"
)

type SourceLinkedinAdsUpdateAuthentication struct {
	SourceLinkedinAdsUpdateOAuth20 *SourceLinkedinAdsUpdateOAuth20
	AccessToken                    *AccessToken

	Type SourceLinkedinAdsUpdateAuthenticationType
}

func CreateSourceLinkedinAdsUpdateAuthenticationSourceLinkedinAdsUpdateOAuth20(sourceLinkedinAdsUpdateOAuth20 SourceLinkedinAdsUpdateOAuth20) SourceLinkedinAdsUpdateAuthentication {
	typ := SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateOAuth20

	return SourceLinkedinAdsUpdateAuthentication{
		SourceLinkedinAdsUpdateOAuth20: &sourceLinkedinAdsUpdateOAuth20,
		Type:                           typ,
	}
}

func CreateSourceLinkedinAdsUpdateAuthenticationAccessToken(accessToken AccessToken) SourceLinkedinAdsUpdateAuthentication {
	typ := SourceLinkedinAdsUpdateAuthenticationTypeAccessToken

	return SourceLinkedinAdsUpdateAuthentication{
		AccessToken: &accessToken,
		Type:        typ,
	}
}

func (u *SourceLinkedinAdsUpdateAuthentication) UnmarshalJSON(data []byte) error {

	accessToken := new(AccessToken)
	if err := utils.UnmarshalJSON(data, &accessToken, "", true, true); err == nil {
		u.AccessToken = accessToken
		u.Type = SourceLinkedinAdsUpdateAuthenticationTypeAccessToken
		return nil
	}

	sourceLinkedinAdsUpdateOAuth20 := new(SourceLinkedinAdsUpdateOAuth20)
	if err := utils.UnmarshalJSON(data, &sourceLinkedinAdsUpdateOAuth20, "", true, true); err == nil {
		u.SourceLinkedinAdsUpdateOAuth20 = sourceLinkedinAdsUpdateOAuth20
		u.Type = SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinAdsUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinAdsUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.SourceLinkedinAdsUpdateOAuth20, "", true)
	}

	if u.AccessToken != nil {
		return utils.MarshalJSON(u.AccessToken, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceLinkedinAdsUpdate struct {
	// Specify the account IDs to pull data from, separated by a space. Leave this field empty if you want to pull the data from all accounts accessible by the authenticated user. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn docs</a> to locate these IDs.
	AccountIds         []int64                                `json:"account_ids,omitempty"`
	AdAnalyticsReports []AdAnalyticsReportConfiguration       `json:"ad_analytics_reports,omitempty"`
	Credentials        *SourceLinkedinAdsUpdateAuthentication `json:"credentials,omitempty"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}

func (s SourceLinkedinAdsUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceLinkedinAdsUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceLinkedinAdsUpdate) GetAccountIds() []int64 {
	if o == nil {
		return nil
	}
	return o.AccountIds
}

func (o *SourceLinkedinAdsUpdate) GetAdAnalyticsReports() []AdAnalyticsReportConfiguration {
	if o == nil {
		return nil
	}
	return o.AdAnalyticsReports
}

func (o *SourceLinkedinAdsUpdate) GetCredentials() *SourceLinkedinAdsUpdateAuthentication {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *SourceLinkedinAdsUpdate) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}
