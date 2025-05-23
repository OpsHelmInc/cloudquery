# Table: aws_glue_triggers

This table shows data for Glue Triggers.

https://docs.aws.amazon.com/glue/latest/webapi/API_Trigger.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|actions|`json`|
|description|`utf8`|
|event_batching_condition|`json`|
|id|`utf8`|
|name|`utf8`|
|predicate|`json`|
|schedule|`utf8`|
|state|`utf8`|
|type|`utf8`|
|workflow_name|`utf8`|