package healthcheck

import (
	"context"
	"time"
)

type unitParamsService interface {
	IsDebug() bool
	GetProbeName() string

	GetHTTPListenAddress() string
	GetHTTPHandlerPath() string
	GetHTTPReadTimeout() time.Duration
	GetHTTPWriteTimeout() time.Duration
}

type probeService interface {
	Do(ctx context.Context) Status
	Init(ctx context.Context) error
	Shutdown(ctx context.Context) error
	ListenAndServe(ctx context.Context) error
}

type appDirectiveExecutionService interface {
	Do(status Status) (httpStatus int, httpBody string)
}
