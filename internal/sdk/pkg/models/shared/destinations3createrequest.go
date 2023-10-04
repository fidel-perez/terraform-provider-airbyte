// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationS3CreateRequest struct {
	Configuration DestinationS3 `json:"configuration"`
	Name          string        `json:"name"`
	WorkspaceID   string        `json:"workspaceId"`
}

func (o *DestinationS3CreateRequest) GetConfiguration() DestinationS3 {
	if o == nil {
		return DestinationS3{}
	}
	return o.Configuration
}

func (o *DestinationS3CreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationS3CreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
