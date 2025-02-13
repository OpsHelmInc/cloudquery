# Table: aws_kinesis_streams

https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html

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
|stream_description_summary|JSON|
|tags|JSON|