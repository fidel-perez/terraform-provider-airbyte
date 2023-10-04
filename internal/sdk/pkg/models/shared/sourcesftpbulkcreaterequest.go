// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceSftpBulkCreateRequest struct {
	Configuration SourceSftpBulk `json:"configuration"`
	Name          string         `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}

func (o *SourceSftpBulkCreateRequest) GetConfiguration() SourceSftpBulk {
	if o == nil {
		return SourceSftpBulk{}
	}
	return o.Configuration
}

func (o *SourceSftpBulkCreateRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceSftpBulkCreateRequest) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceSftpBulkCreateRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
