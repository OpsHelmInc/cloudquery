# Table: aws_sagemaker_endpoint_configurations



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
|creation_time|Timestamp|
|endpoint_config_name|String|
|production_variants|JSON|
|async_inference_config|JSON|
|data_capture_config|JSON|
|enable_network_isolation|Bool|
|execution_role_arn|String|
|explainer_config|JSON|
|kms_key_id|String|
|shadow_production_variants|JSON|
|vpc_config|JSON|
|result_metadata|JSON|