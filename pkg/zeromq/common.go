package zeromq

import (
	"context"
	"time"
)

type config interface {
	GetEndpoint() string
	GetUsername() string
	GetPassword() string
	GetReconnectMaxCount() uint16
	GetReconnectWaitTimeOut() time.Duration
}

type consumerHandler interface {
	Process(ctx context.Context, msg []string) error
}
