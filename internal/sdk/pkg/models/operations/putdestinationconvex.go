// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationConvexRequest struct {
	DestinationConvexPutRequest *shared.DestinationConvexPutRequest `request:"mediaType=application/json"`
	DestinationID               string                              `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationConvexRequest) GetDestinationConvexPutRequest() *shared.DestinationConvexPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationConvexPutRequest
}

func (o *PutDestinationConvexRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationConvexResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationConvexResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationConvexResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationConvexResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
