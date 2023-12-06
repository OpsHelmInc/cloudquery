module github.com/OpsHelmInc/cloudquery

go 1.21

toolchain go1.21.4

require (
	github.com/OpsHelmInc/ohaws v0.0.0-20231127135300-5f0443d71131
	github.com/aws/aws-sdk-go-v2 v1.23.5
	github.com/aws/aws-sdk-go-v2/config v1.25.9
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.15.2
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.26.0
	github.com/aws/aws-sdk-go-v2/service/account v1.14.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.22.0
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.21.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.18.0
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.25.0
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.25.0
	github.com/aws/aws-sdk-go-v2/service/appstream v1.29.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.26.0
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.36.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.41.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.32.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.19.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.35.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.31.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.29.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.26.0
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.22.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.21.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.31.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.43.0
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.35.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.17.0
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.22.0
	github.com/aws/aws-sdk-go-v2/service/docdb v1.29.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.138.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.0
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.21.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.35.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.26.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.35.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.34.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.20.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.21.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.26.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.24.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.35.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.26.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.22.0
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.29.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.39.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.19.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.35.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.28.0
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.21.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.19.0
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.20.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.46.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.28.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.24.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.27.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.49.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.32.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.20.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.22.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.27.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.23.0
	github.com/aws/aws-sdk-go-v2/service/qldb v1.19.0
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.53.0
	github.com/aws/aws-sdk-go-v2/service/ram v1.23.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.64.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.39.0
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.19.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.35.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.20.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.0
	github.com/aws/aws-sdk-go-v2/service/s3control v1.41.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.119.0
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.6.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.25.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.25.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.24.0
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.19.0
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.24.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.24.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.44.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.23.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.0
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.23.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.18.0
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.43.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.35.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.23.0
	github.com/aws/smithy-go v1.18.1
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/gocarina/gocsv v0.0.0-20231116093920-b87c2d0e983a
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.31.0
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.5.0
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.16.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.36.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.0 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/exp v0.0.0-20231127185646-65229373498e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.6 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.6 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.36.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.30.0
	github.com/aws/aws-sdk-go-v2/service/glue v1.71.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.23.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.39.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.19.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
