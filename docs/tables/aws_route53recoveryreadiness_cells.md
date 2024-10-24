# Table: aws_route53recoveryreadiness_cells

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Readiness Cells.

https://docs.aws.amazon.com/recovery-readiness/latest/api/cells.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|cell_arn|`utf8`|
|cell_name|`utf8`|
|cells|`list<item: utf8, nullable>`|
|parent_readiness_scopes|`list<item: utf8, nullable>`|
|tags|`json`|