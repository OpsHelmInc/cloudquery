package ec2

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func CapacityReservations() *schema.Table {
	tableName := "aws_ec2_capacity_reservations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeCapacityReservations.html`,
		Resolver:    fetchEc2CapacityReservations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Transform:   transformers.TransformWithStruct(&types.CapacityReservation{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("CapacityReservationArn"),
				PrimaryKeyComponent: true,
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

func fetchEc2CapacityReservations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	paginator := ec2.NewDescribeCapacityReservationsPaginator(svc, &ec2.DescribeCapacityReservationsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.CapacityReservations
	}

	return nil
}
