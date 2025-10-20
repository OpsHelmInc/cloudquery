module github.com/OpsHelmInc/cloudquery

go 1.25

require (
	github.com/OpsHelmInc/ohaws v0.11.3-0.20251014173354-8d84db7cb1f0
	github.com/aws/aws-sdk-go-v2 v1.39.3
	github.com/aws/aws-sdk-go-v2/config v1.31.8
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.61
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.36.13
	github.com/aws/aws-sdk-go-v2/service/account v1.22.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.37.7
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.35.7
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.32.7
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.34.13
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.32.15
	github.com/aws/aws-sdk-go-v2/service/appstream v1.43.2
	github.com/aws/aws-sdk-go-v2/service/appsync v1.43.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.59.4
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.67.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.55.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.29.7
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.53.7
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.51.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.58.3
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.67.6
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.46.7
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.28.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.57.8
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.12
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.47.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.15
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.12
	github.com/aws/aws-sdk-go-v2/service/docdb v1.40.10
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.51.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.257.2
	github.com/aws/aws-sdk-go-v2/service/ecr v1.50.6
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.37.7
	github.com/aws/aws-sdk-go-v2/service/ecs v1.65.2
	github.com/aws/aws-sdk-go-v2/service/efs v1.40.9
	github.com/aws/aws-sdk-go-v2/service/eks v1.74.3
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.50.6
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.28.16
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.33.7
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.51.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.37.7
	github.com/aws/aws-sdk-go-v2/service/emr v1.47.12
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.45.6
	github.com/aws/aws-sdk-go-v2/service/firehose v1.41.7
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.35.15
	github.com/aws/aws-sdk-go-v2/service/fsx v1.52.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.26.16
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.52.10
	github.com/aws/aws-sdk-go-v2/service/iam v1.47.8
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.33.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.25.15
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.34.9
	github.com/aws/aws-sdk-go-v2/service/iot v1.69.6
	github.com/aws/aws-sdk-go-v2/service/kafka v1.44.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.40.6
	github.com/aws/aws-sdk-go-v2/service/kms v1.46.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.78.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.15
	github.com/aws/aws-sdk-go-v2/service/mq v1.34.5
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.33.10
	github.com/aws/aws-sdk-go-v2/service/neptune v1.35.17
	github.com/aws/aws-sdk-go-v2/service/organizations v1.45.4
	github.com/aws/aws-sdk-go-v2/service/qldb v1.31.0
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.83.6
	github.com/aws/aws-sdk-go-v2/service/ram v1.29.19
	github.com/aws/aws-sdk-go-v2/service/rds v1.108.3
	github.com/aws/aws-sdk-go-v2/service/redshift v1.59.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.27.18
	github.com/aws/aws-sdk-go-v2/service/route53 v1.58.5
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.34.5
	github.com/aws/aws-sdk-go-v2/service/s3 v1.88.5
	github.com/aws/aws-sdk-go-v2/service/s3control v1.66.3
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.215.4
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.17.6
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.39.7
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.32.15
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.30.15
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.25.18
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.41.5
	github.com/aws/aws-sdk-go-v2/service/sfn v1.34.12
	github.com/aws/aws-sdk-go-v2/service/sns v1.38.6
	github.com/aws/aws-sdk-go-v2/service/sqs v1.42.9
	github.com/aws/aws-sdk-go-v2/service/ssm v1.66.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.36.3
	github.com/aws/aws-sdk-go-v2/service/sts v1.38.7
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.29.16
	github.com/aws/aws-sdk-go-v2/service/waf v1.25.15
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.68.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.52.5
	github.com/aws/aws-sdk-go-v2/service/xray v1.30.12
	github.com/aws/smithy-go v1.23.1
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.7.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.34.0
	github.com/stretchr/testify v1.11.1
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.17.0
)

require (
	github.com/OpsHelmInc/pkg v0.0.0-20251017232113-d3a508fc3b33 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.18.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.45.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.38.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v1.58.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.40.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.28.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.24.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.57.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.52.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/personalize v1.46.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.33.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.39.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.34.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.29.6 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/oklog/ulid/v2 v2.1.1 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	golang.org/x/exp v0.0.0-20251017212417-90e834f514db // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.10 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.55.7
	github.com/aws/aws-sdk-go-v2/service/backup v1.49.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.131.1
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.11.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.29.15
	github.com/aws/aws-sdk-go-v2/service/sso v1.29.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.56.3
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.25.15
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/mod v0.29.0 // indirect
	golang.org/x/net v0.46.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	golang.org/x/tools v0.38.0 // indirect
	google.golang.org/grpc v1.75.0 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
