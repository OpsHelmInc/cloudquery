package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/ohaws"
)

func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListRolesInput
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListRoles(ctx, &config)
		if err != nil {
			return err
		}
		wrappedRoles := make([]*ohaws.Role, len(response.Roles))
		for i, role := range response.Roles {
			wrappedRoles[i] = &ohaws.Role{Role: role}
		}
		res <- wrappedRoles
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func getRole(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	role := resource.Item.(*ohaws.Role)
	svc := meta.(*client.Client).Services().Iam
	cl := meta.(*client.Client)
	roleDetails, err := svc.GetRole(ctx, &iam.GetRoleInput{
		RoleName: role.RoleName,
	})
	if err != nil {
		return err
	}

	wrappedRole := &ohaws.Role{Role: *roleDetails.Role}

	input := iam.ListAttachedRolePoliciesInput{
		RoleName: role.RoleName,
	}
	policies := map[string]*string{}
	for {
		response, err := svc.ListAttachedRolePolicies(ctx, &input)
		if err != nil {
			if cl.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		for _, p := range response.AttachedPolicies {
			policies[*p.PolicyArn] = p.PolicyName
		}
		if response.Marker == nil {
			break
		}
		input.Marker = response.Marker
	}

	wrappedRole.Policies = policies
	resource.Item = wrappedRole

	return nil
}

func resolveRolesAssumeRolePolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*ohaws.Role)
	if r.AssumeRolePolicyDocument == nil {
		return nil
	}
	decodedDocument, err := url.QueryUnescape(*r.AssumeRolePolicyDocument)
	if err != nil {
		return err
	}
	var d map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &d)
	if err != nil {
		return err
	}
	return resource.Set("assume_role_policy_document", d)
}

func fetchIamRolePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	role := parent.Item.(*ohaws.Role)
	config := iam.ListRolePoliciesInput{
		RoleName: role.RoleName,
	}
	for {
		output, err := svc.ListRolePolicies(ctx, &config)
		if err != nil {
			if c.IsNotFoundError(err) {
				return nil
			}
			return err
		}
		res <- output.PolicyNames

		if aws.ToString(output.Marker) == "" {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func getRolePolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Iam
	p := resource.Item.(string)
	role := resource.Parent.Item.(*ohaws.Role)

	policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: role.RoleName})
	if err != nil {
		return err
	}
	resource.Item = policyResult
	return nil
}

func resolveRolePoliciesPolicyDocument(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*iam.GetRolePolicyOutput)

	decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
	if err != nil {
		return err
	}

	var document map[string]interface{}
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, document)
}
