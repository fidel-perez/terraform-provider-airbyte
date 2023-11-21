// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourcePrestashopUpdate struct {
	// Your PrestaShop access key. See <a href="https://devdocs.prestashop.com/1.7/webservice/tutorials/creating-access/#create-an-access-key"> the docs </a> for info on how to obtain this.
	AccessKey string `json:"access_key"`
	// The Start date in the format YYYY-MM-DD.
	StartDate types.Date `json:"start_date"`
	// Shop URL without trailing slash.
	URL string `json:"url"`
}

func (s SourcePrestashopUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePrestashopUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePrestashopUpdate) GetAccessKey() string {
	if o == nil {
		return ""
	}
	return o.AccessKey
}

func (o *SourcePrestashopUpdate) GetStartDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.StartDate
}

func (o *SourcePrestashopUpdate) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
