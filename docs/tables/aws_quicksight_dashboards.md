# Table: aws_quicksight_dashboards

This table shows data for QuickSight Dashboards.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DashboardSummary.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|dashboard_id|`utf8`|
|last_published_time|`timestamp[us, tz=UTC]`|
|last_updated_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|published_version_number|`int64`|