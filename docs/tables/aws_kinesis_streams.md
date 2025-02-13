# Table: aws_kinesis_streams

This table shows data for Kinesis Streams.

https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html

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
|enhanced_monitoring|`json`|
|open_shard_count|`int64`|
|retention_period_hours|`int64`|
|stream_arn|`utf8`|
|stream_creation_timestamp|`timestamp[us, tz=UTC]`|
|stream_name|`utf8`|
|stream_status|`utf8`|
|consumer_count|`int64`|
|encryption_type|`utf8`|
|key_id|`utf8`|
|stream_mode_details|`json`|
|tags|`json`|