# Table: aws_resiliencehub_test_recommendations

This table shows data for AWS Resilience Hub Test Recommendations.

https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_TestRecommendation.html

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
|reference_id|`utf8`|
|app_component_name|`utf8`|
|depends_on_alarms|`list<item: utf8, nullable>`|
|description|`utf8`|
|intent|`utf8`|
|items|`json`|
|name|`utf8`|
|prerequisite|`utf8`|
|recommendation_id|`utf8`|
|recommendation_status|`utf8`|
|risk|`utf8`|
|type|`utf8`|