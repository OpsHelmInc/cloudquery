# Table: aws_sagemaker_models



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
|tags|JSON|
|oh_resource_type|String|
|creation_time|Timestamp|
|model_name|String|
|containers|JSON|
|deployment_recommendation|JSON|
|enable_network_isolation|Bool|
|execution_role_arn|String|
|inference_execution_config|JSON|
|primary_container|JSON|
|vpc_config|JSON|
|result_metadata|JSON|