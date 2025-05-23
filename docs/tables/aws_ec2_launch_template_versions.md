# Table: aws_ec2_launch_template_versions

This table shows data for Amazon Elastic Compute Cloud (EC2) Launch Template Versions.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateVersion.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|version_number|`int64`|
|oh_resource_type|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|created_by|`utf8`|
|default_version|`bool`|
|launch_template_data|`json`|
|launch_template_id|`utf8`|
|launch_template_name|`utf8`|
|version_description|`utf8`|