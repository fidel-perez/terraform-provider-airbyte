// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSapFieldglassCreateRequest struct {
	Configuration SourceSapFieldglass `json:"configuration"`
	Name          string              `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceSapFieldglassCreateRequest) GetConfiguration() SourceSapFieldglass {
	if o == nil {
		return SourceSapFieldglass{}
	}
	return o.Configuration
}

func (o *SourceSapFieldglassCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSapFieldglassCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceSapFieldglassCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
