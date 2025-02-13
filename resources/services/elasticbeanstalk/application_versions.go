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
)

func ApplicationVersions() *schema.Table {
	tableName := "aws_elasticbeanstalk_application_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticbeanstalk/latest/api/API_ApplicationVersionDescription.html`,
		Resolver:    fetchElasticbeanstalkApplicationVersions,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticbeanstalk"),
		Transform:   transformers.TransformWithStruct(&types.ApplicationVersionDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ApplicationVersionArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchElasticbeanstalkApplicationVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config elasticbeanstalk.DescribeApplicationVersionsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticbeanstalk).Elasticbeanstalk
	// No paginator available
	for {
		output, err := svc.DescribeApplicationVersions(ctx, &config, func(options *elasticbeanstalk.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- output.ApplicationVersions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
