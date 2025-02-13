# Table: aws_route53_traffic_policies

This table shows data for Amazon Route 53 Traffic Policies.

https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_route53_traffic_policies:
  - [aws_route53_traffic_policy_versions](aws_route53_traffic_policy_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|id|`utf8`|
|latest_version|`int64`|
|name|`utf8`|
|traffic_policy_count|`int64`|
|type|`utf8`|