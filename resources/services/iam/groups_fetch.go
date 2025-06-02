package iam

import (
	"context"
	"fmt"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		wrappedGroups := make([]*ohaws.Group, len(response.Groups))
		for i, group := range response.Groups {
			wrappedGroups[i] = &ohaws.Group{Group: group}
		}
		res <- wrappedGroups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func getGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	group := resource.Item.(*ohaws.Group)
	svc := meta.(*client.Client).Services().Iam
	groupDetail, err := svc.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: aws.String(*group.GroupName),
	})
	if err != nil {
		return err
	}

	wrappedGroup := &ohaws.Group{Group: *groupDetail.Group}

	userARNs := make([]string, len(groupDetail.Users))
	for idx, u := range groupDetail.Users {
		userARNs[idx] = aws.ToString(u.Arn)
	}
	sort.Strings(userARNs)
	wrappedGroup.Users = userARNs

	var policies []string
	input := iam.ListAttachedGroupPoliciesInput{GroupName: group.GroupName}
	paginator := iam.NewListAttachedGroupPoliciesPaginator(svc, &input)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("failed to list attached group policies: %w", err)
		}
		for _, p := range response.AttachedPolicies {
			policies = append(policies, aws.ToString(p.PolicyArn))
		}
	}
	sort.Strings(policies)

	wrappedGroup.Policies = policies

	resource.Item = wrappedGroup
	return nil
}
