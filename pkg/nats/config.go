package nats

import (
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
)

type NatsConfig struct {
	//NatsHost     string `env:"NATS_HOST" envDefault:"nats"`
	//NatsPort     uint16 `env:"NATS_PORT" envDefault:"4222"`
	NatsAddresses string `env:"NATS_ADDRESSES" envDefault:"nats://ns-1:4223,nats://ns-2:4224,nats://na-3:4225"`
	NatsUser      string `env:"NATS_USER" envDefault:"nast"`
	NatsPassword  string `env:"NATS_PASSWORD" envDefault:"password"`

	NatsConnectionRetryOnFailed bool          `env:"NATS_CONNECTION_RETRY" envDefault:"true"`
	NatsConnectionRetryCount    uint16        `env:"NATS_CONNECTION_RETRY_COUNT" envDefault:"30"`
	NatsConnectionRetryTimeout  time.Duration `env:"NATS_CONNECTION_RETRY_TIMEOUT" envDefault:"15s"`

	NatsFlushTimeOut time.Duration `env:"NATS_FLUSH_TIMEOUT" envDefault:"15s"`

	NatsWorkersPerConsumer uint16 `env:"NATS_WORKER_PER_CONSUMER" envDefault:"5s"`

	nastAddresses []string
}

func (c *NatsConfig) GetNatsAddresses() []string {
	return c.nastAddresses
}

func (c *NatsConfig) GetNatsJoinedAddresses() string {
	return c.NatsAddresses
}

//func (c *NatsConfig) GetNatsHost() string {
//	return c.NatsHost
//}
//
//func (c *NatsConfig) GetNatsPort() uint16 {
//	return c.NatsPort
//}

func (c *NatsConfig) GetNatsUser() string {
	return c.NatsUser
}

func (c *NatsConfig) GetNatsPassword() string {
	return c.NatsPassword
}

func (c *NatsConfig) IsRetryOnConnectionFailed() bool {
	return c.NatsConnectionRetryOnFailed
}

func (c *NatsConfig) GetNatsConnectionRetryCount() uint16 {
	return c.NatsConnectionRetryCount
}

func (c *NatsConfig) GetNatsConnectionRetryTimeout() time.Duration {
	return c.NatsConnectionRetryTimeout
}

func (c *NatsConfig) GetFlushTimeout() time.Duration {
	return c.NatsFlushTimeOut
}

func (c *NatsConfig) GetWorkersCountPerConsumer() uint16 {
	return c.NatsWorkersPerConsumer
}

// Prepare variables to static configuration
func (c *NatsConfig) Prepare() error {
	err := env.Parse(c)
	if err != nil {
		return err
	}

	endpoints := strings.Split(c.NatsAddresses, ",")
	length := len(endpoints)
	if length < 1 {
		return nil
	}
	c.nastAddresses = endpoints

	return nil
}
