package apprunner

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func AutoScalingConfigurations() *schema.Table {
	tableName := "aws_apprunner_auto_scaling_configurations"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html`,
		Resolver:            fetchApprunnerAutoScalingConfigurations,
		PreResourceResolver: getAutoScalingConfiguration,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           transformers.TransformWithStruct(&types.AutoScalingConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("AutoScalingConfigurationArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveApprunnerTags("AutoScalingConfigurationArn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchApprunnerAutoScalingConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config apprunner.ListAutoScalingConfigurationsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApprunner).Apprunner
	paginator := apprunner.NewListAutoScalingConfigurationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *apprunner.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.AutoScalingConfigurationSummaryList
	}
	return nil
}

func getAutoScalingConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceApprunner).Apprunner
	asConfig := resource.Item.(types.AutoScalingConfigurationSummary)

	describeTaskDefinitionOutput, err := svc.DescribeAutoScalingConfiguration(ctx, &apprunner.DescribeAutoScalingConfigurationInput{AutoScalingConfigurationArn: asConfig.AutoScalingConfigurationArn}, func(options *apprunner.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = describeTaskDefinitionOutput.AutoScalingConfiguration
	return nil
}
