# Table: aws_acm_certificates

https://docs.aws.amazon.com/acm/latest/APIReference/API_CertificateDetail.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|certificate_detail|JSON|
|tags|JSON|