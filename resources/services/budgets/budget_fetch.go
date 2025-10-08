package budgets

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/ohaws"
)

type WrappedBudget struct {
	*types.Budget
	Tags []types.ResourceTag // This is the same as types.Tag, which doesn't exist in Budgets, because obviously
}

func getBudget(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Budgets

	// We will need this later, so may as well create now for when we need the various parts
	// format: arn:${Partition}:budgets::${Account}:budget/${BudgetName}

	// The resource does not contain it's own ARN, nor does it have an
	// accountID, so we will use the account ID from the billingView that it is
	// associated with (the client account ID does not necessarily match)
	billingViewARN, err := arn.Parse(*resource.Item.(types.Budget).BillingViewArn)
	if err != nil {
		return err
	}
	budgetARNStr := fmt.Sprintf("arn:aws:budgets::%s:budget/%s", billingViewARN.AccountID, *resource.Item.(types.Budget).BudgetName)
	budgetARN, err := arn.Parse(budgetARNStr)
	if err != nil {
		return err
	}

	results, err := svc.DescribeBudget(ctx, &budgets.DescribeBudgetInput{
		AccountId:  &budgetARN.AccountID,
		BudgetName: resource.Item.(types.Budget).BudgetName,
	})
	if err != nil {
		return err
	}

	if results == nil || results.Budget == nil {
		return errors.New("no results")
	}

	budgetOut := ohaws.Budget{
		Budget: results.Budget,
	}

	tags, err := svc.ListTagsForResource(ctx, &budgets.ListTagsForResourceInput{
		ResourceARN: aws.String(budgetARN.String()),
	})
	if err != nil {
		return err
	}

	if tags != nil {
		budgetOut.Tags = tags.ResourceTags
	}

	resource.Item = &budgetOut

	return nil
}

func fetchBudgetsBudgets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Budgets

	var input budgets.DescribeBudgetsInput
	paginator := budgets.NewDescribeBudgetsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- output.Budgets
	}
	return nil
}
