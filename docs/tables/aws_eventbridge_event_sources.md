# Table: aws_eventbridge_event_sources

This table shows data for Amazon EventBridge Event Sources.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventSource.html

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
|created_by|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|expiration_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|state|`utf8`|