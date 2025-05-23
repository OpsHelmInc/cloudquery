# Table: aws_lambda_layer_versions

This table shows data for AWS Lambda Layer Versions.

https://docs.aws.amazon.com/lambda/latest/dg/API_LayerVersionsListItem.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_lambda_layers](aws_lambda_layers.md).

The following tables depend on aws_lambda_layer_versions:
  - [aws_lambda_layer_version_policies](aws_lambda_layer_version_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|layer_arn|`utf8`|
|oh_resource_type|`utf8`|
|compatible_architectures|`list<item: utf8, nullable>`|
|compatible_runtimes|`list<item: utf8, nullable>`|
|created_date|`utf8`|
|description|`utf8`|
|layer_version_arn|`utf8`|
|license_info|`utf8`|
|version|`int64`|