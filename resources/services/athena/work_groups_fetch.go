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

func fetchAthenaWorkGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	input := athena.ListWorkGroupsInput{}
	for {
		response, err := svc.ListWorkGroups(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.WorkGroups
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}

	return nil
}

func getWorkGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena

	wgs := resource.Item.(types.WorkGroupSummary)
	r, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
		WorkGroup: wgs.Name,
	})
	if err != nil {
		return err
	}

	wg := ohaws.AthenaWorkGroup{
		WorkGroup: *r.WorkGroup,
	}

	tags, err := getAthenaWorkGroupTags(ctx, c, svc, createWorkGroupArn(c, *wg.Name))
	if err != nil {
		return err
	}
	wg.Tags = tags

	resource.Item = &wg
	return nil
}

func resolveAthenaWorkGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	dc := resource.Item.(*ohaws.AthenaWorkGroup)
	return resource.Set(c.Name, createWorkGroupArn(cl, *dc.Name))
}

func getAthenaWorkGroupTags(ctx context.Context, c *client.Client, svc services.AthenaClient, workgroupArn string) ([]types.Tag, error) {
	params := athena.ListTagsForResourceInput{ResourceARN: &workgroupArn}

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

func fetchAthenaWorkGroupPreparedStatements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(*ohaws.AthenaWorkGroup)
	input := athena.ListPreparedStatementsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListPreparedStatements(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.PreparedStatements
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getWorkGroupPreparedStatement(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := resource.Parent.Item.(*ohaws.AthenaWorkGroup)

	d := resource.Item.(types.PreparedStatementSummary)
	dc, err := svc.GetPreparedStatement(ctx, &athena.GetPreparedStatementInput{
		WorkGroup:     wg.Name,
		StatementName: d.StatementName,
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.PreparedStatement
	return nil
}

func fetchAthenaWorkGroupQueryExecutions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(*ohaws.AthenaWorkGroup)
	input := athena.ListQueryExecutionsInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListQueryExecutions(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.QueryExecutionIds
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getWorkGroupQueryExecution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena

	d := resource.Item.(string)
	dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
		QueryExecutionId: aws.String(d),
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.QueryExecution
	return nil
}

func fetchAthenaWorkGroupNamedQueries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena
	wg := parent.Item.(*ohaws.AthenaWorkGroup)
	input := athena.ListNamedQueriesInput{WorkGroup: wg.Name}
	for {
		response, err := svc.ListNamedQueries(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.NamedQueryIds
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

func getWorkGroupNamedQuery(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Athena

	d := resource.Item.(string)
	dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
		NamedQueryId: aws.String(d),
	})
	if err != nil {
		return err
	}
	resource.Item = *dc.NamedQuery
	return nil
}

func createWorkGroupArn(cl *client.Client, groupName string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.Athena),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("workgroup/%s", groupName),
	}.String()
}
