# Table: aws_ecs_task_definitions

This table shows data for Amazon Elastic Container Service (ECS) Task Definitions.

https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html

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
|task_definition|`json`|