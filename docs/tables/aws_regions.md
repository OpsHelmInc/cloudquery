# Table: aws_regions

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Region.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|arn|String|
|account_id|String|
|enabled|Bool|
|partition|String|
|region|String|
|oh_resource_type|String|
|ec2_config|JSON|