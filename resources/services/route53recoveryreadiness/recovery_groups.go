package route53recoveryreadiness

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func RecoveryGroups() *schema.Table {
	tableName := "aws_route53recoveryreadiness_recovery_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/recovery-readiness/latest/api/recoverygroups.html`,
		Resolver:    fetchControlPanels,
		Transform:   transformers.TransformWithStruct(&types.RecoveryGroupOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53-recovery-control-config"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("RecoveryGroupArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchControlPanels(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoveryreadiness).Route53recoveryreadiness
	paginator := route53recoveryreadiness.NewListRecoveryGroupsPaginator(svc, &route53recoveryreadiness.ListRecoveryGroupsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoveryreadiness.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.RecoveryGroups
	}
	return nil
}
