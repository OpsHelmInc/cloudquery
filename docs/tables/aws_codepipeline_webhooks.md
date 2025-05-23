# Table: aws_codepipeline_webhooks

This table shows data for Codepipeline Webhooks.

https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html

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
|definition|`json`|
|url|`utf8`|
|error_code|`utf8`|
|error_message|`utf8`|
|last_triggered|`timestamp[us, tz=UTC]`|