# Table: aws_apprunner_connections

This table shows data for AWS App Runner Connections.

https://docs.aws.amazon.com/apprunner/latest/api/API_Connection.html

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
|connection_arn|`utf8`|
|connection_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|provider_type|`utf8`|
|status|`utf8`|