package elasticbeanstalk

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Environments() *schema.Table {
	tableName := "aws_elasticbeanstalk_environments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/APIReference/API_EnvironmentDescription.html`,
		Resolver:    fetchElasticbeanstalkEnvironments,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.EnvironmentDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("EnvironmentArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveElasticbeanstalkEnvironmentTags,
			},
			{
				Name:     "id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("EnvironmentId"),
			},
			{
				Name:     "listeners",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveElasticbeanstalkEnvironmentListeners,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			configurationSettings(),
			configurationOptions(),
		},
	}
}

func fetchElasticbeanstalkEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elasticbeanstalk.DescribeEnvironmentsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticbeanstalk).Elasticbeanstalk
	// No paginator available
	for {
		response, err := svc.DescribeEnvironments(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Environments
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveElasticbeanstalkEnvironmentListeners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EnvironmentDescription)
	if p.Resources == nil || p.Resources.LoadBalancer == nil {
		return nil
	}
	listeners := make(map[int32]*string, len(p.Resources.LoadBalancer.Listeners))
	for _, l := range p.Resources.LoadBalancer.Listeners {
		listeners[l.Port] = l.Protocol
	}
	return resource.Set(c.Name, listeners)
}

func resolveElasticbeanstalkEnvironmentTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.EnvironmentDescription)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticbeanstalk).Elasticbeanstalk
	tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
		ResourceArn: p.EnvironmentArn,
	}, func(o *elasticbeanstalk.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		// It takes a few minutes for an environment to be terminated
		// This ensures we don't error while trying to fetch related resources for a terminated environment
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(tagsOutput.ResourceTags))
}
