# Table: aws_dynamodb_backups

This table shows data for Amazon DynamoDB Backups.

https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_BackupDescription.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|backup_details|`json`|
|source_table_details|`json`|
|source_table_feature_details|`json`|