# Table: aws_glue_security_configurations

This table shows data for Glue Security Configurations.

https://docs.aws.amazon.com/glue/latest/webapi/API_SecurityConfiguration.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|created_time_stamp|`timestamp[us, tz=UTC]`|
|encryption_configuration|`json`|
|name|`utf8`|