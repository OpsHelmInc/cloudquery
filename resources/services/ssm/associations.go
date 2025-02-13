package ssm

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Associations() *schema.Table {
	tableName := "aws_ssm_associations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html`,
		Resolver:    fetchSsmAssociations,
		Transform:   transformers.TransformWithStruct(&types.Association{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "association_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("AssociationId"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSsmAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm

	paginator := ssm.NewListAssociationsPaginator(svc, nil)
	for paginator.HasMorePages() {
		v, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- v.Associations
	}
	return nil
}
