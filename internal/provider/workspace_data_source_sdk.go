// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *WorkspaceDataSourceModel) RefreshFromGetResponse(resp *shared.WorkspaceResponse) {
	if resp.DataResidency != nil {
		r.DataResidency = types.StringValue(string(*resp.DataResidency))
	} else {
		r.DataResidency = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
