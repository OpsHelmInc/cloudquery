package backup

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func fetchBackupRegionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Backup
	input := backup.DescribeRegionSettingsInput{}

	output, err := svc.DescribeRegionSettings(ctx, &input)
	if err != nil {
		return err
	}
	res <- output
	return nil
}
