# Table: aws_shield_protections

This table shows data for Shield Protections.

https://docs.aws.amazon.com/waf/latest/DDOSAPIReference/API_Protection.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|application_layer_automatic_response_configuration|`json`|
|health_check_ids|`list<item: utf8, nullable>`|
|id|`utf8`|
|name|`utf8`|
|protection_arn|`utf8`|
|resource_arn|`utf8`|