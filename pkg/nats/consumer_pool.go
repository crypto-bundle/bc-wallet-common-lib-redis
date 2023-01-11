package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

// consumerWorkerPool is a minimal Worker implementation that simply wraps a
type consumerWorkerPool struct {
	handler consumerHandler
	workers []*consumerWorkerWrapper

	subscriptionSrv subscriptionService

	logger *zap.Logger
}

func (wp *consumerWorkerPool) Healthcheck(ctx context.Context) bool {
	return wp.subscriptionSrv.Healthcheck(ctx)
}

func (wp *consumerWorkerPool) Init(ctx context.Context) error {
	return wp.subscriptionSrv.Init(ctx)
}

func (wp *consumerWorkerPool) Run(ctx context.Context) error {
	wp.run()

	return wp.subscriptionSrv.Run(ctx)
}

func (wp *consumerWorkerPool) run() {
	for _, w := range wp.workers {
		go w.Start()
	}
}

func (wp *consumerWorkerPool) Shutdown(ctx context.Context) error {
	for _, w := range wp.workers {
		w.Stop()
	}

	return nil
}

func NewConsumerWorkersPool(logger *zap.Logger,
	msgChannel chan *nats.Msg,

	workersCount uint16,

	subjectName string,
	groupName string,

	autoReSubscribe bool,
	autoReSubscribeCount uint16,
	autoReSubscribeTimeout time.Duration,

	handler consumerHandler,
	natsConn *nats.Conn,
) *consumerWorkerPool {
	l := logger.Named("consumer_pool")

	subscriptionSrv := newPushSubscriptionService(l, natsConn, subjectName, groupName, autoReSubscribe,
		autoReSubscribeCount, autoReSubscribeTimeout, msgChannel)

	workersPool := &consumerWorkerPool{
		handler: handler,
		logger:  l,

		subscriptionSrv: subscriptionSrv,
	}

	for i := uint16(0); i < workersCount; i++ {
		ww := &consumerWorkerWrapper{
			msgChannel:       msgChannel,
			stopWorkerChanel: make(chan bool),
			handler:          workersPool.handler,
			logger:           l.With(zap.Uint16(WorkerUnitNumberTag, i)),
		}

		workersPool.workers = append(workersPool.workers, ww)
	}

	return workersPool
}
