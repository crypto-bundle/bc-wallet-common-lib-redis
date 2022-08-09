package metric

import (
	"github.com/kelseyhightower/envconfig"
)

type MetricConfig struct {
	HttpPort string `envconfig:"METRIC_HTTP_PORT" default:"8000"`
	Path     string `envconfig:"METRIC_PATH" default:"/metrics"`

	metricNamePrefix string
}

func (c *MetricConfig) GetAddress() string {
	return ":" + c.HttpPort
}

func (c *MetricConfig) GetPath() string {
	return c.Path
}

func (c *MetricConfig) GetMetricNamePrefix() string {
	return c.metricNamePrefix
}

func (c *MetricConfig) SetMetricNamePrefix(metricNamePrefix string) {
	c.metricNamePrefix = metricNamePrefix
}

// Prepare variables to static configuration
func (c *MetricConfig) Prepare() error {
	return envconfig.Process("", c)
}
