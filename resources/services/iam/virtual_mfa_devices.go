// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_virtual_mfa_devices",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html`,
		Resolver:    fetchIamVirtualMfaDevices,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "serial_number",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "base32_string_seed",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("Base32StringSeed"),
			},
			{
				Name:     "enable_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EnableDate"),
			},
			{
				Name:     "qr_code_png",
				Type:     schema.TypeIntArray,
				Resolver: schema.PathResolver("QRCodePNG"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "user",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("User"),
			},
		},
	}
}
