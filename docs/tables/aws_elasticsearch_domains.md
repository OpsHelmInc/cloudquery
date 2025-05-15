# Table: aws_elasticsearch_domains



The composite primary key for this table is (**account_id**, **region**, **id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|id (PK)|String|
|elasticsearch_domain_status|JSON|
|tags|JSON|