// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceLinkedinAdsRequest struct {
	SourceLinkedinAdsPutRequest *shared.SourceLinkedinAdsPutRequest `request:"mediaType=application/json"`
	SourceID                    string                              `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceLinkedinAdsRequest) GetSourceLinkedinAdsPutRequest() *shared.SourceLinkedinAdsPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceLinkedinAdsPutRequest
}

func (o *PutSourceLinkedinAdsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceLinkedinAdsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceLinkedinAdsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceLinkedinAdsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceLinkedinAdsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
