package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
	"github.com/OpsHelmInc/ohaws"
)

func Roles() *schema.Table {
	tableName := "aws_iam_roles"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/IAM/latest/APIReference/API_Role.html`,
		Resolver:            fetchIamRoles,
		PreResourceResolver: getRole,
		Transform:           transformers.TransformWithStruct(&ohaws.Role{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:     "assume_role_policy_document",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveRolesAssumeRolePolicyDocument,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{
			roleAttachedPolicies(),
			// roleLastAccessedDetails(),
			rolePolicies(),
		},
	}
}

func fetchIamRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListRolesInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListRolesPaginator(svc, &config)
	for paginator.HasMorePages() {
		response, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		wrappedRoles := make([]*ohaws.Role, len(response.Roles))
		for i, role := range response.Roles {
			wrappedRoles[i] = &ohaws.Role{Role: role}
		}
		res <- wrappedRoles
	}
	return nil
}

func getRole(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	role := resource.Item.(*ohaws.Role)
	svc := meta.(*client.Client).Services().Iam
	cl := meta.(*client.Client)
	roleDetails, err := svc.GetRole(ctx, &iam.GetRoleInput{
		RoleName: role.RoleName,
	}, func(o *iam.Options) {
		o.Region = cl.Region
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
		response, err := svc.ListAttachedRolePolicies(ctx, &input, func(o *iam.Options) {
			o.Region = cl.Region
		})
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
	var d map[string]any
	err = json.Unmarshal([]byte(decodedDocument), &d)
	if err != nil {
		return err
	}
	return resource.Set("assume_role_policy_document", d)
}
