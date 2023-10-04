// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

type SourceFreshdeskFreshdesk string

const (
	SourceFreshdeskFreshdeskFreshdesk SourceFreshdeskFreshdesk = "freshdesk"
)

func (e SourceFreshdeskFreshdesk) ToPointer() *SourceFreshdeskFreshdesk {
	return &e
}

func (e *SourceFreshdeskFreshdesk) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "freshdesk":
		*e = SourceFreshdeskFreshdesk(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFreshdeskFreshdesk: %v", v)
	}
}

type SourceFreshdesk struct {
	// Freshdesk API Key. See the <a href="https://docs.airbyte.com/integrations/sources/freshdesk">docs</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Freshdesk domain
	Domain string `json:"domain"`
	// The number of requests per minute that this source allowed to use. There is a rate limit of 50 requests per minute per app per account.
	RequestsPerMinute *int64                   `json:"requests_per_minute,omitempty"`
	sourceType        SourceFreshdeskFreshdesk `const:"freshdesk" json:"sourceType"`
	// UTC date and time. Any data created after this date will be replicated. If this parameter is not set, all data will be replicated.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceFreshdesk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFreshdesk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceFreshdesk) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceFreshdesk) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *SourceFreshdesk) GetRequestsPerMinute() *int64 {
	if o == nil {
		return nil
	}
	return o.RequestsPerMinute
}

func (o *SourceFreshdesk) GetSourceType() SourceFreshdeskFreshdesk {
	return SourceFreshdeskFreshdeskFreshdesk
}

func (o *SourceFreshdesk) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
