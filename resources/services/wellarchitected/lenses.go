package wellarchitected

import (
	"context"

	"github.com/apache/arrow/go/v16/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"

	"github.com/OpsHelmInc/cloudquery/v2/client"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/schema"
	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/transformers"
)

// lens info consists of types.LensSummary & types.Lens fields
type lens struct {
	*types.LensSummary
	// extra field from types.Lens
	ShareInvitationId *string
	// extra field from types.Lens
	Tags map[string]string
}

func Lenses() *schema.Table {
	name := "aws_wellarchitected_lenses"
	return &schema.Table{
		Name:                name,
		Description:         `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_Lens.html`,
		Transform:           transformers.TransformWithStruct(new(lens), transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:           client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:            fetchLenses,
		PreResourceResolver: getLens,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("LensArn"),
				PrimaryKeyComponent: true,
			},
			client.OhResourceTypeColumn(),
		},
	}
}

func fetchLenses(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services(client.AWSServiceWellarchitected).Wellarchitected

	// we do fetch for all 3 types
	for _, lensType := range types.LensType("").Values() {
		p := wellarchitected.NewListLensesPaginator(service,
			&wellarchitected.ListLensesInput{
				LensStatus: types.LensStatusTypeAll,
				LensType:   lensType,
				MaxResults: aws.Int32(50),
			},
		)
		for p.HasMorePages() {
			output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}
			res <- output.LensSummaries
		}
	}

	return nil
}

func getLens(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	service := cl.Services(client.AWSServiceWellarchitected).Wellarchitected
	summary := resource.Item.(types.LensSummary)
	l := &lens{LensSummary: &summary}
	resource.SetItem(l)
	input := &wellarchitected.GetLensInput{LensAlias: l.LensAlias, LensVersion: summary.LensVersion}
	if summary.LensType == types.LensTypeAwsOfficial {
		input.LensVersion = nil // official lenses don't support versions
	}
	if l.LensAlias == nil {
		// LensAlias is required so if it's nil we can't do anything
		return nil
	}
	out, err := service.GetLens(ctx, input, func(o *wellarchitected.Options) { o.Region = cl.Region })
	if err != nil {
		cl.Logger().Err(err).Str("table", resource.Table.Name).Msg("Failed to perform get, ignoring...")
		// At the very least we want the summary data to be filled in
		// so don't update the item (leave summary in place)
		return nil
	}

	l.ShareInvitationId = out.Lens.ShareInvitationId
	l.Tags = out.Lens.Tags
	resource.SetItem(l)
	return nil
}
