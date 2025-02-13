package ssoadmin

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func accountAssignments() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_account_assignments"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_AccountAssignment.html
The 'request_account_id' and 'request_region' columns are added to show the account_id and region of where the request was made from.`,
		Resolver:  fetchSsoadminAccountAssignments,
		Transform: transformers.TransformWithStruct(&types.AccountAssignment{}, transformers.WithPrimaryKeyComponents("PermissionSetArn", "PrincipalId", "PrincipalType", "AccountId")),
		Columns: schema.ColumnList{
			client.RequestAccountIDColumn(false),
			client.RequestRegionColumn(false),
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("instance_arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchSsoadminAccountAssignments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin
	configListAccountForPPS := ssoadmin.ListAccountsForProvisionedPermissionSetInput{
		InstanceArn:      parent.Parent.Item.(types.InstanceMetadata).InstanceArn,
		PermissionSetArn: parent.Item.(*types.PermissionSet).PermissionSetArn,
	}

	paginatorListAccountForPPS := ssoadmin.NewListAccountsForProvisionedPermissionSetPaginator(svc, &configListAccountForPPS)
	for paginatorListAccountForPPS.HasMorePages() {
		accounts, err := paginatorListAccountForPPS.NextPage(ctx, func(o *ssoadmin.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		for _, account := range accounts.AccountIds {
			configLAA := ssoadmin.ListAccountAssignmentsInput{
				AccountId:        aws.String(account),
				InstanceArn:      parent.Parent.Item.(types.InstanceMetadata).InstanceArn,
				PermissionSetArn: parent.Item.(*types.PermissionSet).PermissionSetArn,
			}
			paginator := ssoadmin.NewListAccountAssignmentsPaginator(svc, &configLAA)
			for paginator.HasMorePages() {
				page, err := paginator.NextPage(ctx, func(o *ssoadmin.Options) {
					o.Region = cl.Region
				})
				if err != nil {
					return err
				}
				res <- page.AccountAssignments
			}
		}
	}
	return nil
}
