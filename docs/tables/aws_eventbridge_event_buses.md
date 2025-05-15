# Table: aws_eventbridge_event_buses

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_EventBus.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_eventbridge_event_buses:
  - [aws_eventbridge_event_bus_rules](aws_eventbridge_event_bus_rules.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|creation_time|Timestamp|
|dead_letter_config|JSON|
|description|String|
|kms_key_identifier|String|
|last_modified_time|Timestamp|
|name|String|
|policy|String|
|result_metadata|JSON|
|tags|JSON|