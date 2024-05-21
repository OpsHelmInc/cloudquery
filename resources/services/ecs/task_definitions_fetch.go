package ecs

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListTaskDefinitionsInput
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	paginator := ecs.NewListTaskDefinitionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ecs.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.TaskDefinitionArns
	}
	return nil
}

func getTaskDefinition(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ecs
	taskArn := resource.Item.(string)

	describeTaskDefinitionOutput, err := svc.DescribeTaskDefinition(ctx, &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskArn),
		Include:        []types.TaskDefinitionField{types.TaskDefinitionFieldTags},
	}, func(options *ecs.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	if describeTaskDefinitionOutput.TaskDefinition == nil {
		return errors.New("nil TaskDefinition encountered")
	}
	resource.Item = ohaws.TaskDefinition{
		TaskDefinition: describeTaskDefinitionOutput.TaskDefinition,
		Tags:           describeTaskDefinitionOutput.Tags,
	}
	return nil
}

func resolveEcsTaskDefinitionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(ohaws.TaskDefinition)
	return resource.Set(c.Name, client.TagsToMap(r.Tags))
}
