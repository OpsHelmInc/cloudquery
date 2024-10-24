# Table: aws_elasticache_user_groups

This table shows data for Elasticache User Groups.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UserGroup.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|engine|`utf8`|
|minimum_engine_version|`utf8`|
|pending_changes|`json`|
|replication_groups|`list<item: utf8, nullable>`|
|serverless_caches|`list<item: utf8, nullable>`|
|status|`utf8`|
|user_group_id|`utf8`|
|user_ids|`list<item: utf8, nullable>`|