package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func delegatedServices() *schema.Table {
	tableName := "aws_organizations_delegated_services"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedService.html`,
		Resolver:    fetchOrganizationsDelegatedServices,
		Transform:   transformers.TransformWithStruct(&types.DelegatedService{}, transformers.WithPrimaryKeyComponents("ServicePrincipal")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
		},
	}
}

func fetchOrganizationsDelegatedServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	paginator := organizations.NewListDelegatedServicesForAccountPaginator(svc, &organizations.ListDelegatedServicesForAccountInput{
		AccountId: parent.Item.(types.DelegatedAdministrator).Id,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DelegatedServices
	}
	return nil
}
