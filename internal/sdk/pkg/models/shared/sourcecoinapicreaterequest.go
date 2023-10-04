// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceCoinAPICreateRequest struct {
	Configuration SourceCoinAPI `json:"configuration"`
	Name          string        `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceCoinAPICreateRequest) GetConfiguration() SourceCoinAPI {
	if o == nil {
		return SourceCoinAPI{}
	}
	return o.Configuration
}

func (o *SourceCoinAPICreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceCoinAPICreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceCoinAPICreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
