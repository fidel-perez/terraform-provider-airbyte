resource "airbyte_source_aws_cloudtrail" "my_source_awscloudtrail" {
  configuration = {
    aws_key_id      = "...my_aws_key_id..."
    aws_region_name = "...my_aws_region_name..."
    aws_secret_key  = "...my_aws_secret_key..."
    source_type     = "aws-cloudtrail"
    start_date      = "2021-01-01"
  }
  name         = "Leigh Denesik"
  secret_id    = "...my_secret_id..."
  workspace_id = "2257411f-af4b-4754-8e47-2e802857a5b4"
}