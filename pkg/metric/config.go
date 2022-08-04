package metric

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpPort string `envconfig:"METRIC_HTTP_PORT" default:"8000"`
	Path     string `envconfig:"METRIC_PATH" default:"/metrics"`
}

func (c *Config) GetAddress() string {
	return ":" + c.HttpPort
}

func (c *Config) GetPath() string {
	return c.Path
}

// Prepare variables to static configuration
func (c *Config) Prepare() error {
	return envconfig.Process("", c)
}
