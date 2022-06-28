package healthcheck

import (
	"context"
)

type config interface {
	IsDebug() bool
	GetAddress() string
	GetLivenessPath() string
	GetReadinessPath() string
	GetStartupPath() string
}

type Probe interface {
	Do(ctx context.Context) error
}

type HealthcheckService interface {
	Init() error
}
