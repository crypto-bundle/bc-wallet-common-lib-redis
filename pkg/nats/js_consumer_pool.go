package nats

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

var (
	ErrReturnedNilConsumerInfo = errors.New("returned nil consumer info")
	ErrReturnedNilStreamInfo   = errors.New("returned nil stream info")
)

// jsConsumerWorkerPool is a minimal Worker implementation that simply wraps a
type jsConsumerWorkerPool struct {
	subscriptionSrv subscriptionService

	workers []*jsConsumerWorkerWrapper

	logger *zap.Logger
}

func (wp *jsConsumerWorkerPool) Init(ctx context.Context) error {
	return nil
}

func (wp *jsConsumerWorkerPool) Healthcheck(ctx context.Context) bool {
	return wp.subscriptionSrv.Healthcheck(ctx)
}

func (wp *jsConsumerWorkerPool) Run(ctx context.Context) error {
	wp.run(ctx)

	return nil
}

func (wp *jsConsumerWorkerPool) run(ctx context.Context) {
	for _, w := range wp.workers {
		go w.Run(ctx)
	}
}

func (wp *jsConsumerWorkerPool) Shutdown(ctx context.Context) error {
	for _, w := range wp.workers {
		w.Stop()
	}

	return nil
}

func NewConsumerWorkersPool(logger *zap.Logger,
	workersCount uint16,
	subscriptionSrv subscriptionService,
) *jsConsumerWorkerPool {
	l := logger.Named("consumer_worker_pool")

	workersPool := &jsConsumerWorkerPool{
		logger:          l,
		subscriptionSrv: subscriptionSrv,
	}

	for i := uint16(0); i < workersCount; i++ {
		ww := &jsConsumerWorkerWrapper{
			msgChannel:       workersPool.msgChannel,
			stopWorkerChanel: make(chan bool),
			handler:          workersPool.handler,
			logger:           l.With(zap.Uint16(WorkerUnitNumberTag, i)),
		}

		workersPool.workers = append(workersPool.workers, ww)
	}

	return workersPool
}
