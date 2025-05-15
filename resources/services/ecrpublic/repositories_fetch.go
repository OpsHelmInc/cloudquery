package ecrpublic

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchEcrpublicRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	maxResults := int32(1000)
	config := ecrpublic.DescribeRepositoriesInput{
		MaxResults: &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ecrpublic
	for {
		output, err := svc.DescribeRepositories(ctx, &config)
		if err != nil {
			return err
		}

		repos := make([]*ohaws.ECRPublicRepository, len(output.Repositories))
		for idx, r := range output.Repositories {
			repos[idx] = &ohaws.ECRPublicRepository{
				Repository: r,
			}
		}
		res <- repos

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func getRepository(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecrpublic
	repo := resource.Item.(*ohaws.ECRPublicRepository)

	input := ecrpublic.ListTagsForResourceInput{
		ResourceArn: repo.RepositoryArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}

	repo.Tags = client.TagsToMap(output.Tags)
	return nil
}

func fetchEcrpublicRepositoryImages(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	maxResults := int32(1000)
	p := parent.Item.(*ohaws.ECRPublicRepository)
	config := ecrpublic.DescribeImagesInput{
		RepositoryName: p.RepositoryName,
		MaxResults:     &maxResults,
	}
	c := meta.(*client.Client)
	svc := c.Services().Ecrpublic
	for {
		output, err := svc.DescribeImages(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ImageDetails
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveImageArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ImageDetail)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ecr-public",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "repository_image/" + *item.RegistryId + "/" + *item.ImageDigest,
	}
	return resource.Set(c.Name, a.String())
}
