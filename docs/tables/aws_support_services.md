# Table: aws_support_services

This table shows data for Support Services.

https://docs.aws.amazon.com/awssupport/latest/APIReference/API_DescribeServices.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|language_code|`utf8`|
|categories|`json`|
|code|`utf8`|
|name|`utf8`|