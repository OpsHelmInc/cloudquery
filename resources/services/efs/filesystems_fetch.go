package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
)

func fetchEfsFilesystems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config efs.DescribeFileSystemsInput
	c := meta.(*client.Client)
	svc := c.Services().Efs
	paginator := efs.NewDescribeFileSystemsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *efs.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- page.FileSystems
	}
	return nil
}

func ResolveEfsFilesystemBackupPolicyStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.FileSystemDescription)
	config := efs.DescribeBackupPolicyInput{
		FileSystemId: p.FileSystemId,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Efs
	response, err := svc.DescribeBackupPolicy(ctx, &config)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return resource.Set(c.Name, types.StatusDisabled)
		}
		return err
	}
	if response.BackupPolicy == nil {
		return nil
	}

	return resource.Set(c.Name, response.BackupPolicy.Status)
}
