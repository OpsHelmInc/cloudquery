# Table: aws_codepipeline_pipelines

This table shows data for Codepipeline Pipelines.

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_GetPipeline.html

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
|metadata|`json`|
|pipeline|`json`|