# Table: aws_iam_group_attached_policies

This table shows data for IAM Group Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_iam_groups](aws_iam_groups.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|group_arn|`utf8`|
|policy_arn|`utf8`|
|policy_name|`utf8`|