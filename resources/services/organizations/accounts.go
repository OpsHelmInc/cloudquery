package organizations

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Accounts() *schema.Table {
	tableName := "aws_organizations_accounts"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:  fetchOrganizationsAccounts,
		Transform: transformers.TransformWithStruct(&types.Account{}),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "organizations"),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveAccountTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			organizationalAccountParents(),
		},
	}
}

func fetchOrganizationsAccounts(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations
	var input organizations.ListAccountsInput
	paginator := organizations.NewListAccountsPaginator(svc, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Accounts
	}
	return nil
}

func resolveAccountTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	cl := meta.(*client.Client)
	account := resource.Item.(types.Account)
	var tags []types.Tag
	input := organizations.ListTagsForResourceInput{
		ResourceId: account.Id,
	}
	paginator := organizations.NewListTagsForResourcePaginator(cl.Services(client.AWSServiceOrganizations).Organizations, &input)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		tags = append(tags, page.Tags...)
	}
	return resource.Set("tags", client.TagsToMap(tags))
}
