# Table: aws_eventbridge_endpoints

This table shows data for Amazon EventBridge Endpoints.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_Endpoint.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|endpoint_id|`utf8`|
|endpoint_url|`utf8`|
|event_buses|`json`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|replication_config|`json`|
|role_arn|`utf8`|
|routing_config|`json`|
|state|`utf8`|
|state_reason|`utf8`|