// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceSalesforceRequest struct {
	SourceSalesforcePutRequest *shared.SourceSalesforcePutRequest `request:"mediaType=application/json"`
	SourceID                   string                             `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceSalesforceRequest) GetSourceSalesforcePutRequest() *shared.SourceSalesforcePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceSalesforcePutRequest
}

func (o *PutSourceSalesforceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceSalesforceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceSalesforceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceSalesforceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceSalesforceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
