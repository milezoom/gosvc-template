package contract

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type TemplateServiceGPubSubInterface interface {
	Add(ctx context.Context, m *pubsub.Message)
}
