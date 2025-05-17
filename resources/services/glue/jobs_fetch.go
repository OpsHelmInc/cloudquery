package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchGlueJobs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetJobsInput{}
	for {
		result, err := svc.GetJobs(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Jobs
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func resolveGlueJobArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, jobARN(cl, aws.ToString(resource.Item.(*ohaws.GlueJob).Name)))
}

func getGlueJob(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	j := resource.Item.(types.Job)
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(jobARN(cl, aws.ToString(j.Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}

	resource.Item = &ohaws.GlueJob{
		Job:  j,
		Tags: result.Tags,
	}

	return nil
}

func fetchGlueJobRuns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.GetJobRunsInput{
		JobName: parent.Item.(*ohaws.GlueJob).Name,
	}
	for {
		result, err := svc.GetJobRuns(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.JobRuns
		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func jobARN(cl *client.Client, name string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("job/%s", name),
	}.String()
}
