package resiliencehub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func appComponentRecommendations() *schema.Table {
	tableName := "aws_resiliencehub_component_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/resilience-hub/latest/APIReference/API_ComponentRecommendation.html`,
		Resolver:    fetchComponentRecommendations,
		Transform:   transformers.TransformWithStruct(&types.ComponentRecommendation{}, transformers.WithPrimaryKeyComponents("AppComponentName")),
		Columns:     []schema.Column{client.DefaultAccountIDColumn(false), client.DefaultRegionColumn(false), appARN, assessmentARN},
	}
}

func fetchComponentRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceResiliencehub).Resiliencehub
	p := resiliencehub.NewListAppComponentRecommendationsPaginator(svc, &resiliencehub.ListAppComponentRecommendationsInput{AssessmentArn: parent.Item.(*types.AppAssessment).AppArn})
	for p.HasMorePages() {
		out, err := p.NextPage(ctx, func(options *resiliencehub.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- out.ComponentRecommendations
	}
	return nil
}
