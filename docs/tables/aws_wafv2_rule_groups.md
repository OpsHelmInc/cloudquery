# Table: aws_wafv2_rule_groups

https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|policy|JSON|
|rule_group|JSON|
|tags|JSON|