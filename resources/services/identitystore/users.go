package identitystore

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Users() *schema.Table {
	tableName := "aws_identitystore_users"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html`,
		Resolver:    fetchUsers,
		Transform:   transformers.TransformWithStruct(&types.User{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "identitystore"),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(true),
			client.RequestRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveUserARN,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	instances, err := getIamInstances(ctx, meta)
	if err != nil {
		return err
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIdentitystore).Identitystore
	for _, instance := range instances {
		config := identitystore.ListUsersInput{
			IdentityStoreId: instance.IdentityStoreId,
		}
		paginator := identitystore.NewListUsersPaginator(svc, &config)
		for paginator.HasMorePages() {
			page, err := paginator.NextPage(ctx, func(options *identitystore.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- page.Users
		}
	}
	return nil
}

func resolveUserARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, userARN(cl, aws.ToString(resource.Item.(types.User).UserId)))
}

func userARN(cl *client.Client, userID string) string {
	// https: //docs.aws.amazon.com/service-authorization/latest/reference/list_awsidentitystore.html#awsidentitystore-resources-for-iam-policies
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentitystoreService),
		Resource:  "user/" + userID,
	}.String()
}
