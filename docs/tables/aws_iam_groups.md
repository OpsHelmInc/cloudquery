# Table: aws_iam_groups

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_groups:
  - [aws_iam_group_policies](aws_iam_group_policies.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|arn|String|
|account_id (PK)|String|
|id (PK)|String|
|oh_resource_type|String|
|group|JSON|
|users|JSON|
|policies|JSON|