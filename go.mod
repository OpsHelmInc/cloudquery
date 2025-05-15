module github.com/OpsHelmInc/cloudquery

go 1.24

require (
	github.com/OpsHelmInc/ohaws v0.10.0
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.14
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.61
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.36.13
	github.com/aws/aws-sdk-go-v2/service/account v1.22.8
	github.com/aws/aws-sdk-go-v2/service/acm v1.32.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.30.1
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.34.13
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.32.15
	github.com/aws/aws-sdk-go-v2/service/appstream v1.43.2
	github.com/aws/aws-sdk-go-v2/service/appsync v1.43.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.52.4
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.59.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.46.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.29.7
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.48.4
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.44.3
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.45.12
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.60.0
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.41.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.28.5
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.52.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.51.12
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.47.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.23.15
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.30.12
	github.com/aws/aws-sdk-go-v2/service/docdb v1.40.10
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.43.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.218.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.44.0
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.33.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.57.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.35.3
	github.com/aws/aws-sdk-go-v2/service/eks v1.64.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.46.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.28.16
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.29.3
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.45.2
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.33.3
	github.com/aws/aws-sdk-go-v2/service/emr v1.47.12
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.39.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.37.4
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.35.15
	github.com/aws/aws-sdk-go-v2/service/fsx v1.52.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.26.16
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.52.10
	github.com/aws/aws-sdk-go-v2/service/iam v1.42.0
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.28.3
	github.com/aws/aws-sdk-go-v2/service/inspector v1.25.15
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.34.9
	github.com/aws/aws-sdk-go-v2/service/iot v1.64.2
	github.com/aws/aws-sdk-go-v2/service/kafka v1.39.2
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.35.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.38.3
	github.com/aws/aws-sdk-go-v2/service/lambda v1.71.2
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.42.15
	github.com/aws/aws-sdk-go-v2/service/mq v1.29.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.33.10
	github.com/aws/aws-sdk-go-v2/service/neptune v1.35.17
	github.com/aws/aws-sdk-go-v2/service/organizations v1.38.3
	github.com/aws/aws-sdk-go-v2/service/qldb v1.26.2
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.83.6
	github.com/aws/aws-sdk-go-v2/service/ram v1.29.19
	github.com/aws/aws-sdk-go-v2/service/rds v1.95.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.54.3
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.27.18
	github.com/aws/aws-sdk-go-v2/service/route53 v1.51.1
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.29.2
	github.com/aws/aws-sdk-go-v2/service/s3 v1.79.3
	github.com/aws/aws-sdk-go-v2/service/s3control v1.58.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.192.0
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.13.3
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.35.4
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.32.15
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.30.15
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.25.18
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.41.5
	github.com/aws/aws-sdk-go-v2/service/sfn v1.34.12
	github.com/aws/aws-sdk-go-v2/service/sns v1.34.4
	github.com/aws/aws-sdk-go-v2/service/sqs v1.38.5
	github.com/aws/aws-sdk-go-v2/service/ssm v1.59.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.31.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.19
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.29.16
	github.com/aws/aws-sdk-go-v2/service/waf v1.25.15
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.60.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.52.5
	github.com/aws/aws-sdk-go-v2/service/xray v1.30.12
	github.com/aws/smithy-go v1.22.3
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.7.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.34.0
	github.com/stretchr/testify v1.10.0
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.12.0
)

require (
	github.com/OpsHelmInc/pkg v0.0.0-20250407012623-9b679ffd4807 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.67 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.40.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v1.32.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v1.52.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/comprehend v1.36.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/evidently v1.24.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafkaconnect v1.23.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/keyspaces v1.18.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.47.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/opensearch v1.46.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/personalize v1.41.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/schemas v1.29.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.35.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.30.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/verifiedpermissions v1.24.0 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250313205543-e70fdf4c4cb4 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.30 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.50.5
	github.com/aws/aws-sdk-go-v2/service/backup v1.41.2
	github.com/aws/aws-sdk-go-v2/service/glue v1.110.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.29.15
	github.com/aws/aws-sdk-go-v2/service/sso v1.25.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.56.3
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.25.15
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	google.golang.org/grpc v1.71.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
