# Table: aws_scheduler_schedule_groups

This table shows data for Amazon EventBridge Scheduler Schedule Groups.

https://docs.aws.amazon.com/scheduler/latest/APIReference/API_ScheduleGroupSummary.html

The primary key for this table is **_cq_id**.

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
|creation_date|`timestamp[us, tz=UTC]`|
|last_modification_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|state|`utf8`|