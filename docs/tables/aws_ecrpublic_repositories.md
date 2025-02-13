# Table: aws_ecrpublic_repositories

This table shows data for Ecrpublic Repositories.

https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_Repository.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_ecrpublic_repositories:
  - [aws_ecrpublic_repository_images](aws_ecrpublic_repository_images.md)

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
|created_at|`timestamp[us, tz=UTC]`|
|registry_id|`utf8`|
|repository_arn|`utf8`|
|repository_name|`utf8`|
|repository_uri|`utf8`|