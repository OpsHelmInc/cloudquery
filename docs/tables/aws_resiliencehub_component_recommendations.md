# Table: aws_resiliencehub_component_recommendations

This table shows data for AWS Resilience Hub Component Recommendations.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ComponentRecommendation.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_resiliencehub_app_assessments](aws_resiliencehub_app_assessments.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|app_arn|`utf8`|
|assessment_arn|`utf8`|
|app_component_name|`utf8`|
|config_recommendations|`json`|
|recommendation_status|`utf8`|