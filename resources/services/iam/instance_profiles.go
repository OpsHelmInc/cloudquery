package iam

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func InstanceProfiles() *schema.Table {
	tableName := "aws_iam_instance_profiles"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_InstanceProfile.html`,
		Resolver:    fetchIamInstanceProfiles,
		Transform:   transformers.TransformWithStruct(&types.InstanceProfile{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			{
				Name:                "id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("InstanceProfileId"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveIamInstanceProfileTags,
			},
			{
				Name:                "arn",
				Resolver:            schema.PathResolver("Arn"),
				Type:                arrow.BinaryTypes.String,
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchIamInstanceProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := iam.ListInstanceProfilesInput{}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	p := iam.NewListInstanceProfilesPaginator(svc, &config)
	for p.HasMorePages() {
		response, err := p.NextPage(ctx, func(options *iam.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- response.InstanceProfiles
	}
	return nil
}

func resolveIamInstanceProfileTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.InstanceProfile)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceIam).Iam
	response, err := svc.ListInstanceProfileTags(ctx, &iam.ListInstanceProfileTagsInput{InstanceProfileName: r.InstanceProfileName}, func(options *iam.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}
