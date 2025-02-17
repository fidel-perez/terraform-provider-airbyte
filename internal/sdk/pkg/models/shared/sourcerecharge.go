// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type Recharge string

const (
	RechargeRecharge Recharge = "recharge"
)

func (e Recharge) ToPointer() *Recharge {
	return &e
}

func (e *Recharge) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "recharge":
		*e = Recharge(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Recharge: %v", v)
	}
}

type SourceRecharge struct {
	// The value of the Access Token generated. See the <a href="https://docs.airbyte.com/integrations/sources/recharge">docs</a> for more information.
	AccessToken string   `json:"access_token"`
	sourceType  Recharge `const:"recharge" json:"sourceType"`
	// The date from which you'd like to replicate data for Recharge API, in the format YYYY-MM-DDT00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceRecharge) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRecharge) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRecharge) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourceRecharge) GetSourceType() Recharge {
	return RechargeRecharge
}

func (o *SourceRecharge) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
