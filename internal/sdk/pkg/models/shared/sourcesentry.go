// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Sentry string

const (
	SentrySentry Sentry = "sentry"
)

func (e Sentry) ToPointer() *Sentry {
	return &e
}

func (e *Sentry) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sentry":
		*e = Sentry(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sentry: %v", v)
	}
}

type SourceSentry struct {
	// Log into Sentry and then <a href="https://sentry.io/settings/account/api/auth-tokens/">create authentication tokens</a>.For self-hosted, you can find or create authentication tokens by visiting "{instance_url_prefix}/settings/account/api/auth-tokens/"
	AuthToken string `json:"auth_token"`
	// Fields to retrieve when fetching discover events
	DiscoverFields []interface{} `json:"discover_fields,omitempty"`
	// Host name of Sentry API server.For self-hosted, specify your host name here. Otherwise, leave it empty.
	Hostname *string `default:"sentry.io" json:"hostname"`
	// The slug of the organization the groups belong to.
	Organization string `json:"organization"`
	// The name (slug) of the Project you want to sync.
	Project    string `json:"project"`
	sourceType Sentry `const:"sentry" json:"sourceType"`
}

func (s SourceSentry) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceSentry) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceSentry) GetAuthToken() string {
	if o == nil {
		return ""
	}
	return o.AuthToken
}

func (o *SourceSentry) GetDiscoverFields() []interface{} {
	if o == nil {
		return nil
	}
	return o.DiscoverFields
}

func (o *SourceSentry) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *SourceSentry) GetOrganization() string {
	if o == nil {
		return ""
	}
	return o.Organization
}

func (o *SourceSentry) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

func (o *SourceSentry) GetSourceType() Sentry {
	return SentrySentry
}
