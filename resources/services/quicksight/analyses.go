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

func Analyses() *schema.Table {
	tableName := "aws_quicksight_analyses"
	return &schema.Table{
		Name:                tableName,
		Description:         "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Analysis.html",
		Resolver:            fetchQuicksightAnalyses,
		PreResourceResolver: getAnalysis,
		Transform:           transformers.TransformWithStruct(&types.Analysis{}),
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

func fetchQuicksightAnalyses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	input := quicksight.ListAnalysesInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	paginator := quicksight.NewListAnalysesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}
			return err
		}
		res <- result.AnalysisSummaryList
	}
	return nil
}

func getAnalysis(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	item := resource.Item.(types.AnalysisSummary)

	out, err := svc.DescribeAnalysis(ctx, &quicksight.DescribeAnalysisInput{
		AwsAccountId: aws.String(cl.AccountID),
		AnalysisId:   item.AnalysisId,
	}, func(options *quicksight.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = out.Analysis
	return nil
}
