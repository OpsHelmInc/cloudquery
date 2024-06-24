# Table: aws_organization_resource_policies

This table shows data for Organization Resource Policies.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|content|`utf8`|
|resource_policy_summary|`json`|