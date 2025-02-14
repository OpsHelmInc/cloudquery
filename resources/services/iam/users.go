package iam

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/services"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func Users() *schema.Table {
	tableName := "aws_iam_users"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_User.html`,
		Resolver:            fetchIamUsers,
		PreResourceResolver: getUser,
		Transform:           transformers.TransformWithStruct(&ohaws.User{}, transformers.WithUnwrapAllEmbeddedStructs()),
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
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			userAccessKeys(),
			userGroups(),
			userAttachedPolicies(),
			userPolicies(),
			sshPublicKeys(),
			signingCertificates(),
			// userLastAccessedDetails(),
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
	cl := meta.(*client.Client)
	listUser := resource.Item.(*ohaws.User)
	svc := cl.Services(client.AWSServiceIam).Iam
	userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
		UserName: aws.String(*listUser.UserName),
	}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	ohUser := &ohaws.User{User: *userDetail.User}
	userName := userDetail.User.UserName

	// Modifies the given user to include all user metadata used in an OH user (policies, groups, tags, and MFA devices)
	userPolicies, err := svc.ListUserPolicies(ctx, &iam.ListUserPoliciesInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	for _, policyName := range userPolicies.PolicyNames {
		var thisPolicy ohaws.PolicyDocument
		policy, err := svc.GetUserPolicy(ctx, &iam.GetUserPolicyInput{UserName: userName, PolicyName: aws.String(policyName)}, func(o *iam.Options) {
			o.Region = cl.Region
		})
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
	attachedPolicies, err := svc.ListAttachedUserPolicies(ctx, &iam.ListAttachedUserPoliciesInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	ohUser.AttachedPolicies = attachedPolicies.AttachedPolicies

	groups, err := svc.ListGroupsForUser(ctx, &iam.ListGroupsForUserInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	ohUser.Groups = groups.Groups

	tags, err := svc.ListUserTags(ctx, &iam.ListUserTagsInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	ohUser.Tags = tags.Tags

	devices, err := svc.ListMFADevices(ctx, &iam.ListMFADevicesInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	ohUser.SetMFADevices(devices.MFADevices)

	var noUserError *types.NoSuchEntityException
	// GetLoginProfile will return a 404 User not found for any users that don't have a password set. Example error as follows
	// operation error IAM: GetLoginProfile, https response error StatusCode: 404, RequestID: 6158eccc-0d91-4e1c-a243-8cf8562e8e1f, NoSuchEntity: Login Profile for User test-user cannot be found.
	// Therefore, we know the user exists from above, so at this point, we will know if a user has their password set
	profileOutput, err := svc.GetLoginProfile(ctx, &iam.GetLoginProfileInput{UserName: userName}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	// nolint: gocritic
	if errors.As(err, &noUserError) {
		ohUser.SetLoginProfile(nil)
	} else if err != nil {
		return err
	} else {
		ohUser.SetLoginProfile(profileOutput.LoginProfile)
	}

	keysResp, err := svc.ListAccessKeys(ctx, &iam.ListAccessKeysInput{UserName: aws.String(*userName)}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	accessKeys := make([]ohaws.AccessKey, 0, len(keysResp.AccessKeyMetadata))
	for _, md := range keysResp.AccessKeyMetadata {
		k, err := getUserAccessKey(ctx, meta, svc, md)
		if err != nil {
			return err
		}

		accessKeys = append(accessKeys, *k)
	}

	ohUser.AccessKeys = accessKeys
	resource.Item = ohUser

	return nil
}

func getUserAccessKey(ctx context.Context, meta schema.ClientMeta, svc services.IamClient, md types.AccessKeyMetadata) (*ohaws.AccessKey, error) {
	cl := meta.(*client.Client)

	k := ohaws.AccessKey{
		AccessKeyMetadata: md,
	}

	r, err := svc.GetAccessKeyLastUsed(ctx, &iam.GetAccessKeyLastUsedInput{AccessKeyId: k.AccessKeyId}, func(o *iam.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return nil, fmt.Errorf("error getting access key last used: %w", err)
	}

	k.LastUsedDate = r.AccessKeyLastUsed.LastUsedDate
	k.LastUsedRegion = aws.ToString(r.AccessKeyLastUsed.Region)
	k.LastUsedServiceName = aws.ToString(r.AccessKeyLastUsed.ServiceName)

	return &k, nil
}
