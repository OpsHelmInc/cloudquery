# Table: aws_directconnect_virtual_interfaces

This table shows data for AWS Direct Connect Virtual Interfaces.

https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualInterface.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|id|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|address_family|`utf8`|
|amazon_address|`utf8`|
|amazon_side_asn|`int64`|
|asn|`int64`|
|auth_key|`utf8`|
|aws_device_v2|`utf8`|
|aws_logical_device_id|`utf8`|
|bgp_peers|`json`|
|connection_id|`utf8`|
|customer_address|`utf8`|
|customer_router_config|`utf8`|
|direct_connect_gateway_id|`utf8`|
|jumbo_frame_capable|`bool`|
|location|`utf8`|
|mtu|`int64`|
|owner_account|`utf8`|
|region|`utf8`|
|route_filter_prefixes|`json`|
|site_link_enabled|`bool`|
|virtual_gateway_id|`utf8`|
|virtual_interface_id|`utf8`|
|virtual_interface_name|`utf8`|
|virtual_interface_state|`utf8`|
|virtual_interface_type|`utf8`|
|vlan|`int64`|