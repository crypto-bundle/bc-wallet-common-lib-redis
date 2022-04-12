package nats

import (
	"errors"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type ProducerWorkerTask func(msg nats.Msg) error

var (
	ErrNilPubAck = errors.New("nil pub ack received")
)

// producerWorkerWrapper ...
type producerWorkerWrapper struct {
	logger           *zap.Logger
	msgChannel       <-chan *nats.Msg
	jsInfo           *nats.StreamInfo
	streamName       string
	subjects         []string
	natsProducerConn nats.JetStreamContext
	closeChanel      chan bool
	num              uint16
}

func (ww *producerWorkerWrapper) Run() {
	for {
		select {
		case v := <-ww.msgChannel:
			pubAck, err := ww.natsProducerConn.PublishMsg(v)
			if err != nil {
				ww.logger.Error("send message to broker service failed", zap.Error(err),
					zap.String(QueueSubjectNameTag, v.Subject),
					zap.String(QueueStreamNameTag, ww.jsInfo.Config.Name))
			}

			if pubAck == nil {
				ww.logger.Error("received nil pubAck", zap.Error(ErrNilPubAck))
				continue
			}

			ww.logger.Info("received pubAck", zap.String(QueuePubAckStreamNameTag, pubAck.Stream),
				zap.Uint64(QueuePubAckSequenceTag, pubAck.Sequence))

		case <-ww.closeChanel:
			ww.logger.Info("producer worker. received close worker message")
			return
		}
	}
}

func (ww *producerWorkerWrapper) SetStreamInfo(streamInfo *nats.StreamInfo) {
	ww.jsInfo = streamInfo
}

func (ww *producerWorkerWrapper) Stop() {
	ww.closeChanel <- true
}

func newProducerWorker(logger *zap.Logger,
	workerNum uint16,
	msgChannel chan *nats.Msg,
	streamName string,
	subjects []string,
	natsProducerConn nats.JetStreamContext,
	closeChan chan bool,
) *producerWorkerWrapper {
	l := logger.Named("producer.service.worker").
		With(zap.String(QueueStreamNameTag, streamName),
			zap.Strings(QueueSubjectNameTag, subjects),
			zap.Uint16(WorkerUnitNumberTag, workerNum))

	return &producerWorkerWrapper{
		logger:           l,
		msgChannel:       msgChannel,
		streamName:       streamName,
		subjects:         subjects,
		natsProducerConn: natsProducerConn,
		closeChanel:      closeChan,
		num:              0,
	}
}
