package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type configParams interface {
	GetNatsAddresses() []string
	GetNatsJoinedAddresses() string
	//GetNatsHost() string
	//GetNatsPort() uint16
	GetNatsUser() string
	GetNatsPassword() string
	IsRetryOnConnectionFailed() bool
	GetNatsConnectionRetryCount() uint16
	GetNatsConnectionRetryTimeout() time.Duration
	GetFlushTimeout() time.Duration
	GetWorkersCountPerConsumer() uint16
}

type consumerHandler interface {
	Process(ctx context.Context, msg *nats.Msg) (ConsumerDirective, error)
}
