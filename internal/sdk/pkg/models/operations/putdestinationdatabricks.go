// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationDatabricksRequest struct {
	DestinationDatabricksPutRequest *shared.DestinationDatabricksPutRequest `request:"mediaType=application/json"`
	DestinationID                   string                                  `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationDatabricksRequest) GetDestinationDatabricksPutRequest() *shared.DestinationDatabricksPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationDatabricksPutRequest
}

func (o *PutDestinationDatabricksRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationDatabricksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationDatabricksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationDatabricksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationDatabricksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
