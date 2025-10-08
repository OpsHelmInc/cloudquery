# Table: aws_budgets_budgets

https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsbudgetservice.html

The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|budget|JSON|