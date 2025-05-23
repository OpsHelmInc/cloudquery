# Table: aws_fsx_file_systems

This table shows data for Amazon FSx File Systems.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_FileSystem.html

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
|administrative_actions|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|failure_details|`json`|
|file_system_id|`utf8`|
|file_system_type|`utf8`|
|file_system_type_version|`utf8`|
|kms_key_id|`utf8`|
|lifecycle|`utf8`|
|lustre_configuration|`json`|
|network_interface_ids|`list<item: utf8, nullable>`|
|ontap_configuration|`json`|
|open_zfs_configuration|`json`|
|owner_id|`utf8`|
|resource_arn|`utf8`|
|storage_capacity|`int64`|
|storage_type|`utf8`|
|subnet_ids|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|
|windows_configuration|`json`|