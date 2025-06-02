package iam

import (
	"context"

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

	wrappedGroup.Users = userARNs

	policies, err := svc.ListAttachedGroupPolicies(ctx, &iam.ListAttachedGroupPoliciesInput{GroupName: wrappedGroup.GroupName})
	if err != nil {
		return err
	}

	policyMap := map[string]string{}
	for _, p := range policies.AttachedPolicies {
		policyMap[*p.PolicyArn] = aws.ToString(p.PolicyName)
	}

	wrappedGroup.Policies = policyMap

	resource.Item = wrappedGroup
	return nil
}
