# Table: aws_backup_report_plans

This table shows data for Backup Report Plans.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ReportPlan.html

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
|creation_time|`timestamp[us, tz=UTC]`|
|deployment_status|`utf8`|
|last_attempted_execution_time|`timestamp[us, tz=UTC]`|
|last_successful_execution_time|`timestamp[us, tz=UTC]`|
|report_delivery_channel|`json`|
|report_plan_arn|`utf8`|
|report_plan_description|`utf8`|
|report_plan_name|`utf8`|
|report_setting|`json`|