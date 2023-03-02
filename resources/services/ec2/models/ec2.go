package models

import "github.com/aws/aws-sdk-go-v2/service/ec2/types"

type RegionalConfig struct {
	EbsEncryptionEnabledByDefault bool
	EbsDefaultKmsKeyId            *string
}

type LaunchTemplateWrapper struct {
	LaunchTemplateId   *string
	LaunchTemplateData *types.ResponseLaunchTemplateData
}
