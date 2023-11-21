// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceTempoPutRequest struct {
	Configuration SourceTempoUpdate `json:"configuration"`
	Name          string            `json:"name"`
	WorkspaceID   string            `json:"workspaceId"`
}

func (o *SourceTempoPutRequest) GetConfiguration() SourceTempoUpdate {
	if o == nil {
		return SourceTempoUpdate{}
	}
	return o.Configuration
}

func (o *SourceTempoPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceTempoPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
