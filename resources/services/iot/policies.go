package iot

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Policies() *schema.Table {
	tableName := "aws_iot_policies"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/iot/latest/apireference/API_Policy.html`,
		Resolver:            fetchIotPolicies,
		PreResourceResolver: getPolicy,
		Transform:           transformers.TransformWithStruct(&types.Policy{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iot"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: ResolveIotPolicyTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("PolicyArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIotPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIot).Iot
	input := iot.ListPoliciesInput{
		PageSize: aws.Int32(250),
	}
	paginator := iot.NewListPoliciesPaginator(svc, &input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iot.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Policies
	}
	return nil
}

func getPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIot).Iot

	output, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
		PolicyName: resource.Item.(types.Policy).PolicyName,
	}, func(options *iot.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

func ResolveIotPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*iot.GetPolicyOutput)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIot).Iot
	return resolveIotTags(ctx, meta, svc, resource, c, i.PolicyArn)
}
