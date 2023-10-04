// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceGcsRequest struct {
	SourceGcsPutRequest *shared.SourceGcsPutRequest `request:"mediaType=application/json"`
	SourceID            string                      `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceGcsRequest) GetSourceGcsPutRequest() *shared.SourceGcsPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceGcsPutRequest
}

func (o *PutSourceGcsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceGcsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceGcsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceGcsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceGcsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
