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

func fetchIdentitystoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListGroupsInput{}
	config.IdentityStoreId = instance.IdentityStoreId
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Groups

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveIdentityStoreGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	g := resource.Item.(types.Group)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentityStoreService),
		AccountID: "",
		Resource:  "group/" + aws.ToString(g.GroupId),
	}.String())
}
