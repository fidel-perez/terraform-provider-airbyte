---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_amazon_sqs Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAmazonSqs Resource
---

# airbyte_source_amazon_sqs (Resource)

SourceAmazonSqs Resource

## Example Usage

```terraform
resource "airbyte_source_amazon_sqs" "my_source_amazonsqs" {
  configuration = {
    access_key           = "xxxxxHRNxxx3TBxxxxxx"
    attributes_to_return = "attr1,attr2"
    delete_messages      = true
    max_batch_size       = 5
    max_wait_time        = 5
    queue_url            = "https://sqs.eu-west-1.amazonaws.com/1234567890/my-example-queue"
    region               = "ap-northeast-2"
    secret_key           = "hu+qE5exxxxT6o/ZrKsxxxxxxBhxxXLexxxxxVKz"
    visibility_timeout   = 15
  }
  definition_id = "aa9ea927-cae7-4b29-885e-6b85628652e0"
  name          = "Emmett Labadie"
  secret_id     = "...my_secret_id..."
  workspace_id  = "21b517b1-6f1f-4884-abcd-5137451945c4"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `queue_url` (String) URL of the SQS Queue
- `region` (String) must be one of ["us-east-1", "us-east-2", "us-west-1", "us-west-2", "af-south-1", "ap-east-1", "ap-south-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ca-central-1", "cn-north-1", "cn-northwest-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "sa-east-1", "me-south-1", "us-gov-east-1", "us-gov-west-1"]
AWS Region of the SQS Queue

Optional:

- `access_key` (String, Sensitive) The Access Key ID of the AWS IAM Role to use for pulling messages
- `attributes_to_return` (String) Comma separated list of Mesage Attribute names to return
- `delete_messages` (Boolean) Default: false
If Enabled, messages will be deleted from the SQS Queue after being read. If Disabled, messages are left in the queue and can be read more than once. WARNING: Enabling this option can result in data loss in cases of failure, use with caution, see documentation for more detail.
- `max_batch_size` (Number) Max amount of messages to get in one batch (10 max)
- `max_wait_time` (Number) Max amount of time in seconds to wait for messages in a single poll (20 max)
- `secret_key` (String, Sensitive) The Secret Key of the AWS IAM Role to use for pulling messages
- `visibility_timeout` (Number) Modify the Visibility Timeout of the individual message from the Queue's default (seconds).


