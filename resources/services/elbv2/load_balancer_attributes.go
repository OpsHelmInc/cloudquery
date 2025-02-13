package elbv2

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func loadBalancerAttributes() *schema.Table {
	tableName := "aws_elbv2_load_balancer_attributes"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancerAttribute.html`,
		Resolver:    fetchLoadBalancerAttributes,
		Transform:   transformers.TransformWithStruct(&types.LoadBalancerAttribute{}, transformers.WithPrimaryKeyComponents("Key")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "load_balancer_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchLoadBalancerAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	lb := parent.Item.(types.LoadBalancer)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticloadbalancingv2).Elasticloadbalancingv2
	result, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: lb.LoadBalancerArn}, func(options *elbv2.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- result.Attributes
	return nil
}
