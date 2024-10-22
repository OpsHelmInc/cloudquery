package mwaa

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/mwaa"
	"github.com/aws/aws-sdk-go-v2/service/mwaa/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Environments() *schema.Table {
	tableName := "aws_mwaa_environments"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/mwaa/latest/API/API_Environment.html`,
		Resolver:            fetchMwaaEnvironments,
		Transform:           transformers.TransformWithStruct(&types.Environment{}),
		PreResourceResolver: getEnvironment,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "airflow"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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

func fetchMwaaEnvironments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := mwaa.ListEnvironmentsInput{}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceMwaa).Mwaa
	p := mwaa.NewListEnvironmentsPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *mwaa.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.Environments
	}
	return nil
}

func getEnvironment(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceMwaa).Mwaa
	name := resource.Item.(string)

	output, err := svc.GetEnvironment(ctx, &mwaa.GetEnvironmentInput{Name: &name}, func(options *mwaa.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = output.Environment
	return nil
}
