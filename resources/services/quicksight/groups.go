package quicksight

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/pkg/errors"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_quicksight_groups"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Group.html",
		Resolver:    fetchQuicksightGroups,
		Transform:   transformers.TransformWithStruct(&types.Group{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			tagsCol,
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("Arn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
		Relations: []*schema.Table{groupMembers()},
	}
}

const defaultNamespace = "default"

func fetchQuicksightGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	input := quicksight.ListGroupsInput{
		AwsAccountId: aws.String(cl.AccountID),
		Namespace:    aws.String(defaultNamespace),
	}
	var ae smithy.APIError

	for {
		out, err := svc.ListGroups(ctx, &input, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}

			return err
		}
		res <- out.GroupList

		if aws.ToString(out.NextToken) == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}
