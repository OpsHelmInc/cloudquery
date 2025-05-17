# Table: aws_elbv2_target_groups

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_target_groups:
  - [aws_elbv2_target_group_target_health_descriptions](aws_elbv2_target_group_target_health_descriptions.md)

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
|target_group|JSON|
|tags|JSON|
|unknown_attributes|JSON|