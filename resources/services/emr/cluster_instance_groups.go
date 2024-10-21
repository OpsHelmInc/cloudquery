package emr

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func clusterInstanceGroups() *schema.Table {
	tableName := "aws_emr_cluster_instance_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroup.html`,
		Resolver:    fetchClusterInstanceGroups,
		Transform:   transformers.TransformWithStruct(&types.InstanceGroup{}, transformers.WithPrimaryKeyComponents("Id")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "cluster_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchClusterInstanceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cluster := parent.Item.(*types.Cluster)
	// instance fleets and instance groups are mutually exclusive
	if cluster.InstanceCollectionType != types.InstanceCollectionTypeInstanceGroup {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListInstanceGroupsPaginator(svc, &emr.ListInstanceGroupsInput{
		ClusterId: cluster.Id,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.InstanceGroups
	}
	return nil
}
