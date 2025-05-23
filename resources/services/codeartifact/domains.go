package codeartifact

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Domains() *schema.Table {
	tableName := "aws_codeartifact_domains"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html
The 'request_account_id' and 'request_region' columns are added to show the account and region of where the request was made from.`,
		Resolver:            fetchDomains,
		PreResourceResolver: getDomain,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "codeartifact"),
		Transform:           transformers.TransformWithStruct(&types.DomainDescription{}),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCodeartifactTags("Arn"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{},
	}
}

func fetchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	paginator := codeartifact.NewListDomainsPaginator(svc, nil)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codeartifact.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Domains
	}
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	domain := resource.Item.(types.DomainSummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodeartifact).Codeartifact
	domainOutput, err := svc.DescribeDomain(ctx, &codeartifact.DescribeDomainInput{Domain: domain.Name}, func(options *codeartifact.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = domainOutput.Domain
	return nil
}
