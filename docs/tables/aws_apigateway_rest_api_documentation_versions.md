# Table: aws_apigateway_rest_api_documentation_versions

This table shows data for Amazon API Gateway Rest API Documentation Versions.

https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationVersion.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_apigateway_rest_apis](aws_apigateway_rest_apis.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rest_api_arn|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|version|`utf8`|