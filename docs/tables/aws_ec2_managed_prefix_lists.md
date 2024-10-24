# Table: aws_ec2_managed_prefix_lists

This table shows data for Amazon Elastic Compute Cloud (EC2) Managed Prefix Lists.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ManagedPrefixList.html.
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|request_region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|oh_resource_type|`utf8`|
|address_family|`utf8`|
|max_entries|`int64`|
|owner_id|`utf8`|
|prefix_list_arn|`utf8`|
|prefix_list_id|`utf8`|
|prefix_list_name|`utf8`|
|state|`utf8`|
|state_message|`utf8`|
|version|`int64`|