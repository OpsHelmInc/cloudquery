# Table: aws_iam_saml_identity_providers

This table shows data for IAM Saml Identity Providers.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|create_date|`timestamp[us, tz=UTC]`|
|saml_metadata_document|`utf8`|
|tags|`json`|
|valid_until|`timestamp[us, tz=UTC]`|