package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchCloudwatchAlarms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudwatch.DescribeAlarmsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudwatch
	for {
		response, err := svc.DescribeAlarms(ctx, &config)
		if err != nil {
			return err
		}

		alarms := make([]*ohaws.Alarm, len(response.MetricAlarms))
		for idx, a := range response.MetricAlarms {
			alarms[idx] = &ohaws.Alarm{
				MetricAlarm: &a,
			}
		}

		res <- alarms
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func resolveCloudwatchAlarmDimensions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	alarm := resource.Item.(*ohaws.Alarm)
	dimensions := make(map[string]*string)
	for _, d := range alarm.Dimensions {
		dimensions[*d.Name] = d.Value
	}
	return resource.Set("dimensions", dimensions)
}

func getAlarm(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudwatch
	alarm := resource.Item.(*ohaws.Alarm)

	input := cloudwatch.ListTagsForResourceInput{
		ResourceARN: alarm.MetricAlarm.AlarmArn,
	}
	output, err := svc.ListTagsForResource(ctx, &input)
	if err != nil {
		return err
	}

	alarm.Tags = make([]ohaws.Tag, len(output.Tags))
	for i := range output.Tags {
		alarm.Tags[i].Key = output.Tags[i].Key
		alarm.Tags[i].Value = output.Tags[i].Value
	}

	return nil
}
