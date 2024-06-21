module github.com/OpsHelmInc/cloudquery

go 1.22

toolchain go1.22.1

require (
	github.com/OpsHelmInc/ohaws v0.0.0-20240429043244-e94bbe1440a8
	github.com/apache/arrow/go/v15 v15.0.2
	github.com/aws/aws-sdk-go-v2 v1.30.0
	github.com/aws/aws-sdk-go-v2/config v1.27.11
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.15.2
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.26.0
	github.com/aws/aws-sdk-go-v2/service/account v1.14.0
	github.com/aws/aws-sdk-go-v2/service/acm v1.22.0
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.31.1
	github.com/aws/aws-sdk-go-v2/service/amp v1.26.1
	github.com/aws/aws-sdk-go-v2/service/amplify v1.22.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.21.0
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.18.0
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.30.1
	github.com/aws/aws-sdk-go-v2/service/appflow v1.42.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.25.0
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.26.1
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.25.0
	github.com/aws/aws-sdk-go-v2/service/appstream v1.29.0
	github.com/aws/aws-sdk-go-v2/service/appsync v1.26.0
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.34.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.36.0
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.21.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.40.1
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.41.0
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.32.0
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.19.0
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.35.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.31.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.29.0
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.29.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.26.0
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.23.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.22.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.21.0
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.31.0
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.36.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.43.0
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.39.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.35.0
	github.com/aws/aws-sdk-go-v2/service/dax v1.17.0
	github.com/aws/aws-sdk-go-v2/service/detective v1.28.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.22.0
	github.com/aws/aws-sdk-go-v2/service/docdb v1.29.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.21.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.158.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.0
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.21.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.41.7
	github.com/aws/aws-sdk-go-v2/service/efs v1.26.0
	github.com/aws/aws-sdk-go-v2/service/eks v1.35.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.34.0
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.20.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.21.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.26.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.24.0
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.24.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.35.0
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.26.0
	github.com/aws/aws-sdk-go-v2/service/firehose v1.22.0
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.29.0
	github.com/aws/aws-sdk-go-v2/service/fsx v1.39.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.19.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.35.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.32.0
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.21.0
	github.com/aws/aws-sdk-go-v2/service/inspector v1.19.0
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.20.0
	github.com/aws/aws-sdk-go-v2/service/iot v1.46.0
	github.com/aws/aws-sdk-go-v2/service/kafka v1.28.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.27.4
	github.com/aws/aws-sdk-go-v2/service/kms v1.27.0
	github.com/aws/aws-sdk-go-v2/service/lambda v1.49.0
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.32.0
	github.com/aws/aws-sdk-go-v2/service/mq v1.20.0
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.22.0
	github.com/aws/aws-sdk-go-v2/service/neptune v1.27.0
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.38.5
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.27.1
	github.com/aws/aws-sdk-go-v2/service/organizations v1.23.0
	github.com/aws/aws-sdk-go-v2/service/qldb v1.19.0
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.53.0
	github.com/aws/aws-sdk-go-v2/service/ram v1.23.0
	github.com/aws/aws-sdk-go-v2/service/rds v1.77.2
	github.com/aws/aws-sdk-go-v2/service/redshift v1.39.0
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.22.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.19.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.35.0
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.20.0
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.22.1
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.18.1
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.29.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.53.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.44.6
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.119.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.20.1
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.6.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.25.0
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.50.2
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.25.0
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.24.0
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.30.1
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.19.0
	github.com/aws/aws-sdk-go-v2/service/ses v1.23.1
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.24.0
	github.com/aws/aws-sdk-go-v2/service/sfn v1.24.0
	github.com/aws/aws-sdk-go-v2/service/signer v1.23.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.44.0
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.23.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.6
	github.com/aws/aws-sdk-go-v2/service/support v1.23.1
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.23.0
	github.com/aws/aws-sdk-go-v2/service/waf v1.18.0
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.43.0
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.31.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.35.0
	github.com/aws/aws-sdk-go-v2/service/xray v1.23.0
	github.com/aws/smithy-go v1.20.2
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/cloudquery/plugins/source/aws v0.13.24
	github.com/cloudquery/plugin-sdk v1.11.2
	github.com/cloudquery/plugin-sdk/v4 v4.46.0
	github.com/gocarina/gocsv v0.0.0-20231116093920-b87c2d0e983a
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.33.0
	github.com/stretchr/testify v1.9.0
	github.com/thoas/go-funk v0.9.3
	golang.org/x/sync v0.7.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/squirrel v1.5.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apache/arrow/go/v16 v16.1.0 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.4 // indirect
	github.com/cloudquery/cq-provider-sdk v0.14.7 // indirect
	github.com/cloudquery/faker/v3 v3.7.7 // indirect
	github.com/creasty/defaults v1.6.0 // indirect
	github.com/doug-martin/goqu/v9 v9.18.0 // indirect
	github.com/elliotchance/orderedmap v1.4.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/georgysavva/scany v1.1.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v1.2.2 // indirect
	github.com/hashicorp/go-plugin v1.4.4 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jackc/puddle v1.2.1 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/lorenzosaino/go-sysctl v0.3.1 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pbnjay/memory v0.0.0-20210728143218-7b4eea64cf58 // indirect
	github.com/segmentio/stats/v4 v4.6.3 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/cobra v1.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tidwall/gjson v1.17.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xo/dburl v0.11.0 // indirect
	github.com/zclconf/go-cty v1.10.0 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/exp v0.0.0-20240531132922-fd00a4e0eefc // indirect
	golang.org/x/exp/typeparams v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240515191416-fc5f0ca64291 // indirect
	honnef.co/go/tools v0.3.3 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.36.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.30.0
	github.com/aws/aws-sdk-go-v2/service/glue v1.71.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.7 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.23.0
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.39.0
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.19.0
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0
	github.com/iancoleman/strcase v0.2.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
