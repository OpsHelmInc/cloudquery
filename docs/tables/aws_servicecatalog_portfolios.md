# Table: aws_servicecatalog_portfolios

This table shows data for AWS Service Catalog Portfolios.

https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribePortfolio.html

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
|budgets|`json`|
|portfolio_detail|`json`|
|tag_options|`json`|