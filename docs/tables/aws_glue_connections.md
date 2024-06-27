# Table: aws_glue_connections

This table shows data for Glue Connections.

https://docs.aws.amazon.com/glue/latest/webapi/API_Connection.html

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
|authentication_configuration|`json`|
|connection_properties|`json`|
|connection_type|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|last_connection_validation_time|`timestamp[us, tz=UTC]`|
|last_updated_by|`utf8`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|match_criteria|`list<item: utf8, nullable>`|
|name|`utf8`|
|physical_connection_requirements|`json`|
|status|`utf8`|
|status_reason|`utf8`|