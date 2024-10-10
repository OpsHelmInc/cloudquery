package iam

import (
	"context"
	"fmt"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_iam_groups"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Group.html`,
		Resolver:            fetchIamGroups,
		PreResourceResolver: getGroup,
		Transform:           transformers.TransformWithStruct(ohaws.Group{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			groupAttachedPolicies(),
			// groupLastAccessedDetails(),
			groupPolicies(),
		},
	}
}

func fetchIamGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListGroupsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListGroupsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		wrappedGroups := make([]*ohaws.Group, len(page.Groups))
		for i, group := range page.Groups {
			wrappedGroups[i] = &ohaws.Group{Group: group}
		}

		res <- wrappedGroups
	}
	return nil
}

func getGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	group := resource.Item.(*ohaws.Group)
	cl := meta.(*client.Client)
	svc := meta.(*client.Client).Services().Iam
	groupDetail, err := svc.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: aws.String(*group.GroupName),
	},
		func(o *iam.Options) {
			o.Region = cl.Region
		},
	)
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

	policies, err := svc.ListAttachedGroupPolicies(ctx, &iam.ListAttachedGroupPoliciesInput{GroupName: wrappedGroup.GroupName},
		func(o *iam.Options) {
			o.Region = cl.Region
		},
	)
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
