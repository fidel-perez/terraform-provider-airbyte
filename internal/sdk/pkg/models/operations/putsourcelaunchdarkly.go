// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceLaunchdarklyRequest struct {
	SourceLaunchdarklyPutRequest *shared.SourceLaunchdarklyPutRequest `request:"mediaType=application/json"`
	SourceID                     string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceLaunchdarklyRequest) GetSourceLaunchdarklyPutRequest() *shared.SourceLaunchdarklyPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceLaunchdarklyPutRequest
}

func (o *PutSourceLaunchdarklyRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceLaunchdarklyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceLaunchdarklyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceLaunchdarklyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceLaunchdarklyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
