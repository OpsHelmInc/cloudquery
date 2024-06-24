# Table: aws_glacier_data_retrieval_policies

This table shows data for Glacier Data Retrieval Policies.

https://docs.aws.amazon.com/amazonglacier/latest/dev/api-GetDataRetrievalPolicy.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|rules|`json`|