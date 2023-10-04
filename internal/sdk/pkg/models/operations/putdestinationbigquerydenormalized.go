// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationBigqueryDenormalizedRequest struct {
	DestinationBigqueryDenormalizedPutRequest *shared.DestinationBigqueryDenormalizedPutRequest `request:"mediaType=application/json"`
	DestinationID                             string                                            `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationBigqueryDenormalizedRequest) GetDestinationBigqueryDenormalizedPutRequest() *shared.DestinationBigqueryDenormalizedPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationBigqueryDenormalizedPutRequest
}

func (o *PutDestinationBigqueryDenormalizedRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationBigqueryDenormalizedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationBigqueryDenormalizedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationBigqueryDenormalizedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationBigqueryDenormalizedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
