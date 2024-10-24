package config

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

func remediationConfigurations() *schema.Table {
	tableName := "aws_config_remediation_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/config/latest/APIReference/API_RemediationConfiguration.html`,
		Resolver:    fetchRemediationConfigurations,
		Transform:   transformers.TransformWithStruct(&types.RemediationConfiguration{}, transformers.WithPrimaryKeyComponents("Arn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "config_rule_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}

func fetchRemediationConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceConfigservice).Configservice

	configRule := parent.Item.(types.ConfigRule).ConfigRuleName
	input := &configservice.DescribeRemediationConfigurationsInput{
		ConfigRuleNames: []string{*configRule},
	}

	// no pagination for this one
	output, err := svc.DescribeRemediationConfigurations(ctx, input, func(options *configservice.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- output.RemediationConfigurations

	return nil
}
