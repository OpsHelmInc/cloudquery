package quicksight

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/pkg/errors"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Folders() *schema.Table {
	tableName := "aws_quicksight_folders"
	return &schema.Table{
		Name:                tableName,
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Folder.html",
		Resolver:            fetchQuicksightFolders,
		PreResourceResolver: getFolder,
		Transform:           transformers.TransformWithStruct(&types.Folder{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			tagsCol,
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchQuicksightFolders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	input := quicksight.ListFoldersInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError
	// No paginator available
	for {
		out, err := svc.ListFolders(ctx, &input, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}

			return err
		}
		res <- out.FolderSummaryList

		if aws.ToString(out.NextToken) == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}

func getFolder(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	item := resource.Item.(types.FolderSummary)

	out, err := svc.DescribeFolder(ctx, &quicksight.DescribeFolderInput{
		AwsAccountId: aws.String(cl.AccountID),
		FolderId:     item.FolderId,
	}, func(options *quicksight.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = out.Folder
	return nil
}
