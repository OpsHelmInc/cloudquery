package s3

import (
	"testing"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/client/mocks"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildS3Buckets(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockS3Client(ctrl)
	b := s3Types.Bucket{}
	require.NoError(t, faker.FakeObject(&b))
	bloc := s3.GetBucketLocationOutput{}
	require.NoError(t, faker.FakeObject(&bloc))
	blog := s3.GetBucketLoggingOutput{}
	require.NoError(t, faker.FakeObject(&blog))
	bpol := s3.GetBucketPolicyOutput{}
	require.NoError(t, faker.FakeObject(&bpol))
	bpols := s3.GetBucketPolicyStatusOutput{}
	require.NoError(t, faker.FakeObject(&bpols))
	jsonDoc := `{"key":"value"}`
	bpol.Policy = &jsonDoc
	bver := s3.GetBucketVersioningOutput{}
	require.NoError(t, faker.FakeObject(&bver))
	bgrant := s3Types.Grant{}
	require.NoError(t, faker.FakeObject(&bgrant))
	// set up properly
	bgrant.Grantee.EmailAddress = nil
	bgrant.Grantee.ID = nil
	bgrant.Grantee.Type = s3Types.TypeGroup

	bcors := s3Types.CORSRule{}
	require.NoError(t, faker.FakeObject(&bcors))
	bencryption := s3.GetBucketEncryptionOutput{}
	require.NoError(t, faker.FakeObject(&bencryption))

	bpba := s3.GetPublicAccessBlockOutput{}
	require.NoError(t, faker.FakeObject(&bpba))
	btag := s3.GetBucketTaggingOutput{}
	require.NoError(t, faker.FakeObject(&btag))
	bownershipcontrols := s3.GetBucketOwnershipControlsOutput{}
	require.NoError(t, faker.FakeObject(&bownershipcontrols))

	ro := s3.GetBucketReplicationOutput{}
	require.NoError(t, faker.FakeObject(&ro))

	glco := s3.GetBucketLifecycleConfigurationOutput{}
	require.NoError(t, faker.FakeObject(&glco))

	websiteOutput := s3.GetBucketWebsiteOutput{}
	require.NoError(t, faker.FakeObject(&websiteOutput))

	gbnco := s3.GetBucketNotificationConfigurationOutput{}
	require.NoError(t, faker.FakeObject(&gbnco))

	golco := s3.GetObjectLockConfigurationOutput{}
	require.NoError(t, faker.FakeObject(&golco))

	lbitco := s3.ListBucketIntelligentTieringConfigurationsOutput{}
	require.NoError(t, faker.FakeObject(&lbitco))
	lbitco.NextContinuationToken = nil

	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&s3.ListBucketsOutput{Buckets: []s3Types.Bucket{b}}, nil).MinTimes(1)
	m.EXPECT().GetBucketLogging(gomock.Any(), gomock.Any(), gomock.Any()).Return(&blog, nil).MinTimes(1)
	m.EXPECT().GetBucketPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bpol, nil).MinTimes(1)
	m.EXPECT().GetBucketPolicyStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bpols, nil).MinTimes(1)
	m.EXPECT().GetBucketVersioning(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bver, nil).MinTimes(1)
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).Return(&s3.GetBucketAclOutput{Grants: []s3Types.Grant{bgrant}}, nil).MinTimes(1)
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).Return(&s3.GetBucketCorsOutput{CORSRules: []s3Types.CORSRule{bcors}}, nil).MinTimes(1)
	m.EXPECT().GetBucketEncryption(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bencryption, nil).MinTimes(1)
	m.EXPECT().GetPublicAccessBlock(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bpba, nil).MinTimes(1)
	m.EXPECT().GetBucketOwnershipControls(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bownershipcontrols, nil).MinTimes(1)
	m.EXPECT().GetBucketReplication(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ro, nil).MinTimes(1)
	m.EXPECT().GetBucketTagging(gomock.Any(), gomock.Any(), gomock.Any()).Return(&btag, nil).MinTimes(1)
	m.EXPECT().GetBucketLifecycleConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&glco, nil).MinTimes(1)
	m.EXPECT().GetBucketLocation(gomock.Any(), gomock.Any(), gomock.Any()).Return(&bloc, nil).MinTimes(1)
	m.EXPECT().GetBucketWebsite(gomock.Any(), gomock.Any(), gomock.Any()).Return(&websiteOutput, nil).MinTimes(1)
	m.EXPECT().GetBucketNotificationConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&gbnco, nil).MinTimes(1)
	m.EXPECT().GetObjectLockConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&golco, nil).MinTimes(1)
	m.EXPECT().ListBucketIntelligentTieringConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lbitco, nil).MinTimes(1)

	return client.Services{S3: m}
}

func TestS3Buckets(t *testing.T) {
	client.AwsMockTestHelper(t, Buckets(), buildS3Buckets, client.TestOptions{
		SkipEmptyCheckColumns: map[string][]string{"aws_s3_bucket_notification_configurations": {"event_bridge_configuration"}},
	})
}
