# Table: aws_securityhub_enabled_standards

This table shows data for AWS Security Hub Enabled Standards.

https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetEnabledStandards.html

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
|standards_arn|`utf8`|
|standards_input|`json`|
|standards_status|`utf8`|
|standards_subscription_arn|`utf8`|
|standards_status_reason|`json`|