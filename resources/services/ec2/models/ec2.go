package models

import "github.com/aws/aws-sdk-go-v2/service/ec2/types"

type RegionalConfig struct {
	EbsEncryptionEnabledByDefault bool
	EbsDefaultKmsKeyId            *string
}

// Region combines the sdk Region type and the RegionConfig type above
type Region struct {
	types.Region
	EC2Config RegionalConfig
}
