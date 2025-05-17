# Table: aws_kms_keys

https://docs.aws.amazon.com/kms/latest/APIReference/API_KeyMetadata.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_kms_keys:
  - [aws_kms_key_grants](aws_kms_key_grants.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|rotation_enabled|Bool|
|arn (PK)|String|
|replica_keys|JSON|
|key_metadata|JSON|
|tags|JSON|