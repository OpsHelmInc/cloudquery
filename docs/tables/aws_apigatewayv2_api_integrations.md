# Table: aws_apigatewayv2_api_integrations

This table shows data for Amazon API Gateway v2 API Integrations.

https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/apis-apiid-integrations-integrationid.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_apigatewayv2_apis](aws_apigatewayv2_apis.md).

The following tables depend on aws_apigatewayv2_api_integrations:
  - [aws_apigatewayv2_api_integration_responses](aws_apigatewayv2_api_integration_responses.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|api_arn|`utf8`|
|api_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|api_gateway_managed|`bool`|
|connection_id|`utf8`|
|connection_type|`utf8`|
|content_handling_strategy|`utf8`|
|credentials_arn|`utf8`|
|description|`utf8`|
|integration_id|`utf8`|
|integration_method|`utf8`|
|integration_response_selection_expression|`utf8`|
|integration_subtype|`utf8`|
|integration_type|`utf8`|
|integration_uri|`utf8`|
|passthrough_behavior|`utf8`|
|payload_format_version|`utf8`|
|request_parameters|`json`|
|request_templates|`json`|
|response_parameters|`json`|
|template_selection_expression|`utf8`|
|timeout_in_millis|`int64`|
|tls_config|`json`|