# Table: aws_backup_vaults

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupVaultListMember.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_backup_vaults:
  - [aws_backup_vault_recovery_points](aws_backup_vault_recovery_points.md)

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
|access_policy|JSON|
|notifications|JSON|
|backup_vault_list_member|JSON|
|tags|JSON|