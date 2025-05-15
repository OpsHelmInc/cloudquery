# Table: aws_backup_plans

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_backup_plans:
  - [aws_backup_plan_selections](aws_backup_plan_selections.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|get_backup_plan_output|JSON|
|tags|JSON|