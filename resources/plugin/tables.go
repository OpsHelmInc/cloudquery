package plugin

import (
	"fmt"

	"github.com/OpsHelmInc/cloudquery/resources/services/accessanalyzer"
	"github.com/OpsHelmInc/cloudquery/resources/services/acm"
	"github.com/OpsHelmInc/cloudquery/resources/services/acmpca"
	"github.com/OpsHelmInc/cloudquery/resources/services/amp"
	"github.com/OpsHelmInc/cloudquery/resources/services/amplify"
	"github.com/OpsHelmInc/cloudquery/resources/services/apigateway"
	"github.com/OpsHelmInc/cloudquery/resources/services/apigatewayv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/appconfig"
	"github.com/OpsHelmInc/cloudquery/resources/services/appflow"
	"github.com/OpsHelmInc/cloudquery/resources/services/applicationautoscaling"
	"github.com/OpsHelmInc/cloudquery/resources/services/appmesh"
	"github.com/OpsHelmInc/cloudquery/resources/services/apprunner"
	"github.com/OpsHelmInc/cloudquery/resources/services/appstream"
	"github.com/OpsHelmInc/cloudquery/resources/services/appsync"
	"github.com/OpsHelmInc/cloudquery/resources/services/athena"
	"github.com/OpsHelmInc/cloudquery/resources/services/auditmanager"
	"github.com/OpsHelmInc/cloudquery/resources/services/autoscaling"
	"github.com/OpsHelmInc/cloudquery/resources/services/backup"
	"github.com/OpsHelmInc/cloudquery/resources/services/batch"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudformation"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudfront"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudhsmv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudtrail"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudwatch"
	"github.com/OpsHelmInc/cloudquery/resources/services/cloudwatchlogs"
	"github.com/OpsHelmInc/cloudquery/resources/services/codeartifact"
	"github.com/OpsHelmInc/cloudquery/resources/services/codebuild"
	"github.com/OpsHelmInc/cloudquery/resources/services/codecommit"
	"github.com/OpsHelmInc/cloudquery/resources/services/codepipeline"
	"github.com/OpsHelmInc/cloudquery/resources/services/cognito"
	"github.com/OpsHelmInc/cloudquery/resources/services/config"
	"github.com/OpsHelmInc/cloudquery/resources/services/dax"
	"github.com/OpsHelmInc/cloudquery/resources/services/detective"
	"github.com/OpsHelmInc/cloudquery/resources/services/directconnect"
	"github.com/OpsHelmInc/cloudquery/resources/services/dms"
	"github.com/OpsHelmInc/cloudquery/resources/services/docdb"
	"github.com/OpsHelmInc/cloudquery/resources/services/dynamodb"
	"github.com/OpsHelmInc/cloudquery/resources/services/dynamodbstreams"
	"github.com/OpsHelmInc/cloudquery/resources/services/ec2"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecr"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecrpublic"
	"github.com/OpsHelmInc/cloudquery/resources/services/ecs"
	"github.com/OpsHelmInc/cloudquery/resources/services/efs"
	"github.com/OpsHelmInc/cloudquery/resources/services/eks"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticache"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticbeanstalk"
	"github.com/OpsHelmInc/cloudquery/resources/services/elasticsearch"
	"github.com/OpsHelmInc/cloudquery/resources/services/elastictranscoder"
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
	"github.com/OpsHelmInc/cloudquery/resources/services/networkfirewall"
	"github.com/OpsHelmInc/cloudquery/resources/services/networkmanager"
	"github.com/OpsHelmInc/cloudquery/resources/services/organizations"
	"github.com/OpsHelmInc/cloudquery/resources/services/qldb"
	"github.com/OpsHelmInc/cloudquery/resources/services/quicksight"
	"github.com/OpsHelmInc/cloudquery/resources/services/ram"
	"github.com/OpsHelmInc/cloudquery/resources/services/rds"
	"github.com/OpsHelmInc/cloudquery/resources/services/redshift"
	"github.com/OpsHelmInc/cloudquery/resources/services/resiliencehub"
	"github.com/OpsHelmInc/cloudquery/resources/services/resourcegroups"
	"github.com/OpsHelmInc/cloudquery/resources/services/route53"
	"github.com/OpsHelmInc/cloudquery/resources/services/route53recoverycontrolconfig"
	"github.com/OpsHelmInc/cloudquery/resources/services/route53recoveryreadiness"
	"github.com/OpsHelmInc/cloudquery/resources/services/route53resolver"
	"github.com/OpsHelmInc/cloudquery/resources/services/s3"
	"github.com/OpsHelmInc/cloudquery/resources/services/sagemaker"
	"github.com/OpsHelmInc/cloudquery/resources/services/savingsplans"
	"github.com/OpsHelmInc/cloudquery/resources/services/scheduler"
	"github.com/OpsHelmInc/cloudquery/resources/services/secretsmanager"
	"github.com/OpsHelmInc/cloudquery/resources/services/securityhub"
	"github.com/OpsHelmInc/cloudquery/resources/services/servicecatalog"
	"github.com/OpsHelmInc/cloudquery/resources/services/servicediscovery"
	"github.com/OpsHelmInc/cloudquery/resources/services/ses"
	"github.com/OpsHelmInc/cloudquery/resources/services/shield"
	"github.com/OpsHelmInc/cloudquery/resources/services/signer"
	"github.com/OpsHelmInc/cloudquery/resources/services/sns"
	"github.com/OpsHelmInc/cloudquery/resources/services/sqs"
	"github.com/OpsHelmInc/cloudquery/resources/services/ssm"
	"github.com/OpsHelmInc/cloudquery/resources/services/stepfunctions"
	"github.com/OpsHelmInc/cloudquery/resources/services/timestream"
	"github.com/OpsHelmInc/cloudquery/resources/services/transfer"
	"github.com/OpsHelmInc/cloudquery/resources/services/waf"
	"github.com/OpsHelmInc/cloudquery/resources/services/wafregional"
	"github.com/OpsHelmInc/cloudquery/resources/services/wafv2"
	"github.com/OpsHelmInc/cloudquery/resources/services/wellarchitected"
	"github.com/OpsHelmInc/cloudquery/resources/services/workspaces"
	"github.com/OpsHelmInc/cloudquery/resources/services/xray"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func GetTables() schema.Tables {
	t := []*schema.Table{
		accessanalyzer.Analyzers(),
		// account.AlternateContacts(),
		// account.Contacts(),
		acm.Certificates(),
		acmpca.CertificateAuthorities(),
		amp.Workspaces(),
		amplify.Apps(),
		apigateway.ApiKeys(),
		apigateway.ClientCertificates(),
		apigateway.DomainNames(),
		apigateway.RestApis(),
		apigateway.UsagePlans(),
		apigateway.VpcLinks(),
		apigatewayv2.Apis(),
		apigatewayv2.DomainNames(),
		apigatewayv2.VpcLinks(),
		appconfig.Applications(),
		appconfig.DeploymentStrategies(),
		appflow.Flows(),
		appmesh.Meshes(),
		applicationautoscaling.Policies(),
		applicationautoscaling.ScalableTargets(),
		// applicationautoscaling.ScalingActivities(),
		applicationautoscaling.ScheduledActions(),
		apprunner.AutoScalingConfigurations(),
		apprunner.Connections(),
		apprunner.ObservabilityConfigurations(),
		apprunner.Services(),
		apprunner.VpcConnectors(),
		apprunner.VpcIngressConnections(),
		appstream.AppBlocks(),
		appstream.Applications(),
		// appstream.DirectoryConfigs(),
		appstream.Fleets(),
		appstream.ImageBuilders(),
		appstream.Images(),
		appstream.Stacks(),
		// appstream.UsageReportSubscriptions(),
		appstream.Users(),
		appsync.GraphqlApis(),
		athena.DataCatalogs(),
		athena.WorkGroups(),
		auditmanager.Assessments(),
		autoscaling.Groups(),
		autoscaling.LaunchConfigurations(),
		autoscaling.ScheduledActions(),
		// autoscalingplans.Plans(),
		// backup.GlobalSettings(),
		// backup.Jobs(),
		backup.Plans(),
		backup.ProtectedResources(),
		// backup.RegionSettings(),
		backup.ReportPlans(),
		backup.Vaults(),
		batch.ComputeEnvironments(),
		batch.JobDefinitions(),
		batch.JobQueues(),
		cloudformation.Stacks(),
		cloudformation.StackSets(),
		cloudfront.CachePolicies(),
		cloudfront.Distributions(),
		cloudfront.Functions(),
		// cloudfront.OriginAccessIdentities(),
		// cloudfront.OriginRequestPolicies(),
		// cloudfront.ResponseHeaderPolicies(),
		cloudhsmv2.Backups(),
		cloudhsmv2.Clusters(),
		cloudtrail.Channels(),
		// cloudtrail.Events(),
		// cloudtrail.Imports(),
		cloudtrail.Trails(),
		cloudwatch.Alarms(),
		cloudwatchlogs.LogGroups(),
		// cloudwatchlogs.MetricFilters(),
		// cloudwatchlogs.ResourcePolicies(),
		codeartifact.Domains(),
		codeartifact.Repositories(),
		codebuild.Projects(),
		codebuild.SourceCredentials(),
		codecommit.Repositories(),
		codepipeline.Pipelines(),
		codepipeline.Webhooks(),
		cognito.IdentityPools(),
		cognito.UserPools(),
		// computeoptimizer.AutoscalingGroupsRecommendations(),
		// computeoptimizer.EbsVolumeRecommendations(),
		// computeoptimizer.Ec2InstanceRecommendations(),
		// computeoptimizer.EcsServiceRecommendations(),
		// computeoptimizer.EnrollmentStatuses(),
		// computeoptimizer.LambdaFunctionsRecommendations(),
		config.ConfigRules(),
		config.ConfigurationAggregators(),
		config.ConfigurationRecorders(),
		config.ConformancePacks(),
		// config.DeliveryChannels(),
		// config.RetentionConfigurations(),
		// costexplorer.ThirtyDayCost(),
		// costexplorer.ThirtyDayCostForecast(),
		dax.Clusters(),
		detective.Graphs(),
		directconnect.Connections(),
		directconnect.Gateways(),
		directconnect.Lags(),
		// directconnect.Locations(),
		// directconnect.VirtualGateways(),
		directconnect.VirtualInterfaces(),
		dms.ReplicationInstances(),
		docdb.Certificates(),
		docdb.ClusterParameterGroups(),
		docdb.Clusters(),
		// docdb.EngineVersions(),
		// docdb.EventCategories(),
		// docdb.Events(),
		docdb.EventSubscriptions(),
		docdb.GlobalClusters(),
		// docdb.PendingMaintenanceActions(),
		docdb.SubnetGroups(),
		dynamodb.Backups(),
		dynamodb.Exports(),
		dynamodb.GlobalTables(),
		dynamodb.Tables(),
		dynamodbstreams.Streams(),
		// ec2.AccountAttributes(),
		// ec2.AvailabilityZones(),
		// ec2.ByoipCidrs(),
		ec2.CapacityReservations(),
		ec2.CustomerGateways(),
		// ec2.DHCPOptions(),
		ec2.EbsSnapshots(),
		ec2.EbsVolumes(),
		// ec2.EbsVolumesStatuses(),
		ec2.EgressOnlyInternetGateways(),
		// ec2.Eips(),
		ec2.FlowLogs(),
		ec2.Hosts(),
		ec2.Images(),
		ec2.InstanceConnectEndpoints(),
		ec2.Instances(),
		// ec2.InstanceStatuses(),
		// ec2.InstanceTypes(),
		ec2.InternetGateways(),
		ec2.KeyPairs(),
		ec2.LaunchTemplates(),
		ec2.ManagedPrefixLists(),
		ec2.NatGateways(),
		ec2.NetworkAcls(),
		ec2.NetworkInterfaces(),
		// ec2.RegionalConfigs(),
		// ec2.Regions(),
		ec2.ReservedInstances(),
		ec2.RouteTables(),
		ec2.SecurityGroups(),
		// ec2.SpotFleetRequests(),
		// ec2.SpotInstanceRequests(),
		ec2.Subnets(),
		ec2.TransitGateways(),
		// ec2.VpcEndpointConnections(),
		ec2.VpcEndpoints(),
		// ec2.VpcEndpointServiceConfigurations(),
		ec2.VpcEndpointServices(),
		ec2.VpcPeeringConnections(),
		ec2.Vpcs(),
		ec2.VpnConnections(),
		ec2.VpnGateways(),
		// ecr.PullThroughCacheRules(),
		// ecr.Registries(),
		// ecr.RegistryPolicies(),
		ecr.Repositories(),
		ecrpublic.Repositories(),
		ecs.Clusters(),
		ecs.TaskDefinitions(),
		efs.AccessPoints(),
		efs.Filesystems(),
		eks.Clusters(),
		elasticache.Clusters(),
		// elasticache.EngineVersions(),
		// elasticache.Events(),
		elasticache.GlobalReplicationGroups(),
		elasticache.ParameterGroups(),
		elasticache.ReplicationGroups(),
		elasticache.ReservedCacheNodes(),
		// elasticache.ReservedCacheNodesOfferings(),
		elasticache.ServiceUpdates(),
		elasticache.Snapshots(),
		elasticache.SubnetGroups(),
		// elasticache.UpdateActions(),
		elasticache.UserGroups(),
		elasticache.Users(),
		elasticbeanstalk.Applications(),
		elasticbeanstalk.ApplicationVersions(),
		elasticbeanstalk.Environments(),
		elasticsearch.Domains(),
		// elasticsearch.Packages(),
		// elasticsearch.Versions(),
		// elasticsearch.VpcEndpoints(),
		elastictranscoder.Pipelines(),
		elastictranscoder.Presets(),
		elbv1.LoadBalancers(),
		elbv2.LoadBalancers(),
		elbv2.TargetGroups(),
		// emr.BlockPublicAccessConfigs(),
		emr.Clusters(),
		// emr.ReleaseLabels(),
		// emr.SecurityConfigurations(),
		emr.Studios(),
		eventbridge.ApiDestinations(),
		eventbridge.Archives(),
		eventbridge.Connections(),
		eventbridge.Endpoints(),
		eventbridge.EventBuses(),
		eventbridge.EventSources(),
		eventbridge.Replays(),
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
		// fsx.Backups(),
		fsx.DataRepositoryAssociations(),
		fsx.DataRepositoryTasks(),
		fsx.FileCaches(),
		fsx.FileSystems(),
		fsx.Snapshots(),
		fsx.StorageVirtualMachines(),
		fsx.Volumes(),
		// glacier.DataRetrievalPolicies(),
		glacier.Vaults(),
		// glue.Classifiers(),
		glue.Connections(),
		glue.Crawlers(),
		glue.Databases(),
		// glue.DatacatalogEncryptionSettings(),
		glue.DevEndpoints(),
		glue.Jobs(),
		glue.MlTransforms(),
		glue.Registries(),
		// glue.SecurityConfigurations(),
		glue.Triggers(),
		glue.Workflows(),
		guardduty.Detectors(),
		// iam.Accounts(),
		// iam.AccountAuthorizationDetails(),
		iam.CredentialReports(),
		iam.Groups(),
		iam.InstanceProfiles(),
		iam.OpenidConnectIdentityProviders(),
		// iam.PasswordPolicies(),
		iam.Policies(),
		iam.Roles(),
		iam.SamlIdentityProviders(),
		iam.ServerCertificates(),
		iam.Users(),
		// iam.VirtualMfaDevices(),
		identitystore.Groups(),
		identitystore.Users(),
		inspector.Findings(),
		// inspector2.CoveredResources(),
		inspector2.Findings(),
		iot.BillingGroups(),
		iot.CaCertificates(),
		iot.Certificates(),
		iot.Jobs(),
		iot.Policies(),
		iot.SecurityProfiles(),
		iot.Streams(),
		iot.ThingGroups(),
		iot.Things(),
		iot.ThingTypes(),
		iot.TopicRules(),
		kafka.Clusters(),
		kafka.Configurations(),
		kinesis.Streams(),
		kms.Aliases(),
		kms.Keys(),
		lambda.Functions(),
		lambda.Layers(),
		// lambda.Runtimes(),
		lightsail.Alarms(),
		lightsail.Buckets(),
		lightsail.Certificates(),
		lightsail.ContainerServices(),
		lightsail.Databases(),
		lightsail.DatabaseSnapshots(),
		lightsail.Disks(),
		lightsail.Distributions(),
		lightsail.Instances(),
		lightsail.InstanceSnapshots(),
		lightsail.LoadBalancers(),
		lightsail.StaticIps(),
		mq.Brokers(),
		mwaa.Environments(),
		neptune.ClusterParameterGroups(),
		neptune.Clusters(),
		neptune.ClusterSnapshots(),
		neptune.DbParameterGroups(),
		neptune.EventSubscriptions(),
		neptune.GlobalClusters(),
		neptune.Instances(),
		neptune.SubnetGroups(),
		networkfirewall.FirewallPolicies(),
		networkfirewall.Firewalls(),
		networkfirewall.RuleGroups(),
		networkfirewall.TLSInspectionConfigurations(),
		networkmanager.GlobalNetworks(),
		organizations.Accounts(),
		organizations.DelegatedAdministrators(),
		organizations.OrganizationalUnits(),
		organizations.Organizations(),
		organizations.Policies(),
		organizations.ResourcePolicies(),
		organizations.Roots(),
		qldb.Ledgers(),
		quicksight.Analyses(),
		quicksight.Dashboards(),
		quicksight.DataSets(),
		quicksight.DataSources(),
		quicksight.Folders(),
		quicksight.Groups(),
		quicksight.Templates(),
		quicksight.Users(),
		// ram.Principals(),
		ram.Resources(),
		// ram.ResourceShareAssociations(),
		ram.ResourceShareInvitations(),
		ram.ResourceShares(),
		// ram.ResourceTypes(),
		rds.Certificates(),
		rds.ClusterParameterGroups(),
		rds.Clusters(),
		rds.ClusterSnapshots(),
		rds.DbParameterGroups(),
		rds.DbProxies(),
		rds.DbSecurityGroups(),
		rds.DbSnapshots(),
		// rds.EngineVersions(),
		// rds.Events(),
		rds.EventSubscriptions(),
		rds.Instances(),
		rds.OptionGroups(),
		rds.ReservedInstances(),
		rds.SubnetGroups(),
		redshift.Clusters(),
		redshift.DataShares(),
		// redshift.Events(),
		redshift.EventSubscriptions(),
		redshift.SubnetGroups(),
		resiliencehub.Apps(),
		resiliencehub.ResiliencyPolicies(),
		resiliencehub.SuggestedResiliencyPolicies(),
		resourcegroups.ResourceGroups(),
		route53.DelegationSets(),
		// route53.Domains(),
		route53.HealthChecks(),
		route53.HostedZones(),
		// route53.Operations(),
		route53.TrafficPolicies(),
		route53recoverycontrolconfig.Clusters(),
		route53recoverycontrolconfig.ControlPanels(),
		route53recoveryreadiness.Cells(),
		route53recoveryreadiness.ReadinessChecks(),
		route53recoveryreadiness.RecoveryGroups(),
		route53recoveryreadiness.ResourceSets(),
		// route53resolver.FirewallConfigs(),
		route53resolver.FirewallDomainLists(),
		route53resolver.FirewallRuleGroupAssociations(),
		route53resolver.FirewallRuleGroups(),
		route53resolver.ResolverEndpoints(),
		route53resolver.ResolverQueryLogConfigAssociations(),
		route53resolver.ResolverQueryLogConfigs(),
		route53resolver.ResolverRuleAssociations(),
		route53resolver.ResolverRules(),
		s3.AccessGrants(),
		s3.AccessGrantInstances(),
		s3.AccessPoints(),
		// s3.Accounts(),
		s3.Buckets(),
		s3.MultiRegionAccessPoints(),
		sagemaker.Apps(),
		sagemaker.EndpointConfigurations(),
		sagemaker.Models(),
		sagemaker.NotebookInstances(),
		sagemaker.TrainingJobs(),
		savingsplans.Plans(),
		scheduler.ScheduleGroups(),
		scheduler.Schedules(),
		secretsmanager.Secrets(),
		securityhub.EnabledStandards(),
		// securityhub.Findings(),
		securityhub.Hubs(),
		servicecatalog.Portfolios(),
		servicecatalog.Products(),
		servicecatalog.ProvisionedProducts(),
		servicediscovery.Namespaces(),
		servicediscovery.Services(),
		// servicequotas.Services(),
		// ses.ActiveReceiptRuleSets(),
		ses.ConfigurationSets(),
		// ses.ContactLists(),
		ses.CustomVerificationEmailTemplates(),
		ses.Identities(),
		ses.Templates(),
		// shield.Attacks(),
		shield.ProtectionGroups(),
		shield.Protections(),
		shield.Subscriptions(),
		signer.Profiles(),
		sns.Subscriptions(),
		sns.Topics(),
		sqs.Queues(),
		ssm.Associations(),
		ssm.ComplianceSummaryItems(),
		ssm.Documents(),
		ssm.Instances(),
		// ssm.Inventories(),
		// ssm.InventorySchemas(),
		// ssm.Parameters(),
		// ssm.PatchBaselines(),
		// ssm.Sessions(),
		// ssoadmin.Instances(),
		stepfunctions.Activities(),
		stepfunctions.StateMachines(),
		// support.Cases(),
		// support.Services(),
		// support.SeverityLevels(),
		// support.TrustedAdvisorChecks(),
		timestream.Databases(),
		transfer.Servers(),
		waf.RuleGroups(),
		waf.Rules(),
		// waf.SubscribedRuleGroups(),
		waf.WebAcls(),
		wafregional.RateBasedRules(),
		wafregional.RuleGroups(),
		wafregional.Rules(),
		wafregional.WebAcls(),
		wafv2.Ipsets(),
		// wafv2.ManagedRuleGroups(),
		wafv2.RegexPatternSets(),
		wafv2.RuleGroups(),
		wafv2.WebAcls(),
		wellarchitected.Lenses(),
		// wellarchitected.ShareInvitations(),
		wellarchitected.Workloads(),
		workspaces.Directories(),
		workspaces.Workspaces(),
		// xray.EncryptionConfigs(),
		xray.Groups(),
		// xray.ResourcePolicies(),
		xray.SamplingRules(),
	}
	if err := transformers.TransformTables(t); err != nil {
		panic(err)
	}
	transformTitles := titleTransformer()
	for _, table := range t {
		schema.AddCqIDs(table)
		transformTitles(table)
		if err := validateTagsIsJSON(table); err != nil {
			panic(err)
		}
	}
	return t
}

func validateTagsIsJSON(table *schema.Table) error {
	for _, col := range table.Columns {
		if col.Name == "tags" && col.Type != types.ExtensionTypes.JSON {
			return fmt.Errorf("column %s in table %s must be of type %s", col.Name, table.Name, types.ExtensionTypes.JSON)
		}
	}
	for _, rel := range table.Relations {
		if err := validateTagsIsJSON(rel); err != nil {
			return err
		}
	}

	return nil
}
