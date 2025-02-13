package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	"github.com/OpsHelmInc/ohaws"
)

func signingCertificates() *schema.Table {
	tableName := "aws_iam_signing_certificates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html`,
		Resolver:    fetchUserSigningCertificates,
		Transform:   transformers.TransformWithStruct(&types.SigningCertificate{}, transformers.WithPrimaryKeyComponents("CertificateId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "user_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "user_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("user_id"),
			},
		},
	}
}

func fetchUserSigningCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := &iam.ListSigningCertificatesInput{UserName: parent.Item.(*ohaws.User).UserName}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListSigningCertificatesPaginator(svc, config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Certificates
	}
	return nil
}
