package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/iam"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/cloudquery/v2/resources/services/iam/models"
)

func PasswordPolicies() *schema.Table {
	tableName := "aws_iam_password_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_PasswordPolicy.html`,
		Resolver:    fetchIamPasswordPolicies,
		Transform:   transformers.TransformWithStruct(&models.PasswordPolicyWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveIamPasswordPolicyArn,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIamPasswordPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.GetAccountPasswordPolicyInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	response, err := svc.GetAccountPasswordPolicy(ctx, &config, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			res <- models.PasswordPolicyWrapper{PolicyExists: false}
			return nil
		}
		return err
	}
	res <- models.PasswordPolicyWrapper{PasswordPolicy: *response.PasswordPolicy, PolicyExists: true}
	return nil
}

func resolveIamPasswordPolicyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)

	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "iam",
		Region:    "",
		AccountID: cl.AccountID,
		Resource:  "password-policy",
	}
	return resource.Set(c.Name, a.String())
}
