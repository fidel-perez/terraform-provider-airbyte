// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceWebflowPutRequest struct {
	Configuration SourceWebflowUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourceWebflowPutRequest) GetConfiguration() SourceWebflowUpdate {
	if o == nil {
		return SourceWebflowUpdate{}
	}
	return o.Configuration
}

func (o *SourceWebflowPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceWebflowPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
