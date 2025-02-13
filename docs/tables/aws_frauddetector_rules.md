# Table: aws_frauddetector_rules

This table shows data for Amazon Fraud Detector Rules.

https://docs.aws.amazon.com/frauddetector/latest/api/API_RuleDetail.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_frauddetector_detectors](aws_frauddetector_detectors.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|created_time|`utf8`|
|description|`utf8`|
|detector_id|`utf8`|
|expression|`utf8`|
|language|`utf8`|
|last_updated_time|`utf8`|
|outcomes|`list<item: utf8, nullable>`|
|rule_id|`utf8`|
|rule_version|`utf8`|