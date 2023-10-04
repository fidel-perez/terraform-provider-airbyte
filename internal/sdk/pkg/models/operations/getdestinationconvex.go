// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationConvexRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *GetDestinationConvexRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type GetDestinationConvexResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Get a Destination by the id in the path.
	DestinationResponse *shared.DestinationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDestinationConvexResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDestinationConvexResponse) GetDestinationResponse() *shared.DestinationResponse {
	if o == nil {
		return nil
	}
	return o.DestinationResponse
}

func (o *GetDestinationConvexResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDestinationConvexResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
