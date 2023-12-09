package iam

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/OpsHelmInc/ohaws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
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

	userARNs := []arn.ARN{}
	for _, u := range groupDetail.Users {
		userARN, err := arn.Parse(*u.Arn)
		if err != nil {
			return fmt.Errorf("failed to parse user ARN: %w, %v", err, *u.Arn)
		}
		userARNs = append(userARNs, userARN)
	}

	wrappedGroup.Users = userARNs

	policies, err := svc.ListAttachedGroupPolicies(ctx, &iam.ListAttachedGroupPoliciesInput{GroupName: wrappedGroup.GroupName})
	if err != nil {
		return err
	}

	policyMap := map[string]*string{}
	for _, p := range policies.AttachedPolicies {
		policyMap[*p.PolicyArn] = p.PolicyName
	}

	wrappedGroup.Policies = policyMap

	resource.Item = wrappedGroup
	return nil
}
