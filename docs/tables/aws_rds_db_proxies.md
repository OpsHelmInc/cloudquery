# Table: aws_rds_db_proxies

This table shows data for Amazon Relational Database Service (RDS) DB Proxies.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBProxy.html

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
|auth|`json`|
|created_date|`timestamp[us, tz=UTC]`|
|db_proxy_arn|`utf8`|
|db_proxy_name|`utf8`|
|debug_logging|`bool`|
|endpoint|`utf8`|
|engine_family|`utf8`|
|idle_client_timeout|`int64`|
|require_tls|`bool`|
|role_arn|`utf8`|
|status|`utf8`|
|updated_date|`timestamp[us, tz=UTC]`|
|vpc_id|`utf8`|
|vpc_security_group_ids|`list<item: utf8, nullable>`|
|vpc_subnet_ids|`list<item: utf8, nullable>`|