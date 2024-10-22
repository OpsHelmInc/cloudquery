package securityhub

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func EnabledStandards() *schema.Table {
	tableName := "aws_securityhub_enabled_standards"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_GetEnabledStandards.html`,
		Resolver:    fetchEnabledStandards,
		Transform: transformers.TransformWithStruct(&types.StandardsSubscription{},
			transformers.WithTypeTransformer(client.TimestampTypeTransformer),
			transformers.WithResolverTransformer(client.TimestampResolverTransformer),
			transformers.WithPrimaryKeyComponents("StandardsArn", "StandardsSubscriptionArn"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "securityhub"),
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("StandardsArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchEnabledStandards(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSecurityhub).Securityhub
	config := securityhub.GetEnabledStandardsInput{MaxResults: aws.Int32(100)}
	p := securityhub.NewGetEnabledStandardsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(o *securityhub.Options) { o.Region = cl.Region })
		if err != nil {
			return err
		}
		res <- response.StandardsSubscriptions
	}
	return nil
}
