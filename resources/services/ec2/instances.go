package ec2

import (
	"context"
	"regexp"
	"time"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func Instances() *schema.Table {
	tableName := "aws_ec2_instances"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html`,
		Resolver:    fetchEc2Instances,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.Instance{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveInstanceArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "state_transition_reason_time",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: resolveEc2InstanceStateTransitionReasonTime,
			},
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

var stateTransitionReasonTimeRegex = regexp.MustCompile(`\((.*)\)`)

func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	p := ec2.NewDescribeInstancesPaginator(svc,
		&ec2.DescribeInstancesInput{MaxResults: aws.Int32(1000)})

	for p.HasMorePages() {
		output, err := p.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, reservation := range output.Reservations {
			res <- reservation.Instances
		}
	}

	return nil
}

func resolveInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.Instance)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "instance/" + aws.ToString(item.InstanceId),
	}
	return resource.Set(c.Name, a.String())
}

func resolveEc2InstanceStateTransitionReasonTime(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.Instance)
	if instance.StateTransitionReason == nil {
		return nil
	}
	match := stateTransitionReasonTimeRegex.FindStringSubmatch(*instance.StateTransitionReason)
	if len(match) < 2 {
		// failed to get time from message
		return nil
	}
	const layout = "2006-01-02 15:04:05 MST"
	tm, err := time.Parse(layout, match[1])
	if err != nil {
		return err
	}
	return resource.Set(c.Name, tm)
}
