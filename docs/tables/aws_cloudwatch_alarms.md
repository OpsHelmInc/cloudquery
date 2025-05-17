# Table: aws_cloudwatch_alarms

https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html

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
|dimensions|JSON|
|metric_alarm|JSON|
|tags|JSON|