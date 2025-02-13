package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
	sdkTypes "github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/types"
)

func vpcEndpointServicePermissions() *schema.Table {
	tableName := "aws_ec2_vpc_endpoint_service_permissions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AllowedPrincipal.html`,
		Resolver:    fetchEc2VpcEndpointServicePermissions,
		Transform:   transformers.TransformWithStruct(&types.AllowedPrincipal{}, transformers.WithPrimaryKeyComponents("ServiceId", "ServicePermissionId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2VpcEndpointServicePermissions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	endpointService := parent.Item.(types.ServiceDetail)
	if aws.ToString(endpointService.Owner) == "amazon" {
		return nil
	}
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceEc2).Ec2
	paginator := ec2.NewDescribeVpcEndpointServicePermissionsPaginator(svc, &ec2.DescribeVpcEndpointServicePermissionsInput{
		ServiceId: endpointService.ServiceId,
	})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *ec2.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.AllowedPrincipals
	}
	return nil
}
