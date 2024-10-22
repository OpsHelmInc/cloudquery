package ssm

import (
	"context"
	"fmt"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Parameters() *schema.Table {
	tableName := "aws_ssm_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html`,
		Resolver:    fetchSsmParameters,
		Transform:   transformers.TransformWithStruct(&types.ParameterMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "name",
				Type:                arrow.BinaryTypes.String,
				Description:         `The parameter name`,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveParameterTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveSsmParameterArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm
	paginator := ssm.NewDescribeParametersPaginator(svc, &ssm.DescribeParametersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Parameters
	}
	return nil
}

func resolveParameterTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsm).Ssm
	pm := resource.Item.(types.ParameterMetadata)
	resp, err := svc.ListTagsForResource(ctx, &ssm.ListTagsForResourceInput{
		ResourceId:   pm.Name,
		ResourceType: types.ResourceTypeForTaggingParameter,
	}, func(options *ssm.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	tags := make(map[string]string)
	client.TagsIntoMap(resp.TagList, tags)
	return resource.Set(c.Name, tags)
}

func resolveSsmParameterArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	p := resource.Item.(types.ParameterMetadata)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ssm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("parameter/%s", aws.ToString(p.Name)),
	}
	return resource.Set(c.Name, a.String())
}
