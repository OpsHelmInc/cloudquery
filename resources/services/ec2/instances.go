// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instances",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Instance.html`,
		Resolver:    fetchEc2Instances,
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
				Resolver: resolveInstanceArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:          "state_transition_reason_time",
				Type:          schema.TypeTimestamp,
				Resolver:      resolveEc2InstanceStateTransitionReasonTime,
				IgnoreInTests: true,
			},
			{
				Name:     "ami_launch_index",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AmiLaunchIndex"),
			},
			{
				Name:     "architecture",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Architecture"),
			},
			{
				Name:     "block_device_mappings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockDeviceMappings"),
			},
			{
				Name:     "boot_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BootMode"),
			},
			{
				Name:     "capacity_reservation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityReservationId"),
			},
			{
				Name:     "capacity_reservation_specification",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CapacityReservationSpecification"),
			},
			{
				Name:     "client_token",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClientToken"),
			},
			{
				Name:     "cpu_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CpuOptions"),
			},
			{
				Name:     "ebs_optimized",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EbsOptimized"),
			},
			{
				Name:     "elastic_gpu_associations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ElasticGpuAssociations"),
			},
			{
				Name:     "elastic_inference_accelerator_associations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ElasticInferenceAcceleratorAssociations"),
			},
			{
				Name:     "ena_support",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnaSupport"),
			},
			{
				Name:     "enclave_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnclaveOptions"),
			},
			{
				Name:     "hibernation_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HibernationOptions"),
			},
			{
				Name:     "hypervisor",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Hypervisor"),
			},
			{
				Name:     "iam_instance_profile",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IamInstanceProfile"),
			},
			{
				Name:     "image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageId"),
			},
			{
				Name:     "instance_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceId"),
			},
			{
				Name:     "instance_lifecycle",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceLifecycle"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "ipv6_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ipv6Address"),
			},
			{
				Name:     "kernel_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KernelId"),
			},
			{
				Name:     "key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyName"),
			},
			{
				Name:     "launch_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LaunchTime"),
			},
			{
				Name:     "licenses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Licenses"),
			},
			{
				Name:     "maintenance_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MaintenanceOptions"),
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
			},
			{
				Name:     "monitoring",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Monitoring"),
			},
			{
				Name:     "network_interfaces",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NetworkInterfaces"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "placement",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Placement"),
			},
			{
				Name:     "platform",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Platform"),
			},
			{
				Name:     "platform_details",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlatformDetails"),
			},
			{
				Name:     "private_dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateDnsName"),
			},
			{
				Name:     "private_dns_name_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateDnsNameOptions"),
			},
			{
				Name:     "private_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateIpAddress"),
			},
			{
				Name:     "product_codes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProductCodes"),
			},
			{
				Name:     "public_dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicDnsName"),
			},
			{
				Name:     "public_ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PublicIpAddress"),
			},
			{
				Name:     "ramdisk_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RamdiskId"),
			},
			{
				Name:     "root_device_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootDeviceName"),
			},
			{
				Name:     "root_device_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootDeviceType"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "source_dest_check",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SourceDestCheck"),
			},
			{
				Name:     "spot_instance_request_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SpotInstanceRequestId"),
			},
			{
				Name:     "sriov_net_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SriovNetSupport"),
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StateReason"),
			},
			{
				Name:     "state_transition_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateTransitionReason"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tpm_support",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TpmSupport"),
			},
			{
				Name:     "usage_operation",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UsageOperation"),
			},
			{
				Name:     "usage_operation_update_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UsageOperationUpdateTime"),
			},
			{
				Name:     "virtualization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualizationType"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
