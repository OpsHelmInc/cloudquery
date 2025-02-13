package elbv1

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func loadBalancerPolicies() *schema.Table {
	tableName := "aws_elbv1_load_balancer_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html`,
		Resolver:    fetchElbv1LoadBalancerPolicies,
		Transform:   transformers.TransformWithStruct(&types.PolicyDescription{}, transformers.WithPrimaryKeyComponents("PolicyName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "load_balancer_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "load_balancer_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("load_balancer_name"),
			},
			{
				Name:     "policy_attribute_descriptions",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveElbv1loadBalancerPolicyAttributeDescriptions,
			},
		},
	}
}

func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(ohaws.LoadBalancerV1)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticloadbalancing).Elasticloadbalancing
	response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName}, func(options *elbv1.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.PolicyDescriptions
	return nil
}

func resolveElbv1loadBalancerPolicyAttributeDescriptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyDescription)

	response := make(map[string]any, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}
