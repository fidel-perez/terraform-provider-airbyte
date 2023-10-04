// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceBigqueryCreateRequest struct {
	Configuration SourceBigquery `json:"configuration"`
	Name          string         `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceBigqueryCreateRequest) GetConfiguration() SourceBigquery {
	if o == nil {
		return SourceBigquery{}
	}
	return o.Configuration
}

func (o *SourceBigqueryCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceBigqueryCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceBigqueryCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
