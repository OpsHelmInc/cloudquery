# Table: aws_glacier_vaults

This table shows data for Glacier Vaults.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_glacier_vaults:
  - [aws_glacier_vault_access_policies](aws_glacier_vault_access_policies.md)
  - [aws_glacier_vault_lock_policies](aws_glacier_vault_lock_policies.md)
  - [aws_glacier_vault_notifications](aws_glacier_vault_notifications.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|creation_date|`utf8`|
|last_inventory_date|`utf8`|
|number_of_archives|`int64`|
|size_in_bytes|`int64`|
|vault_arn|`utf8`|
|vault_name|`utf8`|