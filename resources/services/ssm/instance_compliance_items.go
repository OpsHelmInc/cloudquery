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

func instanceComplianceItems() *schema.Table {
	tableName := "aws_ssm_instance_compliance_items"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html`,
		Resolver:    fetchSsmInstanceComplianceItems,
		Transform:   transformers.TransformWithStruct(&types.ComplianceItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSsmInstanceComplianceItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance := parent.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm

	input := ssm.ListComplianceItemsInput{
		ResourceIds: []string{*instance.InstanceId},
	}
	paginator := ssm.NewListComplianceItemsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ComplianceItems
	}
	return nil
}
