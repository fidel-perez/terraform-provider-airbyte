// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourcePokeapiPutRequest struct {
	Configuration SourcePokeapiUpdate `json:"configuration"`
	Name          string              `json:"name"`
	WorkspaceID   string              `json:"workspaceId"`
}

func (o *SourcePokeapiPutRequest) GetConfiguration() SourcePokeapiUpdate {
	if o == nil {
		return SourcePokeapiUpdate{}
	}
	return o.Configuration
}

func (o *SourcePokeapiPutRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourcePokeapiPutRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
