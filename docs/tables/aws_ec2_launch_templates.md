# Table: aws_ec2_launch_templates

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ResponseLaunchTemplateData.html

The primary key for this table is **launch_template_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|launch_template_id (PK)|String|
|launch_template_data|JSON|