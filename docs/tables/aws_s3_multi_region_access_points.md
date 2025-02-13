# Table: aws_s3_multi_region_access_points

This table shows data for S3 Multi Region Access Points.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_MultiRegionAccessPointReport.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|alias|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|public_access_block|`json`|
|regions|`json`|
|status|`utf8`|