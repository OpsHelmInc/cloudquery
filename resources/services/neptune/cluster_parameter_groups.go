package neptune

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func ClusterParameterGroups() *schema.Table {
	tableName := "aws_neptune_cluster_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/neptune/latest/userguide/api-parameters.html#DescribeDBParameters`,
		Resolver:    fetchNeptuneClusterParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBClusterParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "neptune"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBClusterParameterGroupArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveNeptuneClusterParameterGroupTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			clusterParameterGroupParameters(),
		},
	}
}

func fetchNeptuneClusterParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNeptune).Neptune
	input := neptune.DescribeDBClusterParameterGroupsInput{
		Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
	}
	paginator := neptune.NewDescribeDBClusterParameterGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *neptune.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBClusterParameterGroups
	}
	return nil
}
