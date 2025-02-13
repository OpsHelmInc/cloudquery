package eks

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func addOns() *schema.Table {
	tableName := "aws_eks_cluster_addons"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/eks/latest/APIReference/API_Addon.html`,
		Resolver:            fetchAddOns,
		PreResourceResolver: getAddOn,
		Transform:           transformers.TransformWithStruct(&types.Addon{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("AddonArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "cluster_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchAddOns(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	cluster := resource.Item.(*types.Cluster)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEks).Eks
	paginator := eks.NewListAddonsPaginator(svc, &eks.ListAddonsInput{ClusterName: cluster.Name})
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx, func(options *eks.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Addons
	}
	return nil
}

func getAddOn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEks).Eks
	name := resource.Item.(string)
	cluster := resource.Parent.Item.(*types.Cluster)
	output, err := svc.DescribeAddon(
		ctx, &eks.DescribeAddonInput{
			ClusterName: cluster.Name,
			AddonName:   &name}, func(options *eks.Options) {
			options.Region = cl.Region
		})
	if err != nil {
		return err
	}
	resource.Item = output.Addon
	return nil
}
