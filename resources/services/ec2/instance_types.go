// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
)

func InstanceTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instance_types",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceTypeInfo.html`,
		Resolver:    fetchEc2InstanceTypes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
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
				Resolver: resolveInstanceTypeArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_recovery_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoRecoverySupported"),
			},
			{
				Name:     "bare_metal",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BareMetal"),
			},
			{
				Name:     "burstable_performance_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BurstablePerformanceSupported"),
			},
			{
				Name:     "current_generation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CurrentGeneration"),
			},
			{
				Name:     "dedicated_hosts_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DedicatedHostsSupported"),
			},
			{
				Name:     "ebs_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EbsInfo"),
			},
			{
				Name:     "fpga_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FpgaInfo"),
			},
			{
				Name:     "free_tier_eligible",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("FreeTierEligible"),
			},
			{
				Name:     "gpu_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GpuInfo"),
			},
			{
				Name:     "hibernation_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HibernationSupported"),
			},
			{
				Name:     "hypervisor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Hypervisor"),
			},
			{
				Name:     "inference_accelerator_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InferenceAcceleratorInfo"),
			},
			{
				Name:     "instance_storage_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceStorageInfo"),
			},
			{
				Name:     "instance_storage_supported",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("InstanceStorageSupported"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "memory_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MemoryInfo"),
			},
			{
				Name:     "network_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkInfo"),
			},
			{
				Name:     "nitro_enclaves_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NitroEnclavesSupport"),
			},
			{
				Name:     "nitro_tpm_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NitroTpmInfo"),
			},
			{
				Name:     "nitro_tpm_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NitroTpmSupport"),
			},
			{
				Name:     "placement_group_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PlacementGroupInfo"),
			},
			{
				Name:     "processor_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProcessorInfo"),
			},
			{
				Name:     "supported_boot_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedBootModes"),
			},
			{
				Name:     "supported_root_device_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedRootDeviceTypes"),
			},
			{
				Name:     "supported_usage_classes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedUsageClasses"),
			},
			{
				Name:     "supported_virtualization_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SupportedVirtualizationTypes"),
			},
			{
				Name:     "v_cpu_info",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("VCpuInfo"),
			},
		},
	}
}
