package writers

import (
	"context"

	"github.com/OpsHelmInc/cloudquery/v2/plugin-sdk/message"
)

type Writer interface {
	Write(ctx context.Context, res <-chan message.WriteMessage) error
}
