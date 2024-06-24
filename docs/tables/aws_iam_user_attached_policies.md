# Table: aws_iam_user_attached_policies

This table shows data for IAM User Attached Policies.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_AttachedPolicy.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|user_arn|`utf8`|
|policy_name|`utf8`|
|user_id|`utf8`|
|policy_arn|`utf8`|