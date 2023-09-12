// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceS3CreateRequest struct {
	// NOTE: When this Spec is changed, legacy_config_transformer.py must also be modified to uptake the changes
	// because it is responsible for converting legacy S3 v3 configs into v4 configs using the File-Based CDK.
	Configuration SourceS3 `json:"configuration"`
	Name          string   `json:"name"`
	// Optional secretID obtained through the public API OAuth redirect flow.
	SecretID    *string `json:"secretId,omitempty"`
	WorkspaceID string  `json:"workspaceId"`
}
