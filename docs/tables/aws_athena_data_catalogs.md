# Table: aws_athena_data_catalogs

This table shows data for Athena Data Catalogs.

https://docs.aws.amazon.com/athena/latest/APIReference/API_DataCatalog.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_athena_data_catalogs:
  - [aws_athena_data_catalog_databases](aws_athena_data_catalog_databases.md)

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
|name|`utf8`|
|type|`utf8`|
|connection_type|`utf8`|
|description|`utf8`|
|error|`utf8`|
|parameters|`json`|
|status|`utf8`|