// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceMailchimpPutRequest struct {
	Configuration SourceMailchimpUpdate `json:"configuration"`
	Name          string                `json:"name"`
	WorkspaceID   string                `json:"workspaceId"`
}

func (o *SourceMailchimpPutRequest) GetConfiguration() SourceMailchimpUpdate {
	if o == nil {
		return SourceMailchimpUpdate{}
	}
	return o.Configuration
}

func (o *SourceMailchimpPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceMailchimpPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
