package dms

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/dms/models"
)

func ReplicationInstances() *schema.Table {
	tableName := "aws_dms_replication_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/dms/latest/APIReference/API_ReplicationInstance.html`,
		Resolver:    fetchDmsReplicationInstances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "dms"),
		Transform:   transformers.TransformWithStruct(&models.ReplicationInstanceWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ReplicationInstanceArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchDmsReplicationInstances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceDatabasemigrationservice).Databasemigrationservice

	config := databasemigrationservice.DescribeReplicationInstancesInput{}
	paginator := databasemigrationservice.NewDescribeReplicationInstancesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *databasemigrationservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if len(page.ReplicationInstances) == 0 {
			continue
		}

		tags, err := getTags(ctx, svc, page.ReplicationInstances, "ReplicationInstanceArn", func(options *databasemigrationservice.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		wrappers := make([]*models.ReplicationInstanceWrapper, len(page.ReplicationInstances))
		for i := range page.ReplicationInstances {
			wrappers[i] = &models.ReplicationInstanceWrapper{
				ReplicationInstance: page.ReplicationInstances[i],
			}
		}

		if err := putTags(wrappers, tags, "ReplicationInstanceArn"); err != nil {
			return err
		}

		res <- wrappers
	}
	return nil
}
