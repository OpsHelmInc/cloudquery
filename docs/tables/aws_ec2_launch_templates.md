# Table: aws_ec2_launch_templates

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateVersion.html

The composite primary key for this table is (**arn**, **launch_template_id**).



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
|launch_template_id (PK)|String|
|oh_resource_type|String|
|create_time|Timestamp|
|created_by|String|
|default_version_number|Int|
|latest_version_number|Int|
|launch_template_name|String|
|tags|JSON|