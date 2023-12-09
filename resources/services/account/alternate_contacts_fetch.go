package account

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/client"
	"github.com/OpsHelmInc/cloudquery/plugin-sdk/schema"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/account/types"
)

func fetchAccountAlternateContacts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Account
	var contactTypes types.AlternateContactType
	for _, acType := range contactTypes.Values() {
		var input account.GetAlternateContactInput
		input.AlternateContactType = acType
		output, err := svc.GetAlternateContact(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.AlternateContact
	}
	return nil
}
