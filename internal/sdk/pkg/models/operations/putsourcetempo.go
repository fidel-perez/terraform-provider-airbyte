// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTempoRequest struct {
	SourceTempoPutRequest *shared.SourceTempoPutRequest `request:"mediaType=application/json"`
	SourceID              string                        `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTempoRequest) GetSourceTempoPutRequest() *shared.SourceTempoPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTempoPutRequest
}

func (o *PutSourceTempoRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTempoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTempoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTempoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTempoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
