// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Dynamodb string

const (
	DynamodbDynamodb Dynamodb = "dynamodb"
)

func (e Dynamodb) ToPointer() *Dynamodb {
	return &e
}

func (e *Dynamodb) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dynamodb":
		*e = Dynamodb(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Dynamodb: %v", v)
	}
}

// DestinationDynamodbDynamoDBRegion - The region of the DynamoDB.
type DestinationDynamodbDynamoDBRegion string

const (
	DestinationDynamodbDynamoDBRegionUnknown      DestinationDynamodbDynamoDBRegion = ""
	DestinationDynamodbDynamoDBRegionUsEast1      DestinationDynamodbDynamoDBRegion = "us-east-1"
	DestinationDynamodbDynamoDBRegionUsEast2      DestinationDynamodbDynamoDBRegion = "us-east-2"
	DestinationDynamodbDynamoDBRegionUsWest1      DestinationDynamodbDynamoDBRegion = "us-west-1"
	DestinationDynamodbDynamoDBRegionUsWest2      DestinationDynamodbDynamoDBRegion = "us-west-2"
	DestinationDynamodbDynamoDBRegionAfSouth1     DestinationDynamodbDynamoDBRegion = "af-south-1"
	DestinationDynamodbDynamoDBRegionApEast1      DestinationDynamodbDynamoDBRegion = "ap-east-1"
	DestinationDynamodbDynamoDBRegionApSouth1     DestinationDynamodbDynamoDBRegion = "ap-south-1"
	DestinationDynamodbDynamoDBRegionApNortheast1 DestinationDynamodbDynamoDBRegion = "ap-northeast-1"
	DestinationDynamodbDynamoDBRegionApNortheast2 DestinationDynamodbDynamoDBRegion = "ap-northeast-2"
	DestinationDynamodbDynamoDBRegionApNortheast3 DestinationDynamodbDynamoDBRegion = "ap-northeast-3"
	DestinationDynamodbDynamoDBRegionApSoutheast1 DestinationDynamodbDynamoDBRegion = "ap-southeast-1"
	DestinationDynamodbDynamoDBRegionApSoutheast2 DestinationDynamodbDynamoDBRegion = "ap-southeast-2"
	DestinationDynamodbDynamoDBRegionCaCentral1   DestinationDynamodbDynamoDBRegion = "ca-central-1"
	DestinationDynamodbDynamoDBRegionCnNorth1     DestinationDynamodbDynamoDBRegion = "cn-north-1"
	DestinationDynamodbDynamoDBRegionCnNorthwest1 DestinationDynamodbDynamoDBRegion = "cn-northwest-1"
	DestinationDynamodbDynamoDBRegionEuCentral1   DestinationDynamodbDynamoDBRegion = "eu-central-1"
	DestinationDynamodbDynamoDBRegionEuNorth1     DestinationDynamodbDynamoDBRegion = "eu-north-1"
	DestinationDynamodbDynamoDBRegionEuSouth1     DestinationDynamodbDynamoDBRegion = "eu-south-1"
	DestinationDynamodbDynamoDBRegionEuWest1      DestinationDynamodbDynamoDBRegion = "eu-west-1"
	DestinationDynamodbDynamoDBRegionEuWest2      DestinationDynamodbDynamoDBRegion = "eu-west-2"
	DestinationDynamodbDynamoDBRegionEuWest3      DestinationDynamodbDynamoDBRegion = "eu-west-3"
	DestinationDynamodbDynamoDBRegionSaEast1      DestinationDynamodbDynamoDBRegion = "sa-east-1"
	DestinationDynamodbDynamoDBRegionMeSouth1     DestinationDynamodbDynamoDBRegion = "me-south-1"
	DestinationDynamodbDynamoDBRegionUsGovEast1   DestinationDynamodbDynamoDBRegion = "us-gov-east-1"
	DestinationDynamodbDynamoDBRegionUsGovWest1   DestinationDynamodbDynamoDBRegion = "us-gov-west-1"
)

func (e DestinationDynamodbDynamoDBRegion) ToPointer() *DestinationDynamodbDynamoDBRegion {
	return &e
}

func (e *DestinationDynamodbDynamoDBRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		fallthrough
	case "us-gov-east-1":
		fallthrough
	case "us-gov-west-1":
		*e = DestinationDynamodbDynamoDBRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDynamodbDynamoDBRegion: %v", v)
	}
}

type DestinationDynamodb struct {
	// The access key id to access the DynamoDB. Airbyte requires Read and Write permissions to the DynamoDB.
	AccessKeyID     string   `json:"access_key_id"`
	destinationType Dynamodb `const:"dynamodb" json:"destinationType"`
	// This is your DynamoDB endpoint url.(if you are working with AWS DynamoDB, just leave empty).
	DynamodbEndpoint *string `default:"" json:"dynamodb_endpoint"`
	// The region of the DynamoDB.
	DynamodbRegion *DestinationDynamodbDynamoDBRegion `default:"" json:"dynamodb_region"`
	// The prefix to use when naming DynamoDB tables.
	DynamodbTableNamePrefix string `json:"dynamodb_table_name_prefix"`
	// The corresponding secret to the access key id.
	SecretAccessKey string `json:"secret_access_key"`
}

func (d DestinationDynamodb) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDynamodb) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDynamodb) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *DestinationDynamodb) GetDestinationType() Dynamodb {
	return DynamodbDynamodb
}

func (o *DestinationDynamodb) GetDynamodbEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.DynamodbEndpoint
}

func (o *DestinationDynamodb) GetDynamodbRegion() *DestinationDynamodbDynamoDBRegion {
	if o == nil {
		return nil
	}
	return o.DynamodbRegion
}

func (o *DestinationDynamodb) GetDynamodbTableNamePrefix() string {
	if o == nil {
		return ""
	}
	return o.DynamodbTableNamePrefix
}

func (o *DestinationDynamodb) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}
