package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/client/services"
	"github.com/OpsHelmInc/ohaws"
)

func fetchBackupPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupPlans.html
	for {
		result, err := svc.ListBackupPlans(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.BackupPlansList

		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func getPlan(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	m := resource.Item.(types.BackupPlansListMember)

	plan, err := svc.GetBackupPlan(
		ctx,
		&backup.GetBackupPlanInput{BackupPlanId: m.BackupPlanId, VersionId: m.VersionId},
	)
	if err != nil {
		return err
	}

	var b ohaws.BackupPlan
	// convert response struct to BackupVaultListMember struct
	if err := mapstructure.Decode(plan, &b.GetBackupPlanOutput); err != nil {
		return err
	}

	tags, err := getPlanTags(ctx, svc, *b.BackupPlanArn)
	if err != nil {
		return err
	}
	b.Tags = tags

	resource.Item = &b
	return nil
}

func getPlanTags(ctx context.Context, svc services.BackupClient, planArn string) (map[string]string, error) {
	params := backup.ListTagsInput{ResourceArn: &planArn}
	paginator := backup.NewListTagsPaginator(svc, &params)
	tags := map[string]string{}
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for k, v := range result.Tags {
			tags[k] = v
		}
	}
	return tags, nil
}

func fetchBackupPlanSelections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	plan := parent.Item.(*ohaws.BackupPlan)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupSelectionsInput{
		BackupPlanId: plan.BackupPlanId,
		MaxResults:   aws.Int32(1000), // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupSelections.html
	}
	for {
		result, err := svc.ListBackupSelections(ctx, &params)
		if err != nil {
			return err
		}
		for _, m := range result.BackupSelectionsList {
			s, err := svc.GetBackupSelection(
				ctx,
				&backup.GetBackupSelectionInput{BackupPlanId: plan.BackupPlanId, SelectionId: m.SelectionId},
				func(o *backup.Options) {
					o.Region = cl.Region
				},
			)
			if err != nil {
				return err
			}
			if s != nil {
				res <- *s
			}
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}
