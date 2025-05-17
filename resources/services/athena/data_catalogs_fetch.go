package athena

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchAthenaDataCatalogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDataCatalogsInput{}
	for {
		response, err := svc.ListDataCatalogs(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.DataCatalogsSummary
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getDataCatalog(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	catalogSummary := resource.Item.(types.DataCatalogSummary)
	dc, err := svc.GetDataCatalog(ctx, &athena.GetDataCatalogInput{
		Name: catalogSummary.CatalogName,
	})
	if err != nil {
		// retrieving of default data catalog (AwsDataCatalog) returns "not found error" (with statuscode 400: InvalidRequestException:...) but it exists and its
		// relations can be fetched by its name
		if client.IsAWSError(err, "InvalidRequestException") && *catalogSummary.CatalogName == "AwsDataCatalog" {
			resource.Item = &ohaws.AthenaDataCatalog{DataCatalog: types.DataCatalog{Name: catalogSummary.CatalogName, Type: catalogSummary.Type}}
			return nil
		}
		return err
	}

	catalog := ohaws.AthenaDataCatalog{
		DataCatalog: *dc.DataCatalog,
	}

	tags, err := getDataCatalogTags(ctx, c, svc, createDataCatalogArn(c, *catalog.Name))
	if err != nil {
		return err
	}
	catalog.Tags = tags

	resource.Item = &catalog

	return nil
}

func resolveAthenaDataCatalogArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(*ohaws.AthenaDataCatalog)
	return resource.Set(c.Name, createDataCatalogArn(cl, *dc.Name))
}

func getDataCatalogTags(ctx context.Context, c *client.Client, svc services.AthenaClient, catalogArn string) ([]types.Tag, error) {
	params := athena.ListTagsForResourceInput{ResourceARN: &catalogArn}

	var tags []types.Tag
	tagsPaginator := athena.NewListTagsForResourcePaginator(svc, &params)
	for tagsPaginator.HasMorePages() {
		result, err := tagsPaginator.NextPage(ctx)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil, nil
			}
			return nil, err
		}

		tags = append(tags, result.Tags...)
	}
	return tags, nil
}

func fetchAthenaDataCatalogDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListDatabasesInput{
		CatalogName: parent.Item.(*ohaws.AthenaDataCatalog).Name,
	}
	for {
		response, err := svc.ListDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.DatabaseList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func fetchAthenaDataCatalogDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Athena
	input := athena.ListTableMetadataInput{
		CatalogName:  parent.Parent.Item.(*ohaws.AthenaDataCatalog).Name,
		DatabaseName: parent.Item.(types.Database).Name,
	}
	for {
		response, err := svc.ListTableMetadata(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.TableMetadataList

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func createDataCatalogArn(cl *client.Client, catalogName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Athena),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("datacatalog/%s", catalogName),
	}.String()
}
