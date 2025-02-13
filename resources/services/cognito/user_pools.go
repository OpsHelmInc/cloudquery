package cognito

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func UserPools() *schema.Table {
	tableName := "aws_cognito_user_pools"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserPoolType.html`,
		Resolver:            fetchCognitoUserPools,
		PreResourceResolver: getUserPool,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cognito-identity"),
		Transform:           transformers.TransformWithStruct(&types.UserPoolType{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("UserPoolTags"),
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},

		Relations: []*schema.Table{
			userPoolIdentityProviders(),
		},
	}
}

func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCognitoidentityprovider).Cognitoidentityprovider
	params := cognitoidentityprovider.ListUserPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: aws.Int32(60),
	}
	paginator := cognitoidentityprovider.NewListUserPoolsPaginator(svc, &params)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cognitoidentityprovider.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.UserPools
	}
	return nil
}

func getUserPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCognitoidentityprovider).Cognitoidentityprovider
	item := resource.Item.(types.UserPoolDescriptionType)

	upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id}, func(options *cognitoidentityprovider.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}

	resource.Item = upo.UserPool
	return nil
}
