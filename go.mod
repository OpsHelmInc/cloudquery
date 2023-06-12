module github.com/OpsHelmInc/cloudquery

go 1.19

require (
	github.com/OpsHelmInc/ohaws v0.0.0-20230526181655-062369d4f0a8
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.25
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.63
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.19.10
	github.com/aws/aws-sdk-go-v2/service/account v1.10.6
	github.com/aws/aws-sdk-go-v2/service/acm v1.17.9
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.16.9
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.13.9
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.19.4
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.17.9
	github.com/aws/aws-sdk-go-v2/service/appstream v1.20.9
	github.com/aws/aws-sdk-go-v2/service/appsync v1.19.9
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.28.4
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.27.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.26.4
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.14.8
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.24.6
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.25.9
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.9
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.20.10
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.14.9
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.15.8
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.22.8
	github.com/aws/aws-sdk-go-v2/service/configservice v1.31.2
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.25.5
	github.com/aws/aws-sdk-go-v2/service/dax v1.12.10
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.18.10
	github.com/aws/aws-sdk-go-v2/service/docdb v1.21.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.19.5
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.98.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.18.9
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.16.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.26.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.19.12
	github.com/aws/aws-sdk-go-v2/service/eks v1.27.10
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.26.8
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.15.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.15.8
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.19.9
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.18.9
	github.com/aws/aws-sdk-go-v2/service/emr v1.24.2
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.19.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.16.10
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.23.6
	github.com/aws/aws-sdk-go-v2/service/fsx v1.28.10
	github.com/aws/aws-sdk-go-v2/service/glacier v1.14.9
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.21.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.19.12
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.16.9
	github.com/aws/aws-sdk-go-v2/service/inspector v1.13.8
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.11.9
	github.com/aws/aws-sdk-go-v2/service/iot v1.36.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.19.10
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.17.10
	github.com/aws/aws-sdk-go-v2/service/kms v1.20.11
	github.com/aws/aws-sdk-go-v2/service/lambda v1.33.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.26.4
	github.com/aws/aws-sdk-go-v2/service/mq v1.14.8
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.15.2
	github.com/aws/aws-sdk-go-v2/service/neptune v1.20.3
	github.com/aws/aws-sdk-go-v2/service/organizations v1.19.6
	github.com/aws/aws-sdk-go-v2/service/qldb v1.15.8
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.34.1
	github.com/aws/aws-sdk-go-v2/service/ram v1.18.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.44.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.27.9
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.14.9
	github.com/aws/aws-sdk-go-v2/service/route53 v1.27.7
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.14.8
	github.com/aws/aws-sdk-go-v2/service/s3 v1.33.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.31.5
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.74.0
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.1.11
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.19.8
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.18.2
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.17.2
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.14.10
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.17.4
	github.com/aws/aws-sdk-go-v2/service/sfn v1.17.9
	github.com/aws/aws-sdk-go-v2/service/sns v1.20.8
	github.com/aws/aws-sdk-go-v2/service/sqs v1.22.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.36.2
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.16.10
	github.com/aws/aws-sdk-go-v2/service/sts v1.19.0
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.16.4
	github.com/aws/aws-sdk-go-v2/service/waf v1.12.8
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.30.0
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.28.9
	github.com/aws/aws-sdk-go-v2/service/xray v1.16.9
	github.com/aws/smithy-go v1.13.5
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/gocarina/gocsv v0.0.0-20230406101422-6445c2b15027
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.5.9
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.29.1
	github.com/stretchr/testify v1.8.3
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.1.0
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.13.24 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.28 // indirect
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.28.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.10 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/getsentry/sentry-go v0.20.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	golang.org/x/exp v0.0.0-20230420155640-133eef4313cb // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.10 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.25.2
	github.com/aws/aws-sdk-go-v2/service/backup v1.20.9
	github.com/aws/aws-sdk-go-v2/service/glue v1.45.3
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.27 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.14.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.18.8
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.28.10
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.13.10
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/iancoleman/strcase v0.2.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.18 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/grpc v1.54.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
