package ssoadmin

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func permissionSets() *schema.Table {
	tableName := "aws_ssoadmin_permission_sets"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_PermissionSet.html.
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.`,
		Resolver:            fetchSsoadminPermissionSets,
		PreResourceResolver: getSsoadminPermissionSet,
		Transform:           transformers.TransformWithStruct(&types.PermissionSet{}, transformers.WithPrimaryKeyComponents("PermissionSetArn")),
		Columns: []schema.Column{
			client.RequestAccountIDColumn(false),
			client.RequestRegionColumn(false),
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("instance_arn"),
				PrimaryKeyComponent: true,
			},
		},

		Relations: []*schema.Table{
			accountAssignments(),
			inlinePolicies(),
			customerManagedPolicies(),
			managedPolicies(),
			permissionsBoundaries(),
		},
	}
}

func getSsoadminPermissionSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin
	permission_set_arn := resource.Item.(string)
	instance_arn := resource.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.DescribePermissionSetInput{
		InstanceArn:      instance_arn,
		PermissionSetArn: &permission_set_arn,
	}

	response, err := svc.DescribePermissionSet(ctx, &config, func(o *ssoadmin.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = response.PermissionSet
	return nil
}

func fetchSsoadminPermissionSets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin
	instance_arn := parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.ListPermissionSetsInput{
		InstanceArn: instance_arn,
	}
	paginator := ssoadmin.NewListPermissionSetsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.PermissionSets
	}

	return nil
}
