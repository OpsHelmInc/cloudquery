# Table: aws_apprunner_vpc_connectors

This table shows data for AWS App Runner VPC Connectors.

https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html

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
|created_at|`timestamp[us, tz=UTC]`|
|deleted_at|`timestamp[us, tz=UTC]`|
|security_groups|`list<item: utf8, nullable>`|
|status|`utf8`|
|subnets|`list<item: utf8, nullable>`|
|vpc_connector_arn|`utf8`|
|vpc_connector_name|`utf8`|
|vpc_connector_revision|`int64`|