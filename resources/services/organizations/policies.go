package organizations

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Policies() *schema.Table {
	tableName := "aws_organizations_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html`,
		Resolver:    fetchOrganizationsPolicies,
		Transform:   transformers.TransformWithStruct(&types.PolicySummary{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			// This is needed as a PK because aws managed policies don't have an account_id in the ARN
			client.DefaultAccountIDColumn(true),
			{
				Name:     "content",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolvePolicyContent,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchOrganizationsPolicies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	for _, policyType := range types.PolicyType("").Values() {
		paginator := organizations.NewListPoliciesPaginator(svc, &organizations.ListPoliciesInput{
			Filter: policyType,
		})

		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Policies
		}
	}
	return nil
}

func resolvePolicyContent(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicySummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	resp, err := svc.DescribePolicy(ctx, &organizations.DescribePolicyInput{
		PolicyId: r.Id,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, resp.Policy.Content)
}
