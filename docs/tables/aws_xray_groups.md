# Table: aws_xray_groups

https://docs.aws.amazon.com/xray/latest/api/API_Group.html

The primary key for this table is **arn**.



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
|tags|JSON|
|oh_resource_type|String|
|filter_expression|String|
|group_name|String|
|insights_configuration|JSON|