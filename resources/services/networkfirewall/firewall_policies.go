package networkfirewall

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/networkfirewall/models"
)

func FirewallPolicies() *schema.Table {
	tableName := "aws_networkfirewall_firewall_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/network-firewall/latest/APIReference/API_FirewallPolicy.html`,
		Resolver:            fetchFirewallPolicies,
		PreResourceResolver: getFirewallPolicy,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "network-firewall"),
		Transform: transformers.TransformWithStruct(
			&models.FirewallPolicyWrapper{},
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("FirewallPolicyArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchFirewallPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input networkfirewall.ListFirewallPoliciesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkfirewall).Networkfirewall
	p := networkfirewall.NewListFirewallPoliciesPaginator(svc, &input)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *networkfirewall.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.FirewallPolicies
	}
	return nil
}

func getFirewallPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceNetworkfirewall).Networkfirewall
	metadata := resource.Item.(types.FirewallPolicyMetadata)

	policy, err := svc.DescribeFirewallPolicy(ctx, &networkfirewall.DescribeFirewallPolicyInput{
		FirewallPolicyArn: metadata.Arn,
	}, func(options *networkfirewall.Options) {
		options.Region = cl.Region
	})
	if err != nil && !cl.IsNotFoundError(err) {
		return err
	}

	resource.Item = &models.FirewallPolicyWrapper{
		FirewallPolicy:         policy.FirewallPolicy,
		FirewallPolicyResponse: policy.FirewallPolicyResponse,
	}
	return nil
}
