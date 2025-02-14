# Table: aws_iam_roles

This table shows data for IAM Roles.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_iam_roles:
  - [aws_iam_role_attached_policies](aws_iam_role_attached_policies.md)
  - [aws_iam_role_policies](aws_iam_role_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|assume_role_policy_document|`json`|
|tags|`json`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|role|`json`|
|policies|`json`|