package iam

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/cloudquery/resources/services/iam/models"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := iam.ListUsersInput{}
	c := meta.(*client.Client)
	svc := c.Services().Iam
	p := iam.NewListUsersPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx)
		if err != nil {
			return err
		}
		wrappedUsers := make([]*ohaws.User, len(response.Users))
		for i, user := range response.Users {
			wrappedUsers[i] = &ohaws.User{User: user}
		}
		res <- wrappedUsers
	}
	return nil
}

func getUser(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	listUser := resource.Item.(*ohaws.User)
	svc := meta.(*client.Client).Services().Iam
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(*listUser.UserName),
	})
	if err != nil {
		return err
	}
	ohUser := &ohaws.User{User: *userDetail.User}
	userName := userDetail.User.UserName

	// Modifies the given user to include all user metadata used in an OH user (policies, groups, tags, and MFA devices)
	userPolicies, err := svc.ListUserPolicies(ctx, &iam.ListUserPoliciesInput{UserName: userName})
	if err != nil {
		return err
	}

	for _, policyName := range userPolicies.PolicyNames {
		var thisPolicy ohaws.PolicyDocument
		policy, err := svc.GetUserPolicy(ctx, &iam.GetUserPolicyInput{UserName: userName, PolicyName: aws.String(policyName)})
		if err != nil {
			return err
		}
		parsedDoc, err := url.QueryUnescape(*policy.PolicyDocument)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(parsedDoc), &thisPolicy)
		if err != nil {
			return err
		}
		ohUser.AddInlinePolicy(policyName, thisPolicy)
	}
	attachedPolicies, err := svc.ListAttachedUserPolicies(ctx, &iam.ListAttachedUserPoliciesInput{UserName: userName})
	if err != nil {
		return err
	}
	ohUser.AttachedPolicies = attachedPolicies.AttachedPolicies

	groups, err := svc.ListGroupsForUser(ctx, &iam.ListGroupsForUserInput{UserName: userName})
	if err != nil {
		return err
	}
	ohUser.Groups = groups.Groups

	tags, err := svc.ListUserTags(ctx, &iam.ListUserTagsInput{UserName: userName})
	if err != nil {
		return err
	}
	ohUser.Tags = tags.Tags

	devices, err := svc.ListMFADevices(ctx, &iam.ListMFADevicesInput{UserName: userName})
	if err != nil {
		return err
	}
	ohUser.SetMFADevices(devices.MFADevices)

	var noUserError *types.NoSuchEntityException
	// GetLoginProfile will return a 404 User not found for any users that don't have a password set. Example error as follows
	// operation error IAM: GetLoginProfile, https response error StatusCode: 404, RequestID: 6158eccc-0d91-4e1c-a243-8cf8562e8e1f, NoSuchEntity: Login Profile for User test-user cannot be found.
	// Therefore, we know the user exists from above, so at this point, we will know if a user has their password set
	profileOutput, err := svc.GetLoginProfile(ctx, &iam.GetLoginProfileInput{UserName: userName})
	// nolint: gocritic
	if errors.As(err, &noUserError) {
		ohUser.SetLoginProfile(nil)
	} else if err != nil {
		return err
	} else {
		ohUser.SetLoginProfile(profileOutput.LoginProfile)
	}

	keysResp, err := svc.ListAccessKeys(ctx, &iam.ListAccessKeysInput{UserName: aws.String(*userName)})
	if err != nil {
		return err
	}

	var accessKeys = make([]ohaws.AccessKey, 0, len(keysResp.AccessKeyMetadata))
	for _, md := range keysResp.AccessKeyMetadata {
		k, err := getUserAccessKey(ctx, svc, *userName, *md.AccessKeyId)
		if err != nil {
			return err
		}

		accessKeys = append(accessKeys, *k)
	}

	ohUser.AccessKeys = accessKeys
	resource.Item = ohUser

	return nil
}

func fetchIamUserGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListGroupsForUserInput
	p := parent.Item.(*ohaws.User)
	svc := meta.(*client.Client).Services().Iam
	config.UserName = p.UserName
	for {
		output, err := svc.ListGroupsForUser(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Groups
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func fetchIamUserAccessKeys(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAccessKeysInput
	p := parent.Item.(*ohaws.User)
	svc := meta.(*client.Client).Services().Iam
	config.UserName = p.UserName
	for {
		output, err := svc.ListAccessKeys(ctx, &config)
		if err != nil {
			return err
		}

		keys := make([]models.AccessKeyWrapper, len(output.AccessKeyMetadata))
		for i, key := range output.AccessKeyMetadata {
			switch i {
			case 0:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
			case 1:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
			default:
				keys[i] = models.AccessKeyWrapper{AccessKeyMetadata: key}
			}
		}
		res <- keys
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func postIamUserAccessKeyResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	r := resource.Item.(models.AccessKeyWrapper)
	if r.AccessKeyId == nil {
		return nil
	}
	svc := meta.(*client.Client).Services().Iam
	output, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: r.AccessKeyId})
	if err != nil {
		return err
	}
	if output.AccessKeyLastUsed != nil {
		if err := resource.Set("last_used", output.AccessKeyLastUsed.LastUsedDate); err != nil {
			return err
		}
		if err := resource.Set("last_used_service_name", output.AccessKeyLastUsed.ServiceName); err != nil {
			return err
		}
	}
	return nil
}

func fetchIamUserAttachedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListAttachedUserPoliciesInput
	p := parent.Item.(*ohaws.User)
	svc := meta.(*client.Client).Services().Iam
	config.UserName = p.UserName
	for {
		output, err := svc.ListAttachedUserPolicies(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.AttachedPolicies
		if output.Marker == nil {
			break
		}
		config.Marker = output.Marker
	}
	return nil
}

func resolveAccessKeyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	key := resource.Item.(models.AccessKeyWrapper)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "iam",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  "access-key/" + aws.ToString(key.UserName) + "/" + aws.ToString(key.AccessKeyId),
	}
	return resource.Set(c.Name, a.String())
}

func getUserAccessKey(ctx context.Context, svc services.IamClient, userName, keyID string) (*ohaws.AccessKey, error) {
	keysResp, err := svc.ListAccessKeys(ctx, &iam.ListAccessKeysInput{UserName: aws.String(userName)})
	if err != nil {
		return nil, fmt.Errorf("error getting user access keys: %w", err)
	}

	var k ohaws.AccessKey
	for _, md := range keysResp.AccessKeyMetadata {
		if md.AccessKeyId != nil && *md.AccessKeyId == keyID {
			k.AccessKeyMetadata = md
			break
		}
	}

	if k.AccessKeyId == nil {
		return nil, fmt.Errorf("key with id not found: %s", keyID)
	}

	r, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: k.AccessKeyId})
	if err != nil {
		return nil, fmt.Errorf("error getting access key last used: %w", err)
	}

	k.LastUsedDate = r.AccessKeyLastUsed.LastUsedDate
	if r.AccessKeyLastUsed.Region != nil {
		k.LastUsedRegion = *r.AccessKeyLastUsed.Region
	}
	if r.AccessKeyLastUsed.ServiceName != nil {
		k.LastUsedServiceName = *r.AccessKeyLastUsed.ServiceName
	}

	return &k, nil
}
