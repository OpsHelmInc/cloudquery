// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/OpsHelmInc/cloudquery/resources/services/accessanalyzer"
	"github.com/OpsHelmInc/cloudquery/resources/services/account"
	"github.com/OpsHelmInc/cloudquery/resources/services/acm"
	"github.com/OpsHelmInc/cloudquery/resources/services/apigateway"
	"github.com/OpsHelmInc/cloudquery/resources/services/apigatewayv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/applicationautoscaling"
	"github.com/OpsHelmInc/cloudquery/resources/services/apprunner"
	"github.com/OpsHelmInc/cloudquery/resources/services/appstream"
	"github.com/OpsHelmInc/cloudquery/resources/services/appsync"
	"github.com/OpsHelmInc/cloudquery/resources/services/athena"
	"github.com/OpsHelmInc/cloudquery/resources/services/autoscaling"
	"github.com/OpsHelmInc/cloudquery/resources/services/backup"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudformation"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudfront"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudhsmv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudtrail"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudwatch"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudwatchlogs"
	"github.com/OpsHelmInc/cloudquery/resources/services/codebuild"
	"github.com/OpsHelmInc/cloudquery/resources/services/codepipeline"
	"github.com/OpsHelmInc/cloudquery/resources/services/cognito"
	"github.com/OpsHelmInc/cloudquery/resources/services/config"
	"github.com/OpsHelmInc/cloudquery/resources/services/dax"
	"github.com/OpsHelmInc/cloudquery/resources/services/directconnect"
	"github.com/OpsHelmInc/cloudquery/resources/services/dms"
	"github.com/OpsHelmInc/cloudquery/resources/services/docdb"
	"github.com/OpsHelmInc/cloudquery/resources/services/dynamodb"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecr"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecrpublic"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecs"
	"github.com/OpsHelmInc/cloudquery/resources/services/efs"
	"github.com/OpsHelmInc/cloudquery/resources/services/eks"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticache"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticbeanstalk"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticsearch"
	"github.com/OpsHelmInc/cloudquery/resources/services/elbv1"
	"github.com/OpsHelmInc/cloudquery/resources/services/elbv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/emr"
	"github.com/OpsHelmInc/cloudquery/resources/services/eventbridge"
	"github.com/OpsHelmInc/cloudquery/resources/services/firehose"
	"github.com/OpsHelmInc/cloudquery/resources/services/frauddetector"
	"github.com/OpsHelmInc/cloudquery/resources/services/fsx"
	"github.com/OpsHelmInc/cloudquery/resources/services/glacier"
	"github.com/OpsHelmInc/cloudquery/resources/services/glue"
	"github.com/OpsHelmInc/cloudquery/resources/services/guardduty"
	"github.com/OpsHelmInc/cloudquery/resources/services/iam"
	"github.com/OpsHelmInc/cloudquery/resources/services/identitystore"
	"github.com/OpsHelmInc/cloudquery/resources/services/inspector"
	"github.com/OpsHelmInc/cloudquery/resources/services/inspector2"
	"github.com/OpsHelmInc/cloudquery/resources/services/iot"
	"github.com/OpsHelmInc/cloudquery/resources/services/kafka"
	"github.com/OpsHelmInc/cloudquery/resources/services/kinesis"
	"github.com/OpsHelmInc/cloudquery/resources/services/kms"
	"github.com/OpsHelmInc/cloudquery/resources/services/lambda"
	"github.com/OpsHelmInc/cloudquery/resources/services/lightsail"
	"github.com/OpsHelmInc/cloudquery/resources/services/mq"
	"github.com/OpsHelmInc/cloudquery/resources/services/mwaa"
	"github.com/OpsHelmInc/cloudquery/resources/services/neptune"
	"github.com/OpsHelmInc/cloudquery/resources/services/organizations"
	"github.com/OpsHelmInc/cloudquery/resources/services/qldb"
	"github.com/OpsHelmInc/cloudquery/resources/services/quicksight"
	"github.com/OpsHelmInc/cloudquery/resources/services/ram"
	"github.com/OpsHelmInc/cloudquery/resources/services/rds"
	"github.com/OpsHelmInc/cloudquery/resources/services/redshift"
	"github.com/OpsHelmInc/cloudquery/resources/services/resourcegroups"
	"github.com/OpsHelmInc/cloudquery/resources/services/route53"
	"github.com/OpsHelmInc/cloudquery/resources/services/s3"
	"github.com/OpsHelmInc/cloudquery/resources/services/sagemaker"
	"github.com/OpsHelmInc/cloudquery/resources/services/scheduler"
	"github.com/OpsHelmInc/cloudquery/resources/services/secretsmanager"
	"github.com/OpsHelmInc/cloudquery/resources/services/servicecatalog"
	"github.com/OpsHelmInc/cloudquery/resources/services/servicequotas"
	"github.com/OpsHelmInc/cloudquery/resources/services/ses"
	"github.com/OpsHelmInc/cloudquery/resources/services/shield"
	"github.com/OpsHelmInc/cloudquery/resources/services/sns"
	"github.com/OpsHelmInc/cloudquery/resources/services/sqs"
	"github.com/OpsHelmInc/cloudquery/resources/services/ssm"
	"github.com/OpsHelmInc/cloudquery/resources/services/ssoadmin"
	"github.com/OpsHelmInc/cloudquery/resources/services/stepfunctions"
	"github.com/OpsHelmInc/cloudquery/resources/services/timestream"
	"github.com/OpsHelmInc/cloudquery/resources/services/transfer"
	"github.com/OpsHelmInc/cloudquery/resources/services/waf"
	"github.com/OpsHelmInc/cloudquery/resources/services/wafregional"
	"github.com/OpsHelmInc/cloudquery/resources/services/wafv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/workspaces"
	"github.com/OpsHelmInc/cloudquery/resources/services/xray"
	"github.com/cloudquery/plugin-sdk/schema"
)

