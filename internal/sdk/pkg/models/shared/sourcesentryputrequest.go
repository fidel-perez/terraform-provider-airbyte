// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSentryPutRequest struct {
	Configuration SourceSentryUpdate `json:"configuration"`
	Name          string             `json:"name"`
	WorkspaceID   string             `json:"workspaceId"`
}

func (o *SourceSentryPutRequest) GetConfiguration() SourceSentryUpdate {
	if o == nil {
		return SourceSentryUpdate{}
	}
	return o.Configuration
}

func (o *SourceSentryPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSentryPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
