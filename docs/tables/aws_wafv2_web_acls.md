# Table: aws_wafv2_web_acls



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
|resources_for_web_acl|StringArray|
|arn (PK)|String|
|web_acl|JSON|
|logging_configuration|JSON|
|tags|JSON|