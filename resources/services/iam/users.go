package iam

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Users() *schema.Table {
	tableName := "aws_iam_users"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
		Transform:           transformers.TransformWithStruct(&types.User{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			userAccessKeys(),
			userGroups(),
			userAttachedPolicies(),
			userPolicies(),
			sshPublicKeys(),
			signingCertificates(),
			userLastAccessedDetails(),
			mfaDevices(),
		},
	}
}

func fetchIamUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.ListUsersInput{}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	p := iam.NewListUsersPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
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
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(*listUser.UserName),
	}, func(options *iam.Options) {
		options.Region = cl.Region
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
