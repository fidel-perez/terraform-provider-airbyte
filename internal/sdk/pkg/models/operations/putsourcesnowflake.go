// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceSnowflakeRequest struct {
	SourceSnowflakePutRequest *shared.SourceSnowflakePutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceSnowflakeRequest) GetSourceSnowflakePutRequest() *shared.SourceSnowflakePutRequest {
	if o == nil {
		return nil
	}
	return o.SourceSnowflakePutRequest
}

func (o *PutSourceSnowflakeRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceSnowflakeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceSnowflakeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceSnowflakeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceSnowflakeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
