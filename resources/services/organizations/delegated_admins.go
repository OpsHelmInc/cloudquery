package organizations

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func DelegatedAdministrators() *schema.Table {
	tableName := "aws_organizations_delegated_administrators"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedAdministrator.html`,
		Resolver:    fetchOrganizationsDelegatedAdmins,
		Transform:   transformers.TransformWithStruct(&types.DelegatedAdministrator{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: schema.Tables{
			delegatedServices(),
		},
	}
}

func fetchOrganizationsDelegatedAdmins(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	var input organizations.ListDelegatedAdministratorsInput
	paginator := organizations.NewListDelegatedAdministratorsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DelegatedAdministrators
	}
	return nil
}
