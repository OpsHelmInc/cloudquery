# Table: aws_xray_groups

This table shows data for AWS X-Ray Groups.

https://docs.aws.amazon.com/xray/latest/api/API_Group.html

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
|filter_expression|`utf8`|
|group_arn|`utf8`|
|group_name|`utf8`|
|insights_configuration|`json`|