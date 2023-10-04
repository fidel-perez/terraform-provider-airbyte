// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
	"time"
)

type SourcePaystackPaystack string

const (
	SourcePaystackPaystackPaystack SourcePaystackPaystack = "paystack"
)

func (e SourcePaystackPaystack) ToPointer() *SourcePaystackPaystack {
	return &e
}

func (e *SourcePaystackPaystack) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "paystack":
		*e = SourcePaystackPaystack(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePaystackPaystack: %v", v)
	}
}

type SourcePaystack struct {
	// When set, the connector will always reload data from the past N days, where N is the value set here. This is useful if your data is updated after creation.
	LookbackWindowDays *int64 `default:"0" json:"lookback_window_days"`
	// The Paystack API key (usually starts with 'sk_live_'; find yours <a href="https://dashboard.paystack.com/#/settings/developer">here</a>).
	SecretKey  string                 `json:"secret_key"`
	sourceType SourcePaystackPaystack `const:"paystack" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourcePaystack) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePaystack) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePaystack) GetLookbackWindowDays() *int64 {
	if o == nil {
		return nil
	}
	return o.LookbackWindowDays
}

func (o *SourcePaystack) GetSecretKey() string {
	if o == nil {
		return ""
	}
	return o.SecretKey
}

func (o *SourcePaystack) GetSourceType() SourcePaystackPaystack {
	return SourcePaystackPaystackPaystack
}

func (o *SourcePaystack) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
