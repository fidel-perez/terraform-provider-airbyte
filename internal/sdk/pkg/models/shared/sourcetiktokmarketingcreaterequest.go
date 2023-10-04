// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceTiktokMarketingCreateRequest struct {
	Configuration SourceTiktokMarketing `json:"configuration"`
	Name          string                `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceTiktokMarketingCreateRequest) GetConfiguration() SourceTiktokMarketing {
	if o == nil {
		return SourceTiktokMarketing{}
	}
	return o.Configuration
}

func (o *SourceTiktokMarketingCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceTiktokMarketingCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceTiktokMarketingCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
