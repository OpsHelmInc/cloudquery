# Table: aws_glue_connections



The primary key for this table is **arn**.



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
|athena_properties|JSON|
|authentication_configuration|JSON|
|compatible_compute_environments|StringArray|
|connection_properties|JSON|
|connection_schema_version|Int|
|connection_type|String|
|creation_time|Timestamp|
|description|String|
|last_connection_validation_time|Timestamp|
|last_updated_by|String|
|last_updated_time|Timestamp|
|match_criteria|StringArray|
|name|String|
|physical_connection_requirements|JSON|
|python_properties|JSON|
|spark_properties|JSON|
|status|String|
|status_reason|String|