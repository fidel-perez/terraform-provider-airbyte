---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_mssql Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMssql Resource
---

# airbyte_source_mssql (Resource)

SourceMssql Resource

## Example Usage

```terraform
resource "airbyte_source_mssql" "my_source_mssql" {
  configuration = {
    database        = "master"
    host            = "...my_host..."
    jdbc_url_params = "...my_jdbc_url_params..."
    password        = "...my_password..."
    port            = 1433
    replication_method = {
      read_changes_using_change_data_capture_cdc = {
        data_to_sync            = "New Changes Only"
        initial_waiting_seconds = 2
        snapshot_isolation      = "Read Committed"
      }
    }
    schemas = [
      "...",
    ]
    ssl_method = {
      source_mssql_encrypted_trust_server_certificate = {}
    }
    tunnel_method = {
      source_mssql_no_tunnel = {}
    }
    username = "Salvatore_Weissnat66"
  }
  definition_id = "b6ad0e44-a4dc-4970-8078-573a20ac990f"
  name          = "Wm Corkery"
  secret_id     = "...my_secret_id..."
  workspace_id  = "7a67a851-50ea-4861-a0cd-618d74280681"
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

- `database` (String) The name of the database.
- `host` (String) The hostname of the database.
- `port` (Number) The port of the database.
- `username` (String) The username which is used to access the database.

Optional:

- `jdbc_url_params` (String) Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
- `password` (String, Sensitive) The password associated with the username.
- `replication_method` (Attributes) Configures how data is extracted from the database. (see [below for nested schema](#nestedatt--configuration--replication_method))
- `schemas` (List of String) The list of schemas to sync from. Defaults to user. Case sensitive.
- `ssl_method` (Attributes) The encryption method which is used when communicating with the database. (see [below for nested schema](#nestedatt--configuration--ssl_method))
- `tunnel_method` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method))

<a id="nestedatt--configuration--replication_method"></a>
### Nested Schema for `configuration.replication_method`

Optional:

- `read_changes_using_change_data_capture_cdc` (Attributes) <i>Recommended</i> - Incrementally reads new inserts, updates, and deletes using the SQL Server's <a href="https://docs.airbyte.com/integrations/sources/mssql/#change-data-capture-cdc">change data capture feature</a>. This must be enabled on your database. (see [below for nested schema](#nestedatt--configuration--replication_method--read_changes_using_change_data_capture_cdc))
- `scan_changes_with_user_defined_cursor` (Attributes) Incrementally detects new inserts and updates using the <a href="https://docs.airbyte.com/understanding-airbyte/connections/incremental-append/#user-defined-cursor">cursor column</a> chosen when configuring a connection (e.g. created_at, updated_at). (see [below for nested schema](#nestedatt--configuration--replication_method--scan_changes_with_user_defined_cursor))

<a id="nestedatt--configuration--replication_method--read_changes_using_change_data_capture_cdc"></a>
### Nested Schema for `configuration.replication_method.read_changes_using_change_data_capture_cdc`

Optional:

- `data_to_sync` (String) must be one of ["Existing and New", "New Changes Only"]; Default: "Existing and New"
What data should be synced under the CDC. "Existing and New" will read existing data as a snapshot, and sync new changes through CDC. "New Changes Only" will skip the initial snapshot, and only sync new changes through CDC.
- `initial_waiting_seconds` (Number) Default: 300
The amount of time the connector will wait when it launches to determine if there is new data to sync or not. Defaults to 300 seconds. Valid range: 120 seconds to 1200 seconds. Read about <a href="https://docs.airbyte.com/integrations/sources/mysql/#change-data-capture-cdc">initial waiting time</a>.
- `snapshot_isolation` (String) must be one of ["Snapshot", "Read Committed"]; Default: "Snapshot"
Existing data in the database are synced through an initial snapshot. This parameter controls the isolation level that will be used during the initial snapshotting. If you choose the "Snapshot" level, you must enable the <a href="https://docs.microsoft.com/en-us/dotnet/framework/data/adonet/sql/snapshot-isolation-in-sql-server">snapshot isolation mode</a> on the database.


<a id="nestedatt--configuration--replication_method--scan_changes_with_user_defined_cursor"></a>
### Nested Schema for `configuration.replication_method.scan_changes_with_user_defined_cursor`



<a id="nestedatt--configuration--ssl_method"></a>
### Nested Schema for `configuration.ssl_method`

Optional:

- `encrypted_trust_server_certificate` (Attributes) Use the certificate provided by the server without verification. (For testing purposes only!) (see [below for nested schema](#nestedatt--configuration--ssl_method--encrypted_trust_server_certificate))
- `encrypted_verify_certificate` (Attributes) Verify and use the certificate provided by the server. (see [below for nested schema](#nestedatt--configuration--ssl_method--encrypted_verify_certificate))

<a id="nestedatt--configuration--ssl_method--encrypted_trust_server_certificate"></a>
### Nested Schema for `configuration.ssl_method.encrypted_trust_server_certificate`


<a id="nestedatt--configuration--ssl_method--encrypted_verify_certificate"></a>
### Nested Schema for `configuration.ssl_method.encrypted_verify_certificate`

Optional:

- `host_name_in_certificate` (String) Specifies the host name of the server. The value of this property must match the subject property of the certificate.



<a id="nestedatt--configuration--tunnel_method"></a>
### Nested Schema for `configuration.tunnel_method`

Optional:

- `no_tunnel` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--no_tunnel))
- `password_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--password_authentication))
- `ssh_key_authentication` (Attributes) Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use. (see [below for nested schema](#nestedatt--configuration--tunnel_method--ssh_key_authentication))

<a id="nestedatt--configuration--tunnel_method--no_tunnel"></a>
### Nested Schema for `configuration.tunnel_method.no_tunnel`


<a id="nestedatt--configuration--tunnel_method--password_authentication"></a>
### Nested Schema for `configuration.tunnel_method.password_authentication`

Required:

- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_user` (String) OS-level username for logging into the jump server host
- `tunnel_user_password` (String, Sensitive) OS-level password for logging into the jump server host

Optional:

- `tunnel_port` (Number) Default: 22
Port on the proxy/jump server that accepts inbound ssh connections.


<a id="nestedatt--configuration--tunnel_method--ssh_key_authentication"></a>
### Nested Schema for `configuration.tunnel_method.ssh_key_authentication`

Required:

- `ssh_key` (String, Sensitive) OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
- `tunnel_host` (String) Hostname of the jump server host that allows inbound ssh tunnel.
- `tunnel_user` (String) OS-level username for logging into the jump server host.

Optional:

- `tunnel_port` (Number) Default: 22
Port on the proxy/jump server that accepts inbound ssh connections.


