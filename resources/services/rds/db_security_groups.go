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

func DbSecurityGroups() *schema.Table {
	tableName := "aws_rds_db_security_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBSecurityGroup.html`,
		Resolver:    fetchRdsDbSecurityGroups,
		Transform: transformers.TransformWithStruct(
			&types.DBSecurityGroup{},
			transformers.WithNameTransformer(client.CreateReplaceTransformer(map[string]string{"e_c2": "ec2"})),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "rds"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("DBSecurityGroupArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRDSTags("DBSecurityGroupArn"),
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchRdsDbSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRds).Rds
	var input rds.DescribeDBSecurityGroupsInput
	paginator := rds.NewDescribeDBSecurityGroupsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *rds.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DBSecurityGroups
	}
	return nil
}
