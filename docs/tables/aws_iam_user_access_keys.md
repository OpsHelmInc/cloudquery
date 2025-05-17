# Table: aws_iam_user_access_keys



The primary key for this table is **_cq_id**.

## Relations
This table depends on [aws_iam_users](aws_iam_users.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|arn|String|
|user_arn|String|
|user_id|String|
|oh_resource_type|String|
|access_key_id|String|
|create_date|Timestamp|
|status|String|
|user_name|String|
|last_used_date|Timestamp|
|last_used_region|String|
|last_used_service_name|String|