---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_stripe Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceStripe Resource
---

# airbyte_source_stripe (Resource)

SourceStripe Resource

## Example Usage

```terraform
resource "airbyte_source_stripe" "my_source_stripe" {
  configuration = {
    account_id           = "...my_account_id..."
    call_rate_limit      = 100
    client_secret        = "...my_client_secret..."
    lookback_window_days = 10
    num_workers          = 3
    slice_range          = 360
    start_date           = "2017-01-25T00:00:00Z"
  }
  definition_id = "46c36bb7-337b-4f0b-aca9-3a8ae78e1e53"
  name          = "Marcella Muller"
  secret_id     = "...my_secret_id..."
  workspace_id  = "b6d5dc1e-250f-480f-bc59-5c3777bccfe7"
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

- `account_id` (String) Your Stripe account ID (starts with 'acct_', find yours <a href="https://dashboard.stripe.com/settings/account">here</a>).
- `client_secret` (String) Stripe API key (usually starts with 'sk_live_'; find yours <a href="https://dashboard.stripe.com/apikeys">here</a>).

Optional:

- `call_rate_limit` (Number) The number of API calls per second that you allow connector to make. This value can not be bigger than real API call rate limit (https://stripe.com/docs/rate-limits). If not specified the default maximum is 25 and 100 calls per second for test and production tokens respectively.
- `lookback_window_days` (Number) Default: 0
When set, the connector will always re-export data from the past N days, where N is the value set here. This is useful if your data is frequently updated after creation. The Lookback Window only applies to streams that do not support event-based incremental syncs: Events, SetupAttempts, ShippingRates, BalanceTransactions, Files, FileLinks, Refunds. More info <a href="https://docs.airbyte.com/integrations/sources/stripe#requirements">here</a>
- `num_workers` (Number) Default: 10
The number of worker thread to use for the sync. The performance upper boundary depends on call_rate_limit setting and type of account.
- `slice_range` (Number) Default: 365
The time increment used by the connector when requesting data from the Stripe API. The bigger the value is, the less requests will be made and faster the sync will be. On the other hand, the more seldom the state is persisted.
- `start_date` (String) Default: "2017-01-25T00:00:00Z"
UTC date and time in the format 2017-01-25T00:00:00Z. Only data generated after this date will be replicated.


