// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
)

type ConnectionCreateRequest struct {
	// A list of configured stream options for a connection.
	Configurations *StreamConfigurations `json:"configurations,omitempty"`
	DataResidency  *GeographyEnum        `default:"auto" json:"dataResidency"`
	DestinationID  string                `json:"destinationId"`
	// Optional name of the connection
	Name *string `json:"name,omitempty"`
	// Define the location where the data will be stored in the destination
	NamespaceDefinition *NamespaceDefinitionEnum `default:"destination" json:"namespaceDefinition"`
	// Used when namespaceDefinition is 'custom_format'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `default:"null" json:"namespaceFormat"`
	// Set how Airbyte handles syncs when it detects a non-breaking schema change in the source
	NonBreakingSchemaUpdatesBehavior *NonBreakingSchemaUpdatesBehaviorEnum `default:"ignore" json:"nonBreakingSchemaUpdatesBehavior"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination (ex. “airbyte_” causes “projects” => “airbyte_projects”).
	Prefix *string `json:"prefix,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	Schedule *ConnectionSchedule   `json:"schedule,omitempty"`
	SourceID string                `json:"sourceId"`
	Status   *ConnectionStatusEnum `json:"status,omitempty"`
}

func (c ConnectionCreateRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionCreateRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionCreateRequest) GetConfigurations() *StreamConfigurations {
	if o == nil {
		return nil
	}
	return o.Configurations
}

func (o *ConnectionCreateRequest) GetDataResidency() *GeographyEnum {
	if o == nil {
		return nil
	}
	return o.DataResidency
}

func (o *ConnectionCreateRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *ConnectionCreateRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionCreateRequest) GetNamespaceDefinition() *NamespaceDefinitionEnum {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionCreateRequest) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionCreateRequest) GetNonBreakingSchemaUpdatesBehavior() *NonBreakingSchemaUpdatesBehaviorEnum {
	if o == nil {
		return nil
	}
	return o.NonBreakingSchemaUpdatesBehavior
}

func (o *ConnectionCreateRequest) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionCreateRequest) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionCreateRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *ConnectionCreateRequest) GetStatus() *ConnectionStatusEnum {
	if o == nil {
		return nil
	}
	return o.Status
}
