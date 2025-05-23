# Table: aws_networkfirewall_tls_inspection_configurations

This table shows data for Networkfirewall TLS Inspection Configurations.

https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_DescribeTLSInspectionConfiguration.html

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
|server_certificate_configurations|`json`|
|tls_inspection_configuration_arn|`utf8`|
|tls_inspection_configuration_id|`utf8`|
|tls_inspection_configuration_name|`utf8`|
|certificate_authority|`json`|
|certificates|`json`|
|description|`utf8`|
|encryption_configuration|`json`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|number_of_associations|`int64`|
|tls_inspection_configuration_status|`utf8`|