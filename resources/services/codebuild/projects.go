package codebuild

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Projects() *schema.Table {
	tableName := "aws_codebuild_projects"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/codebuild/latest/APIReference/API_Project.html`,
		Resolver:    fetchCodebuildProjects,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "codebuild"),
		Transform:   transformers.TransformWithStruct(&types.Project{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: schema.Tables{
			builds(),
		},
	}
}

func fetchCodebuildProjects(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCodebuild).Codebuild
	config := codebuild.ListProjectsInput{}
	paginator := codebuild.NewListProjectsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		if len(page.Projects) == 0 {
			continue
		}
		projectsOutput, err := svc.BatchGetProjects(ctx, &codebuild.BatchGetProjectsInput{Names: page.Projects}, func(options *codebuild.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- projectsOutput.Projects
	}
	return nil
}
