// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceConfigcatCreateRequest struct {
	Configuration SourceConfigcat `json:"configuration"`
	Name          string          `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceConfigcatCreateRequest) GetConfiguration() SourceConfigcat {
	if o == nil {
		return SourceConfigcat{}
	}
	return o.Configuration
}

func (o *SourceConfigcatCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceConfigcatCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceConfigcatCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
