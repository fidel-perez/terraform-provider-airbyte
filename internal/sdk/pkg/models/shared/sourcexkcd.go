// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

type SourceXkcdXkcd string

const (
	SourceXkcdXkcdXkcd SourceXkcdXkcd = "xkcd"
)

func (e SourceXkcdXkcd) ToPointer() *SourceXkcdXkcd {
	return &e
}

func (e *SourceXkcdXkcd) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "xkcd":
		*e = SourceXkcdXkcd(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceXkcdXkcd: %v", v)
	}
}

type SourceXkcd struct {
	sourceType *SourceXkcdXkcd `const:"xkcd" json:"sourceType,omitempty"`
}

func (s SourceXkcd) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceXkcd) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceXkcd) GetSourceType() *SourceXkcdXkcd {
	return SourceXkcdXkcdXkcd.ToPointer()
}
