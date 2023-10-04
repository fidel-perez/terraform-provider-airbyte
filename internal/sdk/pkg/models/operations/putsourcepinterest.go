// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourcePinterestRequest struct {
	SourcePinterestPutRequest *shared.SourcePinterestPutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourcePinterestRequest) GetSourcePinterestPutRequest() *shared.SourcePinterestPutRequest {
	if o == nil {
		return nil
	}
	return o.SourcePinterestPutRequest
}

func (o *PutSourcePinterestRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourcePinterestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourcePinterestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourcePinterestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourcePinterestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
