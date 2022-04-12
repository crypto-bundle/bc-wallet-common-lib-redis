package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go"
)

type Connection struct {
	connection *nats.Conn

	user      string
	password  string
	addresses []string

	retryTimeOut time.Duration
	retryCount   uint16
}

// GetConnection ...
func (c *Connection) GetConnection() *nats.Conn {
	return c.connection
}

func (c *Connection) Close() error {
	c.connection.Close()
	return nil
}

// NewConnection nats connection instance
func NewConnection(ctx context.Context, cfg configParams) (*Connection, error) {
	options := make([]nats.Option, 0)
	if cfg.IsRetryOnConnectionFailed() {
		options = append(options, nats.RetryOnFailedConnect(true),
			nats.MaxReconnects(int(cfg.GetNatsConnectionRetryCount())),
			nats.ReconnectWait(cfg.GetNatsConnectionRetryTimeout()),
		)
	}

	nats.RegisterEncoder(PROTOBUF_ENCODER, &ProtobufEncoder{})

	inst, err := nats.Connect(cfg.GetNatsJoinedAddresses(), options...)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		connection: inst,

		user:      cfg.GetNatsUser(),
		password:  cfg.GetNatsPassword(),
		addresses: cfg.GetNatsAddresses(),

		retryCount:   cfg.GetNatsConnectionRetryCount(),
		retryTimeOut: cfg.GetNatsConnectionRetryTimeout(),
	}

	return conn, nil
}
