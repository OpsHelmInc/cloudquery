package iam

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsInput
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListGroups(ctx, &config)
		if err != nil {
			return err
		}
		wrappedGroups := make([]*ohaws.WrappedGroup, len(response.Groups))
		for i, group := range response.Groups {
			wrappedGroups[i] = &ohaws.WrappedGroup{
				Arn:        group.Arn,
				CreateDate: group.CreateDate,
				GroupId:    group.GroupId,
				GroupName:  group.GroupName,
				Path:       group.Path,
			}
		}
		res <- response.Groups
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func getGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	group := resource.Item.(*ohaws.WrappedGroup)
	svc := meta.(*client.Client).Services().Iam
	userDetail, err := svc.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: aws.String(*group.GroupName),
	})
	if err != nil {
		return err
	}

	wrappedGroup := &ohaws.WrappedGroup{
		Arn:        group.Arn,
		CreateDate: group.CreateDate,
		GroupId:    group.GroupId,
		GroupName:  group.GroupName,
		Path:       group.Path,
	}

	userARNs := []arn.ARN{}
	for _, u := range userDetail.Users {
		userARN, err := arn.Parse(*u.Arn)
		if err != nil {
			return err
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

func resolveIamGroupPolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*ohaws.WrappedGroup)
	svc := meta.(*client.Client).Services().Iam
	config := iam.ListAttachedGroupPoliciesInput{
		GroupName: r.GroupName,
	}
	response, err := svc.ListAttachedGroupPolicies(ctx, &config)
	if err != nil {
		return err
	}
	policyMap := map[string]*string{}
	for _, p := range response.AttachedPolicies {
		policyMap[*p.PolicyArn] = p.PolicyName
	}
	return resource.Set(c.Name, policyMap)
}
