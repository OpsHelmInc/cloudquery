# Table: aws_appmesh_meshes

This table shows data for AWS App Mesh Meshes.

https://docs.aws.amazon.com/app-mesh/latest/APIReference/API_MeshData.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on aws_appmesh_meshes:
  - [aws_appmesh_virtual_gateways](aws_appmesh_virtual_gateways.md)
  - [aws_appmesh_virtual_nodes](aws_appmesh_virtual_nodes.md)
  - [aws_appmesh_virtual_routers](aws_appmesh_virtual_routers.md)
  - [aws_appmesh_virtual_services](aws_appmesh_virtual_services.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|mesh_name|`utf8`|
|metadata|`json`|
|spec|`json`|
|status|`json`|