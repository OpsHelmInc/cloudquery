# Table: aws_xray_sampling_rules

This table shows data for AWS X-Ray Sampling Rules.

https://docs.aws.amazon.com/xray/latest/api/API_SamplingRuleRecord.html

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
|modified_at|`timestamp[us, tz=UTC]`|
|sampling_rule|`json`|