# Table: aws_docdb_certificates

This table shows data for Amazon DocumentDB Certificates.

https://docs.aws.amazon.com/documentdb/latest/developerguide/API_Certificate.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|certificate_arn|`utf8`|
|certificate_identifier|`utf8`|
|certificate_type|`utf8`|
|thumbprint|`utf8`|
|valid_from|`timestamp[us, tz=UTC]`|
|valid_till|`timestamp[us, tz=UTC]`|