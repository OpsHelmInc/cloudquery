# Table: aws_codebuild_source_credentials

This table shows data for AWS CodeBuild Source Credentials.

https://docs.aws.amazon.com/codebuild/latest/APIReference/API_SourceCredentialsInfo.html

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
|auth_type|`utf8`|
|resource|`utf8`|
|server_type|`utf8`|