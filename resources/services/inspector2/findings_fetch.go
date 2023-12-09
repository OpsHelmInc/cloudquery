package inspector2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
)

func fetchInspector2Findings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Inspector2
	input := inspector2.ListFindingsInput{MaxResults: aws.Int32(100)}
	for {
		response, err := svc.ListFindings(ctx, &input)
		if err != nil {
			return err
		}
		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
