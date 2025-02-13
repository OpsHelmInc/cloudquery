package route53recoverycontrolconfig

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func routingControls() *schema.Table {
	tableName := "aws_route53recoverycontrolconfig_routing_controls"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/routing-control/latest/APIReference/API_ListRoutingControls.html`,
		Resolver:    fetchRoutingControls,
		Transform:   transformers.TransformWithStruct(&types.RoutingControl{}, transformers.WithPrimaryKeyComponents("ControlPanelArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("RoutingControlArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchRoutingControls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoverycontrolconfig).Route53recoverycontrolconfig
	controlPanel := parent.Item.(types.ControlPanel)
	paginator := route53recoverycontrolconfig.NewListRoutingControlsPaginator(svc, &route53recoverycontrolconfig.ListRoutingControlsInput{
		ControlPanelArn: controlPanel.ControlPanelArn,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoverycontrolconfig.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.RoutingControls
	}
	return nil
}
