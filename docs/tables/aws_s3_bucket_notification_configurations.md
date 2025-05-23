# Table: aws_s3_bucket_notification_configurations

This table shows data for S3 Bucket Notification Configurations.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_NotificationConfiguration.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_s3_buckets](aws_s3_buckets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|bucket_arn|`utf8`|
|event_bridge_configuration|`json`|
|lambda_function_configurations|`json`|
|queue_configurations|`json`|
|topic_configurations|`json`|