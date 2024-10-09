package organizations

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func organizationalAccountParents() *schema.Table {
	tableName := "aws_organizations_account_parents"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListParents.html
The 'request_account_id' column is added to show from where the request was made.`,
		Resolver:  fetchParents,
		Transform: transformers.TransformWithStruct(&types.Parent{}, transformers.WithPrimaryKeyComponents("Type")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(true),
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "parent_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}
func fetchParents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceOrganizations).Organizations

	resp, err := svc.ListParents(ctx, &organizations.ListParentsInput{
		ChildId: parent.Item.(types.Account).Id,
	}, func(options *organizations.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- resp.Parents

	return nil
}
