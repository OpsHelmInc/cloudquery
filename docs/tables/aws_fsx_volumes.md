# Table: aws_fsx_volumes

This table shows data for Amazon FSx Volumes.

https://docs.aws.amazon.com/fsx/latest/APIReference/API_Volume.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|administrative_actions|`json`|
|tags|`json`|
|oh_resource_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|file_system_id|`utf8`|
|lifecycle|`utf8`|
|lifecycle_transition_reason|`json`|
|name|`utf8`|
|ontap_configuration|`json`|
|open_zfs_configuration|`json`|
|resource_arn|`utf8`|
|volume_id|`utf8`|
|volume_type|`utf8`|