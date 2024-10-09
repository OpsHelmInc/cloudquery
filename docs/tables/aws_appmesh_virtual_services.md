# Table: aws_appmesh_virtual_services

This table shows data for AWS App Mesh Virtual Services.

https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_VirtualServiceData.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_appmesh_meshes](aws_appmesh_meshes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|mesh_arn|`utf8`|
|oh_resource_type|`utf8`|
|mesh_name|`utf8`|
|metadata|`json`|
|spec|`json`|
|status|`json`|
|virtual_service_name|`utf8`|