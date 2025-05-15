# Table: aws_elbv2_load_balancers

https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_elbv2_load_balancers:
  - [aws_elbv2_listeners](aws_elbv2_listeners.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|web_acl_arn|String|
|arn (PK)|String|
|oh_resource_type|String|
|load_balancer|JSON|
|tags|JSON|
|unknown_attributes|JSON|