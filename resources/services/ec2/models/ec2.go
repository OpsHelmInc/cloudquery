package models

type RegionalConfig struct {
	EbsEncryptionEnabledByDefault bool
	EbsDefaultKmsKeyId            *string
}

//type LaunchTemplateWrapper struct {
//	LaunchTemplateId   *string
//	LaunchTemplateData *types.ResponseLaunchTemplateData
//}
