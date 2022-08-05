package metric

import (
	"github.com/kelseyhightower/envconfig"
)

type MetricConfig struct {
	HttpPort string `envconfig:"METRIC_HTTP_PORT" default:"8000"`
	Path     string `envconfig:"METRIC_PATH" default:"/metrics"`
}

func (c *MetricConfig) GetAddress() string {
	return ":" + c.HttpPort
}

func (c *MetricConfig) GetPath() string {
	return c.Path
}

// Prepare variables to static configuration
func (c *MetricConfig) Prepare() error {
	return envconfig.Process("", c)
}
