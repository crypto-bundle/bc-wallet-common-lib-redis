package zeromq

import (
	"context"

	"go.uber.org/zap"
)

type consumerWorker struct {
	msgChannel       <-chan []string
	stopWorkerChanel chan bool

	handler consumerHandler

	logger *zap.Logger
}

func (w *consumerWorker) Start() {
	for {
		select {
		case <-w.stopWorkerChanel:
			w.logger.Info("consumer worker. received close worker message")
			return

		case v, ok := <-w.msgChannel:
			if !ok {
				w.logger.Warn("consumer worker. message channel is closed")
				return
			}

			w.processMsg(v)
		}
	}
}

func (w *consumerWorker) processMsg(msg []string) {
	err := w.handler.Process(context.Background(), msg)
	if err != nil {
		w.logger.Error("unable to process message", zap.Error(err))
	}
}

func (w *consumerWorker) Stop() {
	w.stopWorkerChanel <- true
}
