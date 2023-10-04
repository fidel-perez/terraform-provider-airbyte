// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTwitterRequest struct {
	SourceTwitterPutRequest *shared.SourceTwitterPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTwitterRequest) GetSourceTwitterPutRequest() *shared.SourceTwitterPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTwitterPutRequest
}

func (o *PutSourceTwitterRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTwitterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTwitterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTwitterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTwitterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
