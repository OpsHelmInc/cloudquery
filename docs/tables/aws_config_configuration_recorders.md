# Table: aws_config_configuration_recorders



The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn|String|
|name|String|
|recording_group|JSON|
|recording_mode|JSON|
|recording_scope|String|
|role_arn|String|
|service_principal|String|
|status_last_error_code|String|
|status_last_error_message|String|
|status_last_start_time|Timestamp|
|status_last_status|String|
|status_last_status_change_time|Timestamp|
|status_last_stop_time|Timestamp|
|status_recording|Bool|