# Table: aws_resiliencehub_suggested_resiliency_policies

This table shows data for AWS Resilience Hub Suggested Resiliency Policies.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ResiliencyPolicy.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|data_location_constraint|`utf8`|
|estimated_cost_tier|`utf8`|
|policy|`json`|
|policy_arn|`utf8`|
|policy_description|`utf8`|
|policy_name|`utf8`|
|tags|`json`|
|tier|`utf8`|