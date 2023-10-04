// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceYoutubeAnalyticsRequest struct {
	SourceYoutubeAnalyticsPutRequest *shared.SourceYoutubeAnalyticsPutRequest `request:"mediaType=application/json"`
	SourceID                         string                                   `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceYoutubeAnalyticsRequest) GetSourceYoutubeAnalyticsPutRequest() *shared.SourceYoutubeAnalyticsPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceYoutubeAnalyticsPutRequest
}

func (o *PutSourceYoutubeAnalyticsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceYoutubeAnalyticsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceYoutubeAnalyticsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceYoutubeAnalyticsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceYoutubeAnalyticsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
