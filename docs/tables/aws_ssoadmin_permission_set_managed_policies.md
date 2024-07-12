# Table: aws_ssoadmin_permission_set_managed_policies

This table shows data for Ssoadmin Permission Set Managed Policies.

https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListManagedPoliciesInPermissionSet.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ssoadmin_permission_sets](aws_ssoadmin_permission_sets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|permission_set_arn|`utf8`|
|instance_arn|`utf8`|
|arn|`utf8`|
|name|`utf8`|