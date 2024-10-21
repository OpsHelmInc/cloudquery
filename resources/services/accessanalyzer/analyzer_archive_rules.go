package accessanalyzer

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func analyzerArchiveRules() *schema.Table {
	return &schema.Table{
		Name:        "aws_accessanalyzer_analyzer_archive_rules",
		Description: `https://docs.aws.amazon.com/access-analyzer/latest/APIReference/API_ArchiveRuleSummary.html`,
		Resolver:    fetchAccessanalyzerAnalyzerArchiveRules,
		Transform:   transformers.TransformWithStruct(&types.ArchiveRuleSummary{}, transformers.WithPrimaryKeyComponents("RuleName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "analyzer_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchAccessanalyzerAnalyzerArchiveRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceAccessanalyzer).Accessanalyzer
	config := accessanalyzer.ListArchiveRulesInput{AnalyzerName: analyzer.Name}
	paginator := accessanalyzer.NewListArchiveRulesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *accessanalyzer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.ArchiveRules
	}
	return nil
}
