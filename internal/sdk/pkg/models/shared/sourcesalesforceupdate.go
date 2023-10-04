// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

type SourceSalesforceUpdateAuthType string

const (
	SourceSalesforceUpdateAuthTypeClient SourceSalesforceUpdateAuthType = "Client"
)

func (e SourceSalesforceUpdateAuthType) ToPointer() *SourceSalesforceUpdateAuthType {
	return &e
}

func (e *SourceSalesforceUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceSalesforceUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceUpdateAuthType: %v", v)
	}
}

type SourceSalesforceUpdateStreamsCriteriaSearchCriteria string

const (
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaStartsWith    SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "starts with"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaEndsWith      SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "ends with"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaContains      SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "contains"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaExacts        SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "exacts"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaStartsNotWith SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "starts not with"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaEndsNotWith   SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "ends not with"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaNotContains   SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "not contains"
	SourceSalesforceUpdateStreamsCriteriaSearchCriteriaNotExacts     SourceSalesforceUpdateStreamsCriteriaSearchCriteria = "not exacts"
)

func (e SourceSalesforceUpdateStreamsCriteriaSearchCriteria) ToPointer() *SourceSalesforceUpdateStreamsCriteriaSearchCriteria {
	return &e
}

func (e *SourceSalesforceUpdateStreamsCriteriaSearchCriteria) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "starts with":
		fallthrough
	case "ends with":
		fallthrough
	case "contains":
		fallthrough
	case "exacts":
		fallthrough
	case "starts not with":
		fallthrough
	case "ends not with":
		fallthrough
	case "not contains":
		fallthrough
	case "not exacts":
		*e = SourceSalesforceUpdateStreamsCriteriaSearchCriteria(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceUpdateStreamsCriteriaSearchCriteria: %v", v)
	}
}

type SourceSalesforceUpdateStreamsCriteria struct {
	Criteria *SourceSalesforceUpdateStreamsCriteriaSearchCriteria `default:"contains" json:"criteria"`
	Value    string                                               `json:"value"`
}

func (s SourceSalesforceUpdateStreamsCriteria) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSalesforceUpdateStreamsCriteria) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSalesforceUpdateStreamsCriteria) GetCriteria() *SourceSalesforceUpdateStreamsCriteriaSearchCriteria {
	if o == nil {
		return nil
	}
	return o.Criteria
}

func (o *SourceSalesforceUpdateStreamsCriteria) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type SourceSalesforceUpdate struct {
	authType *SourceSalesforceUpdateAuthType `const:"Client" json:"auth_type,omitempty"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client ID</a>
	ClientID string `json:"client_id"`
	// Enter your Salesforce developer application's <a href="https://developer.salesforce.com/forums/?id=9062I000000DLgbQAG">Client secret</a>
	ClientSecret string `json:"client_secret"`
	// Toggle to use Bulk API (this might cause empty fields for some streams)
	ForceUseBulkAPI *bool `default:"false" json:"force_use_bulk_api"`
	// Toggle if you're using a <a href="https://help.salesforce.com/s/articleView?id=sf.deploy_sandboxes_parent.htm&type=5">Salesforce Sandbox</a>
	IsSandbox *bool `default:"false" json:"is_sandbox"`
	// Enter your application's <a href="https://developer.salesforce.com/docs/atlas.en-us.mobile_sdk.meta/mobile_sdk/oauth_refresh_token_flow.htm">Salesforce Refresh Token</a> used for Airbyte to access your Salesforce account.
	RefreshToken string `json:"refresh_token"`
	// Enter the date (or date-time) in the YYYY-MM-DD or YYYY-MM-DDTHH:mm:ssZ format. Airbyte will replicate the data updated on and after this date. If this field is blank, Airbyte will replicate the data for last two years.
	StartDate *time.Time `json:"start_date,omitempty"`
	// Add filters to select only required stream based on `SObject` name. Use this field to filter which tables are displayed by this connector. This is useful if your Salesforce account has a large number of tables (>1000), in which case you may find it easier to navigate the UI and speed up the connector's performance if you restrict the tables displayed by this connector.
	StreamsCriteria []SourceSalesforceUpdateStreamsCriteria `json:"streams_criteria,omitempty"`
}

func (s SourceSalesforceUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSalesforceUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSalesforceUpdate) GetAuthType() *SourceSalesforceUpdateAuthType {
	return SourceSalesforceUpdateAuthTypeClient.ToPointer()
}

func (o *SourceSalesforceUpdate) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *SourceSalesforceUpdate) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *SourceSalesforceUpdate) GetForceUseBulkAPI() *bool {
	if o == nil {
		return nil
	}
	return o.ForceUseBulkAPI
}

func (o *SourceSalesforceUpdate) GetIsSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.IsSandbox
}

func (o *SourceSalesforceUpdate) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

func (o *SourceSalesforceUpdate) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}

func (o *SourceSalesforceUpdate) GetStreamsCriteria() []SourceSalesforceUpdateStreamsCriteria {
	if o == nil {
		return nil
	}
	return o.StreamsCriteria
}
