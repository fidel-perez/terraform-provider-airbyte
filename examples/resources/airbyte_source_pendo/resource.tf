resource "airbyte_source_pendo" "my_source_pendo" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "6503c474-3ee7-49bd-93e2-04659bbdc56c"
  name          = "Mandy Conroy"
  secret_id     = "...my_secret_id..."
  workspace_id  = "0259c6b1-3998-4d3f-8543-0ae066d4a91b"
}