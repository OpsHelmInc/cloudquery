package ssm

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func fetchSsmComplianceSummaryItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ssm

	params := ssm.ListComplianceSummariesInput{
		MaxResults: aws.Int32(50),
	}
	for {
		output, err := svc.ListComplianceSummaries(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.ComplianceSummaryItems

		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
