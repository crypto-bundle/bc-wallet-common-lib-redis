package zeromq

import (
	"errors"
	"time"

	zmq "github.com/pebbe/zmq4"
)

var (
	ErrConnectionClosed            = errors.New("connection closed")
	ErrReachedMaxReconnectionCount = errors.New("reached max reconnection count")
)

type Connection struct {
	socket *zmq.Socket

	reconnectMaxCount    uint16
	reconnectWaitTimeOut time.Duration

	endpoint string
}

func (c *Connection) Init() error {
	err := c.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) Connect() error {
	var reconnectCnt uint16
	for {
		err := c.socket.Connect(c.endpoint)
		if err == nil {
			return nil
		}

		if c.reconnectMaxCount != 0 {
			reconnectCnt++
			if reconnectCnt > c.reconnectMaxCount {
				return ErrReachedMaxReconnectionCount
			}
		}

		time.Sleep(c.reconnectWaitTimeOut)
	}

}

func (c *Connection) GetSocket() *zmq.Socket {
	return c.socket
}

func (c *Connection) Receive() ([]string, error) {
	msg, err := c.GetSocket().RecvMessage(0)
	if err == zmq.ErrorSocketClosed {
		return nil, ErrConnectionClosed
	} else if err != nil {
		return nil, err
	}

	return msg, nil
}

func (c *Connection) Close() error {
	return c.socket.Close()
}

// NewSubConnection zeromq subscriber connection instance
func NewSubConnection(cfg config) (*Connection, error) {
	socket, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		return nil, err
	}

	err = socket.SetPlainUsername(cfg.GetUsername())
	if err != nil {
		return nil, err
	}

	err = socket.SetPlainPassword(cfg.GetPassword())
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		socket:               socket,
		endpoint:             cfg.GetEndpoint(),
		reconnectMaxCount:    cfg.GetReconnectMaxCount(),
		reconnectWaitTimeOut: cfg.GetReconnectWaitTimeOut(),
	}

	return conn, nil
}
