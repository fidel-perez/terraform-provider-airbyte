// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceGetlagoCreateRequest struct {
	Configuration SourceGetlago `json:"configuration"`
	Name          string        `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceGetlagoCreateRequest) GetConfiguration() SourceGetlago {
	if o == nil {
		return SourceGetlago{}
	}
	return o.Configuration
}

func (o *SourceGetlagoCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceGetlagoCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceGetlagoCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
