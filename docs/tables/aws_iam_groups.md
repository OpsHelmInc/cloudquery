# Table: aws_iam_groups

This table shows data for IAM Groups.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_iam_groups:
  - [aws_iam_group_attached_policies](aws_iam_group_attached_policies.md)
  - [aws_iam_group_policies](aws_iam_group_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|group|`json`|
|users|`json`|
|policies|`json`|