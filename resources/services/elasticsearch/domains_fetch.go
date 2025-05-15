package elasticsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchElasticsearchDomains(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticsearchservice
	out, err := svc.ListDomainNames(ctx, &elasticsearchservice.ListDomainNamesInput{})
	if err != nil {
		return err
	}

	res <- out.DomainNames
	return nil
}

func getDomain(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Elasticsearchservice

	info := resource.Item.(types.DomainInfo)

	domainOutput, err := svc.DescribeElasticsearchDomain(ctx, &elasticsearchservice.DescribeElasticsearchDomainInput{DomainName: info.DomainName})
	if err != nil {
		return nil
	}

	domain := ohaws.ElasticsearchDomain{
		ElasticsearchDomainStatus: *domainOutput.DomainStatus,
	}

	tags, err := getElasticsearchDomainTags(ctx, svc, *domain.ARN)
	if err != nil {
		return err
	}
	domain.Tags = tags

	resource.Item = &domain
	return nil
}

func getElasticsearchDomainTags(ctx context.Context, svc services.ElasticsearchserviceClient, domainARN string) ([]types.Tag, error) {
	tagsOutput, err := svc.ListTags(ctx, &elasticsearchservice.ListTagsInput{
		ARN: &domainARN,
	})
	if err != nil {
		return nil, err
	}
	if len(tagsOutput.TagList) == 0 {
		return nil, nil
	}

	return tagsOutput.TagList, nil
}
