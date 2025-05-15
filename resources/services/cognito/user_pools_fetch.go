package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchCognitoUserPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	params := cognitoidentityprovider.ListUserPoolsInput{
		// we want max results to reduce List calls as much as possible, services limited to less than or equal to 60"
		MaxResults: aws.Int32(60),
	}
	for {
		out, err := svc.ListUserPools(ctx, &params)
		if err != nil {
			return err
		}
		res <- out.UserPools

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}

func getUserPool(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	item := resource.Item.(types.UserPoolDescriptionType)

	upo, err := svc.DescribeUserPool(ctx, &cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: item.Id})
	if err != nil {
		return err
	}

	pool := ohaws.CognitoUserPool{
		UserPoolType: *upo.UserPool,
	}

	tags, err := getUserPoolTags(ctx, svc, *pool.Arn)
	if err != nil {
		return err
	}
	pool.Tags = tags

	resource.Item = &pool
	return nil
}

func getUserPoolTags(ctx context.Context, svc services.CognitoidentityproviderClient, poolArn string) (map[string]string, error) {
	tags, err := svc.ListTagsForResource(ctx, &cognitoidentityprovider.ListTagsForResourceInput{
		ResourceArn: &poolArn,
	})
	if err != nil {
		return nil, err
	}

	return tags.Tags, nil
}

func fetchCognitoUserPoolIdentityProviders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	pool := parent.Item.(*ohaws.CognitoUserPool)
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider

	params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
	for {
		out, err := svc.ListIdentityProviders(ctx, &params)
		if err != nil {
			return err
		}
		res <- out.Providers

		if aws.ToString(out.NextToken) == "" {
			break
		}
		params.NextToken = out.NextToken
	}
	return nil
}

func getUserPoolIdentityProvider(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cognitoidentityprovider
	item := resource.Item.(types.ProviderDescription)
	pool := resource.Parent.Item.(*ohaws.CognitoUserPool)

	pd, err := svc.DescribeIdentityProvider(ctx, &cognitoidentityprovider.DescribeIdentityProviderInput{
		ProviderName: item.ProviderName,
		UserPoolId:   pool.Id,
	})
	if err != nil {
		return err
	}

	resource.Item = pd.IdentityProvider
	return nil
}
