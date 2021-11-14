package websocket

import (
	"context"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/connectors/websocket/httphandler"
	"github.com/yngvark/gridwalls3/source/zombie-go/pkg/pubsub"
	"go.uber.org/zap"
)

// New returns a new instance
func New(
	ctx context.Context,
	cancelFn context.CancelFunc,
	logger *zap.SugaredLogger,
	subscriber chan string,
	allowedCorsOrigins map[string]bool,
) (pubsub.Publisher, pubsub.Consumer) {
	httpHandler := httphandler.New(cancelFn, logger, allowedCorsOrigins, subscriber)

	publisher := newPublisher(logger, httpHandler)
	consumer := newConsumer(ctx, logger, subscriber, httpHandler)

	return publisher, consumer
}
