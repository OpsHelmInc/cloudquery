# Table: aws_wafregional_rules

This table shows data for Wafregional Rules.

https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_Rule.html

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
|predicates|`json`|
|rule_id|`utf8`|
|metric_name|`utf8`|
|name|`utf8`|