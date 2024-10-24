package emr

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Studios() *schema.Table {
	tableName := "aws_emr_studios"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/emr/latest/APIReference/API_Studio.html`,
		Resolver:            fetchEmrStudios,
		PreResourceResolver: getStudio,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "elasticmapreduce"),
		Transform:           transformers.TransformWithStruct(&types.Studio{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Description:         `The Amazon Resource Name (ARN) of the EMR Studio.`,
				Resolver:            schema.PathResolver("StudioArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			studioSessionMapping(),
		},
	}
}

func fetchEmrStudios(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	config := emr.ListStudiosInput{}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	paginator := emr.NewListStudiosPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *emr.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Studios
	}
	return nil
}

func getStudio(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEmr).Emr
	response, err := svc.DescribeStudio(ctx, &emr.DescribeStudioInput{StudioId: resource.Item.(types.StudioSummary).StudioId}, func(options *emr.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.Studio
	err = resource.Set("tags", client.TagsToMap(response.Studio.Tags))
	if err != nil {
		return err
	}
	return nil
}
