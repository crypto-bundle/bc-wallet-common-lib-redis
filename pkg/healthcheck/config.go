package healthcheck

import (
	"github.com/kelseyhightower/envconfig"
)

const ConfigPrefix = "HEALTH_CHECK"

type Config struct {
	HttpPort      string `envconfig:"HTTP_PORT" default:"8081"`
	LivenessPath  string `envconfig:"LIVENESS_PATH" default:"/liveness"`
	ReadinessPath string `envconfig:"READINESS_PATH" default:"/readiness"`
	StartupPath   string `envconfig:"STARTUP_PATH" default:"/startup"`
	Debug         bool   `envconfig:"DEBUG"`
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
	return envconfig.Process(ConfigPrefix, c)
}
