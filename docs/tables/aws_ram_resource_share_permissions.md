# Table: aws_ram_resource_share_permissions

This table shows data for RAM Resource Share Permissions.

https://docs.aws.amazon.com/ram/latest/APIReference/API_ResourceSharePermissionSummary.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_ram_resource_shares](aws_ram_resource_shares.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|resource_share_arn|`utf8`|
|permission|`json`|
|tags|`json`|
|arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|default_version|`bool`|
|feature_set|`utf8`|
|is_resource_type_default|`bool`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|permission_type|`utf8`|
|resource_type|`utf8`|
|status|`utf8`|
|version|`utf8`|