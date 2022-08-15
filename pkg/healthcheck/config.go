package healthcheck

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug         bool   `envconfig:"HEALTH_CHECK_DEBUG"`
	HttpPort      string `envconfig:"HEALTH_CHECK_HTTP_PORT" default:"8081"`
	LivenessPath  string `envconfig:"HEALTH_CHECK_LIVENESS_PATH" default:"/liveness"`
	ReadinessPath string `envconfig:"HEALTH_CHECK_READINESS_PATH" default:"/readiness"`
	StartupPath   string `envconfig:"HEALTH_CHECK_STARTUP_PATH" default:"/startup"`
}

func (c *Config) IsDebug() bool {
	return c.Debug
}

func (c *Config) GetAddress() string {
	return ":" + c.HttpPort
}

func (c *Config) GetLivenessPath() string {
	return c.LivenessPath
}

func (c *Config) GetReadinessPath() string {
	return c.ReadinessPath
}

func (c *Config) GetStartupPath() string {
	return c.StartupPath
}

// Prepare variables to static configuration
func (c *Config) Prepare() error {
	return envconfig.Process("", c)
}
