# Table: aws_appstream_app_blocks

https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html

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
|oh_resource_type|String|
|name|String|
|app_block_errors|JSON|
|created_time|Timestamp|
|description|String|
|display_name|String|
|packaging_type|String|
|post_setup_script_details|JSON|
|setup_script_details|JSON|
|source_s3_location|JSON|
|state|String|