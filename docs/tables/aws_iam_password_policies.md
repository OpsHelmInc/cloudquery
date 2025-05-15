# Table: aws_iam_password_policies



The primary key for this table is **account_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|arn|String|
|account_id (PK)|String|
|oh_resource_type|String|
|password_policy|JSON|
|policy_exists|Bool|