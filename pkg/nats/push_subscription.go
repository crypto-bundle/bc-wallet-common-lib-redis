package nats

import (
	"context"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

type pushSubscription struct {
	natsSubs *nats.Subscription
	natsConn *nats.Conn

	subjectName string
	groupName   string

	autoReSubscribe        bool
	autoReSubscribeCount   uint16
	autoReSubscribeTimeout time.Duration

	msgChannel chan *nats.Msg

	logger *zap.Logger
}

func (s *pushSubscription) Healthcheck(ctx context.Context) bool {
	if !s.natsConn.IsConnected() {
		s.logger.Warn("consumer lost nats connection")

		return false
	}

	if !s.natsSubs.IsValid() {
		s.logger.Warn("consumer lost nats subscription")

		if s.autoReSubscribe {
			return s.tryResubscribe(ctx)
		}

		return false
	}

	return true
}

func (s *pushSubscription) Init(ctx context.Context) error {
	return nil
}

func (s *pushSubscription) Run(ctx context.Context) error {
	subs, err := s.natsConn.ChanQueueSubscribe(s.subjectName, s.groupName, s.msgChannel)
	if err != nil {
		return err
	}

	s.natsSubs = subs

	return nil
}

func (s *pushSubscription) tryResubscribe(ctx context.Context) bool {
	var isSubscribed = false

	for i := uint16(0); i != s.autoReSubscribeCount; i++ {
		subs, err := s.natsConn.ChanQueueSubscribe(s.subjectName, s.groupName, s.msgChannel)
		if err != nil {
			s.logger.Warn("unable to re-subscribe", zap.Error(err),
				zap.Uint16(ResubscribeTag, i))

			time.Sleep(s.autoReSubscribeTimeout)
			continue
		}

		s.natsSubs = subs
		isSubscribed = true

		s.logger.Info("re-subscription success")
		break
	}

	return isSubscribed
}

func newPushSubscriptionService(logger *zap.Logger,
	natsConn *nats.Conn,

	subjectName string,
	groupName string,

	autoReSubscribe bool,
	autoReSubscribeCount uint16,
	autoReSubscribeTimeout time.Duration,

	msgChannel chan *nats.Msg,
) *pushSubscription {
	l := logger.Named("subscription")

	return &pushSubscription{
		natsConn: natsConn,
		natsSubs: nil, // it will be set @ run stage

		subjectName: subjectName,
		groupName:   groupName,

		autoReSubscribe:        autoReSubscribe,
		autoReSubscribeCount:   autoReSubscribeCount,
		autoReSubscribeTimeout: autoReSubscribeTimeout,

		msgChannel: msgChannel,
		logger:     l,
	}
}
