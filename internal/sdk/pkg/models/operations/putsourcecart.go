// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceCartRequest struct {
	SourceCartPutRequest *shared.SourceCartPutRequest `request:"mediaType=application/json"`
	SourceID             string                       `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceCartRequest) GetSourceCartPutRequest() *shared.SourceCartPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceCartPutRequest
}

func (o *PutSourceCartRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceCartResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceCartResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceCartResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceCartResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
