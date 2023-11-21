// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationKinesisUpdate struct {
	// Generate the AWS Access Key for current user.
	AccessKey string `json:"accessKey"`
	// Buffer size for storing kinesis records before being batch streamed.
	BufferSize *int64 `default:"100" json:"bufferSize"`
	// AWS Kinesis endpoint.
	Endpoint string `json:"endpoint"`
	// The AWS Private Key - a string of numbers and letters that are unique for each account, also known as a "recovery phrase".
	PrivateKey string `json:"privateKey"`
	// AWS region. Your account determines the Regions that are available to you.
	Region string `json:"region"`
	// Number of shards to which the data should be streamed.
	ShardCount *int64 `default:"5" json:"shardCount"`
}

func (d DestinationKinesisUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationKinesisUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationKinesisUpdate) GetAccessKey() string {
	if o == nil {
		return ""
	}
	return o.AccessKey
}

func (o *DestinationKinesisUpdate) GetBufferSize() *int64 {
	if o == nil {
		return nil
	}
	return o.BufferSize
}

func (o *DestinationKinesisUpdate) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *DestinationKinesisUpdate) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *DestinationKinesisUpdate) GetRegion() string {
	if o == nil {
		return ""
	}
	return o.Region
}

func (o *DestinationKinesisUpdate) GetShardCount() *int64 {
	if o == nil {
		return nil
	}
	return o.ShardCount
}
