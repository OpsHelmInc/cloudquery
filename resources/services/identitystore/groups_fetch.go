package identitystore

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIdentitystoreGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	instance, err := getIamInstance(ctx, meta)
	if err != nil {
		return err
	}
	svc := meta.(*client.Client).Services().Identitystore
	config := identitystore.ListGroupsInput{
		IdentityStoreId: instance.IdentityStoreId,
		MaxResults:      aws.Int32(100),
	}
	paginator := identitystore.NewListGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		wrappedGroups := make([]*ohaws.IdentityStoreGroup, len(response.Groups))
		for i := range response.Groups {
			wrappedGroups[i] = &ohaws.IdentityStoreGroup{
				Group: response.Groups[i],
			}
		}

		res <- wrappedGroups
	}
	return nil
}

func resolveIdentityStoreGroupArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	g := resource.Item.(*ohaws.IdentityStoreGroup)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.IdentityStoreService),
		AccountID: "",
		Resource:  "group/" + aws.ToString(g.GroupId),
	}.String())
}

func getGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := meta.(*client.Client).Services().Identitystore
	group := resource.Item.(*ohaws.IdentityStoreGroup)

	config := identitystore.ListGroupMembershipsInput{
		GroupId:         group.GroupId,
		IdentityStoreId: group.IdentityStoreId,
		MaxResults:      aws.Int32(100),
	}
	paginator := identitystore.NewListGroupMembershipsPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		for _, m := range response.GroupMemberships {
			if mid, ok := m.MemberId.(*types.MemberIdMemberUserId); ok {
				group.Users = append(group.Users, arn.ARN{
					Partition: cl.Partition,
					Service:   string(client.IdentityStoreService),
					AccountID: "",
					Resource:  "user/" + mid.Value,
				}.String())
			}
		}
	}

	if group.Users == nil {
		group.Users = make([]string, 0)
	}
	sort.Strings(group.Users)

	return nil
}
