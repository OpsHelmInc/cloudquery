package rds

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func DbParameterGroups() *schema.Table {
	tableName := "aws_rds_db_parameter_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBParameterGroup.html`,
		Resolver:    fetchRdsDbParameterGroups,
		Transform:   transformers.TransformWithStruct(&types.DBParameterGroup{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBParameterGroupArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRDSTags("DBParameterGroupArn"),
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			dbParameterGroupDbParameters(),
		},
	}
}

func fetchRdsDbParameterGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	var input rds.DescribeDBParameterGroupsInput
	paginator := rds.NewDescribeDBParameterGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBParameterGroups
	}
	return nil
}

func fetchRdsDbParameterGroupDbParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	g := parent.Item.(types.DBParameterGroup)
	input := rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName}
	paginator := rds.NewDescribeDBParametersPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "DBParameterGroupNotFound") {
				cl.Logger().Warn().Err(err).Msg("received DBParameterGroupNotFound on DescribeDBParameters")
				return nil
			}
			return err
		}
		res <- page.Parameters
	}
	return nil
}
