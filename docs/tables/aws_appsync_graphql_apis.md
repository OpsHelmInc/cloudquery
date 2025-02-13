# Table: aws_appsync_graphql_apis

This table shows data for Appsync Graphql APIs.

https://docs.aws.amazon.com/appsync/latest/APIReference/API_GraphqlApi.html

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
|additional_authentication_providers|`json`|
|api_id|`utf8`|
|api_type|`utf8`|
|authentication_type|`utf8`|
|dns|`json`|
|enhanced_metrics_config|`json`|
|introspection_config|`utf8`|
|lambda_authorizer_config|`json`|
|log_config|`json`|
|merged_api_execution_role_arn|`utf8`|
|name|`utf8`|
|open_id_connect_config|`json`|
|owner|`utf8`|
|owner_contact|`utf8`|
|query_depth_limit|`int64`|
|resolver_count_limit|`int64`|
|tags|`json`|
|uris|`json`|
|user_pool_config|`json`|
|visibility|`utf8`|
|waf_web_acl_arn|`utf8`|
|xray_enabled|`bool`|