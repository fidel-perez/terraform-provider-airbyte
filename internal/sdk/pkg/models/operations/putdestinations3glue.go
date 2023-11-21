// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationS3GlueRequest struct {
	DestinationS3GluePutRequest *shared.DestinationS3GluePutRequest `request:"mediaType=application/json"`
	DestinationID               string                              `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationS3GlueRequest) GetDestinationS3GluePutRequest() *shared.DestinationS3GluePutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationS3GluePutRequest
}

func (o *PutDestinationS3GlueRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationS3GlueResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationS3GlueResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationS3GlueResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationS3GlueResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
