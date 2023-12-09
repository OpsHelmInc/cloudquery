package ec2

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/OpsHelmInc/cloudquery/client"
)

func fetchEc2Eips(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	output, err := svc.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{
		Filters: []types.Filter{{Name: aws.String("domain"), Values: []string{"vpc"}}},
	})
	if err != nil {
		return err
	}
	res <- output.Addresses
	return nil
}

func resolveEipArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	eip := resource.Item.(types.Address)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "elastic-ip/" + aws.ToString(eip.AllocationId),
	}
	return resource.Set(c.Name, a.String())
}
