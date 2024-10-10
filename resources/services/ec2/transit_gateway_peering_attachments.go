package ec2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func transitGatewayPeeringAttachments() *schema.Table {
	tableName := "aws_ec2_transit_gateway_peering_attachments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachment.html`,
		Resolver:    fetchEc2TransitGatewayPeeringAttachments,
		Transform:   transformers.TransformWithStruct(&types.TransitGatewayPeeringAttachment{}, transformers.WithResolverTransformer(client.TagsResolverTransformer)),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "transit_gateway_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TransitGatewayAttachmentId"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchEc2TransitGatewayPeeringAttachments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(types.TransitGateway)

	config := ec2.DescribeTransitGatewayPeeringAttachmentsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("transit-gateway-id"),
				Values: []string{*r.TransitGatewayId},
			},
		},
	}

	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	paginator := ec2.NewDescribeTransitGatewayPeeringAttachmentsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TransitGatewayPeeringAttachments
	}
	return nil
}
