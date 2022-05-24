package zeromq

import (
	"context"

	"go.uber.org/zap"
)

type ConsumerWorkerPool struct {
	conn *Connection

	msgChannel chan []string

	subjectName string

	workers []*consumerWorker

	logger *zap.Logger
}

func (p *ConsumerWorkerPool) Run(ctx context.Context) error {
	err := p.conn.GetSocket().SetSubscribe(p.subjectName)
	if err != nil {
		return err
	}

	go p.runReceiver(ctx)

	p.runWorkers()

	return nil
}

func (p *ConsumerWorkerPool) runWorkers() {
	for _, w := range p.workers {
		w.msgChannel = p.msgChannel

		go w.Start()
	}
}

func (p *ConsumerWorkerPool) runReceiver(ctx context.Context) {
	for {
		msg, err := p.conn.Receive()
		if err == ErrConnectionClosed {
			reconnectErr := p.conn.Connect()
			if reconnectErr != nil {
				p.logger.Error("unable to reconnect", zap.Error(err), zap.String(SubjectTag, p.subjectName))

				p.Shutdown(ctx)
				return
			}
		} else if err != nil {
			p.logger.Error("unable to get message", zap.Error(err), zap.String(SubjectTag, p.subjectName))
			continue
		}

		p.msgChannel <- msg
	}
}

func (p *ConsumerWorkerPool) Shutdown(_ context.Context) {
	for _, w := range p.workers {
		w.Stop()
	}
}

func NewConsumerWorkersPool(logger *zap.Logger,
	conn *Connection,
	subjectName string,
	workersCount uint16,
	handler consumerHandler,
) *ConsumerWorkerPool {
	workersPool := &ConsumerWorkerPool{
		conn:        conn,
		logger:      logger,
		subjectName: subjectName,
	}

	for i := uint16(0); i < workersCount; i++ {
		ww := &consumerWorker{
			msgChannel:       workersPool.msgChannel,
			stopWorkerChanel: make(chan bool),
			handler:          handler,
			logger:           logger,
		}

		workersPool.workers = append(workersPool.workers, ww)
	}

	return workersPool
}
