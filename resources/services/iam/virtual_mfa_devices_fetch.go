package iam

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func fetchIamVirtualMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config iam.ListVirtualMFADevicesInput
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListVirtualMFADevices(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.VirtualMFADevices
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
