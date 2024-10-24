package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
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
		res <- response.Users
	}
	return nil
}

func getUser(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	listUser := resource.Item.(types.User)
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
	resource.Item = userDetail.User
	return nil
}
