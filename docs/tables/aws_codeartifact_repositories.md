# Table: aws_codeartifact_repositories

This table shows data for AWS CodeArtifact Repositories.

https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_RepositoryDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|administrator_account|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|domain_name|`utf8`|
|domain_owner|`utf8`|
|external_connections|`json`|
|name|`utf8`|
|upstreams|`json`|