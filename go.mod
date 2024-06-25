module github.com/OpsHelmInc/cloudquery

go 1.22

toolchain go1.22.1

require (
	github.com/apache/arrow/go/v16 v16.1.0
	github.com/aws/aws-sdk-go-v2 v1.30.0
	github.com/aws/aws-sdk-go-v2/config v1.27.21
	github.com/aws/aws-sdk-go-v2/credentials v1.17.21
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.31.1
	github.com/aws/aws-sdk-go-v2/service/account v1.18.1
	github.com/aws/aws-sdk-go-v2/service/acm v1.27.1
	github.com/aws/aws-sdk-go-v2/service/acmpca v1.31.1
	github.com/aws/aws-sdk-go-v2/service/amp v1.26.1
	github.com/aws/aws-sdk-go-v2/service/amplify v1.22.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.24.1
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.21.1
	github.com/aws/aws-sdk-go-v2/service/appconfig v1.30.1
	github.com/aws/aws-sdk-go-v2/service/appflow v1.42.1
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.28.1
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.26.1
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.29.1
	github.com/aws/aws-sdk-go-v2/service/appstream v1.35.1
	github.com/aws/aws-sdk-go-v2/service/appsync v1.33.1
	github.com/aws/aws-sdk-go-v2/service/auditmanager v1.34.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.42.0
	github.com/aws/aws-sdk-go-v2/service/autoscalingplans v1.21.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.40.1
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.52.1
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.37.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.23.1
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.41.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.39.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.36.1
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.29.0
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.39.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.23.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.29.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v1.24.1
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.40.1
	github.com/aws/aws-sdk-go-v2/service/computeoptimizer v1.36.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.47.1
	github.com/aws/aws-sdk-go-v2/service/costexplorer v1.39.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.39.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.20.1
	github.com/aws/aws-sdk-go-v2/service/detective v1.28.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.26.0
	github.com/aws/aws-sdk-go-v2/service/docdb v1.35.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.33.2
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.21.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.166.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.29.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.24.1
	github.com/aws/aws-sdk-go-v2/service/ecs v1.43.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.30.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.44.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.39.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.24.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.25.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.32.1
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.29.1
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.24.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.40.1
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.32.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.30.1
	github.com/aws/aws-sdk-go-v2/service/frauddetector v1.32.1
	github.com/aws/aws-sdk-go-v2/service/fsx v1.45.1
	github.com/aws/aws-sdk-go-v2/service/glacier v1.23.1
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.44.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.33.1
	github.com/aws/aws-sdk-go-v2/service/identitystore v1.24.1
	github.com/aws/aws-sdk-go-v2/service/inspector v1.22.1
	github.com/aws/aws-sdk-go-v2/service/inspector2 v1.27.1
	github.com/aws/aws-sdk-go-v2/service/iot v1.54.1
	github.com/aws/aws-sdk-go-v2/service/kafka v1.34.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.28.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.34.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.55.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.39.1
	github.com/aws/aws-sdk-go-v2/service/mq v1.23.1
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.28.1
	github.com/aws/aws-sdk-go-v2/service/neptune v1.32.1
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.39.1
	github.com/aws/aws-sdk-go-v2/service/networkmanager v1.28.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.28.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.22.1
	github.com/aws/aws-sdk-go-v2/service/quicksight v1.65.1
	github.com/aws/aws-sdk-go-v2/service/ram v1.26.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.80.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.45.1
	github.com/aws/aws-sdk-go-v2/service/resiliencehub v1.22.1
	github.com/aws/aws-sdk-go-v2/service/resourcegroups v1.23.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.41.1
	github.com/aws/aws-sdk-go-v2/service/route53domains v1.24.1
	github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig v1.22.1
	github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness v1.18.1
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.29.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.56.1
	github.com/aws/aws-sdk-go-v2/service/s3control v1.45.1
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.147.0
	github.com/aws/aws-sdk-go-v2/service/savingsplans v1.20.1
	github.com/aws/aws-sdk-go-v2/service/scheduler v1.9.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.31.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.50.2
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.29.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalogappregistry v1.27.1
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.30.1
	github.com/aws/aws-sdk-go-v2/service/servicequotas v1.22.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.23.1
	github.com/aws/aws-sdk-go-v2/service/sesv2 v1.31.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.28.1
	github.com/aws/aws-sdk-go-v2/service/signer v1.23.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.30.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.33.1
	github.com/aws/aws-sdk-go-v2/service/ssm v1.51.1
	github.com/aws/aws-sdk-go-v2/service/ssoadmin v1.26.1
	github.com/aws/aws-sdk-go-v2/service/sts v1.29.1
	github.com/aws/aws-sdk-go-v2/service/support v1.23.1
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.26.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.22.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.50.1
	github.com/aws/aws-sdk-go-v2/service/wellarchitected v1.31.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.40.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.26.1
	github.com/aws/smithy-go v1.20.2
	github.com/basgys/goxml2json v1.1.0
	github.com/cloudquery/codegen v0.3.12
	github.com/cloudquery/plugin-sdk/v4 v4.48.0
	github.com/gertd/go-pluralize v0.2.1
	github.com/gocarina/gocsv v0.0.0-20240520201108-78e41c74b4b1
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/invopop/jsonschema v0.12.0
	github.com/mitchellh/hashstructure/v2 v2.0.2
	github.com/mitchellh/mapstructure v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.33.0
	github.com/stretchr/testify v1.9.0
	github.com/thoas/go-funk v0.9.3
	github.com/wk8/go-ordered-map/v2 v2.1.8
	golang.org/x/sync v0.7.0
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/Joker/jade v1.1.3 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/apache/arrow/go/v13 v13.0.0 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.25.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/sonic v1.11.9 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cloudquery/cloudquery-api-go v1.11.3 // indirect
	github.com/cloudquery/plugin-pb-go v1.20.2 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/deepmap/oapi-codegen v1.16.3 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.4 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.10.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.22.0 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomarkdown/markdown v0.0.0-20240419095408-642f0ee99ae2 // indirect
	github.com/google/flatbuffers v24.3.25+incompatible // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/iris-contrib/schema v0.0.6 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jpillora/longestcommon v0.0.0-20161227235612-adb9d91ee629 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/blocks v0.0.8 // indirect
	github.com/kataras/golog v0.1.12 // indirect
	github.com/kataras/iris/v12 v12.2.11 // indirect
	github.com/kataras/pio v0.0.13 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/labstack/echo/v4 v4.12.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailgun/raymond/v2 v2.0.48 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/microcosm-cc/bluemonday v1.0.26 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tdewolff/minify/v2 v2.20.34 // indirect
	github.com/tdewolff/parse/v2 v2.7.15 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240624140628-dc46fd24d27d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v1.43.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.35.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.87.1
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/shield v1.26.1
	github.com/aws/aws-sdk-go-v2/service/sso v1.21.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/transfer v1.49.1
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.22.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/google/uuid v1.6.0
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.18.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.22.0 // indirect
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// github.com/cloudquery/jsonschema @ cqmain
replace github.com/invopop/jsonschema => github.com/cloudquery/jsonschema v0.0.0-20231018073309-6c617a23d42f
