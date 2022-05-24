package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Endpoint string `env:"ENDPOINT" required:"true"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`

	ReconnectMaxCount    uint16        `env:"RECONNECT_MAX_COUNT" envDefault:"30"`
	ReconnectWaitTimeOut time.Duration `env:"RECONNECT_WAIT_TIMEOUT" envDefault:"15s"`
}

func (c *Config) GetEndpoint() string {
	return c.Endpoint
}

func (c *Config) GetUsername() string {
	return c.Username
}

func (c *Config) GetPassword() string {
	return c.Password
}

func (c *Config) GetReconnectMaxCount() uint16 {
	return c.ReconnectMaxCount
}

func (c *Config) GetReconnectWaitTimeOut() time.Duration {
	return c.ReconnectWaitTimeOut
}

// Prepare variables to static configuration
func (c *Config) Prepare(prefix string) error {
	return envconfig.Process(prefix, c)
}
