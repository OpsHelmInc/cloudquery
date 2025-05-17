package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchGlueDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetDatabasesInput{}
	for {
		result, err := svc.GetDatabases(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.DatabaseList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func resolveGlueDatabaseArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, databaseARN(cl, aws.ToString(resource.Item.(*ohaws.GlueDatabase).Name)))
}

func getGlueDatabase(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	d := resource.Item.(types.Database)
	input := glue.GetTagsInput{
		ResourceArn: aws.String(databaseARN(cl, aws.ToString(d.Name))),
	}

	response, err := svc.GetTags(ctx, &input)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	resource.Item = &ohaws.GlueDatabase{
		Database: d,
		Tags:     response.Tags,
	}

	return nil
}

func fetchGlueDatabaseTables(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*ohaws.GlueDatabase)
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetTablesInput{
		DatabaseName: r.Name,
	}
	for {
		result, err := svc.GetTables(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.TableList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func getGlueTable(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	t := resource.Item.(types.Table)

	input := glue.GetTagsInput{
		ResourceArn: aws.String(tableARN(cl, aws.ToString(t.DatabaseName), aws.ToString(t.Name))),
	}

	response, err := svc.GetTags(ctx, &input)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	resource.Item = &ohaws.GlueTable{
		Table: t,
		Tags:  response.Tags,
	}

	return nil
}

func fetchGlueDatabaseTableIndexes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	d := parent.Parent.Item.(*ohaws.GlueDatabase)
	t := parent.Item.(*ohaws.GlueTable)
	input := glue.GetPartitionIndexesInput{DatabaseName: d.Name, CatalogId: d.CatalogId, TableName: t.Name}
	for {
		result, err := svc.GetPartitionIndexes(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.PartitionIndexDescriptorList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func databaseARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("database/%s", name),
	}.String()
}

func tableARN(cl *client.Client, dbName, tableName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("table/%s/%s", dbName, tableName),
	}.String()
}
