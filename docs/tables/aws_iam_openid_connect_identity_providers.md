# Table: aws_iam_openid_connect_identity_providers

This table shows data for IAM Openid Connect Identity Providers.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetOpenIDConnectProvider.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|client_id_list|`list<item: utf8, nullable>`|
|create_date|`timestamp[us, tz=UTC]`|
|thumbprint_list|`list<item: utf8, nullable>`|
|url|`utf8`|