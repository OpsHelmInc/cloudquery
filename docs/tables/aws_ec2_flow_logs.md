# Table: aws_ec2_flow_logs

This table shows data for Amazon Elastic Compute Cloud (EC2) Flow Logs.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FlowLog.html

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
|creation_time|`timestamp[us, tz=UTC]`|
|deliver_cross_account_role|`utf8`|
|deliver_logs_error_message|`utf8`|
|deliver_logs_permission_arn|`utf8`|
|deliver_logs_status|`utf8`|
|destination_options|`json`|
|flow_log_id|`utf8`|
|flow_log_status|`utf8`|
|log_destination|`utf8`|
|log_destination_type|`utf8`|
|log_format|`utf8`|
|log_group_name|`utf8`|
|max_aggregation_interval|`int64`|
|resource_id|`utf8`|
|traffic_type|`utf8`|