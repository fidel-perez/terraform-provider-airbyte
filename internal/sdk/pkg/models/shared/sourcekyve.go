// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Kyve string

const (
	KyveKyve Kyve = "kyve"
)

func (e Kyve) ToPointer() *Kyve {
	return &e
}

func (e *Kyve) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kyve":
		*e = Kyve(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Kyve: %v", v)
	}
}

type SourceKyve struct {
	// The maximum amount of pages to go trough. Set to 'null' for all pages.
	MaxPages *int64 `json:"max_pages,omitempty"`
	// The pagesize for pagination, smaller numbers are used in integration tests.
	PageSize *int64 `default:"100" json:"page_size"`
	// The IDs of the KYVE storage pool you want to archive. (Comma separated)
	PoolIds    string `json:"pool_ids"`
	sourceType Kyve   `const:"kyve" json:"sourceType"`
	// The start-id defines, from which bundle id the pipeline should start to extract the data (Comma separated)
	StartIds string `json:"start_ids"`
	// URL to the KYVE Chain API.
	URLBase *string `default:"https://api.korellia.kyve.network" json:"url_base"`
}

func (s SourceKyve) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceKyve) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceKyve) GetMaxPages() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxPages
}

func (o *SourceKyve) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SourceKyve) GetPoolIds() string {
	if o == nil {
		return ""
	}
	return o.PoolIds
}

func (o *SourceKyve) GetSourceType() Kyve {
	return KyveKyve
}

func (o *SourceKyve) GetStartIds() string {
	if o == nil {
		return ""
	}
	return o.StartIds
}

func (o *SourceKyve) GetURLBase() *string {
	if o == nil {
		return nil
	}
	return o.URLBase
}
