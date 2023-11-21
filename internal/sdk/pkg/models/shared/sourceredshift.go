// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceRedshiftRedshift string

const (
	SourceRedshiftRedshiftRedshift SourceRedshiftRedshift = "redshift"
)

func (e SourceRedshiftRedshift) ToPointer() *SourceRedshiftRedshift {
	return &e
}

func (e *SourceRedshiftRedshift) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "redshift":
		*e = SourceRedshiftRedshift(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRedshiftRedshift: %v", v)
	}
}

type SourceRedshift struct {
	// Name of the database.
	Database string `json:"database"`
	// Host Endpoint of the Redshift Cluster (must include the cluster-id, region and end with .redshift.amazonaws.com).
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password string `json:"password"`
	// Port of the database.
	Port *int64 `default:"5439" json:"port"`
	// The list of schemas to sync from. Specify one or more explicitly or keep empty to process all schemas. Schema names are case sensitive.
	Schemas    []string               `json:"schemas,omitempty"`
	sourceType SourceRedshiftRedshift `const:"redshift" json:"sourceType"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (s SourceRedshift) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceRedshift) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceRedshift) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceRedshift) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceRedshift) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *SourceRedshift) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceRedshift) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *SourceRedshift) GetSchemas() []string {
	if o == nil {
		return nil
	}
	return o.Schemas
}

func (o *SourceRedshift) GetSourceType() SourceRedshiftRedshift {
	return SourceRedshiftRedshiftRedshift
}

func (o *SourceRedshift) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
