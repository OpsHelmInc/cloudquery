package appstream

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func stackEntitlements() *schema.Table {
	tableName := "aws_appstream_stack_entitlements"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_Entitlement.html`,
		Resolver:    fetchAppstreamStackEntitlements,
		Transform:   transformers.TransformWithStruct(&types.Entitlement{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "stack_name",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("StackName"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "name",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Name"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAppstreamStackEntitlements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input appstream.DescribeEntitlementsInput
	input.StackName = parent.Item.(types.Stack).Name
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAppstream).Appstream
	// No paginator available
	for {
		response, err := svc.DescribeEntitlements(ctx, &input, func(options *appstream.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Entitlements
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}
