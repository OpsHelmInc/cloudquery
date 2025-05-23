package servicecatalog

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func provisioningArtifact() *schema.Table {
	tableName := "aws_servicecatalog_provisioning_artifacts"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/servicecatalog/latest/dg/API_DescribeProvisioningArtifact.html`,
		Resolver:    fetchProvisioningArtifacts,
		Transform:   transformers.TransformWithStruct(&servicecatalog.DescribeProvisioningArtifactOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "provisioned_product_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "product_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("product_id"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "provisioning_artifact_id",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("provisioning_artifact_id"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchProvisioningArtifacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceServicecatalog).Servicecatalog
	p := parent.Item.(types.ProvisionedProductAttribute)

	input := servicecatalog.DescribeProvisioningArtifactInput{
		ProductId:                             p.ProductId,
		ProvisioningArtifactId:                p.ProvisioningArtifactId,
		IncludeProvisioningArtifactParameters: true,
	}

	resp, err := svc.DescribeProvisioningArtifact(ctx, &input, func(o *servicecatalog.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- resp
	return nil
}
