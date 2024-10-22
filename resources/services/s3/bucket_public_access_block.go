package s3

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/OpsHelmInc/ohaws"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func bucketPublicAccessBlock() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_public_access_blocks",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html`,
		Resolver:    fetchBucketPublicAccessBlock,
		Transform:   transformers.TransformWithStruct(&s3.GetPublicAccessBlockOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "bucket_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchBucketPublicAccessBlock(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*ohaws.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	publicAccessOutput, err := svc.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = r.Region
	})
	if err != nil {
		// If we received any error other than NoSuchPublicAccessBlockConfiguration, we return and error
		if isBucketNotFoundError(cl, err) {
			return nil
		}
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	res <- publicAccessOutput
	return nil
}
