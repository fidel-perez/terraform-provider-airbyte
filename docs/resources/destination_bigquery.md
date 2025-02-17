---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_bigquery Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationBigquery Resource
---

# airbyte_destination_bigquery (Resource)

DestinationBigquery Resource

## Example Usage

```terraform
resource "airbyte_destination_bigquery" "my_destination_bigquery" {
  configuration = {
    big_query_client_buffer_size_mb = 15
    credentials_json                = "...my_credentials_json..."
    dataset_id                      = "...my_dataset_id..."
    dataset_location                = "me-central2"
    disable_type_dedupe             = true
    loading_method = {
      gcs_staging = {
        credential = {
          destination_bigquery_hmac_key = {
            hmac_key_access_id = "1234567890abcdefghij1234"
            hmac_key_secret    = "1234567890abcdefghij1234567890ABCDEFGHIJ"
          }
        }
        gcs_bucket_name          = "airbyte_sync"
        gcs_bucket_path          = "data_sync/test"
        keep_files_in_gcs_bucket = "Keep all tmp files in GCS"
      }
    }
    project_id              = "...my_project_id..."
    raw_data_dataset        = "...my_raw_data_dataset..."
    transformation_priority = "batch"
  }
  definition_id = "2d142842-c5e9-475e-80d1-1a3c6d933cc0"
  name          = "Miss Celia Moore"
  workspace_id  = "2d2700dc-d43a-4c80-9ede-88b16b5e1575"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the destination e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided.

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `dataset_id` (String) The default BigQuery Dataset ID that tables are replicated to if the source does not specify a namespace. Read more <a href="https://cloud.google.com/bigquery/docs/datasets#create-dataset">here</a>.
- `dataset_location` (String) must be one of ["US", "EU", "asia-east1", "asia-east2", "asia-northeast1", "asia-northeast2", "asia-northeast3", "asia-south1", "asia-south2", "asia-southeast1", "asia-southeast2", "australia-southeast1", "australia-southeast2", "europe-central1", "europe-central2", "europe-north1", "europe-southwest1", "europe-west1", "europe-west2", "europe-west3", "europe-west4", "europe-west6", "europe-west7", "europe-west8", "europe-west9", "europe-west12", "me-central1", "me-central2", "me-west1", "northamerica-northeast1", "northamerica-northeast2", "southamerica-east1", "southamerica-west1", "us-central1", "us-east1", "us-east2", "us-east3", "us-east4", "us-east5", "us-south1", "us-west1", "us-west2", "us-west3", "us-west4"]
The location of the dataset. Warning: Changes made after creation will not be applied. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
- `project_id` (String) The GCP project ID for the project containing the target BigQuery dataset. Read more <a href="https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects">here</a>.

Optional:

- `big_query_client_buffer_size_mb` (Number) Default: 15
Google BigQuery client's chunk (buffer) size (MIN=1, MAX = 15) for each table. The size that will be written by a single RPC. Written data will be buffered and only flushed upon reaching this size or closing the channel. The default 15MB value is used if not set explicitly. Read more <a href="https://googleapis.dev/python/bigquery/latest/generated/google.cloud.bigquery.client.Client.html">here</a>.
- `credentials_json` (String) The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/bigquery#service-account-key">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
- `disable_type_dedupe` (Boolean) Default: false
Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions
- `loading_method` (Attributes) The way data will be uploaded to BigQuery. (see [below for nested schema](#nestedatt--configuration--loading_method))
- `raw_data_dataset` (String) The dataset to write raw tables into (default: airbyte_internal)
- `transformation_priority` (String) must be one of ["interactive", "batch"]; Default: "interactive"
Interactive run type means that the query is executed as soon as possible, and these queries count towards concurrent rate limit and daily limit. Read more about interactive run type <a href="https://cloud.google.com/bigquery/docs/running-queries#queries">here</a>. Batch queries are queued and started as soon as idle resources are available in the BigQuery shared resource pool, which usually occurs within a few minutes. Batch queries don’t count towards your concurrent rate limit. Read more about batch queries <a href="https://cloud.google.com/bigquery/docs/running-queries#batch">here</a>. The default "interactive" value is used if not set explicitly.

<a id="nestedatt--configuration--loading_method"></a>
### Nested Schema for `configuration.loading_method`

Optional:

- `gcs_staging` (Attributes) <i>(recommended)</i> Writes large batches of records to a file, uploads the file to GCS, then uses COPY INTO to load your data into BigQuery. Provides best-in-class speed, reliability and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--gcs_staging))
- `standard_inserts` (Attributes) <i>(not recommended)</i> Direct loading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In all other cases, you should use GCS staging. (see [below for nested schema](#nestedatt--configuration--loading_method--standard_inserts))

<a id="nestedatt--configuration--loading_method--gcs_staging"></a>
### Nested Schema for `configuration.loading_method.gcs_staging`

Required:

- `credential` (Attributes) An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--gcs_staging--credential))
- `gcs_bucket_name` (String) The name of the GCS bucket. Read more <a href="https://cloud.google.com/storage/docs/naming-buckets">here</a>.
- `gcs_bucket_path` (String) Directory under the GCS bucket where data will be written.

Optional:

- `keep_files_in_gcs_bucket` (String) must be one of ["Delete all tmp files from GCS", "Keep all tmp files in GCS"]; Default: "Delete all tmp files from GCS"
This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.

<a id="nestedatt--configuration--loading_method--gcs_staging--credential"></a>
### Nested Schema for `configuration.loading_method.gcs_staging.keep_files_in_gcs_bucket`

Optional:

- `hmac_key` (Attributes) An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--gcs_staging--keep_files_in_gcs_bucket--hmac_key))

<a id="nestedatt--configuration--loading_method--gcs_staging--keep_files_in_gcs_bucket--hmac_key"></a>
### Nested Schema for `configuration.loading_method.gcs_staging.keep_files_in_gcs_bucket.hmac_key`

Required:

- `hmac_key_access_id` (String, Sensitive) HMAC key access ID. When linked to a service account, this ID is 61 characters long; when linked to a user account, it is 24 characters long.
- `hmac_key_secret` (String, Sensitive) The corresponding secret for the access ID. It is a 40-character base-64 encoded string.




<a id="nestedatt--configuration--loading_method--standard_inserts"></a>
### Nested Schema for `configuration.loading_method.standard_inserts`


