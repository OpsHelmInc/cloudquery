# Table: aws_savingsplans_plans

This table shows data for Savingsplans Plans.

https://docs.aws.amazon.com/savingsplans/latest/APIReference/API_SavingsPlan.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|oh_resource_type|`utf8`|
|commitment|`utf8`|
|currency|`utf8`|
|description|`utf8`|
|ec2_instance_family|`utf8`|
|end|`utf8`|
|offering_id|`utf8`|
|payment_option|`utf8`|
|product_types|`list<item: utf8, nullable>`|
|recurring_payment_amount|`utf8`|
|region|`utf8`|
|returnable_until|`utf8`|
|savings_plan_arn|`utf8`|
|savings_plan_id|`utf8`|
|savings_plan_type|`utf8`|
|start|`utf8`|
|state|`utf8`|
|tags|`json`|
|term_duration_in_seconds|`int64`|
|upfront_payment_amount|`utf8`|