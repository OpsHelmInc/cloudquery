package mwaa

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
)

func fetchMwaaEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := mwaa.ListEnvironmentsInput{}
	c := meta.(*client.Client)
	svc := c.Services().Mwaa
	p := mwaa.NewListEnvironmentsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- response.Environments
	}
	return nil
}

func getEnvironment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	svc := meta.(*client.Client).Services().Mwaa
	name := resource.Item.(string)

	output, err := svc.GetEnvironment(ctx, &mwaa.GetEnvironmentInput{Name: &name})
	if err != nil {
		return err
	}

	resource.Item = output.Environment
	return nil
}
