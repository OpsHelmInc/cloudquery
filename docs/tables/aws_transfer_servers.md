# Table: aws_transfer_servers

This table shows data for Transfer Servers.

https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html

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
|as2_service_managed_egress_ip_addresses|`list<item: utf8, nullable>`|
|certificate|`utf8`|
|domain|`utf8`|
|endpoint_details|`json`|
|endpoint_type|`utf8`|
|host_key_fingerprint|`utf8`|
|identity_provider_details|`json`|
|identity_provider_type|`utf8`|
|logging_role|`utf8`|
|post_authentication_login_banner|`utf8`|
|pre_authentication_login_banner|`utf8`|
|protocol_details|`json`|
|protocols|`list<item: utf8, nullable>`|
|s3_storage_options|`json`|
|security_policy_name|`utf8`|
|server_id|`utf8`|
|state|`utf8`|
|structured_log_destinations|`list<item: utf8, nullable>`|
|user_count|`int64`|
|workflow_details|`json`|