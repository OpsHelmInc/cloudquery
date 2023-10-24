# Table: aws_iam_roles

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The composite primary key for this table is (**account_id**, **id**).

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_policies](aws_iam_role_policies.md)

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
|assume_role_policy_document|JSON|
|oh_resource_type|String|
|role|JSON|
|policies|JSON|