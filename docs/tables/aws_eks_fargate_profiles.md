# Table: aws_eks_fargate_profiles

This table shows data for Amazon Elastic Kubernetes Service (EKS) Fargate Profiles.

https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|cluster_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|fargate_profile_arn|`utf8`|
|fargate_profile_name|`utf8`|
|health|`json`|
|pod_execution_role_arn|`utf8`|
|selectors|`json`|
|status|`utf8`|
|subnets|`list<item: utf8, nullable>`|
|tags|`json`|