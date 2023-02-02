// Code generated by codegen; DO NOT EDIT.

package kinesis

import (
	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Streams() *schema.Table {
	return &schema.Table{
		Name:                "aws_kinesis_streams",
		Description:         `https://docs.aws.amazon.com/kinesis/latest/APIReference/API_StreamDescriptionSummary.html`,
		Resolver:            fetchKinesisStreams,
		PreResourceResolver: getStream,
		Multiplex:           client.ServiceAccountRegionMultiplexer("kinesis"),
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
				Resolver: schema.PathResolver("StreamARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveKinesisStreamTags,
			},
			{
				Name:     "enhanced_monitoring",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnhancedMonitoring"),
			},
			{
				Name:     "open_shard_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("OpenShardCount"),
			},
			{
				Name:     "retention_period_hours",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("RetentionPeriodHours"),
			},
			{
				Name:     "stream_creation_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StreamCreationTimestamp"),
			},
			{
				Name:     "stream_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamName"),
			},
			{
				Name:     "stream_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StreamStatus"),
			},
			{
				Name:     "consumer_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ConsumerCount"),
			},
			{
				Name:     "encryption_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EncryptionType"),
			},
			{
				Name:     "key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyId"),
			},
			{
				Name:     "stream_mode_details",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StreamModeDetails"),
			},
		},
	}
}
