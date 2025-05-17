package iot

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

func fetchIotCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListCertificatesInput{
		PageSize: aws.Int32(250),
	}

	for {
		response, err := svc.ListCertificates(ctx, &input)
		if err != nil {
			return err
		}

		for _, ct := range response.Certificates {
			resp, err := svc.DescribeCertificate(ctx, &iot.DescribeCertificateInput{
				CertificateId: ct.CertificateId,
			}, func(options *iot.Options) {
				options.Region = cl.Region
			})
			if err != nil {
				return err
			}
			cert := &ohaws.IoTCertificate{
				CertificateDescription: *resp.CertificateDescription,
			}
			res <- cert
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return nil
}

func getIotCertificate(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	cert := resource.Item.(*ohaws.IoTCertificate)

	tags, err := getResourceTags(ctx, svc, aws.ToString(cert.CertificateArn))
	if err != nil {
		return fmt.Errorf("error listing tags: %w", err)
	}

	cert.Tags = tags
	resource.Item = cert

	return nil
}

func ResolveIotCertificatePolicies(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	i := resource.Item.(*ohaws.IoTCertificate)
	cl := meta.(*client.Client)
	svc := cl.Services().Iot
	input := iot.ListAttachedPoliciesInput{
		Target:   i.CertificateArn,
		PageSize: aws.Int32(250),
	}

	var policies []string
	for {
		response, err := svc.ListAttachedPolicies(ctx, &input)
		if err != nil {
			return err
		}

		for _, p := range response.Policies {
			policies = append(policies, *p.PolicyArn)
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		input.Marker = response.NextMarker
	}
	return resource.Set(c.Name, policies)
}
