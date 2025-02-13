# Table: aws_elbv1_load_balancers

This table shows data for Amazon Elastic Load Balancer (ELB) v1 Load Balancers.

https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_LoadBalancerDescription.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_elbv1_load_balancers:
  - [aws_elbv1_load_balancer_policies](aws_elbv1_load_balancer_policies.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|availability_zones|`list<item: utf8, nullable>`|
|backend_server_descriptions|`json`|
|canonical_hosted_zone_name|`utf8`|
|canonical_hosted_zone_name_id|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|dns_name|`utf8`|
|health_check|`json`|
|instances|`json`|
|listener_descriptions|`json`|
|load_balancer_name|`utf8`|
|policies|`json`|
|scheme|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|source_security_group|`json`|
|subnets|`list<item: utf8, nullable>`|
|vpc_id|`utf8`|
|access_log|`json`|
|additional_attributes|`json`|
|connection_draining|`json`|
|connection_settings|`json`|
|cross_zone_load_balancing|`json`|
|tags|`json`|