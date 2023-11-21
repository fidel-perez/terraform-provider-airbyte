// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTrelloRequest struct {
	SourceTrelloPutRequest *shared.SourceTrelloPutRequest `request:"mediaType=application/json"`
	SourceID               string                         `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTrelloRequest) GetSourceTrelloPutRequest() *shared.SourceTrelloPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTrelloPutRequest
}

func (o *PutSourceTrelloRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTrelloResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTrelloResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTrelloResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTrelloResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
