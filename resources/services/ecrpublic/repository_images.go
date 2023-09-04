// Code generated by codegen; DO NOT EDIT.

package ecrpublic

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RepositoryImages() *schema.Table {
	return &schema.Table{
		Name:        "aws_ecrpublic_repository_images",
		Description: `https://docs.aws.amazon.com/AmazonECRPublic/latest/APIReference/API_ImageDetail.html`,
		Resolver:    fetchEcrpublicRepositoryImages,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "oh_resource_type",
				Type:     schema.TypeString,
				Resolver: client.StaticValueResolver("AWS::ECR::Image"),
			},
			{
				Name:     "artifact_media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ArtifactMediaType"),
			},
			{
				Name:     "image_digest",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageDigest"),
			},
			{
				Name:     "image_manifest_media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageManifestMediaType"),
			},
			{
				Name:     "image_pushed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ImagePushedAt"),
			},
			{
				Name:     "image_size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ImageSizeInBytes"),
			},
			{
				Name:     "image_tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ImageTags"),
			},
			{
				Name:     "registry_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryId"),
			},
			{
				Name:     "repository_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryName"),
			},
		},
	}
}
