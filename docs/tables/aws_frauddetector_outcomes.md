# Table: aws_frauddetector_outcomes

This table shows data for Amazon Fraud Detector Outcomes.

https://docs.aws.amazon.com/frauddetector/latest/api/API_Outcome.html

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
|created_time|`utf8`|
|description|`utf8`|
|last_updated_time|`utf8`|
|name|`utf8`|