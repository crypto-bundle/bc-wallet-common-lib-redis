package nats

import (
	"context"
	"errors"
	"go.uber.org/zap"

	"github.com/nats-io/nats.go"
)

// ProducerWorkerPool is a minimal Worker implementation that simply wraps a
type ProducerWorkerPool struct {
	logger           *zap.Logger
	msgChannel       chan *nats.Msg
	handler          ProducerWorkerTask
	streamName       string
	subject          []string
	storage          nats.StorageType
	jsInfo           *nats.StreamInfo
	jsConfig         *nats.StreamConfig
	natsProducerConn nats.JetStreamContext
	workers          []*producerWorkerWrapper
}

func (wp *ProducerWorkerPool) Init(ctx context.Context) error {
	streamInfo, err := wp.getOrCreateStream(ctx)
	if err != nil {
		return err
	}

	if streamInfo == nil {
		return ErrReturnedNilStreamInfo
	}

	//addStreamErr := wp.natsProducerConn.DeleteStream(wp.streamName)
	//if addStreamErr != nil {
	//	return addStreamErr
	//}

	if len(streamInfo.Config.Subjects) != len(wp.subject) {
		oldConf := streamInfo.Config
		oldConf.Subjects = append(wp.subject, oldConf.Subjects...)

		stream, addStreamErr := wp.natsProducerConn.UpdateStream(&oldConf)
		if addStreamErr != nil {
			return addStreamErr
		}

		streamInfo = stream
	}

	for i, _ := range wp.workers {
		wp.workers[i].SetStreamInfo(streamInfo)
	}

	return nil
}

func (wp *ProducerWorkerPool) getOrCreateStream(ctx context.Context) (*nats.StreamInfo, error) {
	streamInfo, err := wp.natsProducerConn.StreamInfo(wp.jsConfig.Name)
	if err != nil {
		if errors.Is(err, nats.ErrStreamNotFound) {
			stream, addStreamErr := wp.natsProducerConn.AddStream(wp.jsConfig)
			if addStreamErr != nil {
				return nil, addStreamErr
			}

			streamInfo = stream
		}

		return nil, err
	}

	return streamInfo, nil
}

func (wp *ProducerWorkerPool) Run(ctx context.Context) error {
	wp.run()

	return nil
}

func (wp *ProducerWorkerPool) run() {
	for i, _ := range wp.workers {
		go wp.workers[i].Run()
	}
}

func (wp *ProducerWorkerPool) Shutdown(ctx context.Context) error {
	for _, w := range wp.workers {
		w.Stop()
	}

	return nil
}

func (wp *ProducerWorkerPool) Produce(ctx context.Context, msg *nats.Msg) {
	wp.msgChannel <- msg
}

func NewProducerWorkersPool(
	logger *zap.Logger,
	workersCount uint16,
	streamName string,
	subjects []string,
	storage nats.StorageType,
	natsProducerConn nats.JetStreamContext,
) (*ProducerWorkerPool, error) {
	l := logger.Named("producer.service").
		With(zap.String(QueueStreamNameTag, streamName))

	streamChannel := make(chan *nats.Msg, workersCount)

	jsConfig := &nats.StreamConfig{
		Name:     streamName,
		Subjects: subjects,
		Storage:  storage,
	}

	workersPool := &ProducerWorkerPool{
		logger:           l,
		msgChannel:       streamChannel,
		jsConfig:         jsConfig,
		streamName:       streamName,
		subject:          subjects,
		storage:          storage,
		natsProducerConn: natsProducerConn,
	}

	for i := uint16(0); i < workersCount; i++ {
		ww := newProducerWorker(logger, i, streamChannel, streamName, subjects,
			natsProducerConn, make(chan bool))

		workersPool.workers = append(workersPool.workers, ww)
	}

	return workersPool, nil
}