func tables() []*schema.Table {
	return []*schema.Table{
		accessanalyzer.Analyzers(),
		account.AlternateContacts(),
		account.Contacts(),
		acm.Certificates(),
		apigateway.ApiKeys(),
		apigateway.ClientCertificates(),
		apigateway.DomainNames(),
		apigateway.RestApis(),
		apigateway.UsagePlans(),
		apigateway.VpcLinks(),
		apigatewayv2.Apis(),
		apigatewayv2.DomainNames(),
		apigatewayv2.VpcLinks(),
		applicationautoscaling.Policies(),
		apprunner.AutoScalingConfigurations(),
		apprunner.Connections(),
		apprunner.ObservabilityConfigurations(),
		apprunner.Services(),
		apprunner.VpcConnectors(),
		apprunner.VpcIngressConnections(),
		appstream.AppBlocks(),
		appstream.Applications(),
		appstream.DirectoryConfigs(),
		appstream.Fleets(),
		appstream.ImageBuilders(),
		appstream.Images(),
		appstream.Stacks(),
		appstream.UsageReportSubscriptions(),
		appstream.Users(),
		appsync.GraphqlApis(),
		athena.DataCatalogs(),
		athena.WorkGroups(),
		autoscaling.LaunchConfigurations(),
		autoscaling.Groups(),
		autoscaling.ScheduledActions(),
		backup.GlobalSettings(),
		backup.Plans(),
		backup.RegionSettings(),
		backup.Vaults(),
		cloudformation.Stacks(),
		cloudfront.CachePolicies(),
		cloudfront.Distributions(),
		cloudhsmv2.Clusters(),
		cloudhsmv2.Backups(),
		cloudtrail.Trails(),
		cloudwatchlogs.ResourcePolicies(),
		cloudwatchlogs.MetricFilters(),
		cloudwatchlogs.LogGroups(),
		cloudwatch.Alarms(),
		codebuild.Projects(),
		codepipeline.Webhooks(),
		codepipeline.Pipelines(),
		cognito.IdentityPools(),
		cognito.UserPools(),
		config.ConfigurationRecorders(),
		config.ConformancePacks(),
		config.ConfigRules(),
		dax.Clusters(),
		directconnect.Connections(),
		directconnect.Gateways(),
		directconnect.Lags(),
		directconnect.VirtualGateways(),
		directconnect.VirtualInterfaces(),
		dms.ReplicationInstances(),
		docdb.Clusters(),
		docdb.ClusterParameterGroups(),
		docdb.Certificates(),
		docdb.EngineVersions(),
		docdb.SubnetGroups(),
		docdb.GlobalClusters(),
		docdb.Events(),
		docdb.EventSubscriptions(),
		docdb.EventCategories(),
		docdb.PendingMaintenanceActions(),
		dynamodb.Tables(),
		ec2.ByoipCidrs(),
		ec2.CustomerGateways(),
		ec2.EbsSnapshots(),
		ec2.EbsVolumes(),
		ec2.EgressOnlyInternetGateways(),
		ec2.Eips(),
		ec2.FlowLogs(),
		ec2.Hosts(),
		ec2.Images(),
		ec2.InstanceStatuses(),
		ec2.Instances(),
		ec2.InstanceTypes(),
		ec2.InternetGateways(),
		ec2.KeyPairs(),
		ec2.LaunchTemplates(),
		ec2.LaunchTemplateVersions(),
		ec2.NatGateways(),
		ec2.NetworkAcls(),
		ec2.NetworkInterfaces(),
		ec2.Regions(),
		ec2.RegionalConfigs(),
		ec2.ReservedInstances(),
		ec2.RouteTables(),
		ec2.SecurityGroups(),
		ec2.Subnets(),
		ec2.TransitGateways(),
		ec2.VpcEndpointServiceConfigurations(),
		ec2.VpcEndpointServices(),
		ec2.VpcEndpoints(),
		ec2.VpcPeeringConnections(),
		ec2.Vpcs(),
		ec2.VpnGateways(),
		ecrpublic.Repositories(),
		ecr.Registries(),
		ecr.RegistryPolicies(),
		ecr.Repositories(),
		ecs.Clusters(),
		ecs.TaskDefinitions(),
		efs.Filesystems(),
		eks.Clusters(),
		elasticache.Clusters(),
		elasticache.EngineVersions(),
		elasticache.GlobalReplicationGroups(),
		elasticache.ParameterGroups(),
		elasticache.ReplicationGroups(),
		elasticache.ReservedCacheNodesOfferings(),
		elasticache.ReservedCacheNodes(),
		elasticache.ServiceUpdates(),
		elasticache.Snapshots(),
		elasticache.UserGroups(),
		elasticache.Users(),
		elasticache.SubnetGroups(),
		elasticbeanstalk.ApplicationVersions(),
		elasticbeanstalk.Applications(),
		elasticbeanstalk.Environments(),
		elasticsearch.Domains(),
		elbv1.LoadBalancers(),
		elbv2.LoadBalancers(),
		elbv2.TargetGroups(),
		emr.BlockPublicAccessConfigs(),
		emr.Clusters(),
		eventbridge.EventBuses(),
		eventbridge.ApiDestinations(),
		eventbridge.Archives(),
		eventbridge.Connections(),
		eventbridge.EventSources(),
		eventbridge.Replays(),
		eventbridge.Endpoints(),
		firehose.DeliveryStreams(),
		frauddetector.BatchImports(),
		frauddetector.BatchPredictions(),
		frauddetector.Detectors(),
		frauddetector.EntityTypes(),
		frauddetector.EventTypes(),
		frauddetector.ExternalModels(),
		frauddetector.Labels(),
		frauddetector.Models(),
		frauddetector.Outcomes(),
		frauddetector.Variables(),
		fsx.Backups(),
		fsx.DataRepositoryAssociations(),
		fsx.DataRepositoryTasks(),
		fsx.FileCaches(),
		fsx.FileSystems(),
		fsx.Snapshots(),
		fsx.StorageVirtualMachines(),
		fsx.Volumes(),
		glacier.Vaults(),
		glacier.DataRetrievalPolicies(),
		glue.Classifiers(),
		glue.Connections(),
		glue.Crawlers(),
		glue.Databases(),
		glue.DatacatalogEncryptionSettings(),
		glue.DevEndpoints(),
		glue.Jobs(),
		glue.MlTransforms(),
		glue.Registries(),
		glue.SecurityConfigurations(),
		glue.Triggers(),
		glue.Workflows(),
		guardduty.Detectors(),
		iam.Accounts(),
		iam.CredentialReports(),
		iam.Groups(),
		iam.OpenidConnectIdentityProviders(),
		iam.PasswordPolicies(),
		iam.Policies(),
		iam.Roles(),
		iam.SamlIdentityProviders(),
		iam.ServerCertificates(),
		iam.Users(),
		iam.VirtualMfaDevices(),
		identitystore.Groups(),
		identitystore.Users(),
		inspector2.Findings(),
		inspector.Findings(),
		iot.BillingGroups(),
		iot.CaCertificates(),
		iot.Certificates(),
		iot.Jobs(),
		iot.Policies(),
		iot.SecurityProfiles(),
		iot.Streams(),
		iot.ThingGroups(),
		iot.ThingTypes(),
		iot.Things(),
		iot.TopicRules(),
		kafka.Clusters(),
		kafka.Configurations(),
		kinesis.Streams(),
		kms.Aliases(),
		kms.Keys(),
		lambda.Functions(),
		lambda.Layers(),
		lambda.Runtimes(),
		lightsail.Alarms(),
		lightsail.Buckets(),
		lightsail.Certificates(),
		lightsail.ContainerServices(),
		lightsail.DatabaseSnapshots(),
		lightsail.Databases(),
		lightsail.Disks(),
		lightsail.Distributions(),
		lightsail.InstanceSnapshots(),
		lightsail.Instances(),
		lightsail.LoadBalancers(),
		lightsail.StaticIps(),
		mq.Brokers(),
		mwaa.Environments(),
		neptune.ClusterParameterGroups(),
		neptune.ClusterSnapshots(),
		neptune.Clusters(),
		neptune.DbParameterGroups(),
		neptune.GlobalClusters(),
		neptune.EventSubscriptions(),
		neptune.Instances(),
		neptune.SubnetGroups(),
		organizations.Accounts(),
		qldb.Ledgers(),
		quicksight.Analyses(),
		quicksight.Dashboards(),
		quicksight.DataSets(),
		quicksight.DataSources(),
		quicksight.Folders(),
		quicksight.Groups(),
		quicksight.Templates(),
		quicksight.Users(),
		ram.Principals(),
		ram.Resources(),
		ram.ResourceShares(),
		ram.ResourceShareAssociations(),
		ram.ResourceShareInvitations(),
		ram.ResourceTypes(),
		rds.Certificates(),
		rds.EngineVersions(),
		rds.ClusterParameterGroups(),
		rds.ClusterSnapshots(),
		rds.Clusters(),
		rds.DbParameterGroups(),
		rds.DbSecurityGroups(),
		rds.DbSnapshots(),
		rds.EventSubscriptions(),
		rds.Instances(),
		rds.SubnetGroups(),
		redshift.Clusters(),
		redshift.EventSubscriptions(),
		redshift.SubnetGroups(),
		resourcegroups.ResourceGroups(),
		route53.DelegationSets(),
		route53.Domains(),
		route53.HealthChecks(),
		route53.HostedZones(),
		route53.TrafficPolicies(),
		s3.Accounts(),
		s3.Buckets(),
		sagemaker.EndpointConfigurations(),
		sagemaker.Models(),
		sagemaker.NotebookInstances(),
		sagemaker.TrainingJobs(),
		scheduler.ScheduleGroups(),
		scheduler.Schedules(),
		secretsmanager.Secrets(),
		servicecatalog.Portfolios(),
		servicecatalog.Products(),
		servicecatalog.ProvisionedProducts(),
		servicequotas.Services(),
		ses.Templates(),
		ses.ConfigurationSets(),
		ses.ContactLists(),
		ses.Identities(),
		shield.Attacks(),
		shield.ProtectionGroups(),
		shield.Protections(),
		shield.Subscriptions(),
		sns.Subscriptions(),
		sns.Topics(),
		sqs.Queues(),
		ssm.Documents(),
		ssm.Instances(),
		ssm.Parameters(),
		ssm.ComplianceSummaryItems(),
		ssm.Associations(),
		ssm.Inventories(),
		ssm.InventorySchemas(),
		ssm.PatchBaselines(),
		ssoadmin.Instances(),
		stepfunctions.StateMachines(),
		timestream.Databases(),
		transfer.Servers(),
		wafregional.RateBasedRules(),
		wafregional.RuleGroups(),
		wafregional.Rules(),
		wafregional.WebAcls(),
		waf.RuleGroups(),
		waf.Rules(),
		waf.SubscribedRuleGroups(),
		waf.WebAcls(),
		wafv2.Ipsets(),
		wafv2.ManagedRuleGroups(),
		wafv2.RegexPatternSets(),
		wafv2.RuleGroups(),
		wafv2.WebAcls(),
		workspaces.Workspaces(),
		workspaces.Directories(),
		xray.EncryptionConfigs(),
		xray.Groups(),
		xray.SamplingRules(),
	}
}
