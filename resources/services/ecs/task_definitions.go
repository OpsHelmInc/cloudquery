package ecs

import (
	"context"
	"errors"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func TaskDefinitions() *schema.Table {
	tableName := "aws_ecs_task_definitions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_TaskDefinition.html`,
		Resolver:            fetchEcsTaskDefinitions,
		PreResourceResolver: getTaskDefinition,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "ecs"),
		Transform:           transformers.TransformWithStruct(&ohaws.TaskDefinition{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("TaskDefinitionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveEcsTaskDefinitionTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchEcsTaskDefinitions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ecs.ListTaskDefinitionsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEcs).Ecs
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
	svc := cl.Services(client.AWSServiceEcs).Ecs
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
