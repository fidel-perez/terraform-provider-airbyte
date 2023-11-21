// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// SourceCoinAPIEnvironment - The environment to use. Either sandbox or production.
type SourceCoinAPIEnvironment string

const (
	SourceCoinAPIEnvironmentSandbox    SourceCoinAPIEnvironment = "sandbox"
	SourceCoinAPIEnvironmentProduction SourceCoinAPIEnvironment = "production"
)

func (e SourceCoinAPIEnvironment) ToPointer() *SourceCoinAPIEnvironment {
	return &e
}

func (e *SourceCoinAPIEnvironment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sandbox":
		fallthrough
	case "production":
		*e = SourceCoinAPIEnvironment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCoinAPIEnvironment: %v", v)
	}
}

type CoinAPI string

const (
	CoinAPICoinAPI CoinAPI = "coin-api"
)

func (e CoinAPI) ToPointer() *CoinAPI {
	return &e
}

func (e *CoinAPI) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "coin-api":
		*e = CoinAPI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CoinAPI: %v", v)
	}
}

type SourceCoinAPI struct {
	// API Key
	APIKey string `json:"api_key"`
	// The end date in ISO 8601 format. If not supplied, data will be returned
	// from the start date to the current time, or when the count of result
	// elements reaches its limit.
	//
	EndDate *string `json:"end_date,omitempty"`
	// The environment to use. Either sandbox or production.
	//
	Environment *SourceCoinAPIEnvironment `default:"sandbox" json:"environment"`
	// The maximum number of elements to return. If not supplied, the default
	// is 100. For numbers larger than 100, each 100 items is counted as one
	// request for pricing purposes. Maximum value is 100000.
	//
	Limit *int64 `default:"100" json:"limit"`
	// The period to use. See the documentation for a list. https://docs.coinapi.io/#list-all-periods-get
	Period     string  `json:"period"`
	sourceType CoinAPI `const:"coin-api" json:"sourceType"`
	// The start date in ISO 8601 format.
	StartDate string `json:"start_date"`
	// The symbol ID to use. See the documentation for a list.
	// https://docs.coinapi.io/#list-all-symbols-get
	//
	SymbolID string `json:"symbol_id"`
}

func (s SourceCoinAPI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceCoinAPI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceCoinAPI) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceCoinAPI) GetEndDate() *string {
	if o == nil {
		return nil
	}
	return o.EndDate
}

func (o *SourceCoinAPI) GetEnvironment() *SourceCoinAPIEnvironment {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *SourceCoinAPI) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *SourceCoinAPI) GetPeriod() string {
	if o == nil {
		return ""
	}
	return o.Period
}

func (o *SourceCoinAPI) GetSourceType() CoinAPI {
	return CoinAPICoinAPI
}

func (o *SourceCoinAPI) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}

func (o *SourceCoinAPI) GetSymbolID() string {
	if o == nil {
		return ""
	}
	return o.SymbolID
}
