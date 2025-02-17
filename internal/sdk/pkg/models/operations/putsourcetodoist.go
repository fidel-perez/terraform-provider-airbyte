// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTodoistRequest struct {
	SourceTodoistPutRequest *shared.SourceTodoistPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *PutSourceTodoistRequest) GetSourceTodoistPutRequest() *shared.SourceTodoistPutRequest {
	if o == nil {
		return nil
	}
	return o.SourceTodoistPutRequest
}

func (o *PutSourceTodoistRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

type PutSourceTodoistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutSourceTodoistResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutSourceTodoistResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutSourceTodoistResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
