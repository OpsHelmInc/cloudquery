package workspaces

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Workspaces() *schema.Table {
	tableName := "aws_workspaces_workspaces"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/workspaces/latest/api/API_Workspace.html`,
		Resolver:    fetchWorkspacesWorkspaces,
		Transform:   transformers.TransformWithStruct(&types.Workspace{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "workspaces"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveWorkspaceArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchWorkspacesWorkspaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceWorkspaces).Workspaces
	input := workspaces.DescribeWorkspacesInput{}
	paginator := workspaces.NewDescribeWorkspacesPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *workspaces.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Workspaces
	}
	return nil
}

func resolveWorkspaceArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Workspace)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "workspaces",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "workspaces/" + *item.WorkspaceId,
	}
	return resource.Set(c.Name, a.String())
}
