// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceGetlagoPutRequest struct {
	Configuration SourceGetlagoUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourceGetlagoPutRequest) GetConfiguration() SourceGetlagoUpdate {
	if o == nil {
		return SourceGetlagoUpdate{}
	}
	return o.Configuration
}

func (o *SourceGetlagoPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGetlagoPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
