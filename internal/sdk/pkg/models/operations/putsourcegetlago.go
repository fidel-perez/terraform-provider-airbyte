// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceGetlagoRequest struct {
	SourceGetlagoPutRequest *shared.SourceGetlagoPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceGetlagoRequest) GetSourceGetlagoPutRequest() *shared.SourceGetlagoPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceGetlagoPutRequest
}

func (o *PutSourceGetlagoRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceGetlagoResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceGetlagoResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceGetlagoResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceGetlagoResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
