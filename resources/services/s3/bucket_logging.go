package s3

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func bucketLogging() *schema.Table {
	return &schema.Table{
		Name:        "aws_s3_bucket_loggings",
		Description: `https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html`,
		Resolver:    fetchBucketLogging,
		Transform:   transformers.TransformWithStruct(&s3.GetBucketLoggingOutput{}, transformers.WithSkipFields("ResultMetadata")),
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
func fetchBucketLogging(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	r := parent.Item.(*ohaws.WrappedBucket)
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceS3).S3
	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: r.Name}, func(o *s3.Options) {
		o.Region = r.Region
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) {
			return nil
		}
		return err
	}
	res <- loggingOutput
	return nil
}
