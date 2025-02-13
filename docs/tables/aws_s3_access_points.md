# Table: aws_s3_access_points

This table shows data for S3 Access Points.

https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AccessPoint.html

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
|bucket|`utf8`|
|name|`utf8`|
|network_origin|`utf8`|
|access_point_arn|`utf8`|
|alias|`utf8`|
|bucket_account_id|`utf8`|
|vpc_configuration|`json`|