package organizations

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func ResourcePolicies() *schema.Table {
	tableName := "aws_organization_resource_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DescribeResourcePolicy.html`,
		Resolver:    fetchOrganizationsResourcePolicies,
		Transform:   transformers.TransformWithStruct(&types.ResourcePolicy{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ResourcePolicySummary.Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchOrganizationsResourcePolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	o, err := svc.DescribeResourcePolicy(ctx, &organizations.DescribeResourcePolicyInput{}, func(options *organizations.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- o.ResourcePolicy
	return nil
}
