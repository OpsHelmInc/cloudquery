package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func recommendationTemplates() *schema.Table {
	tableName := "aws_resiliencehub_recommendation_templates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_RecommendationTemplate.html`,
		Resolver:    fetchRecommendationTemplates,
		Transform:   transformers.TransformWithStruct(&types.RecommendationTemplate{}, transformers.WithPrimaryKeyComponents("AppArn", "AssessmentArn")),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), arnColumn("RecommendationTemplateArn")},
	}
}

func fetchRecommendationTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	p := resiliencehub.NewListRecommendationTemplatesPaginator(svc, &resiliencehub.ListRecommendationTemplatesInput{AssessmentArn: parent.Item.(*types.AppAssessment).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- out.RecommendationTemplates
	}
	return nil
}
