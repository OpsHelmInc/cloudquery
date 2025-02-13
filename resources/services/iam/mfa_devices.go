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

func mfaDevices() *schema.Table {
	tableName := "aws_iam_mfa_devices"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_MFADevice.html`,
		Resolver:    fetchIamMfaDevices,
		Transform:   transformers.TransformWithStruct(&types.MFADevice{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "serial_number",
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchIamMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*ohaws.User)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	paginator := iam.NewListMFADevicesPaginator(svc, &iam.ListMFADevicesInput{
		UserName: r.UserName,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.MFADevices
	}

	return nil
}
