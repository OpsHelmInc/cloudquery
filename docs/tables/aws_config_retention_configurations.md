# Table: aws_config_retention_configurations

This table shows data for Config Retention Configurations.

https://docs.aws.amazon.com/config/latest/APIReference/API_RetentionConfiguration.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|name|`utf8`|
|retention_period_in_days|`int64`|