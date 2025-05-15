# Table: aws_firehose_delivery_streams

https://docs.aws.amazon.com/firehose/latest/APIReference/API_DeliveryStreamDescription.html

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
|delivery_stream_description|JSON|
|tags|JSON|