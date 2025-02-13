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

func managedPolicies() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_managed_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ListManagedPoliciesInPermissionSet.html`,
		Resolver:    fetchManagedPolicies,
		Transform:   transformers.TransformWithStruct(&types.AttachedManagedPolicy{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			{
				Name:                "permission_set_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("permission_set_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("instance_arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchManagedPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin

	permissionSetARN := parent.Item.(*types.PermissionSet).PermissionSetArn
	instanceARN := parent.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.ListManagedPoliciesInPermissionSetInput{
		InstanceArn:      instanceARN,
		PermissionSetArn: permissionSetARN,
	}
	paginator := ssoadmin.NewListManagedPoliciesInPermissionSetPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AttachedManagedPolicies
	}
	return nil
}
