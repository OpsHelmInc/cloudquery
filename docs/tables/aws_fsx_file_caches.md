# Table: aws_fsx_file_caches

This table shows data for Amazon FSx File Caches.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileCache.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|data_repository_association_ids|`list<item: utf8, nullable>`|
|failure_details|`json`|
|file_cache_id|`utf8`|
|file_cache_type|`utf8`|
|file_cache_type_version|`utf8`|
|kms_key_id|`utf8`|
|lifecycle|`utf8`|
|lustre_configuration|`json`|
|network_interface_ids|`list<item: utf8, nullable>`|
|owner_id|`utf8`|
|resource_arn|`utf8`|
|storage_capacity|`int64`|
|subnet_ids|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|