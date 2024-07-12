# Table: aws_sagemaker_models

This table shows data for Amazon SageMaker Models.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeModel.html

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
|creation_time|`timestamp[us, tz=UTC]`|
|model_arn|`utf8`|
|model_name|`utf8`|
|containers|`json`|
|deployment_recommendation|`json`|
|enable_network_isolation|`bool`|
|execution_role_arn|`utf8`|
|inference_execution_config|`json`|
|primary_container|`json`|
|vpc_config|`json`|