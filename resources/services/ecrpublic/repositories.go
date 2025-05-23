package ecrpublic

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Repositories() *schema.Table {
	tableName := "aws_ecrpublic_repositories"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_Repository.html`,
		Resolver:    fetchEcrpublicRepositories,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "api.ecr-public"),
		Transform:   transformers.TransformWithStruct(&types.Repository{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("RepositoryArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRepositoryTags,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			repositoryImages(),
		},
	}
}

func fetchEcrpublicRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcrpublic).Ecrpublic
	paginator := ecrpublic.NewDescribeRepositoriesPaginator(svc, &ecrpublic.DescribeRepositoriesInput{
		MaxResults: aws.Int32(1000),
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecrpublic.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "UnsupportedCommandException") {
				return nil
			}
			return err
		}
		res <- page.Repositories
	}
	return nil
}

func resolveRepositoryTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcrpublic).Ecrpublic
	repo := resource.Item.(types.Repository)

	input := ecrpublic.ListTagsForResourceInput{
		ResourceArn: repo.RepositoryArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input, func(options *ecrpublic.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(output.Tags))
}
