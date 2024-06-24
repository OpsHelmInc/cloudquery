# Table: aws_rds_cluster_parameter_groups

This table shows data for Amazon Relational Database Service (RDS) Cluster Parameter Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBClusterParameterGroup.html

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_rds_cluster_parameter_groups:
  - [aws_rds_cluster_parameter_group_parameters](aws_rds_cluster_parameter_group_parameters.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|db_cluster_parameter_group_arn|`utf8`|
|db_cluster_parameter_group_name|`utf8`|
|db_parameter_group_family|`utf8`|
|description|`utf8`|