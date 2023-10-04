// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/utils"
	"encoding/json"
	"fmt"
)

// SourcePocketContentType - Select the content type of the items to retrieve.
type SourcePocketContentType string

const (
	SourcePocketContentTypeArticle SourcePocketContentType = "article"
	SourcePocketContentTypeVideo   SourcePocketContentType = "video"
	SourcePocketContentTypeImage   SourcePocketContentType = "image"
)

func (e SourcePocketContentType) ToPointer() *SourcePocketContentType {
	return &e
}

func (e *SourcePocketContentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "article":
		fallthrough
	case "video":
		fallthrough
	case "image":
		*e = SourcePocketContentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePocketContentType: %v", v)
	}
}

// SourcePocketDetailType - Select the granularity of the information about each item.
type SourcePocketDetailType string

const (
	SourcePocketDetailTypeSimple   SourcePocketDetailType = "simple"
	SourcePocketDetailTypeComplete SourcePocketDetailType = "complete"
)

func (e SourcePocketDetailType) ToPointer() *SourcePocketDetailType {
	return &e
}

func (e *SourcePocketDetailType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "simple":
		fallthrough
	case "complete":
		*e = SourcePocketDetailType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePocketDetailType: %v", v)
	}
}

// SourcePocketSortBy - Sort retrieved items by the given criteria.
type SourcePocketSortBy string

const (
	SourcePocketSortByNewest SourcePocketSortBy = "newest"
	SourcePocketSortByOldest SourcePocketSortBy = "oldest"
	SourcePocketSortByTitle  SourcePocketSortBy = "title"
	SourcePocketSortBySite   SourcePocketSortBy = "site"
)

func (e SourcePocketSortBy) ToPointer() *SourcePocketSortBy {
	return &e
}

func (e *SourcePocketSortBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "newest":
		fallthrough
	case "oldest":
		fallthrough
	case "title":
		fallthrough
	case "site":
		*e = SourcePocketSortBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePocketSortBy: %v", v)
	}
}

type SourcePocketPocket string

const (
	SourcePocketPocketPocket SourcePocketPocket = "pocket"
)

func (e SourcePocketPocket) ToPointer() *SourcePocketPocket {
	return &e
}

func (e *SourcePocketPocket) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pocket":
		*e = SourcePocketPocket(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePocketPocket: %v", v)
	}
}

// SourcePocketState - Select the state of the items to retrieve.
type SourcePocketState string

const (
	SourcePocketStateUnread  SourcePocketState = "unread"
	SourcePocketStateArchive SourcePocketState = "archive"
	SourcePocketStateAll     SourcePocketState = "all"
)

func (e SourcePocketState) ToPointer() *SourcePocketState {
	return &e
}

func (e *SourcePocketState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unread":
		fallthrough
	case "archive":
		fallthrough
	case "all":
		*e = SourcePocketState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePocketState: %v", v)
	}
}

type SourcePocket struct {
	// The user's Pocket access token.
	AccessToken string `json:"access_token"`
	// Your application's Consumer Key.
	ConsumerKey string `json:"consumer_key"`
	// Select the content type of the items to retrieve.
	ContentType *SourcePocketContentType `json:"content_type,omitempty"`
	// Select the granularity of the information about each item.
	DetailType *SourcePocketDetailType `json:"detail_type,omitempty"`
	// Only return items from a particular `domain`.
	Domain *string `json:"domain,omitempty"`
	// Retrieve only favorited items.
	Favorite *bool `default:"false" json:"favorite"`
	// Only return items whose title or url contain the `search` string.
	Search *string `json:"search,omitempty"`
	// Only return items modified since the given timestamp.
	Since *string `json:"since,omitempty"`
	// Sort retrieved items by the given criteria.
	Sort       *SourcePocketSortBy `json:"sort,omitempty"`
	sourceType SourcePocketPocket  `const:"pocket" json:"sourceType"`
	// Select the state of the items to retrieve.
	State *SourcePocketState `json:"state,omitempty"`
	// Return only items tagged with this tag name. Use _untagged_ for retrieving only untagged items.
	Tag *string `json:"tag,omitempty"`
}

func (s SourcePocket) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePocket) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePocket) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *SourcePocket) GetConsumerKey() string {
	if o == nil {
		return ""
	}
	return o.ConsumerKey
}

func (o *SourcePocket) GetContentType() *SourcePocketContentType {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *SourcePocket) GetDetailType() *SourcePocketDetailType {
	if o == nil {
		return nil
	}
	return o.DetailType
}

func (o *SourcePocket) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *SourcePocket) GetFavorite() *bool {
	if o == nil {
		return nil
	}
	return o.Favorite
}

func (o *SourcePocket) GetSearch() *string {
	if o == nil {
		return nil
	}
	return o.Search
}

func (o *SourcePocket) GetSince() *string {
	if o == nil {
		return nil
	}
	return o.Since
}

func (o *SourcePocket) GetSort() *SourcePocketSortBy {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *SourcePocket) GetSourceType() SourcePocketPocket {
	return SourcePocketPocketPocket
}

func (o *SourcePocket) GetState() *SourcePocketState {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *SourcePocket) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}
