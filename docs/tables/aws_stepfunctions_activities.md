# Table: aws_stepfunctions_activities

This table shows data for Stepfunctions Activities.

https://docs.aws.amazon.com/step-functions/latest/apireference/API_ListActivities.html

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
|activity_arn|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|