# Table: aws_shield_protection_groups

This table shows data for Shield Protection Groups.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_ProtectionGroup.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|aggregation|`utf8`|
|members|`list<item: utf8, nullable>`|
|pattern|`utf8`|
|protection_group_id|`utf8`|
|protection_group_arn|`utf8`|
|resource_type|`utf8`|