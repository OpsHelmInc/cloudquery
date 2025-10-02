package identitystore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
)

func fetchIdentitystoreUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListUsersInput{}
	config.IdentityStoreId = instance.IdentityStoreId
	for {
		response, err := svc.ListUsers(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Users

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveIdentityStoreUserArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	u := resource.Item.(types.User)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentityStoreService),
		AccountID: "",
		Resource:  "user/" + aws.ToString(u.UserId),
	}.String())
}
