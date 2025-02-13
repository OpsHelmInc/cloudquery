# Table: aws_eventbridge_event_buses

This table shows data for Amazon EventBridge Event Buses.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_eventbridge_event_buses:
  - [aws_eventbridge_event_bus_rules](aws_eventbridge_event_bus_rules.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|policy|`utf8`|