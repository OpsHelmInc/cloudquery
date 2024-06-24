# Table: aws_ecr_registries

This table shows data for Amazon Elastic Container Registry (ECR) Registries.

https://docs.aws.amazon.com/AmazonECR/latest/APIReference/API_DescribeRegistry.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|registry_id|`utf8`|
|replication_configuration|`json`|