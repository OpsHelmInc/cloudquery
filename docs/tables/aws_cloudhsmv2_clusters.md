# Table: aws_cloudhsmv2_clusters

This table shows data for AWS CloudHSM v2 Clusters.

https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_Cluster.html

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
|backup_policy|`utf8`|
|backup_retention_policy|`json`|
|certificates|`json`|
|cluster_id|`utf8`|
|create_timestamp|`timestamp[us, tz=UTC]`|
|hsm_type|`utf8`|
|hsms|`json`|
|mode|`utf8`|
|pre_co_password|`utf8`|
|security_group|`utf8`|
|source_backup_id|`utf8`|
|state|`utf8`|
|state_message|`utf8`|
|subnet_mapping|`json`|
|vpc_id|`utf8`|