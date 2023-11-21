// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceAwsCloudtrailRequest struct {
	SourceAwsCloudtrailPutRequest *shared.SourceAwsCloudtrailPutRequest `request:"mediaType=application/json"`
	SourceID                      string                                `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceAwsCloudtrailRequest) GetSourceAwsCloudtrailPutRequest() *shared.SourceAwsCloudtrailPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceAwsCloudtrailPutRequest
}

func (o *PutSourceAwsCloudtrailRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceAwsCloudtrailResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceAwsCloudtrailResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceAwsCloudtrailResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceAwsCloudtrailResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
