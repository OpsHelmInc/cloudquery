package elbv2

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func TargetGroups() *schema.Table {
	tableName := "aws_elbv2_target_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html`,
		Resolver:    fetchTargetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.TargetGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveTargetGroupTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TargetGroupArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			targetGroupTargetHealthDescriptions(),
		},
	}
}

func fetchTargetGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceElasticloadbalancingv2).Elasticloadbalancingv2
	paginator := elbv2.NewDescribeTargetGroupsPaginator(svc, &elbv2.DescribeTargetGroupsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *elbv2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TargetGroups
	}
	return nil
}

func resolveTargetGroupTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	region := cl.Region
	svc := cl.Services(client.AWSServiceElasticloadbalancingv2).Elasticloadbalancingv2
	targetGroup := resource.Item.(types.TargetGroup)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*targetGroup.TargetGroupArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, s := range tagsOutput.TagDescriptions[0].Tags {
		tags[*s.Key] = s.Value
	}
	return resource.Set(c.Name, tags)
}
